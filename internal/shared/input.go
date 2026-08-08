package shared

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrNoTextInput = errors.New("no positional input and stdin is not redirected")

// ReadTextInput prefers positional arguments and falls back to redirected stdin.
func ReadTextInput(args []string, stdin *os.File) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	if stdin == nil {
		return "", ErrNoTextInput
	}

	info, err := stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("inspect stdin: %w", err)
	}
	if info.Mode()&os.ModeCharDevice != 0 {
		return "", ErrNoTextInput
	}

	data, err := io.ReadAll(stdin)
	if err != nil {
		return "", fmt.Errorf("read stdin: %w", err)
	}

	text := strings.TrimSpace(string(data))
	if text == "" {
		return "", ErrNoTextInput
	}

	return text, nil
}
