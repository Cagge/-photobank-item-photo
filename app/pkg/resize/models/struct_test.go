package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidConfigSuccess(t *testing.T) {
	m := ValidConfig{Size: 123}
	err := m.ValidConfigSize()
	assert.Equal(t, err, nil)
}
func TestValidConfigFail(t *testing.T) {
	m := ValidConfig{Size: 0}
	err := m.ValidConfigSize()
	assert.NotNil(t, err)
}