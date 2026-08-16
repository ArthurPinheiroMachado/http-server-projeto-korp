package util

import (
	"fmt"
	"os"
	"strconv"
)

func EnvAsResult(name string) (string, error) {
	ok, _exists := os.LookupEnv(name)

	if _exists {
		return ok, nil
	}

	return "", fmt.Errorf("a variável %s não foi declarada", name)
}

func EnvAsIntegerResult(name string) (int, error) {
	trace := CreateErrorContext("EnvAsIntegerResult")

	numstr, numstrE := EnvAsResult(name)
	if numstrE != nil {
		return 0, trace.Apply(numstrE)
	}

	atoi, atoiE := strconv.Atoi(numstr)

	if atoiE != nil {
		return 0, trace.Apply(atoiE)
	}

	return atoi, nil
}
