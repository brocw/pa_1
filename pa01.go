/*============================================================================
| Assignment: pa01 - Calculate the checksum of an input file given:
|                  -> the name of the input file,
|                  -> the checksum size of either 8, 16, or 32 bits
| Author: Broc Weselmann
| Language: go
|                 -> tested on go version go1.24.4 linux/amd64
| To Compile: go build pa01.go
| To Execute: ./pa01 inputFilename.txt checksumSize
|                  -> where inputFilename.txt is the input file
|                  -> and checksumSize is either 8, 16, or 32
| Note:
| All input files are simple 8 bit ASCII input
| All execute commands above have been tested on Eustis
| Class: CIS3360 - Security in Computing - Summer 2025
| Instructor: McAlpin
| Due Date: 06/22/25
+===========================================================================*/

package main

import (
	"fmt"
	"os"
	"strconv"
	//"math/bits"
)

// File check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Generic error functiion (prints to stderr)
func errorOut(s string) {
	err := fmt.Errorf("%s\n", s)
	fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(1)
}

// 8-bit checksum calculation
func check8(file []byte) (uint8, int) {
	var count int
	var running_sum uint8
	for _, value := range file {
		count += 1
		running_sum += value
	}

	return running_sum, count
}

// 16-bit checksum calculation
func check16(file []byte) (int16, int) {
	return 0, 0
}

// 32-bit checksum calculation
func check32(file []byte) (int32, int) {
	return 0, 0
}

func main() {
	if len(os.Args) != 3 {
		errorOut("Must have two arguments: File and Checksum size")
	}
	
	fileName := os.Args[1]
	checksumSizeStr := os.Args[2]

	checksumSize, checksumErr := strconv.Atoi(checksumSizeStr) 
	if checksumErr != nil {
		errorOut("Error converting Checksum size to int")
	}

	// Open file
	fileContents, err := os.ReadFile(fileName)
	check(err)

	switch checksumSize {
	case 8:
		checksum, characterCount := check8(fileContents)
		defer fmt.Printf("%2d bit checksum is %8X for all %4d chars\n", checksumSize, checksum, characterCount)
	case 16:
		checksum, characterCount := check16(fileContents)
		defer fmt.Printf("%2d bit checksum is %8X for all %4d chars\n", checksumSize, checksum, characterCount)
	case 32:
		checksum, characterCount := check32(fileContents)
		defer fmt.Printf("%2d bit checksum is %8X for all %4d chars\n", checksumSize, checksum, characterCount)
	default:
		errorOut("Valid checksum sizes are 8, 16, or 32")
	}
	
	for count, value := range fileContents {
		if count % 80 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", value)
	}
	fmt.Printf("\n")
}

/*=============================================================================
| I Broc Weselmann (br142931) affirm that this program is
| entirely my own work and that I have neither developed my code together with
| any another person, nor copied any code from any other person, nor permitted
| my code to be copied or otherwise used by any other person, nor have I
| copied, modified, or otherwise used programs created by others. I acknowledge
| that any violation of the above terms will be treated as academic dishonesty.
+=============================================================================*/
