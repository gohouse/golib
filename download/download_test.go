package download

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	var url = "https://novelbase.oss-cn-hongkong.aliyuncs.com/book_images/00726111fd90f720fad4d6ab5f83dede.jpeg"
	//var savepath = "./abc/a.jpeg"
	err := Download(url)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestRename(t *testing.T) {
	fullFilename := "xxx/abc/test.txt"
	fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	fileDir := path.Dir(fullFilename)
	fmt.Println("fileDir = ", fileDir)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)
}
