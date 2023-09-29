package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define flags
	zeroPad := flag.Bool("z", false, "Zero-pad line numbers")
	width := flag.Int("w", 6, "Number of digits in line number")
	separator := flag.String("s", " ", "Separator between line number and content")
	flag.Parse()

	var scanner *bufio.Scanner

	// Use flag.Arg(0) to get the file name argument
	fileName := flag.Arg(0)

	// Determine if there is data available on stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// If stdin has data, use it
		scanner = bufio.NewScanner(os.Stdin)
	} else if fileName != "" {
		// If a file name was provided, try to open it
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		// Use stdin if no file name argument is provided
		scanner = bufio.NewScanner(os.Stdin)
	}

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		lineNumberStr := strconv.Itoa(lineNumber)

		if *zeroPad {
			lineNumberStr = fmt.Sprintf("%0*d", *width, lineNumber)
		} else {
			lineNumberStr = fmt.Sprintf("%*s", *width, lineNumberStr)
		}

		fmt.Printf("%s%s %s\n", lineNumberStr, *separator, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}