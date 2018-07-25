package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

// allFlags is a struct to hold all of the flags that will be used
type allFlags struct {
	fromPtr *string
	toPtr   *string
	debug   *bool
	tail    []string
}

// parseFlags returns the all of the flags parsed
func parseFlags() allFlags {
	flags := allFlags{
		fromPtr: flag.String("from", "rar", "input type"),
		toPtr:   flag.String("to", "mkv", "output type"),
		debug:   flag.Bool("debug", false, "enable for extra debug output"),
		tail:    flag.Args()}
	flag.Parse()
	return flags
}

// printFlags prints out debugging info for users and developers
func printFlags(inFlags allFlags) {
	fmt.Println(reflect.TypeOf(inFlags.fromPtr))
	fmt.Println("from:", *(inFlags.fromPtr))
	fmt.Println("to:", *(inFlags.toPtr))
	fmt.Println("tail", inFlags.tail)
	fmt.Println(reflect.TypeOf(inFlags.tail))
}

// findFiles returns []string of filenames that match the *.pattern by using Glob
func findFiles(pattern string) (matches []string, err error) {
	return filepath.Glob("*." + pattern)
}

// renameFiles takes in fileNames and a newExt and renames the files with the new ext
func renameFiles(fileNames []string, newExt string) (newFileNames []string, err error) {
	for _, fileName := range fileNames {
		ext := path.Ext(fileName)
		outfile := fileName[0:len(fileName)-len(ext)] + "." + newExt
		err := os.Rename(fileName, outfile)
		if err != nil {
			return newFileNames, nil
		}
		newFileNames = append(newFileNames, outfile)
	}

	return newFileNames, nil
}

// main is main yo
func main() {

	// parse the flags
	flags := parseFlags()
	if *(flags.debug) {
		printFlags(flags)
	}

	// find all files with the desired "from" file extension
	fileNames, err := findFiles(*(flags.fromPtr))
	if err != nil {
		log.Fatal(err)
	}

	// rename all found files with the desired "to" file extension
	renameFiles(fileNames, *(flags.toPtr))

}
