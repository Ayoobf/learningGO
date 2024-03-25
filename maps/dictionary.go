package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word is not in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
