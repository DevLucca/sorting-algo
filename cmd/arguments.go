package cmd

import prompt "github.com/c-bata/go-prompt"

var commands = []prompt.Suggest{
	{Text: "help"},
}

func init() {
	for key := range mapSupportedAlgorithms {
		commands = append(commands, prompt.Suggest{Text: key})
	}
}
