package main

import (
	"fmt"
	cobra "github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "becmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	},
}
var genCmd = &cobra.Command{
	Use: "gen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello gen")
	},
}

func main() {
	rootCmd.AddCommand(genCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
