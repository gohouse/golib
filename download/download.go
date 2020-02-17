package download

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gohouse/file"
	"github.com/gohouse/random"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var UrlEmpty = errors.New("url can't be empty.")
var UrlFormatError = errors.New("url format is wrong.")
var UrlNotFound = errors.New("url not found.")

// ExistsOper 文件存在时,需要如何处理
type ExistsOperType int

const (
	EO_Cover  ExistsOperType = iota // 覆盖(默认)
	EO_Rename                       // 重命名
	EO_Skip                         // 跳过,不做任何处理
)

type Option struct {
	Dir        string
	FileName   string
	ExistsOper ExistsOperType
}

// Download 下载文件
// @param url string 下载地址
// @param savepath string 保存路径.
// 		如果未传参,则默认保存当前运行目录,文件名为原文件名
// 		如果传入的是文件名(可以带路径),则使用传入的文件名, 如: "abc"或"abc.jpg"或"/abc/xxx.jpg"代表文件, 会自动创建目录和文件,不存在的话
// 		如果传入的是目录,则默认使用原文件名命名,如: "abc/"代表目录, 会自动创建目录,不存在的话
func Download(url string, opts ...Option) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp == nil {
		return UrlNotFound
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)

	return DownloadSave(url, &body, opts...)
}

func DownloadSave(url string, body *[]byte, opts ...Option) (err error) {
	var opt Option
	if len(opts) > 0 {
		opt = opts[0]
	}
	if url == "" {
		return UrlEmpty
	}
	if !strings.Contains(url, "/") {
		return UrlFormatError
	}
	var paths = strings.Split(url, "/")
	if len(paths) == 0 {
		return UrlFormatError
	}
	if opt.FileName == "" {
		opt.FileName = paths[len(paths)-1]
	}
	var name = strings.Replace(fmt.Sprintf("%s/%s", opt.Dir, opt.FileName), "//", "/", -1)
	if file.FileExists(name) {
		switch opt.ExistsOper {
		case EO_Skip:
			return
		case EO_Rename:
			name = Rename(name)
		case EO_Cover:
		}
	}

	out := file.NewFile(name).OpenFile()
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(*body))
	return
}

func GetDownloadName(url string,prefix ...string) string {
	var paths = strings.Split(url, "/")
	if len(paths) == 0 {
		return ""
	}
	name := paths[len(paths)-1]
	if len(prefix)>0{
		return strings.Replace(fmt.Sprintf("%s/%s",prefix[0], name),"//","/",-1)
	}
	return name
}

func Rename(src string, dst ...string) string {
	var realname string
	if len(dst) > 0 {
		realname = dst[0]
	}

	if realname == "" {
		var dir = path.Dir(src)
		var ext = path.Ext(src)
		if dir == "" {
			dir = "./"
		}
		randname := random.Random(32)
		if ext == "" {
			realname = fmt.Sprintf("%s/%s", dir, randname)
		} else {
			realname = fmt.Sprintf("%s/%s.%s", dir, randname, ext)
		}
		realname = strings.Replace(realname, "//", "/", -1)
	}

	if file.FileExists(realname) {
		return Rename(src)
	}

	return realname
}
