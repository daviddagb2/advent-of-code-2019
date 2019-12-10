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
	var totalvalue int = 0
    for s.Scan() {
		f, err := strconv.ParseFloat(s.Text(), 64)
		if err == nil{
			var calculatedval =  getCalculatedValue(f)
			var calculatedvalfuel =  calculateDoubleCheker(calculatedval)
			fmt.Println("Valor calculado  ", calculatedvalfuel)
			totalvalue = totalvalue + calculateDoubleCheker(calculatedval)
		}
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println("--------------------------------------")
	fmt.Println("Resultado = ", totalvalue)
	fmt.Println("--------------------------------------")
}


func getCalculatedValue(mass float64) int {
	var result = mass / 3
	var roundedresult int = 0
	roundedresult = int(math.Trunc(result)) 	
	roundedresult = roundedresult - 2
	return roundedresult
}


func calculateDoubleCheker(initialvalue int) int{
	var total int = initialvalue
	i := initialvalue
    for i > 0 {
		i = getCalculatedValue(float64(i))
		if i > 0{
			total = total + i
		}
	}
	return total
}