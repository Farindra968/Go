package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err) // Panic if the file cannot be opened
	}

	//f.name() returns the name of the file as a string, and f.Stat() returns a FileInfo object that contains information about the file, such as its size and permissions. We handle any errors that may occur when retrieving this information by panicking if an error occurs. Finally, we print the file name, size, and mode to the console.
	fileName := f.Name()


	fileInfo, err := f.Stat()
	if err != nil {
		panic(err) // Panic if the file information cannot be retrieved
	}
	fmt.Println("File Information:", fileName, fileInfo.Size(), fileInfo.Mode())
	fmt.Println("Is Directory:", fileInfo.IsDir()) // Check if the file is a directory
	fmt.Println("Last Modification Time:", fileInfo.ModTime()) // Print the last modification time of the file
	fmt.Println("System-specific Information:", fileInfo.Sys()) // Print system-specific information about the file
	fmt.Println("Size (bytes):", fileInfo.Size()) // Print the size of the file in bytes
	fmt.Println("Permissions:", fileInfo.Mode()) // Print the file mode (permissions)
	fmt.Println("Name:", fileInfo.Name()) // Print the name of the file
	fmt.Println("Last Modification Time + 24 hours:", fileInfo.ModTime().Add(24 * time.Hour)) // Print the last modification time of the file plus 24 hours
	fmt.Println("Is After Current Time:", fileInfo.ModTime().After(time.Now())) // Print whether the last modification time of the file is after the current time
	fmt.Println("Last Modification Time + 1 year, 1 month, 1 day:", fileInfo.ModTime().AddDate(1, 1, 1)) // Print the last modification time of the file plus one year, one month, and one day

	// Reading the file content
	content, err := os.Open("example.txt")
	if err != nil {
		panic(err) // Panic if the file cannot be opened for reading
	}

	defer content.Close() // Ensure the file is closed after reading

	buffer:= make([]byte, 100) // Create a buffer to hold the file content
	d, err := content.Read(buffer)
	if err != nil {
		panic(err) // Panic if there is an error reading the file
	}

	// Print each byte of the buffer as a character
	for i:=0; i < len(buffer); i++ {
		fmt.Printf("%c", string(buffer[i])) // Print each byte of the buffer as a character
	}

	// Print the file content as a string, trimming any whitespace
	fmt.Println("File Content", d, strings.TrimSpace(string(buffer))) 

	// Other ways to read the file content
	// Read the entire file content into a byte slice
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err) // Panic if there is an error reading the file
	}
	fmt.Println("File Content using ReadFile:", string(data)) // Print the file content as a string

}