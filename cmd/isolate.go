package main

import (
	"fmt"

	"rogchap.com/v8go"
)

func main() {
	iso, _ := v8go.NewIsolate() // creates a new JavaScript VM
	ctx1, _ := v8go.NewContext(iso) // new context within the VM

	ctx1.RunScript("const multiply = (a, b) => a * b", "math.js")
	ctx1.RunScript("const result = multiply(3, 4)", "main.js") // any functions previously added to the context can be called

	val, _ := ctx1.RunScript("result", "value.js")        // return a value in JavaScript back to Go
	fmt.Printf("multiply result: %s\n", val)

	ctx2, _ := v8go.NewContext(iso) // another context on the same VM
	if _, err := ctx2.RunScript("multiply(3, 4)", "main.js"); err != nil {
	  // this will error as multiply is not defined in this context
	  fmt.Println("Error in context:", err)
	}
}
