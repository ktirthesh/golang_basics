package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func checke(e error) {
	if e != nil {
		panic(e)
	}
}

func direcrtories_golang() {
	err := os.Mkdir("subdir", 0755)
	checke(err)

	defer os.RemoveAll("subdir")
	createEmptyFile := func(name string) {
		d := []byte("")
		checke(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")
	err = os.MkdirAll("subdir/parent/child", 0755)
	checke(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file")
	createEmptyFile("subdir/parent/child/file4")
	c, err := os.ReadDir("subdir/parent")
	checke(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	err = os.Chdir("subdir/parent/child")
	checke(err)
	c, err = os.ReadDir(".")
	checke(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	err = os.Chdir("../../..")
	checke(err)

	fmt.Println("visiter  subdir")
	err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
