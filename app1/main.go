package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func remainderOf7(goTo int) {
	for i := 7; i < goTo; i++ {
		fmt.Println(i % 7)
	}
}

func printPrimeNumber(goTo int) {
	primeNumberSlice := make([]int, 0, goTo)
	for i := 2; i < goTo; i++ {
		for k := 2; k <= i; k++ {
			if k == i {
				primeNumberSlice = append(primeNumberSlice, i)
				//fmt.Println(i, " Is Prime")
				//fmt.Println(" Remainder is ", i%7)
				break
			}
			if i%k == 0 {
				break
			}
		}
	}

	for i, primeNum := range primeNumberSlice {
		fmt.Println(i, " - ", primeNum)
	}
}

func webRequest() {
	response, err := http.Get("http://golang.org/")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}

// The "..." is the "Variadic" argument
// meaning we will take as many arguments
// as provided
func average1(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	// Could also write the below statement as
	// var total = float64, which would initialize
	// to the zero value
	total := 0.0

	// Used the "_" to indicate that we want the value
	// and not the index
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func average2(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	printPrimeNumber(1000)
	//n := average1(43, 56, 87, 12, 45, 57)
	//data := []float64{43, 56, 87, 12, 45, 57}
	//o := average2(data)
	//fmt.Println("n is ", n)
	//fmt.Println("o is ", o)
	//webRequest()
}
