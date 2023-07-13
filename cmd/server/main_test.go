package main

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"testing"
)

func TestNewApp(t *testing.T) {
	assert.NoError(t, fx.ValidateApp(NewApp()))
}
