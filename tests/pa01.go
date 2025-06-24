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
	"encoding/binary"
	"bytes"
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
	var running_sum uint8
	for _, value := range file {
		running_sum += value
	}

	return running_sum, len(file)
}

// 16-bit checksum calculation
func check16(file []string) (uint16, int) {
	var running_sum uint16
	for _, value := range file {
		running_sum += binary.BigEndian.Uint16([]byte(value))
	}

	return running_sum, len(file) * 2
}

// Splitting byte array into string slices
// of two characters - includes the padded X's
func split_2(file []byte) ([]string, int) {
	numPairs := (len(file) + 1) / 2
	arr := make([]string, numPairs)
	var index int
	var num_x int

	for i := 0; i < len(file); i += 2 {
		var buffer bytes.Buffer
		buffer.WriteByte(file[i])

		if i + 1 < len(file) {
			buffer.WriteByte(file[i + 1])
		} else {
			buffer.WriteByte('X')
			num_x += 1
		}

		arr[index] = buffer.String()
		index += 1
	}

	return arr, num_x
}

// 32-bit checksum calculation
func check32(file []string) (uint32, int) {
	var running_sum uint32
	for _, value := range file {
		running_sum += binary.BigEndian.Uint32([]byte(value))
	}

	return running_sum, len(file) * 4
}

// Splitting byte array into string slices
// of four characters - includes the padded X's
func split_4(file []byte) ([]string, int) {
	numPairs := (len(file) + 3) / 4
	arr := make([]string, numPairs)
	var index int
	var num_x int

	for i := 0; i < len(file); i+= 4 {
		var buffer bytes.Buffer

		for j := 0; j < 4; j++ {
			byteIndex := i + j

			if byteIndex < len(file) {
				buffer.WriteByte(file[byteIndex])
			} else {
				buffer.WriteByte('X')
				num_x += 1
			}
		}

		arr[index] = buffer.String()
		index += 1
	}

	return arr, num_x
}

// Prints the file from the byte array itself
func print_file(file []byte, num_x int) {
	for count, value := range file {
		if count % 80 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", value)
	}

	for i := 0; i < num_x; i++ {
		fmt.Printf("X")
	}
	fmt.Printf("\n")
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
		print_file(fileContents, 0)
		fmt.Printf("%2d bit checksum is %8x for all %4d chars\n", checksumSize, checksum, characterCount)
	case 16:
		file_arr, num_x := split_2(fileContents)
		checksum, characterCount := check16(file_arr)
		print_file(fileContents, num_x)
		fmt.Printf("%2d bit checksum is %8x for all %4d chars\n", checksumSize, checksum, characterCount)
	case 32:
		file_arr, num_x := split_4(fileContents)
		checksum, characterCount := check32(file_arr)
		print_file(fileContents, num_x)
		fmt.Printf("%2d bit checksum is %8x for all %4d chars\n", checksumSize, checksum, characterCount)
	default:
		errorOut("Valid checksum sizes are 8, 16, or 32")
	}
}

/*=============================================================================
| I Broc Weselmann (br142931) affirm that this program is
| entirely my own work and that I have neither developed my code together with
| any another person, nor copied any code from any other person, nor permitted
| my code to be copied or otherwise used by any other person, nor have I
| copied, modified, or otherwise used programs created by others. I acknowledge
| that any violation of the above terms will be treated as academic dishonesty.
+=============================================================================*/
