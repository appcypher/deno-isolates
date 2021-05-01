package main

import (
	"fmt"

	"rogchap.com/v8go"
)

func main() {
	iso, _ := v8go.NewIsolate()
	ctx1, _ := v8go.NewContext(iso)

	if _, err := ctx1.RunScript("const multiply = (a, b) => a * b", "math.js"); err != nil {
		fmt.Println(err)
		return
	}

	val, err := ctx1.RunScript("multiply(3, 4)", "result.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("result: %s\n", val)

	ctx2, _ := v8go.NewContext(iso)
	if _, err := ctx2.RunScript("multiply(3, 4)", "main.js"); err != nil {
		fmt.Println("Error in context:", err)
	}
}
