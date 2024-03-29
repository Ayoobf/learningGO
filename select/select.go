package _select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeOut = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer2(a, b, tenSecondTimeOut)
}

func ConfigurableRacer2(a string, b string, out time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(out):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
