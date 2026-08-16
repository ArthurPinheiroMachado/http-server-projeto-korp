package util

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func ErrorWithContext(err error, context string) error {
	if err != nil {
		return fmt.Errorf("%s\n%w", context, err)
	}

	return nil
}

type ErrorContext struct{ context string }

func CreateErrorContext(context string) ErrorContext {
	return ErrorContext{context}
}

func (context *ErrorContext) Apply(err error) error {
	if err != nil {
		return fmt.Errorf("%s\n\ncaused by:\n%w", context.context, err)
	}

	return nil
}

func SendHttpError(resp http.ResponseWriter, code int, err error) {
	http.Error(resp, ErrorInspect(err).Error(), code)
}

func ErrorInspect(err error) error {
	if err != nil {
		timestamp := time.Now().Format("02/01/2006 15:04:05")
		fmt.Fprintf(os.Stderr, "\n[%s] %s\n------------------------------\n\n", timestamp, err)
	}
	return err
}
