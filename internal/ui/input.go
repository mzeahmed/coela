package ui

import "github.com/manifoldco/promptui"

// Input prompts the user for a single line of free-form text.
func Input(label string) (string, error) {

	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	return result, err
}
