package util

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

// GetUserHome 获取运行用户home目录 跨os支持
func GetUserHome() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support
	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix / Unix-like
	return homeUnix()
}

func homeUnix() (string, error) {
	// 找环境变量 $HOME
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// 失败后 尝试shell命令
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
