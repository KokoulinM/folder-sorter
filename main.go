package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func getFormat(fileName string) string {
	format := regexp.MustCompile(`^.*(.*\.)`)

	result := format.ReplaceAllString(fileName, "")

	return result
}

func moveFile(dst, trg string) error {
	err := os.Chmod(dst, 0777)
	if err != nil {
		return err
	}

	err = os.Rename(dst, trg)
	if err != nil {
		return err
	}

	return nil
}

func makeDir(name string) error {
	err := os.Mkdir(name, os.ModeDir)
	if err != nil {
		return err
	}

	err = os.Chmod(name, 0777)
	if err != nil {
		return err
	}

	return nil
}

func scan(str *string, info string) {
	fmt.Print(info)
	_, err := fmt.Scanln(str)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var fromPath string
	var toPath string

	scan(&fromPath, "Source directory: ")
	scan(&toPath, "Target directory: ")

	var dumpName = strconv.Itoa(int(time.Now().Unix()))
	var targetPath = filepath.Join(toPath, dumpName)

	files, err := ioutil.ReadDir(fromPath)
	if err != nil {
		log.Fatal(err)
	}

	err = makeDir(targetPath)
	if err != nil {
		log.Println(err)
	}

	for _, v := range files {
		if !v.IsDir() {
			formatDir := getFormat(v.Name())

			srcFile := filepath.Join(fromPath, v.Name())
			finalDir := filepath.Join(targetPath, formatDir)

			err := makeDir(finalDir)
			if err != nil {
				log.Println(err)
			}

			moveFile(srcFile, filepath.Join(finalDir, v.Name()))
		}
	}

	// If the folder is empty, delete it
	err = os.Remove(targetPath)
	if err != nil {
		log.Fatal(err)
	}
}
