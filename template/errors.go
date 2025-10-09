package template

import (
	"errors"
)

var (
	ErrNoGitRepo                                        = errors.New("not in a git repository")
	ErrBase64EncodeIsMandatoryWhenGzipIsEnabled         = errors.New("base64_encode is mandatory when gzip is enabled")
	ErrNoPartsFoundInCloudinitResourceDeclaration       = errors.New("no parts found in the cloudinit resource declaration")
	ErrUnableToParsePartsInCloudinitResourceDeclaration = errors.New("unable to parse parts in cloudinit resource declaration")
)
