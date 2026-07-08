package traefik

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mzeahmed/coela/internal/project"
)

// RegisterHosts appends any of p's Traefik domains that are missing from
// /etc/hosts, mirroring the reference project's `make hosts` target. Each
// missing entry is appended via `sudo tee -a /etc/hosts`, which prompts for
// the user's password on the terminal.
func RegisterHosts(p *project.Project) error {
	for _, domain := range Domains(p) {
		present, err := hostsHas(domain)
		if err != nil {
			return err
		}

		if present {
			fmt.Printf("%s already present in /etc/hosts\n", domain)
			continue
		}

		if err := appendHost(domain); err != nil {
			return err
		}

		fmt.Printf("Added %s to /etc/hosts\n", domain)
	}

	return nil
}

func hostsHas(domain string) (bool, error) {
	f, err := os.Open("/etc/hosts")
	if err != nil {
		return false, err
	}
	defer f.Close()

	pattern := regexp.MustCompile(`^127\.0\.0\.1\s+` + regexp.QuoteMeta(domain) + `\s*$`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if pattern.MatchString(scanner.Text()) {
			return true, nil
		}
	}

	return false, scanner.Err()
}

func appendHost(domain string) error {
	cmd := exec.Command("sudo", "tee", "-a", "/etc/hosts")
	cmd.Stdin = strings.NewReader("127.0.0.1 " + domain + "\n")
	cmd.Stdout = io.Discard // tee echoes the appended line back; we print our own message instead
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
