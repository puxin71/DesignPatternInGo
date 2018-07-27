package headFirstPatterns_test

import (
	"testing"

	"github.com/puxin71/DesignPatternInGo/headFirstPatterns"
	"github.com/stretchr/testify/require"
)

func TestNewSubject(t *testing.T) {
	aSubject, err := headFirstPatterns.NewSubject(-1)
	require.NotNilf(t, err, "error message: %s", "expect an error")
	require.Truef(t, aSubject == nil, "error message: %s", "expect nil object")

	aSubject, err = headFirstPatterns.NewSubject(headFirstPatterns.MaxNumOfObservers)
	require.Nil(t, err, "error message: %s", "expect no error")
	require.Truef(t, aSubject != nil, "error message: %s", "expect concrete object")
}

func TestNewObserver(t *testing.T) {
	aObserver := headFirstPatterns.NewObserver()
	require.NotNilf(t, aObserver, "error message: %s", "expect concrete object")
}
