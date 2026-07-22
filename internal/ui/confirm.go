// Package ui provides small terminal-interaction primitives (text input,
// single-select, and yes/no confirmation) built on top of promptui.
// It knows nothing about Project, stacks, or scaffolding — stacks compose
// these primitives into their own wizards.
package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Confirm asks a yes/no question and returns true if the user picked "Yes".
func Confirm(label string) (bool, error) {

	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
		Templates: &promptui.SelectTemplates{
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ "%s:" | bold }} {{ . | cyan }}`, promptui.IconGood, label),
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return false, err
	}

	return result == "Yes", nil
}
