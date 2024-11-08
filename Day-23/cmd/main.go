package main

import (
	"fmt"
)

/*
1. What buffers are and why they’re used
2. Using bytes.Buffer for in-memory data buffering
3. Using bufio for efficient file and I/O operations
4. Best practices and performance tips
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nUsing Buffers in Go")

	inMemoryDataManipulation()

	efficientFileAndIOoperations()

	writingFilesWithBufferedWriter()
}

/*
BEST PRACTICES

Use Buffers for Repeated Writes: When writing frequently to an I/O source, like a file or network, use a buffer to accumulate data before committing it in bulk.
								 This reduces the number of writes and improves performance.

Buffer Size: For large files, a buffer size of 4 KB or 8 KB is often a good choice, as it aligns with disk block sizes on many systems.

Always Flush Buffers: When writing with bufio.Writer, always call Flush to ensure data in the buffer is saved. Not flushing can lead to data loss.

Using bytes.Buffer for String Concatenation: If you are frequently concatenating strings, a bytes.Buffer is more memory-efficient than using + or fmt.Sprintf.

Error Handling: Always check for errors after operations, especially when dealing with buffers and I/O. This helps catch any read/write issues
				that could lead to incomplete data processing.
*/
