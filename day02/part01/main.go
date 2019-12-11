package main

import (  
	"fmt"
)

func main() {
	//input := []int{1,0,0,0,99}
	//input := []int{2,3,0,3,99}
	//input := []int{2,4,4,5,99,0}
	//input := []int{1,1,1,4,99,5,6,0,99}
	input := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,5,19,23,1,6,23,27,1,27,10,31,1,31,5,35,2,10,35,39,1,9,39,43,1,43,5,47,1,47,6,51,2,51,6,55,1,13,55,59,2,6,59,63,1,63,5,67,2,10,67,71,1,9,71,75,1,75,13,79,1,10,79,83,2,83,13,87,1,87,6,91,1,5,91,95,2,95,9,99,1,5,99,103,1,103,6,107,2,107,13,111,1,111,10,115,2,10,115,119,1,9,119,123,1,123,9,127,1,13,127,131,2,10,131,135,1,135,5,139,1,2,139,143,1,143,5,0,99,2,0,14,0}
	//Restore gravity assist program
	input[1] = 12
	input[2] = 2
	calculateIntCode(input)
	printArray(input)
}

func calculateIntCode(input[] int){
	var outputresult int
	i := 0
    for i < len(input) {
		fmt.Println("--------------------------------------")
		fmt.Println("Command: ", input[i])
				
		switch input[i] {
		case 1:
			fmt.Println("adds")
			outputresult = input[input[i+1]] + input[input[i+2]]
			fmt.Println("result = ", outputresult)
			input[input[i+3]] = outputresult
		case 2:
			fmt.Println("multiplies")
			outputresult = input[input[i+1]] * input[input[i+2]]
			fmt.Println("result = ", outputresult)
			input[input[i+3]] = outputresult
		case 99:
			fmt.Println("exit")
			i = i + 4
			break
		default:
			fmt.Println("Error")
			i = i + 4
			break
		}
		fmt.Println("--------------------------------------")
		i = i + 4
	}
}

func printArray(input[] int){
	fmt.Println("--------------------------------------")
	fmt.Print("\nRESULTADO: ")
	i := 0
    for i < len(input) {
		fmt.Print("", input[i])
		i = i + 1
		if i < len(input){
			fmt.Print(",")
		}
	}
	fmt.Println("\n--------------------------------------")

	fmt.Print("\nRESULTADO POS 0: ", input[0], "\n\n\n")
}