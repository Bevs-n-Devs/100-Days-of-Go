package main

import "fmt"

func Example2() {
	fmt.Println("\nExample 2: Separation in a CLI Tool")

	var scenario = `
	Suppose you're building a CLI tool to read a file and process its contents.

	Without Separation of Concerns:
		- All functionality is in one function

	Code Example:
	func ProcessFile(filePath string) {
		// Read file
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Process content
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			fmt.Println("Processed:", strings.ToUpper(line))
		}
	}

	Problems:
		1. Mixing file reading and content processing.
		2. Hard to test content processing without actual files.

	With Separation of Concerns:
		- Separate file reading and content processing

	Code Example:
	// FileReader: Handles file operations
	type FileReader struct{}

	func (fr *FileReader) Read(filePath string) (string, error) {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}

	// Processor: Handles content processing
	type Processor struct{}

	func (p *Processor) Process(content string) []string {
		lines := strings.Split(content, "\n")
		processed := []string{}
		for _, line := range lines {
			processed = append(processed, strings.ToUpper(line))
		}
		return processed
	}

	// Main Function
	func main() {
		fileReader := &FileReader{}
		processor := &Processor{}

		// Read file
		content, err := fileReader.Read("example.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process content
		results := processor.Process(content)
		for _, result := range results {
			fmt.Println("Processed:", result)
		}
	}

	Benefits:
		1. Independent Components: You can test the Process without needing to deal with the file reading.
		2. Reusability: The Processor can process strings from any source, not just files.
	`
	fmt.Println(scenario)
}
