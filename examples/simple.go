package main

import (
	"fmt"
	"github.com/alecaivazis/probe"
)

var qs = []*probe.Question{
	{
		Name:   "name",
		Prompt: &probe.Input{"What is your name?"},
	},
	{
		Name: "color",
		Prompt: &probe.Choice{
			Question: "Choose a color:",
			Choices:  []string{"red", "blue", "green"},
		},
	},
}

func main() {
	answers, err := probe.Ask(qs)
	if err != nil {
		fmt.Println("\n", err.Error())
		return
	}

	fmt.Printf("%s chose %s.\n", answers["name"], answers["color"])
}