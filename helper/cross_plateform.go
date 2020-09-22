package helper

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func OpenUrl(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "cmd"
		args = append(args, `/c`, `start`)
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func Notice(msg string) error {
	var shellfile string
	var cmd []string
	switch runtime.GOOS {
	case "windows":
		home,_:=Home()
		shellfile = fmt.Sprintf("%s/tmp/notice_file.vbs",home)
		content := `mshta vbscript:msgbox("do sports now",64,"notice")(window.close())`
		cmd = append(cmd, `cmd`, `/c`)
		// content 写入文件
		ioutil.WriteFile(shellfile, []byte(content), 777)
	default:
		home,_:=Home()
		shellfile = fmt.Sprintf("%s/tmp/notice_file.sh",home)
		content := `#!/bin/env sh
title="日常提醒"
content="记得喝水活动一下"
subtitle="记得喝水"
sound="Pon"
cmd=$(printf 'display notification "%s" with title "%s" subtitle "%s" sound name "%s"' "$content" "$title" "$subtitle" "$sound")
osascript -e "$cmd"
say -v Ting-ting $content`
		cmd = append(cmd, `sh`)
		// content 写入文件
		ioutil.WriteFile(shellfile, []byte(content), 777)
	}
	return exec.Command(strings.Join(cmd," "), shellfile).Start()
	//return exec.Command("sh", shellfile).Start()
}


// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}