// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package govulncheck provides functionality to support the govulncheck command.
package govulncheck

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/packages"
	"github.com/julieqiu/vuln/internal/result"
)

// LoadMode is the level of information needed for each package
// for running golang.org/x/tools/go/packages.Load.
var LoadMode = packages.NeedName | packages.NeedImports | packages.NeedTypes |
	packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedDeps |
	packages.NeedModule

// IsCalled reports whether the vulnerability is called, therefore
// affecting the target source code or binary.
func IsCalled(v *result.Vuln) bool {
	for _, m := range v.Modules {
		for _, p := range m.Packages {
			if len(p.CallStacks) > 0 {
				return true
			}
		}
	}
	return false
}

// FuncName returns the full qualified function name from sf,
// adjusted to remove pointer annotations.
func FuncName(sf *result.StackFrame) string {
	var n string
	if sf.Receiver == "" {
		n = fmt.Sprintf("%s.%s", sf.Package, sf.Function)
	} else {
		n = fmt.Sprintf("%s.%s", sf.Receiver, sf.Function)
	}
	return strings.TrimPrefix(n, "*")
}

// Pos returns the position of the call in sf as string.
// If position is not available, return "".
func Pos(sf *result.StackFrame) string {
	if sf.Position.IsValid() {
		return sf.Position.String()
	}
	return ""
}
