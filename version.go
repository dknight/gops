package gops

// Don't want to use https://pkg.go.dev/github.com/hashicorp/go-version
// this is overkill for such small project where only me is maintainer.

// Version gets the version of gops. Maybe do better structure
// major, minor, patch.
func Version() string {
	return "1.0.11"
}
