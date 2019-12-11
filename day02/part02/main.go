package main

import (  
	"fmt"
)

func main() {
	input := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,5,19,23,1,6,23,27,1,27,10,31,1,31,5,35,2,10,35,39,1,9,39,43,1,43,5,47,1,47,6,51,2,51,6,55,1,13,55,59,2,6,59,63,1,63,5,67,2,10,67,71,1,9,71,75,1,75,13,79,1,10,79,83,2,83,13,87,1,87,6,91,1,5,91,95,2,95,9,99,1,5,99,103,1,103,6,107,2,107,13,111,1,111,10,115,2,10,115,119,1,9,119,123,1,123,9,127,1,13,127,131,2,10,131,135,1,135,5,139,1,2,139,143,1,143,5,0,99,2,0,14,0}
	var output int = 19690720
	fmt.Println("\nRESULT PARTE 1: ", calculateIntCode(input, 12, 2))
	fmt.Println("\nRESULT PARTE 2: ", CalculateNounAndVerbOutput(input, output) )
}

func calculateIntCode(inputcode []int, noun int, verb int) int{

	var input[] int = getInput(inputcode)
	var outputresult int
	input[1] = noun
	input[2] = verb
	var i int = 0
	var opcode = 0

	for opcode != 99{
		opcode = input[i]
		if input[i] == 1 {
			outputresult = input[input[i+1]] + input[input[i+2]]
			input[input[i+3]] = outputresult
		}else{
			if input[i] == 2 {
				outputresult = input[input[i+1]] * input[input[i+2]]
				if input[input[i+3]] >= len(input) {
					break
				}else{
					input[input[i+3]] = outputresult
				}
			}
		}
		i = i + 4
	}
	return input[0]
}


func CalculateNounAndVerbOutput(input []int, output int) int{
	var newinput []int
	var vi int = 0
	var vj int = 0
	var result int = 0
	i := 0
	var intcodecalc int = 0
	for i <= 99 {
		j := 0
		//fmt.Print("I:", i)
		for j <= 99 {
			newinput = getInput(input)
			/*fmt.Println("--------------------------------------")
			fmt.Println("noun: ", i, "verb: ", j)
			fmt.Println("--------------------------------------") */
			intcodecalc = calculateIntCode(newinput, i, j)
			if intcodecalc == output{
				vi = i
				vj = j
				result = 100 * i + j
				//return result
			}
			j = j+1
		}
		i = i +1
	}

	fmt.Println("\nENCONTRADO: ", result, " EN LA COMBINACION NOUN: ",  vi, " VERB: ", vj)
	return result
}



func getInput(arr []int) []int{
	var input []int
	for i := 0; i < len(arr); i++ {
		input = append(input, arr[i])
	}
	return input
}