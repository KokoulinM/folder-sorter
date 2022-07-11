package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	var fromPath string
	var toPath string

	fmt.Print("Source directory: ")
	_, err := fmt.Scanln(&fromPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Target directory: ")
	_, err = fmt.Scanln(&toPath)
	if err != nil {
		log.Fatal(err)
	}

	var dumpName = strconv.Itoa(int(time.Now().Unix()))
	var targetPath = fmt.Sprintf("%s/%s", toPath, dumpName)

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

			srcFile := fmt.Sprintf("%s/%s", fromPath, v.Name())
			finalDir := fmt.Sprintf("%s/%s", targetPath, formatDir)

			err := makeDir(finalDir)
			if err != nil {
				log.Println(err)
			}

			moveFile(srcFile, finalDir+"/"+v.Name())
		}
	}

	// If the folder is empty, delete it
	err = os.Remove(targetPath)
	if err != nil {
		log.Fatal(err)
	}
}
