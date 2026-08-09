package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunSkipsGitDirectoryForAbsoluteStartPath(t *testing.T) {
	dir := t.TempDir()

	codeownersPath := filepath.Join(dir, "CODEOWNERS")
	if err := os.WriteFile(codeownersPath, []byte("* @example/all\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "visible.txt"), []byte("visible\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	gitDir := filepath.Join(dir, ".git")
	if err := os.Mkdir(gitDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(gitDir, "config"), []byte("[core]\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr bytes.Buffer
	exitCode := Run([]string{"--file", codeownersPath, dir}, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("Run() exit code = %d, stderr = %q", exitCode, stderr.String())
	}

	output := stdout.String()
	if !strings.Contains(output, "visible.txt") {
		t.Fatalf("Run() output %q does not include visible file", output)
	}
	if strings.Contains(output, ".git") {
		t.Fatalf("Run() output %q includes .git directory contents", output)
	}
}
