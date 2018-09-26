package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/guard/pkg/service/types"
)

func TestID(t *testing.T) {
	tests := []struct {
		name    string
		id      ID
		isValid bool
	}{
		{"ID is empty", "", false},
		{"ID is invalid", "abc-def-ghi", false},
		{"ID is not UUID v4", "41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false},
		{"ID in lowercase", "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true},
		{"ID in uppercase", "41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.isValid, test.id.IsValid(), test.name)
		assert.Equal(t, test.id, ID(test.id.String()), test.name)
	}
}
