package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use a verbosity flag
	verbose := flag.Bool("verbose", false, "Print verbose information.")
	flag.Parse()
	// Read a single line from standard input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
        
	// Convert gron output to a jq query
	jqQuery := gronLineToJq(line)


	if *verbose {
	    fmt.Printf("✨ The INCOMING gron line looks like this: \r\n")
            fmt.Println(line)
	    fmt.Printf("🎊 The 'jq' query to use would look like this: \r\n")
	    fmt.Println(jqQuery)
	} else {
	    fmt.Println(jqQuery)

    }
}



func splitSlice(input []string, delimiter string) []string {
	var result []string

	// Iterate over each element in the []string
	for _, item := range input {
		// Split the element by the specified delimiter
		parts := strings.Split(item, delimiter)

		// Append the split parts to the result
		result = append(result, parts...)
	}

	return result
}


func gronLineToJq(line string) string {
	tokens := strings.Fields(line)
	midwayQuery := splitSlice(tokens,"=")[0]
	jqFilter := strings.TrimPrefix(midwayQuery, "json") 
	jqQuery := "jq '" + jqFilter + "'"
	return jqQuery
}

