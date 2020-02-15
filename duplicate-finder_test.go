package main

import (
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var files []filedesc
		_ = SearchFiles("/mnt/Music/Музыка/", &files)
	}

}

func TestSearch(b *testing.T) {
	var files []filedesc
	_ = SearchFiles("/mnt/Music/Музыка/", &files)
}
