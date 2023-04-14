package valuedurliterator

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestValuedUrlIteratorHappyPath(t *testing.T) {
	dummyFileContent := strings.NewReader(`http://api.tech.com/item/411102 336
	http://api.tech.com/item/152121 930`)

	vui := NewValuedUrlIterator(dummyFileContent)
	if !vui.HasNextValuedUrl() {
		t.Errorf("Expected NextValuedUrl to be exist")
	}
	if _, val := vui.GetNextValuedUrl(); val != 336 {
		t.Errorf("Expected valued url to be in order with buffer reader, got %d", val)
	}
	if !vui.HasNextValuedUrl() {
		t.Errorf("Expected NextValuedUrl to be exist")
	}
	if _, val := vui.GetNextValuedUrl(); val != 930 {
		t.Errorf("Expected valued url to be in order with buffer reader, got %d", val)
	}
	if vui.HasNextValuedUrl() {
		t.Errorf("Expected to finish reading from buffer reader")
	}
}

func TestErrorLogging(t *testing.T) {
	dummyFileContent := strings.NewReader(`http://api.tech.com/item/411102 336
	http://api.tech.com/item/152121 abc`)

	vui := NewValuedUrlIterator(dummyFileContent)
	if !vui.HasNextValuedUrl() {
		t.Errorf("Expected NextValuedUrl to be exist")
	}
	if _, val := vui.GetNextValuedUrl(); val != 336 {
		t.Errorf("Expected valued url to be in order with buffer reader, got %d", val)
	}
	if !vui.HasNextValuedUrl() {
		t.Errorf("Expected NextValuedUrl to be exist")
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	if _, val := vui.GetNextValuedUrl(); val != 0 {
		t.Errorf("Expected valued url to be in the wrong format, got %d", val)
	}
	logBuffer := buf.String()
	if !strings.Contains(logBuffer, "Failed to parse valued url from line") {
		t.Errorf("Expected to log error when line in the wrong format, got: %s", logBuffer)
	}
	if vui.HasNextValuedUrl() {
		t.Errorf("Expected to finish reading from buffer reader")
	}
}
