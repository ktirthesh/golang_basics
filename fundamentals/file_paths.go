package main

import (
	"fmt"
	"path/filepath"
)

func file_paths_golang() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	fmt.Println("p:", p)
}
