package main

import (
	"fmt"

	"github.com/DevLucca/sorting-algo/cmd"
	"github.com/c-bata/go-prompt"
)

var (
	version string
)

func main() {
	fmt.Printf("Sorting-Algorithms v%s\n", version)
	fmt.Printf("Please use `exit`, `quit` or `Ctrl-D` to exit this program.\n")
	defer fmt.Println("\nBye!")

	p := prompt.New(
		cmd.Executor,
		cmd.Complete,
		prompt.OptionTitle("sorting-algos: interactive sorting algorithms"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(" "),
	)
	p.Run()
}
