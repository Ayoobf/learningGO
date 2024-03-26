package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())

		if err != nil {
			return // todo: log error here
		}

		fmt.Fprint(writer, data)

	}
}

type StubStore struct {
	response string
}

func (s *StubStore) Cancel() {
	//TODO implement me
	panic("implement me")
}

func (s *StubStore) Fetch() string {
	return s.response
}
