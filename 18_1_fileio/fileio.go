package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/*
Most of the time we'll use `os.Open` or `os.Create`,
but we can open a file using `os.OpenFile(name, flag, perm)`
We can then specify:
	Exactly one of theses:
	- O_RDONLY	Open the file read only
	- O_WRONLY	Open the file write only
	- O_RDWR	Open the file read-write

	Or many of theses or'ed:
	- O_APPEND	Append data to the file
	- O_CREATE	Create a new file if none exists
	- O_EXCL	used with O_CREATE, file must not exists
	- O_SYNC	open for synchonous I/O
	- O_trunc	truncate regular writable file when opened

Examples:
f,err := os.Openfile("file.txt", O_WRONLY, 0644)

f,err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

*/

func main() {
	_, err := os.Stat("data.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

}
