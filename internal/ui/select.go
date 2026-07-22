package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Select asks the user to pick one item from items and returns the chosen
// value.
func Select(label string, items []string) (string, error) {

	prompt := promptui.Select{
		Label: label,
		Items: items,
		Templates: &promptui.SelectTemplates{
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ "%s:" | bold }} {{ . | cyan }}`, promptui.IconGood, label),
		},
	}

	_, result, err := prompt.Run()

	return result, err
}
