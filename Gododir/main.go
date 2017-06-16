package main

import (
	"fmt"
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	do.Env = `GOPATH=vendor::$GOPATH`

	/*
		p.Task("default", do.S{"hello", "server"}, nil)
		p.Task("hello", nil, func(c *do.Context) {
			name := c.Args.AsString("name", "n")
			if name == "" {
				c.Bash("echo Hello $USER!")
			} else {
				fmt.Println("Hello", name)
			}
		})
	*/

	p.Task("server", nil, func(c *do.Context) {
		fmt.Println("server start!")
		// c.Start("main.go", do.M{"$in": "./"})
	}).Src("*.go,toml", "**/*.go").
		Debounce(3000) // タスクが再実行されるまでのms
}

func main() {
	do.Godo(tasks)
}
