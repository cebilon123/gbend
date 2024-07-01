/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gbend/internal/command"
	"gbend/internal/configuration"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the backend based on the config file.",
	Long: `Builds the backend based on the config file.

	Firstly, it validates if given config is valid and then uses it to create
	the whole backend code with simple rest api.`,
	Run: func(cmd *cobra.Command, args []string) {
		out := "./testapp"

		cfg, err := configuration.LoadConfiguration(fileName)
		if err != nil {
			panic(err.Error())
		}

		if err := command.NewBuildCommand(out, cfg).Execute(); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
