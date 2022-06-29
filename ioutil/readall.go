package ioutil

import (
	"io"
	"io/ioutil"
)

// ReadAll is now just a shim for ``ioutil.ReadAll``.
// An interesting data race caused a bunch of examination that showed it
// unsafe with small buffers under high rates of churn, and fixing it
// made performance worse than the stock.
func ReadAll(r io.Reader) (b []byte, err error) {
	return ioutil.ReadAll(r)
}
