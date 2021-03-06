package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	randArray "github.com/DevLucca/sorting-algo/utils/random-array"
	"github.com/c-bata/go-prompt"
)

var (
	ctx *context
)

const (
	defaultLenght = "10"
)

func init() {
	ctx = &context{}
}

func Executor(s string) {
	s = strings.TrimSpace(s)

	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	args := strings.Split(s, " ")
	selectedAlgo := args[0]
	if len(args) == 1 {
		args = append(args, defaultLenght)
	}
	n, _ := strconv.Atoi(args[1])

	if args[0] == "help" {
		fmt.Println("Available commands:")
		for key := range mapSupportedAlgorithms {
			fmt.Println("-", key)
		}
		return
	}

	if err := checkCommand(args); err != nil {
		fmt.Println(err)
		return
	}

	ctx.SetStrategy(selectedAlgo)
	newArray := randArray.GenerateArray(n)

	done, changed := make(chan bool), make(chan bool)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	data := &newArray
	// fmt.Println(data)
	go ctx.ExecuteStrategy(done, changed, data)
	go uIRenderer(wg, selectedAlgo, done, changed, data)

	wg.Wait()
	// fmt.Println(data)
}

func checkCommand(args []string) error {
	if _, found := mapSupportedAlgorithms[args[0]]; !found {
		return errors.New("oops... command not found :( \nTry 'help'")
	}
	return nil
}

func Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	return prompt.FilterHasPrefix(commands, args[0], true)
}
