package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getConfigDirPath(homedir string) string {
	return filepath.Join(homedir, ".config/themeology")
}

func getConfigFilePath(configPath string) string {
	path := filepath.Join(configPath, "themeology.conf")
	fmt.Printf("Config File: %s\n", path)
	return path
}

func ensureDir(dirpath string) {
	err := os.MkdirAll(dirpath, os.ModePerm)
	checkError(err)
}

func ensureFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	checkError(err)

	_, err = f.Seek(0, io.SeekStart)
	checkError(err)

	return f
}
