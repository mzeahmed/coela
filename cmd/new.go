/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mzeahmed/coela/internal/project"
	"github.com/mzeahmed/coela/internal/scaffold"
	"github.com/mzeahmed/coela/internal/stacks/symfony"
	"github.com/mzeahmed/coela/internal/stacks/wordpress"
	"github.com/mzeahmed/coela/internal/traefik"
	"github.com/mzeahmed/coela/internal/ui"
)

// newCmd implements `coela new`.
//
// The flow is: ask which stack to use, run that stack's interactive Wizard
// to build a *project.Project, render the stack's templates on disk with
// scaffold.Generate, install the framework itself via the stack's Install
// func, then, if Traefik was enabled, provision local HTTPS (mkcert
// certificate + /etc/hosts entries) via the traefik package.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		choice, err := ui.Select("Stack", []string{"Symfony", "WordPress (Bedrock)"})
		if err != nil {
			return err
		}

		// Each stack is a plain package (no shared interface): the switch
		// below is the only place that knows both stacks exist.
		var (
			p            *project.Project
			templatesDir string
			install      func(string) error
		)

		switch choice {
		case "Symfony":
			p, err = symfony.Wizard()
			templatesDir = symfony.TemplatesDir()
			install = symfony.Install
		case "WordPress (Bedrock)":
			p, err = wordpress.Wizard()
			templatesDir = wordpress.TemplatesDir()
			install = wordpress.Install
		}
		if err != nil {
			return err
		}

		if err := scaffold.Generate(p, templatesDir); err != nil {
			return err
		}

		if err := install(filepath.Join(p.Name, "app")); err != nil {
			return err
		}

		if p.Traefik {
			if err := traefik.GenerateCert(p); err != nil {
				return err
			}

			if err := traefik.RegisterHosts(p); err != nil {
				return err
			}
		}

		fmt.Println("Project ready.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
