package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func temp_file_and_dirs_golang() {
	f, err := os.CreateTemp("", "Samples")
	checkerr(err)
	fmt.Println("Temp file name is ", f.Name())

	defer os.Remove(f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	checkerr(err)

	dname, err := os.MkdirTemp("", "sampledir")
	checkerr(err)
	fmt.Println("temp dir name", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file2")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	checkerr(err)

}
