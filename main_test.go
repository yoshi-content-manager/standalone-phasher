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

var expectedAHashes = []string{
	"a:ffff3f030703c1f0",
	"a:003c7e7e7e7e3c18",
}

var expectedDHashes = []string{
	"d:e14f6f67af0f0b89",
	"d:f0f0f4f2f0f8f0f0",
}

func TestSampleImagesPhash(t *testing.T) {
	samples := []string{"samples/sample1.jpg", "samples/sample2.jpg"}
	var gotP, gotA, gotD []string
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
		ph, err := goimagehash.PerceptionHash(img)
		if err != nil {
			t.Fatalf("PerceptionHash %s: %v", path, err)
		}
		ah, err := goimagehash.AverageHash(img)
		if err != nil {
			t.Fatalf("AverageHash %s: %v", path, err)
		}
		dh, err := goimagehash.DifferenceHash(img)
		if err != nil {
			t.Fatalf("DifferenceHash %s: %v", path, err)
		}
		gotP = append(gotP, ph.ToString())
		gotA = append(gotA, ah.ToString())
		gotD = append(gotD, dh.ToString())
	}

	assertSortedMatch(t, "phash", gotP, expectedPhashes)
	assertSortedMatch(t, "ahash", gotA, expectedAHashes)
	assertSortedMatch(t, "dhash", gotD, expectedDHashes)
}

func assertSortedMatch(t *testing.T, name string, got, expected []string) {
	t.Helper()
	exp := append([]string(nil), expected...)
	sort.Strings(exp)
	g := append([]string(nil), got...)
	sort.Strings(g)
	if len(g) != len(exp) {
		t.Fatalf("%s count: got %v, expected %v", name, g, exp)
	}
	for i := range exp {
		if g[i] != exp[i] {
			t.Errorf("%s mismatch: got %v, expected %v", name, g, exp)
			return
		}
	}
}
