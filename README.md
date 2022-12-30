# go-utils
Go utility functions

* ioutil.ReadAll(r io.Reader) (b []byte, err error) // ~~Zero allocs and 7x better performance than stock~~ Use the stdlib
* recyclablebuffer.RecyclableBuffer // reusable io.Reader/Writer
* writeatbuffer.Buffer // Efficient io.WriterAt by AWS
