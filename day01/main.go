package main

import (  
	"fmt"
	"math"
	"bufio"
    "flag"
    "log"
	"os"
	"strconv"
)


func main() {
	fmt.Println("Hello, mundo")
	
	fptr := flag.String("fpath", "input.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
        log.Fatal(err)
    }
    }()
	s := bufio.NewScanner(f)
	var resultvalue int = 0
    for s.Scan() {
		fmt.Println(s.Text())

		f, err := strconv.ParseFloat(s.Text(), 64)
		if err == nil{
			resultvalue = resultvalue + getCalculatedValue(f)
		}

    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println("--------------------------------------")
	fmt.Println("Resultado = ", resultvalue)
	fmt.Println("--------------------------------------")
}


func getCalculatedValue(mass float64) int {
	//var mass float64
	//fmt.Scan(&mass)
	fmt.Println("......................................")
	var result = mass / 3
	var roundedresult int = 0
	fmt.Println(mass, "/ 3  = ", result)

	roundedresult = int(math.Trunc(result)) 
	fmt.Println("redondeado = ", roundedresult)
	
	roundedresult = roundedresult - 2
	fmt.Println("menos 2 =", roundedresult)
	fmt.Println("......................................")

	return roundedresult
}