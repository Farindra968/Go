package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	baseDir := findWorkingDir()
	examplePath := filepath.Join(baseDir, "example.txt")

	fmt.Println("=== File metadata and open methods ===")
	file, err := os.Open(examplePath)
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size(), "bytes")
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Last Modified:", fileInfo.ModTime().Format(time.RFC3339))
	fmt.Println("Current mode string:", fileInfo.Mode().String())
	fmt.Println("System info:", fileInfo.Sys())
	_ = file.Close()

	fmt.Println("\n=== Reading methods ===")
	contentFile, err := os.Open(examplePath)
	if err != nil {
		panic(err)
	}
	defer contentFile.Close()

	buffer := make([]byte, 128)
	n, err := contentFile.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	fmt.Printf("Read using File.Read: %q\n", strings.TrimSpace(string(buffer[:n])))

	data, err := os.ReadFile(examplePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read using os.ReadFile: %q\n", strings.TrimSpace(string(data)))

	fmt.Println("\n=== Directory methods ===")
	dir, err := os.Open(baseDir)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fmt.Printf("Entry: %s | IsDir: %v | Type: %v\n", entry.Name(), entry.IsDir(), entry.Type())
	}

	fmt.Println("\n=== Write methods ===")
	writePath := filepath.Join(baseDir, "write_example.txt")
	writeContent := "Hello from os.WriteFile!\nThis is the first write example.\n"
	if err := os.WriteFile(writePath, []byte(writeContent), 0644); err != nil {
		panic(err)
	}
	fmt.Println("os.WriteFile created:", writePath)
	fmt.Println(string(mustReadFile(writePath)))

	updatePath := filepath.Join(baseDir, "update_example.txt")
	if err := os.WriteFile(updatePath, []byte("Old content\n"), 0644); err != nil {
		panic(err)
	}
	fileToUpdate, err := os.OpenFile(updatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := fileToUpdate.WriteString("Updated content using O_TRUNC\n"); err != nil {
		panic(err)
	}
	if err := fileToUpdate.Close(); err != nil {
		panic(err)
	}
	fmt.Println("os.OpenFile with O_TRUNC updated:", updatePath)
	fmt.Println(string(mustReadFile(updatePath)))

	appendPath := filepath.Join(baseDir, "append_example.txt")
	if err := os.WriteFile(appendPath, []byte("First line\n"), 0644); err != nil {
		panic(err)
	}
	fileToAppend, err := os.OpenFile(appendPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := fileToAppend.WriteString("Second line appended\nThird line appended\n"); err != nil {
		panic(err)
	}
	if err := fileToAppend.Close(); err != nil {
		panic(err)
	}
	fmt.Println("os.OpenFile with O_APPEND appended:", appendPath)
	fmt.Println(string(mustReadFile(appendPath)))

	modifyPath := filepath.Join(baseDir, "modify_example.txt")
	if err := os.WriteFile(modifyPath, []byte("Go is fun and easy"), 0644); err != nil {
		panic(err)
	}
	currentData, err := os.ReadFile(modifyPath)
	if err != nil {
		panic(err)
	}
	updatedText := strings.Replace(string(currentData), "easy", "powerful", 1)
	if err := os.WriteFile(modifyPath, []byte(updatedText), 0644); err != nil {
		panic(err)
	}
	fmt.Println("Content modified using strings.Replace:", string(mustReadFile(modifyPath)))

	bytesPath := filepath.Join(baseDir, "bytes_example.txt")
	bytesFile, err := os.Create(bytesPath)
	if err != nil {
		panic(err)
	}
	if _, err := bytesFile.Write([]byte("This file was written using File.Write")); err != nil {
		panic(err)
	}
	if err := bytesFile.Close(); err != nil {
		panic(err)
	}
	fmt.Println("File.Write example:", string(mustReadFile(bytesPath)))

	textPath := filepath.Join(baseDir, "text.txt")
	textFile, err := os.Create(textPath)
	if err != nil {
		panic(err)
	}
	if _, err := textFile.WriteString("Welcome to Go!\nThis is written with WriteString.\n"); err != nil {
		panic(err)
	}
	if err := textFile.Close(); err != nil {
		panic(err)
	}
	fmt.Println("File.WriteString example:", string(mustReadFile(textPath)))

	fmt.Println("\n=== File permissions and cleanup ===")
	tempPath := filepath.Join(baseDir, "temp_demo.txt")
	if err := os.WriteFile(tempPath, []byte("temporary file"), 0644); err != nil {
		panic(err)
	}
	if err := os.Chmod(tempPath, 0600); err != nil {
		panic(err)
	}
	info, err := os.Stat(tempPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("File permission after Chmod:", info.Mode().String())
	if err := os.Remove(tempPath); err != nil {
		panic(err)
	}
	fmt.Println("Temporary file removed successfully.")
}

func mustReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func findWorkingDir() string {
	candidatePaths := []string{".", filepath.Join(".", "files"), filepath.Join("..", "files")}

	for _, p := range candidatePaths {
		if info, err := os.Stat(filepath.Join(p, "example.txt")); err == nil && !info.IsDir() {
			return p
		}
	}

	return "."
}
