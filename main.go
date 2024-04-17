package main

import (
	"fmt"
	"wp/pkg/sillyapi"
)

func main() {
	joke, err := sillyapi.GetJoke()
	if err != nil {
		panic(err)
	}

	fmt.Println(joke)
}
