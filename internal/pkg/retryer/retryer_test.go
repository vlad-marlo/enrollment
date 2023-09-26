package retryer

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTryWithAttempts_NoErrAfterFirstCall(t *testing.T) {
	err := TryWithAttempts(func() error { return nil }, 0, time.Microsecond)
	assert.NoError(t, err)
}

func TestTryWithAttempts_Negative(t *testing.T) {
	i := 0
	err := TryWithAttempts(func() error {
		if i == 0 {
			return errors.New("")
		}
		return nil
	}, 1, time.Microsecond)
	assert.Error(t, err)
}

type testStruct struct {
	counter int
	want    int
}

func (t *testStruct) ok() error {
	if t.counter != t.want {
		t.counter++
		return errors.New("")
	}
	return nil
}

func TestOK(t *testing.T) {
	str := testStruct{0, 1}
	assert.Error(t, str.ok())
	assert.NoError(t, str.ok())
}

func TestTryWithAttempts_Positive(t *testing.T) {
	tests := testStruct{0, 2}
	err := TryWithAttempts(tests.ok, 3, time.Microsecond)
	assert.NoError(t, err)
}

func TestTryWithAttempts_Negative_expired(t *testing.T) {
	tests := testStruct{0, 3}
	err := TryWithAttempts(tests.ok, 3, time.Microsecond)
	assert.Error(t, err)
}

func TestTryWithAttemptsCtx(t *testing.T) {
	err := TryWithAttemptsCtx(context.Background(), func(ctx context.Context) error {
		return nil
	}, 0, time.Second)
	assert.NoError(t, err)
}
