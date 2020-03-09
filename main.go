package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type tree struct {
	depth []int
}

func sortFiles(info []os.FileInfo) {
	var back int
	var temp os.FileInfo
	for i := 1; i < len(info); i++ {
		back = i
		for back > 0 && info[back-1].Name() > info[back].Name() {
			temp = info[back-1]
			info[back-1] = info[back]
			info[back] = temp
		}

	}
}

func findLastIndex(info []os.FileInfo, printFiles bool) int {
	var lastIndex int
	if printFiles {
		lastIndex = len(info) - 1
	} else {
		for i := len(info) - 1; i > -1; i-- {
			if info[i].IsDir() {
				lastIndex = i
				break
			}
		}
	}
	return lastIndex
}

func printSize(out io.Writer, info os.FileInfo) {
	if sizeToPrint := info.Size(); sizeToPrint == 0 {
		fmt.Fprintln(out, " (empty)")
	} else {
		fmt.Fprint(out, " (", sizeToPrint, "b)\n")
	}
}

func formattedPrint(out io.Writer, last int, depth []int, info os.FileInfo) {
	for _, d := range depth {
		if d == 0 {
			fmt.Fprint(out, "│\t")
		} else {
			fmt.Fprint(out, "\t")
		}
	}
	if last == 1 {
		fmt.Fprint(out, "└───")
	} else {
		fmt.Fprint(out, "├───")
	}
	fmt.Fprint(out, info.Name())
	if info.IsDir() {
		fmt.Fprint(out, "\n")
	} else {
		printSize(out, info)
	}
}

func recursiveTree(out io.Writer, path string, printFiles bool, tr *tree) error {
	var last int
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Error with ReadDir")
	}
	sortFiles(files)
	lastIndex := findLastIndex(files, printFiles)
	for i, f := range files {
		if i == lastIndex {
			last = 1
		} else {
			last = 0
		}
		if f.IsDir() {
			formattedPrint(out, last, tr.depth, f)
			tr.depth = append(tr.depth, last)
			if err := recursiveTree(out, path+string(os.PathSeparator)+f.Name(), printFiles, tr); err != nil {
				return fmt.Errorf("Error with ReadDir")
			}
			tr.depth = tr.depth[:len(tr.depth)-1]
		} else if printFiles {
			formattedPrint(out, last, tr.depth, f)
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	tr := new(tree)
	if err := recursiveTree(out, path, printFiles, tr); err != nil {
		return fmt.Errorf("Error with ReadDir")
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
