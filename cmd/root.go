/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/

// Package cmd implements the Coela command-line interface using Cobra.
//
// It contains no business logic: each command only collects user input
// (via internal/ui and a stack's Wizard) and delegates the actual work to
// internal/scaffold and the chosen stack's Install function.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command invoked when coela is called without any subcommand.
var rootCmd = &cobra.Command{
	Use:   "coela",
	Short: "From zero to a ready-to-code PHP project",
	Long: `Coela scaffolds complete, Docker-based PHP development environments.

Instead of hand-writing Docker Compose files, Dockerfiles, Traefik
configuration, and project structure every time, Coela generates all of it —
and installs your framework of choice — with a single command.

Run "coela new" to start the interactive wizard.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coela.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
