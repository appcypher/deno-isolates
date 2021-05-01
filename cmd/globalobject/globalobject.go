package main

import (
	"fmt"

	"rogchap.com/v8go"
)

const code = `version`

func main() {
	iso, _ := v8go.NewIsolate()
	ctx, _ := v8go.NewContext(iso)

	global := ctx.Global()
	global.Set("version", "v1.0.1")

	val, err := ctx.RunScript(code, "result.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("result: %s\n", val)
}
