package headFirstPatterns_test

import (
	"sync"
	"testing"

	"github.com/puxin71/DesignPatternInGo/headFirstPatterns"
	"github.com/stretchr/testify/require"
)

// Subject test
func TestNewSubject(t *testing.T) {
	aSubject := headFirstPatterns.NewSubject()
	require.NotNilf(t, aSubject, "error message: %s", "expect concrete object")
}

func TestNumOfObservers(t *testing.T) {
	aSubject := headFirstPatterns.NewSubject()
	require.EqualValuesf(t, 0, aSubject.NumOfObservers(), "error message: %d", aSubject.NumOfObservers())
}

/*
You needn't close every channel when you've finished with it.
It's only necessary to close a channel when it is important to tell the receiving goroutines
that all data have been sent. A channel that the garbage collector determines to be unreachable
will have its resources reclaimed whether or not it is closed.
(Don't confuse this with the close operation for open files. It is important to call the Close
method on every file when you've finished with it.)
*/

func TestRegisterObserver(t *testing.T) {
	aSubject := headFirstPatterns.NewSubject()
	for i := 0; i < 10; i++ {
		aObserver := headFirstPatterns.NewObserver()
		err := aSubject.RegisterObserver(aObserver)
		require.Nilf(t, err, "error message: %s", "expect no error since the slice grows")
	}
	require.EqualValuesf(t, 10, aSubject.NumOfObservers(), "error message: %d", aSubject.NumOfObservers())
}

func TestRegisterObserver_concurrent(t *testing.T) {
	var wg sync.WaitGroup

	aSubject := headFirstPatterns.NewSubject()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			aObserver := headFirstPatterns.NewObserver()
			err := aSubject.RegisterObserver(aObserver)
			require.Nilf(t, err, "error message: %s", "expect no error since the slice grows")
		}()
	}
	wg.Wait()
	require.EqualValuesf(t, 10, aSubject.NumOfObservers(), "error message: %d", aSubject.NumOfObservers())
}

func TestRegisterObserver_noDuplicateObserver(t *testing.T) {
	var wg sync.WaitGroup

	aSubject := headFirstPatterns.NewSubject()
	aObserver := headFirstPatterns.NewObserver()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := aSubject.RegisterObserver(aObserver)
			require.Nilf(t, err, "error message: %s", "expect no error since the slice grows")
		}()
	}
	wg.Wait()
	require.EqualValuesf(t, 1, aSubject.NumOfObservers(), "error message: %d", aSubject.NumOfObservers())
}

// Observer test

func TestNewObserver(t *testing.T) {
	aObserver := headFirstPatterns.NewObserver()
	require.NotNilf(t, aObserver, "error message: %s", "expect concrete object")
	require.NotNilf(t, aObserver.Channel(), "error message: %s", "expect concrete channel")
}
