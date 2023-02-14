package muserrs

import "errors"

// ErrSmallBuf means that an Unmarshal requires a longer buffer than was
// provided.
var ErrSmallBuf = errors.New("buf is too small")

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow = errors.New("overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength = errors.New("negative length")

// ErrWrongByte happens on Unmarshal when unexpected byte was caught.
var ErrWrongByte = errors.New("wrong byte")
