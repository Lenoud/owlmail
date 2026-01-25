// Package maildev tests.
//
// Note: The environment variable compatibility functions have been moved to
// the internal/config package. See internal/config/config_test.go for those tests.
//
// This package now only contains API documentation comments.
package maildev

import "testing"

func TestPackageExists(t *testing.T) {
	// This test verifies the package can be imported.
	// The actual environment variable handling tests are in internal/config.
	t.Log("maildev package exists and can be imported")
}
