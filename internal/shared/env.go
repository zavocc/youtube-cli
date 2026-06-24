package shared

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnvironment loads environment variables from ~/.youtube.env when the file exists.
func LoadEnvironment() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("get user home directory: %w", err)
	}

	// return errors for malformed .youtube.env files, but ignore errors for non-existent files
	if err := godotenv.Load(filepath.Join(homeDir, ".youtube.env")); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("load .youtube.env: %w", err)
	}

	return nil
}
