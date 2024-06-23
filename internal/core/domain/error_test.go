package domain

import (
	"fmt"
	"testing"
)

func TestAppErrorCode_IsErrEqual(t *testing.T) {
	tc := []struct {
		code     AppErrorCode
		err      error
		expected bool
	}{
		{code: ErrCodeNotFound, err: fmt.Errorf("generic error"), expected: false},
		{code: ErrCodeNotFound, err: FromError(ErrCodeNotFound, fmt.Errorf("generic error")), expected: true},
		{code: ErrCodeInternal, err: FromError(ErrCodeNotFound, fmt.Errorf("generic error")), expected: false},
	}

	for i, s := range tc {
		res := s.code.IsErrEqual(s.err)
		if res != s.expected {
			t.Errorf("%d: expected %v, got %v", i, s.expected, res)
		}
	}
}
