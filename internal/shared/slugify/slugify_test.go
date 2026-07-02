package slugify_test

import (
	"testing"

	"github.com/wigata-intech/logres/internal/shared/slugify"
)

func TestSlugify(t *testing.T) {
	testCases := []struct {
		testName        string
		testInput       string
		testExpectation string
	}{
		{
			testName:        "internal/shared/slugify/slugify.go: slugify strings \"Produk Kecantikan Terbaru\"",
			testInput:       "Produk Kecantikan Terbaru",
			testExpectation: "produk-kecantikan-terbaru",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			slug := slugify.Slugify(tc.testInput)
			if slug != tc.testExpectation {
				t.Errorf("result: %s | expectation: %s", slug, tc.testExpectation)
			}
		})
	}
}
