package retry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type RetryConfig[T, R any] struct {
	attempts        int
	delay           time.Duration
	retryAbleErrors []error

	ctx context.Context
	fn  func(T) (R, error)
}

func Func[T, R any](ctx context.Context, fn func(T) (R, error)) *RetryConfig[T, R] {
	return &RetryConfig[T, R]{
		attempts:        3,
		delay:           1 * time.Second,
		retryAbleErrors: []error{},

		ctx: ctx,
		fn:  fn,
	}
}

func (r *RetryConfig[T, R]) Attempts(attempts int) *RetryConfig[T, R] {
	r.attempts = attempts
	return r
}

func (r *RetryConfig[T, R]) Delay(delay time.Duration) *RetryConfig[T, R] {
	r.delay = delay
	return r
}

func (r *RetryConfig[T, R]) RetryAbleErrors(retryAbleErrors []error) *RetryConfig[T, R] {
	r.retryAbleErrors = retryAbleErrors
	return r
}

func (r *RetryConfig[T, R]) Do(args T) (R, error) {
	var err error
	var result R
	for i := 0; i < r.attempts; i++ {
		result, err = r.fn(args)
		if err == nil {
			return result, nil
		}
		if !r.isRetryAbleError(err) {
			return result, err
		}
		fmt.Printf("function %v failed with error: %v, retrying after %v \n", getFunctionName(r.fn), err, r.delay)
		time.Sleep(r.delay)
	}
	return result, err
}

func (r *RetryConfig[T, R]) isRetryAbleError(err error) bool {
	if len(r.retryAbleErrors) == 0 {
		return true
	}
	for _, retryError := range r.retryAbleErrors {
		if errors.Is(retryError, err) {
			return true
		}
	}
	return false
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
