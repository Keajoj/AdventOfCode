package cmd

import "fmt"
import "bufio"
import "strconv"
import "os"

func dayOnePart1 (inputFile string) {
	fmt.Println(inputFile)
		fmt.Println(InputFilepath)
		file, err := os.Open(InputFilepath)
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			//fmt.Println(scanner.Text())
			value, err := strconv.Atoi(scanner.Text())
			if err != nil { os.Exit(1) }
			fmt.Printf("Value as int: %d\n", value)
			
		}

		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}
    }
