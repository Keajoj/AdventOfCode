package cmd

import (
"fmt"
"bufio"
"strconv"
"os"
"github.com/spf13/cobra"
)

func d1p1 () {
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


// d1p1Cmd represents the d1p1 command
var d1p1Cmd = &cobra.Command{
	Use:   "d1p1",
	Short: "Advent of Code 22 Day 1 Part 1",
	Run: func(cmd *cobra.Command, args []string) {
		d1p1()	},
}

func init() {
	rootCmd.AddCommand(d1p1Cmd)
}
