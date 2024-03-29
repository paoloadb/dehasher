package main

import (
	"bufio"
	"dehasher/functions"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: dehash [wordlist] [algorithm] [hash]")
	}
	fileNamePath := os.Args[1]
	algo := os.Args[2]
	referenceHash := os.Args[3]

	fmt.Println("Wordlist is", fileNamePath)
	fmt.Println("Hash is", referenceHash)
	fmt.Println("Algorithm is", algo)

	readFile, err := os.Open(fileNamePath)
	if err != nil {
		log.Fatal("File access error!")
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)
	fileScan.Split(bufio.ScanLines)

	fmt.Println("Starting decryption...")
	var retval bool
	for fileScan.Scan() {
		retval = functions.SelectAlgorithm(referenceHash, fileScan.Text(), algo)
		if retval {
			break
		}
	}
	if retval {
		fmt.Println("Match! >> ", fileScan.Text())
	} else {
		fmt.Println("No Match Found")
	}
}
