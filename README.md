# Empty File Scanner

A simple file scanner which scans directory for empty files and writes the summary of result to a `.txt` file.

## How to Use

Pass the directory name to command line. There are 2 ways of doing it

1. Run the code directly

   ```none
   go run main.go {directory_name}
   ```

2. Or compile the code into an executable using `go build` command.

   On Windows

   ```none
   ./empty-file-scanner.exe {directory_file}
   ```

   On Linux

   ```none
   ./empty-file-scanner {directory_name}
   ```

## Example

```none
Allocate 219 bytes in memory...

Scan completed!
12 empty files found in D:/Documents
Result stored in 'summary.txt'
File size:  219 bytes
```
