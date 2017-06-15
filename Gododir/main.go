package main

import (
	"fmt"
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	// do.Env = `GOPATH=vendor::$GOPATH`

	p.Task("default", do.S{"server"}, nil)

	p.Task("server", nil, func(c *do.Context) {
		fmt.Println("server start!")
		c.Start("main.go", do.M{"$in": "./"})
	}).Src("main.go").
		Debounce(5000) // タスクが再実行されるまでのms
}

func main() {
	do.Godo(tasks)
}
