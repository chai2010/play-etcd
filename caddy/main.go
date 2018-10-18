package main

import (
	"github.com/mholt/caddy"
)

func main() {
	caddy.Start(caddy.DefaultInput("http"))
}
