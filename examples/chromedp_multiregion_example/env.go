package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var requiredEnv = []string{
	"STAGEHAND_API_URL",
	"MODEL_API_KEY",
	"BROWSERBASE_API_KEY",
	"BROWSERBASE_PROJECT_ID",
}

func loadExampleEnv() {
	envPath, ok := findEnvPath()
	if !ok {
		panic("Missing examples/.env (expected in repo examples/ directory).")
	}

	file, err := os.Open(envPath)
	if err != nil {
		panic(fmt.Sprintf("failed to read examples/.env: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		if _, ok := os.LookupEnv(key); !ok {
			os.Setenv(key, value)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to parse examples/.env: %v", err))
	}

	missing := []string{}
	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		panic("Missing required env vars: " + strings.Join(missing, ", ") + " (from examples/.env)")
	}

	if os.Getenv("STAGEHAND_BASE_URL") == "" {
		os.Setenv("STAGEHAND_BASE_URL", os.Getenv("STAGEHAND_API_URL"))
	}
}

func findEnvPath() (string, bool) {
	current, err := os.Getwd()
	if err != nil {
		return "", false
	}
	for {
		candidate := filepath.Join(current, "examples", ".env")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", false
		}
		current = parent
	}
}
