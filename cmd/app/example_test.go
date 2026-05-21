package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	tests := map[string]struct {
		giveInput  string
		wantResult string
		wantErr    error
	}{
		"success": {giveInput: "valid", wantResult: "valid"},
		"error":   {giveInput: "invalid", wantErr: assert.AnError},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := func() (string, error) {
				if tt.giveInput == "valid" {
					return tt.giveInput, nil
				}

				return "", assert.AnError
			}()

			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
