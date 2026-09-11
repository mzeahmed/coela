package symfony

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Install downloads and installs a fresh Symfony skeleton into dir by
// shelling out to `composer create-project symfony/skeleton`. Composer's
// own output is streamed directly to the terminal.
func Install(dir, phpVersion, projectName string) error {
	cmd := exec.Command("composer", "create-project", "symfony/skeleton", dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("composer", "require", "--dev",
		"friendsofphp/php-cs-fixer:^3.95.15",
		"phpstan/phpstan:^2.2.5",
		"phpstan/phpstan-doctrine:^2.0.28",
		"rector/rector:^2.5.7",
	)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := configureComposer(filepath.Join(dir, "composer.json")); err != nil {
		return err
	}
	if err := writeQualityFiles(dir, phpVersion, projectName); err != nil {
		return err
	}

	cmd = exec.Command("composer", "dump-autoload")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func configureComposer(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var composer map[string]any
	if err := json.Unmarshal(content, &composer); err != nil {
		return fmt.Errorf("parse composer.json: %w", err)
	}
	autoloadDev, ok := composer["autoload-dev"].(map[string]any)
	if !ok {
		autoloadDev = make(map[string]any)
		composer["autoload-dev"] = autoloadDev
	}
	psr4, ok := autoloadDev["psr-4"].(map[string]any)
	if !ok {
		psr4 = make(map[string]any)
		autoloadDev["psr-4"] = psr4
	}
	psr4["App\\CsFixer\\"] = "tools/CsFixer/"

	scripts, ok := composer["scripts"].(map[string]any)
	if !ok {
		scripts = make(map[string]any)
		composer["scripts"] = scripts
	}
	scripts["rector"] = "@php vendor/bin/rector process --dry-run"
	scripts["rector:fix"] = "@php vendor/bin/rector process"
	scripts["lint"] = "@php vendor/bin/php-cs-fixer fix --dry-run --diff"
	scripts["lint:fix"] = "@php vendor/bin/php-cs-fixer fix --diff"
	scripts["stan"] = "@php vendor/bin/phpstan analyse -c phpstan.neon"

	updated, err := json.MarshalIndent(composer, "", "    ")
	if err != nil {
		return fmt.Errorf("encode composer.json: %w", err)
	}
	return os.WriteFile(path, append(updated, '\n'), 0644)
}

func writeQualityFiles(dir, phpVersion, projectName string) error {
	fixerNamespace := fixerNamespace(projectName)
	files := map[string]string{
		"phpstan.neon": qualityPHPStan,
		"rector.php": strings.NewReplacer(
			"__PHP_VERSION_ID__", rectorPHPVersion(phpVersion),
			"__PHP_MINOR__", strings.ReplaceAll(phpVersion, ".", ""),
		).Replace(qualityRector),
		".php-cs-fixer.dist.php": strings.ReplaceAll(
			qualityPHPcsFixer, "__FIXER_NAMESPACE__", fixerNamespace,
		),
		"tools/CsFixer/SplitMethodAttributeArgsFixer.php": strings.ReplaceAll(
			splitMethodAttributeArgsFixer, "__FIXER_NAMESPACE__", fixerNamespace,
		),
		"tools/CsFixer/BlankLineAfterControlStructureFixer.php": strings.ReplaceAll(
			blankLineAfterControlStructureFixer, "__FIXER_NAMESPACE__", fixerNamespace,
		),
	}
	for name, content := range files {
		path := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}

func rectorPHPVersion(phpVersion string) string {
	if phpVersion == "8.3" {
		return "80300"
	}
	return "80400"
}

// fixerNamespace turns the free-form project name into the vendor part of a
// PHP-CS-Fixer custom rule name. PHP-CS-Fixer requires an ASCII identifier
// beginning with an uppercase letter, so separators are removed and each word
// is capitalized. A stable fallback keeps even an empty or punctuation-only
// project name valid.
func fixerNamespace(projectName string) string {
	var b strings.Builder
	capitalizeNext := true

	for _, r := range projectName {
		switch {
		case r >= 'a' && r <= 'z':
			if capitalizeNext {
				b.WriteRune(r - ('a' - 'A'))
			} else {
				b.WriteRune(r)
			}
			capitalizeNext = false
		case r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			b.WriteRune(r)
			capitalizeNext = false
		default:
			capitalizeNext = true
		}
	}

	identifier := b.String()
	if identifier == "" {
		return "Project"
	}
	if identifier[0] >= '0' && identifier[0] <= '9' {
		return "Project" + identifier
	}

	return identifier
}
