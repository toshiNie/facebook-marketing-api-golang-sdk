package fb

import (
	"fmt"
	"github.com/cenk/backoff"
	"net/http"
)

type retryTransport struct {
	next http.RoundTripper
	bo   backoff.BackOff
}

func NewRetryTransport(bo backoff.BackOff, next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}
	if bo == nil {
		bo = backoff.NewExponentialBackOff()
	}

	return &retryTransport{
		next: next,
		bo:   bo,
	}
}

func (t *retryTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	var resp *http.Response
	var attempt int
	err := backoff.Retry(func() error {
		attempt++
		var e error

		resp, e = t.next.RoundTrip(r) // nolint:bodyclose // not a correct linter detection

		if e != nil {
			return e
		} else if resp.StatusCode >= 500 {
			resp.Body.Close()

			return fmt.Errorf("unexpected status %s from facebook, attempt %d", resp.Status, attempt)
		}

		return nil
	}, backoff.WithContext(t.bo, r.Context()))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
