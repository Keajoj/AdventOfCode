package cmd

import "fmt"
import "bufio"
import "strconv"
import "os"

func dayOnePart1 () {
	file, err := os.Open(InputFilepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var curr int = 0
	var max int = 0
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			curr += value
		} else {
			if curr > max { 
				max = curr
			}
			curr = 0
		}
	}

	fmt.Printf("Max value: %d",max);

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
