package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameFromGitURI(t *testing.T) {
	tests := []struct {
		gitURI string
		name   string
	}{
		{
			gitURI: "https://github.com/harryzcy/gmg.git",
			name:   "gmg",
		},
		{
			gitURI: "https://github.com/harryzcy/test.git",
			name:   "test",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			name := getNameFromGitURI(test.gitURI)
			assert.Equal(t, test.name, name)
		})
	}
}

func TestValidateGitURI(t *testing.T) {
	tests := []struct {
		gitURI string
		valid  bool
	}{
		{
			gitURI: "https://go.zcy.dev/gmg.git",
			valid:  true,
		},
		{
			gitURI: "",
			valid:  false,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			valid := validateGitURI(test.gitURI)
			assert.Equal(t, test.valid, valid)
		})
	}
}
