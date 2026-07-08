/*
Copyright © 2026 Mze Ahmed <houdjiva@gmail.com>

*/

// Coela is a CLI that scaffolds complete Docker-based development
// environments for PHP projects (Symfony, WordPress/Bedrock, ...).
//
// Usage:
//
//	coela new
package main

import "github.com/mzeahmed/coela/cmd"

// main delegates to the cmd package, which owns the actual Cobra command tree.
func main() {
	cmd.Execute()
}
