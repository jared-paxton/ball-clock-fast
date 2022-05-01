package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jared-paxton/ball-clock-fast/pkg/clock"
)

func main() {

	fmt.Println("\nBall Clock!")
	for {
		mode, err := getNumberInput("Enter 1 for Mode 1: Cycle Days, or \nenter 2 for Mode 2: Clock State, or \nenter 9 to quit: ")
		printIfError(err)
		if mode == 9 {
			fmt.Println("Quitting...")
			break
		}

		switch mode {
		case 1:
			numBalls, err := getNumberInput("Enter the number of clock balls between 27 and 127 (or 9 to quit): ")
			printIfError(err)
			if numBalls == 9 {
				fmt.Println("Quitting...")
				break
			}

			fmt.Println()
			err = clock.CycleDays(numBalls)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			numBalls, err := getNumberInput("Enter the number of clock balls between 27 and 127 (or 9 to quit): ")
			printIfError(err)
			if numBalls == 9 {
				fmt.Println("Quitting...")
				break
			}
			minutes, err := getNumberInput("Enter the number of minutes for which the clock should run (or 9 to quit): ")
			printIfError(err)
			if minutes == 9 {
				fmt.Println("Quitting...")
				break
			}

			fmt.Println()
			err = clock.State(numBalls, minutes)
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			fmt.Println("Error: Not a valid option.")
		}
		fmt.Println()
	}
	fmt.Println()
}

func getNumberInput(instruction string) (int, error) {
	fmt.Print(instruction)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	trimmedInput := strings.TrimSuffix(input, "\n")
	num, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return num, nil
}

func printIfError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
