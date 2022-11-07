package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	filewithpath := make([]string, 0)
	//read the file file.conf in the same directory line by line and store it in filewithpath
	file, err := os.Open("file.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bufio.NewScanner(file)
	for buf.Scan() {
		filewithpath = append(filewithpath, buf.Text())
	}

	defer file.Close()

	//loop and sleep for 10 minutes
	for {

		fmt.Println(filewithpath)
		for _, file2 := range filewithpath {

			//delete the file
			files, err := filepath.Glob(file2)
			if err != nil {
				panic(err)
			}
			fmt.Println(files)
			for _, f := range files {
				if erro := os.Remove(f); erro != nil {
					fmt.Println(erro)
				}
			}

		}

		//sleep for 10 minutes
		time.Sleep(10 * time.Minute)

	}

}
