package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sort"
	"testing"

	"github.com/corona10/goimagehash"
)

var expectedPhashes = []string{
	"p:af97d2205c6b1f82",
	"p:95274a996be4349e",
}

func TestSampleImagesPhash(t *testing.T) {
	samples := []string{"samples/sample1.jpg", "samples/sample2.jpg"}
	var computed []string
	for _, path := range samples {
		f, err := os.Open(path)
		if err != nil {
			t.Fatalf("open %s: %v", path, err)
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			t.Fatalf("decode %s: %v", path, err)
		}
		hash, err := goimagehash.PerceptionHash(img)
		if err != nil {
			t.Fatalf("PerceptionHash %s: %v", path, err)
		}
		computed = append(computed, hash.ToString())
	}

	expected := make([]string, len(expectedPhashes))
	copy(expected, expectedPhashes)
	sort.Strings(expected)
	got := make([]string, len(computed))
	copy(got, computed)
	sort.Strings(got)

	if len(got) != len(expected) {
		t.Fatalf("phash count: got %v, expected %v", got, expected)
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("phashes mismatch: got %v, expected %v", got, expected)
			break
		}
	}
}
