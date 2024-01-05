// SPDX-License-Identifier: BSD-3-Clause

package version

import "fmt"

const (
	SemVer  = "0.0.1"
	GitHash = ""
	BuiltBy = ""
	BuiltAt = ""
)

func Version() string {
	v := SemVer

	if GitHash != "" {
		v = fmt.Sprintf("%s-%s", v, GitHash)
	}

	if BuiltBy != "" {
		v = fmt.Sprintf("%s-%s", v, BuiltBy)
	}

	if BuiltAt != "" {
		v = fmt.Sprintf("%s-%s", v, BuiltAt)
	}

	return v
}
