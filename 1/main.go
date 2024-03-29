package main

import (
	"fmt"
	"os"
)

func dirTree(out *os.File, path string, printFiles bool) error {

	if !printFiles {
		treeDirFile, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		fmt.Println(".")
		for _, v := range treeDirFile {
			fmt.Println(v)
			// fmt.Println(v.Name())
		}
	} else {
		fmt.Println("Here files tree.")
	}

	return nil
}

func main() {
	outLook := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Println("Error, need write 1 or 2 parametrs, example: \"go run main.go . (-f).\"")
		os.Exit(0)
	}

	path := os.Args[1]

	printFiles := false
	if len(os.Args) == 3 && os.Args[2] == "-f" {
		printFiles = true
	}

	err := dirTree(outLook, path, printFiles)
	if err != nil {
		fmt.Println("Error read files from directory:", err)
		os.Exit(0)
	}
}

// func dirTree1(in string) {
// 	treeDir, err := os.ReadDir(in)
// 	if err != nil {
// 		fmt.Println("Error read directory:", err)
// 	}

// 	for _, v := range treeDir {
// 		path := ""
// 		fmt.Println(v.Name())
// 		if v.IsDir() {
// 			path += v.Name() + "/"
// 			dirTree1(path)
// 		}
// 	}
// }

// out := os.Stdout
// if !(len(os.Args) == 2 || len(os.Args) == 3) {
// 	panic("usage go run main.go . [-f]")
// }
// path := os.Args[1]
// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
// err := dirTree(out, path, printFiles)
// if err != nil {
// 	panic(err.Error())
// }

// switch os.Args[1] {
// case "-f":
// 	dirTree()
// case ".":
// 	dirTree1(".")
// }
// fmt.Println(os.ReadDir(input[2]))
// fmt.Println(os.ReadFile(input[2]))
// }

// Утилита tree.
// Выводит дерево каталогов и файлов (если указана опция -f).
// Необходимо реализовать функцию `dirTree` внутри `main.go`.

// Мои комментарии:
// Проверка на некорректный ввод.

/*
const testFullResult = `├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│		└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
└───zzfile.txt (empty)
`
*/
