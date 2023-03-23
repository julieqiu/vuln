package vuln

import "github.com/julieqiu/vuln/internal/govulncheck"

type Cmd = govulncheck.Cmd

var Command = govulncheck.Command

var (
	ErrMissingArgPatters    = govulncheck.ErrMissingArgPatterns
	ErrVulnerabilitiesFound = govulncheck.ErrVulnerabilitiesFound
)
