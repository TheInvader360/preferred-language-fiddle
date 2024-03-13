package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPreferredCy(t *testing.T) {
	type test struct {
		acceptLanguageHeader string
		expected             bool
	}
	tests := []test{
		{"", false},
		{"cy", true},
		{"de", false},
		{"en", false},
		{"fr", false},
		{"cy,de", true},
		{"de,cy", true},
		{"en,fr", false},
		{"fr,en", false},
		{"cy,de,en,fr", true},
		{"fr,en,de,cy", false},
		{"cy,de;q=0.9,en;q=0.8,fr;q=0.7", true},
		{"de,en;q=0.9,fr;q=0.8,cy;q=0.7", false},
		{"en,fr;q=0.9,cy;q=0.8,de;q=0.7", false},
		{"fr,cy;q=0.9,de;q=0.8,en;q=0.7", true},
		{"en;q=0.8,fr;q=0.7,cy,de;q=0.9", true},
		{"fr;q=0.8,cy;q=0.7,de,en;q=0.9", false},
		{"cy;q=0.8,de;q=0.7,en,fr;q=0.9", false},
		{"de;q=0.8,en;q=0.7,fr,cy;q=0.9", true},
		{"cy,en;q=0.95,en-AU;q=0.9,en-BZ;q=0.85,en-CA;q=0.8,en-GB;q=0.75,en-IE;q=0.7,en-IN;q=0.65,en-JM;q=0.6,en-NZ;q=0.55,en-PH;q=0.5,en-TT;q=0.45,en-US;q=0.4,en-ZA;q=0.35,en-ZW;q=0.3", true},
		{"en,en-AU;q=0.95,en-BZ;q=0.9,en-CA;q=0.85,en-GB;q=0.8,en-IE;q=0.75,en-IN;q=0.7,en-JM;q=0.65,en-NZ;q=0.6,en-PH;q=0.55,en-TT;q=0.5,en-US;q=0.45,en-ZA;q=0.4,en-ZW;q=0.35,cy;q=0.3", false},
		{"cy,en;q=0.5", true},
		{"cy,en-AU;q=0.5", true},
		{"cy,en-BZ;q=0.5", true},
		{"cy,en-CA;q=0.5", true},
		{"cy,en-GB;q=0.5", true},
		{"cy,en-IE;q=0.5", true},
		{"cy,en-IN;q=0.5", true},
		{"cy,en-JM;q=0.5", true},
		{"cy,en-NZ;q=0.5", true},
		{"cy,en-PH;q=0.5", true},
		{"cy,en-TT;q=0.5", true},
		{"cy,en-US;q=0.5", true},
		{"cy,en-ZA;q=0.5", true},
		{"cy,en-ZW;q=0.5", true},
		{"en,cy;q=0.5", false},
		{"en-AU,cy;q=0.5", false},
		{"en-BZ,cy;q=0.5", false},
		{"en-CA,cy;q=0.5", false},
		{"en-GB,cy;q=0.5", false},
		{"en-IE,cy;q=0.5", false},
		{"en-IN,cy;q=0.5", false},
		{"en-JM,cy;q=0.5", false},
		{"en-NZ,cy;q=0.5", false},
		{"en-PH,cy;q=0.5", false},
		{"en-TT,cy;q=0.5", false},
		{"en-US,cy;q=0.5", false},
		{"en-ZA,cy;q=0.5", false},
		{"en-ZW,cy;q=0.5", false},
	}
	for _, tc := range tests {
		assert.Equal(t, isPreferredCy(tc.acceptLanguageHeader), tc.expected)
	}
}
