package process


import "errors"


var ErrAtomAlreadyExists = errors.New("atom already exists")

var ErrAtomNotFound = errors.New("atom is not found")