package main

import (
	"errors"
	"log"
)

func Division(numerator,denominator int64) (float64, error){
	if denominator == 0 {
		return 0, errors.New("division by zero")
	}
	return float64(numerator) / float64(denominator), nil
}

func main() {
	log.Println(Division(1,2))
}