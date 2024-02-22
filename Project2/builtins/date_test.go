package builtins_test

import (
	"regexp"
	"testing"

	"github.com/sriram-krishna/CSCE-4600/Project2/builtins"
)

func TestDate(t *testing.T) {
	got := builtins.Date()

	// Basic check for date format "2006-01-02 15:04:05 MST"
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [A-Z]{3}$`, got)
	if !match {
		t.Errorf("Date() got = %v, which does not match the expected format", got)
	}
}
