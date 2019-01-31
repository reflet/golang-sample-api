package main

import (
    godo "gopkg.in/godo.v2"
)

func tasks(p *godo.Project) {

    p.Task("server", nil, func(c *godo.Context) {
        // rebuilds and restarts when a watched file changes
        c.Start("main.go", godo.M{"$in": "./"})
    }).Src("*.go", "**/*.go").
        Debounce(3000)
}

func main() {
    godo.Godo(tasks)
}

