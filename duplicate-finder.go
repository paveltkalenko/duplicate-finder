package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type filedesc struct {
	path     string
	fileInfo os.FileInfo
}

var rootDir string = "/home/pavel/test/"

func main() {
	var files []filedesc
	fmt.Println(files)
	err := SearchFiles(rootDir, &files)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, f := range files {
		fmt.Println(filepath.Join(f.path, f.fileInfo.Name()), f.fileInfo.Size())
	}
	fmt.Println("hello test world")
}

func SearchFiles(root string, files *[]filedesc) error {
	fmt.Println(root)
	f, err := os.Open(root)
	defer f.Close()
	if err != nil {
		return err
	}

	fileInfo, err := f.Readdir(-1)

	if err != nil {
		return err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			fmt.Println("try to enter directory ", file.Name(), *files)
			SearchFiles(filepath.Join(root, file.Name()), files)
		}

		//filepath.Join(root, file.Name())
		desc := filedesc{
			path:     filepath.ToSlash(root),
			fileInfo: file,
		}
		*files = append(*files, desc)
	}
	return nil
}

/*
117  func SameFile(fi1, fi2 FileInfo) bool {
	118  	fs1, ok1 := fi1.(*fileStat)
	119  	fs2, ok2 := fi2.(*fileStat)
	120  	if !ok1 || !ok2 {
	121  		return false
	122  	}
	123  	return sameFile(fs1, fs2)
	124  }*/
