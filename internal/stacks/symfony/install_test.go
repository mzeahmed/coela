package symfony

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigureComposerAddsQualityAutoloadAndScripts(t *testing.T) {
	path := filepath.Join(t.TempDir(), "composer.json")
	if err := os.WriteFile(path, []byte(`{"name":"example/app"}`), 0644); err != nil {
		t.Fatal(err)
	}

	if err := configureComposer(path); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var composer map[string]any
	if err := json.Unmarshal(content, &composer); err != nil {
		t.Fatal(err)
	}
	autoloadDev := composer["autoload-dev"].(map[string]any)
	psr4 := autoloadDev["psr-4"].(map[string]any)
	if got := psr4["App\\CsFixer\\"]; got != "tools/CsFixer/" {
		t.Fatalf("custom fixer autoload = %v", got)
	}
	scripts := composer["scripts"].(map[string]any)
	for _, name := range []string{"rector", "rector:fix", "lint", "lint:fix", "stan"} {
		if _, ok := scripts[name]; !ok {
			t.Errorf("missing Composer script %q", name)
		}
	}
}

func TestWriteQualityFilesUsesSelectedPHPVersion(t *testing.T) {
	for _, phpVersion := range []string{"8.3", "8.4"} {
		t.Run(phpVersion, func(t *testing.T) {
			dir := t.TempDir()
			if err := writeQualityFiles(dir, phpVersion, "booking-app"); err != nil {
				t.Fatal(err)
			}

			rector, err := os.ReadFile(filepath.Join(dir, "rector.php"))
			if err != nil {
				t.Fatal(err)
			}
			versionID := strings.ReplaceAll(phpVersion, ".", "0") + "00"
			if !strings.Contains(string(rector), "withPhpVersion("+versionID+")") {
				t.Fatalf("Rector configuration does not target PHP %s", phpVersion)
			}
			if !strings.Contains(string(rector), "withPhpSets(php"+strings.ReplaceAll(phpVersion, ".", "")+": true)") {
				t.Fatalf("Rector PHP set does not target PHP %s", phpVersion)
			}
			for _, name := range []string{
				"phpstan.neon",
				".php-cs-fixer.dist.php",
				"tools/CsFixer/SplitMethodAttributeArgsFixer.php",
				"tools/CsFixer/BlankLineAfterControlStructureFixer.php",
			} {
				if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
					t.Errorf("missing generated quality file %s: %v", name, err)
				}
			}

			for _, name := range []string{
				".php-cs-fixer.dist.php",
				"tools/CsFixer/SplitMethodAttributeArgsFixer.php",
				"tools/CsFixer/BlankLineAfterControlStructureFixer.php",
			} {
				content, err := os.ReadFile(filepath.Join(dir, name))
				if err != nil {
					t.Fatal(err)
				}
				if !strings.Contains(string(content), "BookingApp/") {
					t.Errorf("%s does not use the project fixer namespace", name)
				}
			}
		})
	}
}

func TestFixerNamespace(t *testing.T) {
	tests := map[string]string{
		"jobscan":      "Jobscan",
		"booking-app":  "BookingApp",
		"my project":   "MyProject",
		"123-bookings": "Project123Bookings",
		"---":          "Project",
	}

	for projectName, want := range tests {
		if got := fixerNamespace(projectName); got != want {
			t.Errorf("fixerNamespace(%q) = %q, want %q", projectName, got, want)
		}
	}
}
