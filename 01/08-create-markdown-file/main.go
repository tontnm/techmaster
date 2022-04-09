package main

import "os"

func main() {
	createMarkdownFile("Ton.md","Thanh.md")
	createMarkdownFile()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createMarkdownFile(fileName ...string) {
	var filename string

	if len(fileName) == 0 {
		filename = "ReadMe.md"
	} else {
		filename = fileName[0]
	}

	f, err := os.Create(filename)
	checkError(err)

	defer f.Close()

	_, err = f.WriteString("Hello Markdown\n")
	checkError(err)

	_, err = f.WriteString("Introduction")
	checkError(err)

	err = f.Sync()
	checkError(err)
}
