package main

import (
	"errors"
	"fmt"
)

//NUMBERODEFAULT Numero por defecto para poder sumar
const NUMBERODEFAULT int = 0

//suma Esta funciona es para sumar @autor Victor Barrera
func suma(numberOne int, numberTwo int) (int, error) {
	if numberOne <= NUMBERODEFAULT {
		return NUMBERODEFAULT, errors.New("Numero menor o igual a 0")
	}
	return numberOne + numberTwo, nil
}

func main() {
	i, err := suma(1, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
