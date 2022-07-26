package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const readFolder string = "./data"
const resultFolder string = "./result"
const pattern string = "*.html"

type CopyFile struct {
	sourceFile          string
	destinationFileName string
}

func main() {

	start := time.Now()

	fmt.Println("Gooo...")

	os.RemoveAll("./" + resultFolder)
	os.MkdirAll("./"+resultFolder, 0755)

	read(func(copyFile CopyFile) { copy(copyFile) })

	elapsed := time.Since(start)

	fmt.Println(elapsed)

	fmt.Println("Done")

}

func read(callback func(CopyFile)) {
	err := filepath.Walk(readFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
				return err
			} else if matched {
				copyFile := CopyFile{
					sourceFile:          path,
					destinationFileName: info.Name(),
				}
				callback(copyFile)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func copy(copyFile CopyFile) {
	input, err := ioutil.ReadFile(copyFile.sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(resultFolder+"/"+copyFile.destinationFileName, input, 0644)
	if err != nil {
		fmt.Println("Error creating", copyFile.destinationFileName)
		fmt.Println(err)
		return
	}
}
