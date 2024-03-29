package argutil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUri(t *testing.T) {
	tests := []struct {
		args []string
		uri  string
		err  error
	}{
		{
			args: []string{},
			err:  ErrInvalidArgument,
		},
		{
			args: []string{"http://gitea.com", "https://github.com"},
			err:  ErrInvalidArgument,
		},
		{
			args: []string{"https://github.com/harryzcy/gmg"},
			uri:  "https://github.com/harryzcy/gmg",
		},
		{
			args: []string{"http://github.com/harryzcy/gmg"},
			uri:  "http://github.com/harryzcy/gmg",
		},
		{
			args: []string{"git://github.com:harryzcy/gmg"},
			uri:  "git://github.com:harryzcy/gmg",
		},
		{
			args: []string{"git@github.com:harryzcy/gmg"},
			uri:  "git@github.com:harryzcy/gmg",
		},
		{
			args: []string{"https://git.custom.com/harryzcy/gmg"},
			uri:  "https://git.custom.com/harryzcy/gmg",
		},
		{
			args: []string{"invalid@github.com:harryzcy/gmg"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com:harryzcy"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"https://github.com:harryzcy"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com/harryzcy/"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com:harryzcy/gmg/path"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@invalid~domain.com:harryzcy/gmg"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com:invalid.username/gmg"},
			err:  ErrInvalidURI,
		},
		{
			args: []string{"git@github.com:harryzcy/invalid~repo"},
			err:  ErrInvalidURI,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			uri, err := GetURI(test.args)
			assert.Equal(t, test.uri, uri)
			assert.Equal(t, test.err, err)
		})
	}
}
