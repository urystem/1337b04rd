package config

import (
	"fmt"
	"os"
	"strconv"

	myerrors "1337b04rd/internal/core/domain/myErrors"
)

func mustGetEnvString(key string) (string, error) {
	str := os.Getenv(key)
	if len(str) == 0 {
		return "", fmt.Errorf("%w: env: %s - not found", myerrors.ErrConfNotFound, key)
	}
	return str, nil
}

func mustGetEnvInt(key string) (int, error) {
	if str := os.Getenv(key); len(str) == 0 {
		return 0, fmt.Errorf("%w: env: %s - not found", myerrors.ErrConfNotFound, key)
	} else if n, err := strconv.Atoi(str); err != nil {
		return 0, fmt.Errorf("%w: key: %s - invalid env, %w", myerrors.ErrConfInvalid, key, err)
	} else {
		return n, nil
	}
}

