package main

import (
	"fmt"
	"../Hello_Go/Shape"
)

func main() {

	shapes := make([]Shape.Shape,10)

	for i:=0; i<len(shapes); i++ {
		if(i%2==0) {
			shapes[i] = Shape.Square{Height:float32(i), Width:float32(i*2)}
		}else {
			shapes[i] = Shape.Circle{Radius : float32(i)}
		}

	}

	for _, value := range shapes {
		fmt.Printf("Type : %T \t Area : %v, \t Perimeter : %v \n",value, value.Area(), value.Perimeter())
	}
}