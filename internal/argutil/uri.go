package argutil

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrInvalidArgument = fmt.Errorf("invalid argument")
	ErrInvalidUri      = fmt.Errorf("invalid uri")
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
	var trimmedUri string
	if strings.HasPrefix(uri, "http://") {
		trimmedUri = strings.TrimPrefix(uri, "http://")
	} else if strings.HasPrefix(uri, "https://") {
		trimmedUri = strings.TrimPrefix(uri, "https://")
	} else if strings.HasPrefix(uri, "git@") {
		trimmedUri = strings.TrimPrefix(uri, "git@")
		isGit = true
	} else if strings.HasPrefix(uri, "git://") {
		trimmedUri = strings.TrimPrefix(uri, "git://")
		isGit = true
	} else {
		return "", ErrInvalidUri
	}

	var parts []string
	if isGit {
		parts = strings.SplitN(trimmedUri, ":", 2)
	} else {
		parts = strings.SplitN(trimmedUri, "/", 2)
	}

	if len(parts) != 2 {
		return "", ErrInvalidUri
	}

	domain := parts[0]
	path := parts[1]

	pattern := regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]*\\.)?[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\\.[a-zA-Z]{2,}$`)
	if !pattern.MatchString(domain) {
		return "", ErrInvalidUri
	}

	parts = strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		return "", ErrInvalidUri
	}

	username := parts[0]
	pattern = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !pattern.MatchString(username) {
		return "", ErrInvalidUri
	}

	repo := parts[1]
	pattern = regexp.MustCompile(`^[a-zA-Z0-9-_\\.]+$`)
	if !pattern.MatchString(repo) {
		return "", ErrInvalidUri
	}

	return uri, nil
}
