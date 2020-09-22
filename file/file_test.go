package file

import "testing"

func newfile() *File {
	return NewFile("tmp/xxx.log")
}

func TestFile_Exists(t *testing.T) {
	t.Log(newfile())
}

func TestFile_IsDir(t *testing.T) {
	t.Log(newfile().IsDir())
}

func TestFilePutContents(t *testing.T) {
	t.Log(FilePutContents("tmp/xxx.log", []byte("xxx")))
}

func TestFileGetContents(t *testing.T) {
	t.Log(FileGetContents("tmp/xxx.log"))
}
