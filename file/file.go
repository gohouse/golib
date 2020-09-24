package file

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// File 文件对象
type File struct {
	file string
}

// NewFile 初始化文件对象
func NewFile(file string) *File {
	f := &File{file: file}
	err := FileInit(file)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// FileInit 初始化文件目录
func FileInit(file string) (err error) {
	if !FileExists(file) {
		// 检查是否结尾是否存在点, 如果有点则吧最后一个当做文件处理
		dir := filepath.Dir(file)
		if dir != "." {
			// 递归创建目录
			if runtime.GOOS == "darwin" {
				mask := syscall.Umask(0)
				defer syscall.Umask(mask)
				err = os.MkdirAll(dir, 0766)
			} else {
				err = os.MkdirAll(dir, 0766)
			}
		}
	}
	return
}

// FileExists 判断所给路径文件/文件夹是否存在
func FileExists(file string) bool {
	_, err := os.Stat(file) //os.Stat获取文件信息
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

// File.IsDir 判断所给路径是否为文件夹
func (f *File) IsDir() bool {
	s, err := os.Stat(f.file)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func (f *File) IsFile() bool {
	if FileExists(f.file) {
		return !f.IsDir()
	}
	return false
}

// Write 写入文件内容
func (f *File) Write(p []byte, mods ...int) (n int, err error) {
	fp := f.OpenFile(mods...)
	defer fp.Close()
	return fp.Write(p)
}

// ReadFile 读取文件内容
func (f *File) ReadFile() ([]byte, error) {
	file := f.OpenFile()
	return ioutil.ReadAll(file)
}

// OpenFile os.OpenFile()
func (f *File) OpenFile(mods ...int) *os.File {
	var mod = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	//mod = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if len(mods) > 0 {
		mod = mods[0]
	}
	fp, err := os.OpenFile(f.file, mod, 0766)
	if err != nil {
		panic(err.Error())
	}
	return fp
}

// ReplaceFileContent 替换文件内容
func ReplaceFileContent(filePath, search, replacement string) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	newFile := strings.Replace(string(file), search, replacement, -1)

	fp, err := os.Create(filePath)
	if err != nil {
		return
	}
	_, err = fp.Write([]byte(newFile))
	if err != nil {
		return
	}
	return fp.Close()
}

// GetAllFiles 递归获取指定目录下的所有go文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, strings.Replace(dirPth+PthSep+fi.Name(), "//", "/", -1))
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			//ok := strings.HasSuffix(fi.Name(), ".go")
			//if ok {
			files = append(files, strings.Replace(dirPth+PthSep+fi.Name(), "//", "/", -1))
			//}
		}
	}

	// read child file
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

// FilePutContents 替换或创建文件内容
func FilePutContents(file string, data []byte, mods ...int) (n int, err error) {
	return NewFile(file).Write(data, mods...)
}

// FileGetContents 获取文件内容
func FileGetContents(file string) (data string, err error) {
	res, err := NewFile(file).ReadFile()
	if err != nil {
		return
	}
	if res == nil {
		return
	}
	return string(res), nil
}

func Tail_f(filename string, lines int64) (rows string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	fileInfo, _ := file.Stat()
	buf := bufio.NewReader(file)
	offset := fileInfo.Size() % 8192
	data := make([]byte, 8192) // 一行的数据
	totalByte := make([][][]byte, 0)
	readLines := int64(0)
	for i := int64(0); i <= fileInfo.Size()/8192; i++ {
		readByte := make([][]byte, 0) // 读取一页的数据
		file.Seek(fileInfo.Size()-offset-8192*i, io.SeekStart)
		data = make([]byte, 8192)
		n, err := buf.Read(data)
		if err == io.EOF {
			if strings.TrimSpace(string(bytes.Trim(data, "\x00"))) != "" {
				readLines++
				readByte = append(readByte, data)
				totalByte = append(totalByte, readByte)
			}
			if readLines > lines {
				break
			}
			continue
		}
		if err != nil {
			log.Println("Read file error:", err)
			return
		}
		strs := strings.Split(string(data[:n]), "\n")
		if len(strs) == 1 {
			b := bytes.Trim([]byte(strs[0]), "\x00")
			if len(b) == 0 {
				continue
			}
		}
		if (readLines + int64(len(strs))) > lines {
			strs = strs[int64(len(strs))-lines+readLines:]
		}
		for j := 0; j < len(strs); j++ {
			readByte = append(readByte, bytes.Trim([]byte(strs[j]+"\n"), "\x00"))
		}
		readByte[len(readByte)-1] = bytes.TrimSuffix(readByte[len(readByte)-1], []byte("\n"))
		totalByte = append(totalByte, readByte)
		readLines += int64(len(strs))

		if readLines >= lines {
			break
		}
	}
	totalByte = ReverseByteArray(totalByte)
	return ByteArrayToString(totalByte)
}

func ReverseByteArray(s [][][]byte) [][][]byte {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
	return s
}

func ByteArrayToString(buf [][][]byte) string {
	str := make([]string, 0)
	for _, v := range buf {
		for _, vv := range v {
			str = append(str, string(vv))
		}
	}
	return strings.Join(str, "")
}
