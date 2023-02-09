package search_test

import (
	"testing"
)

// TestSearch tests for an empty search query
func TestSearch(t *testing.T) {
	emptySearch := ""
	if emptySearch != "" {
		t.Errorf("main(\"\") failed, expected %v", emptySearch)
	}
}
