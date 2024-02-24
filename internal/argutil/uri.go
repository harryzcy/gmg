package argutil

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrInvalidArgument = fmt.Errorf("invalid argument")
	ErrInvalidURI      = fmt.Errorf("invalid uri")
)

func GetURI(args []string) (string, error) {
	if len(args) == 0 {
		fmt.Println("Please specify the repository URL")
		return "", ErrInvalidArgument
	} else if len(args) > 1 {
		fmt.Println("Please specify only one repository URL")
		return "", ErrInvalidArgument
	}

	uri := args[0]

	var isGit bool
	var trimmedURI string
	switch {
	case strings.HasPrefix(uri, "http://"):
		trimmedURI = strings.TrimPrefix(uri, "http://")
	case strings.HasPrefix(uri, "https://"):
		trimmedURI = strings.TrimPrefix(uri, "https://")
	case strings.HasPrefix(uri, "git@"):
		trimmedURI = strings.TrimPrefix(uri, "git@")
		isGit = true
	case strings.HasPrefix(uri, "git://"):
		trimmedURI = strings.TrimPrefix(uri, "git://")
		isGit = true
	default:
		return "", ErrInvalidURI
	}

	var parts []string
	if isGit {
		parts = strings.SplitN(trimmedURI, ":", 2)
	} else {
		parts = strings.SplitN(trimmedURI, "/", 2)
	}

	if len(parts) != 2 {
		return "", ErrInvalidURI
	}

	domain := parts[0]
	path := parts[1]

	pattern := regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]*\.)?[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.[a-zA-Z]{2,}$`)
	if !pattern.MatchString(domain) {
		return "", ErrInvalidURI
	}

	parts = strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		return "", ErrInvalidURI
	}

	username := parts[0]
	pattern = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !pattern.MatchString(username) {
		return "", ErrInvalidURI
	}

	repo := parts[1]
	pattern = regexp.MustCompile(`^[a-zA-Z0-9-_\.]+$`)
	if !pattern.MatchString(repo) {
		return "", ErrInvalidURI
	}

	return uri, nil
}
