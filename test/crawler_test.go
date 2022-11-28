package test

import (
	"testing"
)

// ScanBlocks test
func TestScanBlocks(t *testing.T) {
	c := crawler()
	_, err := c.ScanBlocks(5)
	if err != nil {
		t.Errorf("ScanBlocks error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// ScanBlocksFromTo test
func TestScanBlocksFromTo(t *testing.T) {
	c := crawler()

	res, err := c.ScanBlocksFromTo(28905235, 28905236)
	if err != nil {
		t.Errorf("ScanBlocks error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(res) == 0 {
		t.Errorf("ScanBlocks res was incorect, got: %q, want: %q.", res, "not empty")
	}
}
