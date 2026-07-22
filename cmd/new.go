/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mzeahmed/coela/internal/project"
	"github.com/mzeahmed/coela/internal/scaffold"
	"github.com/mzeahmed/coela/internal/stacks/symfony"
	"github.com/mzeahmed/coela/internal/stacks/wordpress"
	"github.com/mzeahmed/coela/internal/traefik"
	"github.com/mzeahmed/coela/internal/ui"
)

// devOutputDir, when non-empty, is the directory (relative to the current
// working directory) every `coela new` run generates its project under,
// instead of directly at the working directory. It exists so running the
// CLI from its own source tree (`go run .`, a plain `go build`) confines
// every generated project to one single, gitignored directory instead of
// scattering them across the coela repository itself.
//
// Release builds force this back to "" via ldflags (see .goreleaser.yaml),
// so an installed `coela` binary always generates in the current working
// directory — this redirection is a development-only convenience.
var devOutputDir = "tmp"

// newCmd implements `coela new`.
//
// The flow is: ask which stack to use, run that stack's interactive Wizard
// to build a *project.Project, render the stack's templates on disk with
// scaffold.Generate, install the framework itself via the stack's Install
// func, optionally point the installed framework at the scaffolded Docker
// services (configureEnv), then, if Traefik was enabled, provision local
// HTTPS (mkcert certificate + /etc/hosts entries) via the traefik package.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Scaffold a new PHP project",
	Long: `Runs an interactive wizard to scaffold a complete, Docker-based PHP
development environment.

It asks for the framework (Symfony, WordPress/Bedrock), PHP version,
database, and optional Redis, Mailpit, and Traefik support, then generates
the Docker Compose setup, Nginx and PHP-FPM configuration, and installs the
selected framework — ready to use.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("> coela new")

		if devOutputDir != "" {
			if err := os.MkdirAll(devOutputDir, 0755); err != nil {
				return err
			}

			if err := os.Chdir(devOutputDir); err != nil {
				return err
			}
		}

		choice, err := ui.Select("Stack", []string{"Symfony", "WordPress (Bedrock)"})
		if err != nil {
			return err
		}

		// Each stack is a plain package (no shared interface): the switch
		// below is the only place that knows both stacks exist.
		var (
			p            *project.Project
			templatesDir fs.FS
			install      func(string) error
			configureEnv func(*project.Project) error
		)

		switch choice {
		case "Symfony":
			p, err = symfony.Wizard()
			templatesDir = symfony.TemplatesDir()
			install = symfony.Install
			configureEnv = symfony.ConfigureEnv
		case "WordPress (Bedrock)":
			p, err = wordpress.Wizard()
			templatesDir = wordpress.TemplatesDir()
			install = wordpress.Install
			configureEnv = wordpress.ConfigureEnv
		}
		if err != nil {
			return err
		}

		fmt.Println("Creating project...")

		if err := scaffold.Generate(p, templatesDir); err != nil {
			return err
		}

		if err := install(filepath.Join(p.Name, "app")); err != nil {
			return err
		}

		if configureEnv != nil {
			if err := configureEnv(p); err != nil {
				return err
			}
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
