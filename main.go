package main

import "github.com/dwisiswant0/galer/internal/runner"

func main() {
	options := runner.Parse()
	runner.New(options)
}
