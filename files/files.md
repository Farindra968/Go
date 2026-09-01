# Files in Go

Go provides powerful built-in support for working with files and directories through the `os` package and the `File` type. These tools let you open, read, write, update, delete, and inspect files efficiently.

## File basics

A file in Go is usually handled in one of these ways:

- `os.Open` and `os.Create` for opening or creating a file
- `os.ReadFile` and `os.WriteFile` for quick full-file reads and writes
- `File.Read`, `File.Write`, and `File.WriteString` for streaming content
- `os.OpenFile` for advanced control such as append and truncate modes
- `os.Stat` and `File.Stat` for file metadata such as size, permissions, and modification time
- `os.ReadDir` to list directory entries

## Common file operations

### 1. Open a file

```go
file, err := os.Open("example.txt")
if err != nil {
    panic(err)
}
defer file.Close()
```

`os.Open` opens an existing file for reading. It returns a `*File` and an error.

### 2. Read file content

#### Read a small chunk

```go
buffer := make([]byte, 128)
n, err := file.Read(buffer)
if err != nil && err != io.EOF {
    panic(err)
}
fmt.Println(string(buffer[:n]))
```

#### Read the whole file

```go
content, err := os.ReadFile("example.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(content))
```

`os.ReadFile` is the simplest way to read the entire file at once.

### 3. Write a file

```go
err := os.WriteFile("write_example.txt", []byte("Hello from Go!\n"), 0644)
if err != nil {
    panic(err)
}
```

This creates or overwrites a file with the given bytes.

### 4. Open with flags and append

```go
file, err := os.OpenFile("append_example.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
if err != nil {
    panic(err)
}
_, err = file.WriteString("Second line\n")
if err != nil {
    panic(err)
}
file.Close()
```

`os.OpenFile` is useful when you need to control file flags such as:

- `os.O_CREATE` — create if not exists
- `os.O_WRONLY` — write-only mode
- `os.O_APPEND` — append to the end of the file
- `os.O_TRUNC` — clear the file before writing

### 5. Truncate and overwrite content

```go
file, err := os.OpenFile("update_example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
if err != nil {
    panic(err)
}
_, err = file.WriteString("Updated content\n")
if err != nil {
    panic(err)
}
file.Close()
```

Using `os.O_TRUNC` removes the previous content before writing new data.

### 6. Create a file and write text

```go
file, err := os.Create("text.txt")
if err != nil {
    panic(err)
}
defer file.Close()

_, err = file.WriteString("Welcome to Go!\n")
if err != nil {
    panic(err)
}
```

`os.Create` opens or creates a file and truncates it by default.

### 7. Write raw bytes

```go
file, err := os.Create("bytes_example.txt")
if err != nil {
    panic(err)
}
defer file.Close()

_, err = file.Write([]byte("This file was written using Write(byte[])"))
if err != nil {
    panic(err)
}
```

This method is useful when you already have a byte slice and want to write it directly.

### 8. Read file metadata

```go
info, err := os.Stat("example.txt")
if err != nil {
    panic(err)
}

fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("Mode:", info.Mode())
fmt.Println("Modified:", info.ModTime())
```

`os.Stat` gives information about the file, including size, permissions, and last modified time.

### 9. Read directory entries

```go
entries, err := os.ReadDir(".")
if err != nil {
    panic(err)
}

for _, entry := range entries {
    fmt.Println(entry.Name(), entry.IsDir())
}
```

`os.ReadDir` reads all entries in a directory and returns a slice of `DirEntry` values.

### 10. File permissions

```go
err := os.Chmod("example.txt", 0600)
if err != nil {
    panic(err)
}
```

Permissions are represented using Unix-style numeric values such as:

- `0644` — owner read/write, group/others read
- `0600` — owner read/write only
- `0755` — executable for owner and read/execute for others

## Example program

The complete example for this topic is in [files/files.go](files/files.go). It demonstrates:

- `os.Open`
- `File.Stat`
- `File.Read`
- `os.ReadFile`
- `os.ReadDir`
- `os.WriteFile`
- `os.OpenFile` with `O_TRUNC`
- `os.OpenFile` with `O_APPEND`
- `os.Create`
- `File.Write`
- `File.WriteString`
- `os.Chmod`
- `os.Remove`

## Best practices

- Always close files with `defer file.Close()` after opening them.
- Use `os.ReadFile` for small files when you want the whole content at once.
- Use `os.OpenFile` when you need explicit write flags.
- Check errors immediately after file operations.
- Use `os.ReadDir` for directory listing and metadata inspection.

## Summary

The Go file system API is simple and flexible. The most common operations are:

- open a file
- read or write content
- inspect metadata
- manage directories
- modify permissions
- append or truncate content

These operations are the foundation for working with configuration files, logs, JSON, text files, and user data in real Go programs.
