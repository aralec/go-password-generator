package main

import (
	"fmt"
	generator "go-password-generator/password-generator"
)

func main() {
	pwd, err := generator.Generate(12)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated password is : ", pwd)
}
