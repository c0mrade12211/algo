// Кодирование длин серий (Run-length encoding)
// Run-length encoding (RLE) is a form of lossless data compression in which runs of data (sequences in which the same data value occurs in many consecutive data elements) are stored as a single data value and count, rather than as the original run. 
//This is most efficient on data that contains many such runs, for example, simple graphic images such as icons, line drawings, Conway's Game of Life, and animations. 
//For files that do not have many runs, RLE could increase the file size.

// Example: WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW
// Output: 12W1B12W3B24W1B14W

package main

import (
	"fmt"
	"io/ioutil"
)

func rleEncode(data []byte) []byte {
	var result []byte
	i := 0
	for i < len(data) {
		count := 1
		for i+1 < len(data) && data[i] == data[i+1] {
			count++
			i++
		}
		result = append(result, byte(count))
		result = append(result, data[i])
		i++
	}
	return result
}

func rleDecode(data []byte) []byte {
	var result []byte
	i := 0
	for i < len(data) {
		count := int(data[i])
		char := data[i+1]
		for j := 0; j < count; j++ {
			result = append(result, char)
		}
		i += 2
	}
	fmt.Println("[+] You're compressed data : ", result)
	return result
}
func main() {

	fmt.Println("[+] Start compressed data by RLE Algo")
	fileData, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	compressedData := rleEncode(fileData)

	decodedData := rleDecode(compressedData)
	fmt.Println(string(decodedData))

	err = ioutil.WriteFile("output.txt", compressedData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[+] End compressed data by RLE Algo")
}
