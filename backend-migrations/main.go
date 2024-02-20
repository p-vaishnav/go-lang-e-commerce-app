package main

import (
	"fmt"
	"migrations/commands"
	"migrations/configs"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("in main function...")

	configs.LoadConfigs()
	cmd := &cobra.Command{}

	cmd.AddCommand(commands.DropTables())
	cmd.AddCommand(commands.Migrate())

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
