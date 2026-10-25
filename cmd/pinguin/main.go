package main

import (
	"pinguin/internal/pinger"
)

func main() {
	println("pinger v.0.0.0")

	var pinger = pinger.Pinger{}

	//
	pinger.Ping()

}
