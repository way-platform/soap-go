//go:build mage

package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// Build runs a full CI build.
func Build() {
	mg.SerialDeps(
		Download,
		Generate,
		Lint,
		Test,
		Tidy,
		Diff,
	)
}

// Lint runs the Go linter.
func Lint() error {
	return forEachGoMod(func(dir string) error {
		return tool(dir, "golangci-lint", "run", "--fix", "--path-prefix", dir, "--build-tags", "mage").Run()
	})
}

// Test runs the Go unit tests.
func Test() error {
	if err := os.MkdirAll("build", 0o700); err != nil {
		return err
	}
	return cmd(
		root(),
		"go",
		"test",
		"-cover",
		"-tags=synctest",
		"./...",
		"-coverprofile",
		"build/cover.out",
	).Run()
}

// IntegrationTest runs the Go integration tests.
func IntegrationTest() error {
	if err := os.MkdirAll("build", 0o700); err != nil {
		return err
	}
	return cmd(
		root(),
		"go",
		"test",
		"-v",
		"-tags",
		"integration,synctest",
		"./...",
		"-coverprofile",
		"build/integration-cover.out",
	).Run()
}

// Download downloads the Go dependencies.
func Download() error {
	return forEachGoMod(func(dir string) error {
		return cmd(dir, "go", "mod", "download").Run()
	})
}

// Tidy tidies the Go mod files.
func Tidy() error {
	return forEachGoMod(func(dir string) error {
		return cmd(dir, "go", "mod", "tidy", "-v").Run()
	})
}

// Diff checks for git diffs.
func Diff() error {
	return cmd(root(), "git", "diff", "--exit-code").Run()
}

// Generate runs all code generators.
func Generate() error {
	return forEachGoMod(func(dir string) error {
		return cmd(dir, "go", "generate", "-v", "./...").Run()
	})
}

func forEachGoMod(f func(dir string) error) error {
	return filepath.WalkDir(root(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || d.Name() != "go.mod" {
			return nil
		}
		return f(filepath.Dir(path))
	})
}

func cmd(dir string, command string, args ...string) *exec.Cmd {
	c := exec.Command(command, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	// Strip GOWORK=off inherited from the mage bootstrap script so that go
	// commands in targets auto-detect the go.work workspace file.
	env := make([]string, 0, len(os.Environ()))
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, "GOWORK=") {
			env = append(env, e)
		}
	}
	c.Env = env
	return c
}

func tool(dir string, tool string, args ...string) *exec.Cmd {
	cmdArgs := []string{"tool", "-modfile", filepath.Join(root(), "tools", "go.mod"), tool}
	c := cmd(dir, "go", append(cmdArgs, args...)...)
	// -modfile is incompatible with workspace mode; restore GOWORK=off.
	c.Env = append(c.Env, "GOWORK=off")
	return c
}

func root(subdirs ...string) string {
	result, err := sh.Output("git", "rev-parse", "--show-toplevel")
	if err != nil {
		panic(err)
	}
	return filepath.Join(append([]string{result}, subdirs...)...)
}
