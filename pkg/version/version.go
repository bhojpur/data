package version

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"
	"regexp"
	"strings"

	pb "github.com/bhojpur/data/pkg/api/v1/version"
)

const (
	// MajorVersion is the current major version for Bhojpur Data.
	MajorVersion = 1
	// MinorVersion is the current minor version for Bhojpur Data.
	MinorVersion = 2
	// MicroVersion is the patch number for Bhojpur Data.
	MicroVersion = 0
)

var (
	// AdditionalVersion is the string provided at release time
	// The value is passed to the linker at build time
	//
	// DO NOT set the value of this variable here. For some reason, if
	// AdditionalVersion is set here, the go linker will not overwrite it.
	AdditionalVersion string

	// Version is the semver release name of this build
	Version = &pb.Version{
		Major:      MajorVersion,
		Minor:      MinorVersion,
		Micro:      MicroVersion,
		Additional: AdditionalVersion,
	}
	// Commit is the commit hash this build was created from
	Commit string
	// Date is the time when this build was created
	Date string

	// GitCommit will be overwritten automatically by the build system
	BuildTime string
	// BuildCommit will be overwritten automatically by the build system
	BuildCommit = "HEAD"

	// Custom release have a 40 character commit hash build into the version string
	customReleaseRegex = regexp.MustCompile(`[0-9a-f]{40}`)
)

// Print writes the version info to stdout
func Print() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Commit:     %s\n", Commit)
	fmt.Printf("Build Date: %s\n", Date)
}

// FullVersion formats the version to be printed
func FullVersion() string {
	return fmt.Sprintf("%s (%s, build %s)", Version, BuildTime, BuildCommit)
}

// RC checks if the Bhojpur Data version is a release candidate or not
func RC() bool {
	return strings.Contains(AdditionalVersion, "rc")
}

// IsUnstable will return true for alpha or beta builds, and false otherwise.
func IsUnstable() bool {
	return strings.Contains(Version.Additional, "beta") || strings.Contains(Version.Additional, "alpha")
}

// PrettyPrintVersion returns a version string optionally tagged with metadata.
// For example: "1.2.3", or "1.2.3rc1" if version.Additional is "rc1".
func PrettyPrintVersion(version *pb.Version) string {
	result := PrettyPrintVersionNoAdditional(version)
	if version.Additional != "" {
		result += version.Additional
	}
	return result
}

// IsAtLeast returns true if Bhojpur Data is at least at the given version. This
// allows us to gate backwards-incompatible features on release boundaries.
func IsAtLeast(major, minor int) bool {
	return MajorVersion > major || (MajorVersion == major && MinorVersion >= minor)
}

// IsCustomRelease returns true if versionAdditional is a hex commit hash that is
// 40 characters long
func IsCustomRelease(version *pb.Version) bool {
	if version.Additional != "" && customReleaseRegex.MatchString(version.Additional) {
		return true
	}
	return false
}

// BranchFromVersion returns version string for the release branch
// patch release of .0 is always from the master. Others are from the M.m.x branch
func BranchFromVersion(version *pb.Version) string {
	if version.Micro == 0 {
		return "master"
	}
	return fmt.Sprintf("%d.%d.x", version.Major, version.Minor)
}

// PrettyVersion calls PrettyPrintVersion on Version and returns the result.
func PrettyVersion() string {
	return PrettyPrintVersion(Version)
}

// PrettyPrintVersionNoAdditional returns a version string without
// version.Additional.
func PrettyPrintVersionNoAdditional(version *pb.Version) string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Micro)
}
