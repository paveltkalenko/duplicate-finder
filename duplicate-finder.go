package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Filedesc struct {
	Path     string
	FileInfo os.FileInfo
	Hash     string
}

var rootDir string = "/mnt/Music/Музыка"

func main() {
	var files []Filedesc

	err := SearchFiles(rootDir, &files)
	if err != nil {
		log.Fatal(err)
		return
	}

	GetSameFiles(files)

	// простой алгоритм поиска одинаковых значений
	// 1. Сортируем
	// 2. Группируем по размеру.
	//
	/*
		sort.Slice(files, func(i, j int) bool { return files[i].fileInfo.Size() > files[j].fileInfo.Size() })
		for _, f := range files {

			if last != nil && last.Size() == f.fileInfo.Size() && !os.SameFile(last, f.fileInfo) {

				}

				fmt.Println(filepath.Join(f.path, f.fileInfo.Name()), "\t\t\t\t\t\t\t\t\t\t", f.fileInfo.Size())
			}

			last = f.fileInfo
		}
	*/
	fmt.Println("Total ", len(files))

}

func GetSameFiles(files []Filedesc) map[string][]Filedesc {
	//	filesSizes := make(map[int64]*list.List)
	filesSizes := make(map[int64][]Filedesc)
	for _, f := range files {
		element, ok := filesSizes[f.FileInfo.Size()]
		if !ok {
			//	element = list.New()
			element = make([]Filedesc, 0, 2)
			filesSizes[f.FileInfo.Size()] = element
		}
		//element.PushBack(f)
		filesSizes[f.FileInfo.Size()] = append(element, f)
	}

	for k, v := range filesSizes {
		if len(v) > 1 {
			for _, v2 := range v {
				v2.Hash = FindHashFiles(filepath.Join(v2.Path, v2.FileInfo.Name()))
				fmt.Println("v2.Path = ", filepath.Join(v2.Path, v2.FileInfo.Name()), "  v2.hash = ", v2.Hash)
			}
			fmt.Println("key = ", k, "v = ", v)
			//fmt.Println("key = ", k, "len = ", v.Len(), " value = ", v)
		} else {
			delete(filesSizes, k)
		}

	}

	return nil

}
func FindHashFiles(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SearchFiles(root string, files *[]Filedesc) error {
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
			SearchFiles(filepath.Join(root, file.Name()), files)
		} else {
			desc := Filedesc{
				Path:     filepath.ToSlash(root),
				FileInfo: file,
			}
			*files = append(*files, desc)
		}
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
