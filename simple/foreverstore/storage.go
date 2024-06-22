package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathKey struct {
	Pathname string
	Original string
}

func (p PathKey) Filename() string {
	return fmt.Sprintf("%s%s", p.Pathname, p.Original)
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashstr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashstr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashstr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Original: hashstr,
	}
}

type PathTransformFunc func(string) PathKey

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	pathAndFileName := pathKey.Filename()

	f, err := os.OpenFile(pathAndFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("writen (%d) bytes to disk : %s", n, pathAndFileName)

	return nil
}
