package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	var result string
	if len(os.Args) == 2 {
		result = os.Args[1]
	} else {
		// list all files in current folder
		files, err := filepath.Glob("*.log")
		if err != nil {
			log.Fatalln(err)
		}

		// prompt for a file
		prompt := promptui.Select{
			Label: "Select a file",
			Items: files,
		}

		_, result, err = prompt.Run()
		if err != nil {
			log.Fatalln("Prompt failed", err)
		}
	}

	// open file for reading
	fin, err := os.Open(result)
	if err != nil {
		log.Fatalln(err)
	}
	defer fin.Close()

	// create file for writing
	name := strings.TrimSuffix(result, filepath.Ext(result)) + ".txt"
	fout, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	// encryption mask
	var mask = []byte{0x7C, 0xBD, 0x81, 0x9F, 0x3D, 0x93, 0xE2, 0x56, 0x2A, 0x73, 0xD2, 0x3E, 0xF2, 0x83, 0x95, 0xBF}
	buf := make([]byte, 1024)

	br := bufio.NewReader(fin)
	bw := bufio.NewWriter(fout)
	for {
		n, err := br.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		// apply mask
		for i := 0; i < n; i++ {
			buf[i] = buf[i] ^ mask[i%16]
		}
		bw.Write(buf[:n])
	}
	bw.Flush()

}
