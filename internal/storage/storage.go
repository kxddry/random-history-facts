package storage

import "errors"

var (
    ErrFactAlreadyExists = errors.New("such fact already exists")
    ErrNoFacts           = errors.New("no facts found")
)
