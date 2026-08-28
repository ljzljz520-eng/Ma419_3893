package content

import "testing"

func TestSeedContent(t *testing.T) {
	if !ValidateSeed() {
		t.Fatal("seed")
	}
	if len(CategoryNames()) != 4 {
		t.Fatal("categories")
	}
}
