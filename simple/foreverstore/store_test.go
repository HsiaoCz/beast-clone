package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	fmt.Println(pathKey.Pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg data"))

	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Fatal(err)
	}
}
