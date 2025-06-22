package ahautil

import (
	"fmt"
	"strings"
)

func newErrorBadStatusCode(statusCode int) error {
	return fmt.Errorf("bad response status code (%d)", statusCode)
}

func newErrorDateNotFound(dates []string) error {
	return fmt.Errorf("date not found (%s)", strings.Join(dates, ","))
}
