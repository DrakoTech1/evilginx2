// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package packagesdriver fetches type sizes for go/packages and go/analysis.
package packagesdriver

import (
	"context"
	"fmt"
<<<<<<< HEAD
=======
	"go/types"
>>>>>>> deathstrox/main
	"strings"

	"golang.org/x/tools/internal/gocommand"
)

<<<<<<< HEAD
func GetSizesForArgsGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (string, string, error) {
=======
var debug = false

func GetSizesGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (types.Sizes, error) {
>>>>>>> deathstrox/main
	inv.Verb = "list"
	inv.Args = []string{"-f", "{{context.GOARCH}} {{context.Compiler}}", "--", "unsafe"}
	stdout, stderr, friendlyErr, rawErr := gocmdRunner.RunRaw(ctx, inv)
	var goarch, compiler string
	if rawErr != nil {
<<<<<<< HEAD
		rawErrMsg := rawErr.Error()
		if strings.Contains(rawErrMsg, "cannot find main module") ||
			strings.Contains(rawErrMsg, "go.mod file not found") {
			// User's running outside of a module.
			// All bets are off. Get GOARCH and guess compiler is gc.
=======
		if rawErrMsg := rawErr.Error(); strings.Contains(rawErrMsg, "cannot find main module") || strings.Contains(rawErrMsg, "go.mod file not found") {
			// User's running outside of a module. All bets are off. Get GOARCH and guess compiler is gc.
>>>>>>> deathstrox/main
			// TODO(matloob): Is this a problem in practice?
			inv.Verb = "env"
			inv.Args = []string{"GOARCH"}
			envout, enverr := gocmdRunner.Run(ctx, inv)
			if enverr != nil {
<<<<<<< HEAD
				return "", "", enverr
			}
			goarch = strings.TrimSpace(envout.String())
			compiler = "gc"
		} else if friendlyErr != nil {
			return "", "", friendlyErr
		} else {
			// This should be unreachable, but be defensive
			// in case RunRaw's error results are inconsistent.
			return "", "", rawErr
=======
				return nil, enverr
			}
			goarch = strings.TrimSpace(envout.String())
			compiler = "gc"
		} else {
			return nil, friendlyErr
>>>>>>> deathstrox/main
		}
	} else {
		fields := strings.Fields(stdout.String())
		if len(fields) < 2 {
<<<<<<< HEAD
			return "", "", fmt.Errorf("could not parse GOARCH and Go compiler in format \"<GOARCH> <compiler>\":\nstdout: <<%s>>\nstderr: <<%s>>",
=======
			return nil, fmt.Errorf("could not parse GOARCH and Go compiler in format \"<GOARCH> <compiler>\":\nstdout: <<%s>>\nstderr: <<%s>>",
>>>>>>> deathstrox/main
				stdout.String(), stderr.String())
		}
		goarch = fields[0]
		compiler = fields[1]
	}
<<<<<<< HEAD
	return compiler, goarch, nil
=======
	return types.SizesFor(compiler, goarch), nil
>>>>>>> deathstrox/main
}
