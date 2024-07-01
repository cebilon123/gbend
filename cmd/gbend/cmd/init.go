/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const (
	fileName = "./config.toml"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration file.",
	Long:  `Initializes the configuration file with a state in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.WriteFile(fileName, []byte("test"), 0777); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
