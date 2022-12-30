

# recyclablebuffer
`import "github.com/cognusion/go-utils/recyclablebuffer"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package recyclablebuffer provides the RecyclableBuffer, a multiuse buffer that very reusable,
supports re-reading the contained buffer, and when Close()d, will return home to its sync.Pool




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type RecyclableBuffer](#RecyclableBuffer)
  * [func NewRecyclableBuffer(home *sync.Pool, bytes []byte) *RecyclableBuffer](#NewRecyclableBuffer)
  * [func (r *RecyclableBuffer) Bytes() []byte](#RecyclableBuffer.Bytes)
  * [func (r *RecyclableBuffer) Close() error](#RecyclableBuffer.Close)
  * [func (r *RecyclableBuffer) Error() string](#RecyclableBuffer.Error)
  * [func (r *RecyclableBuffer) ResetFromLimitedReader(reader io.Reader, max int64) error](#RecyclableBuffer.ResetFromLimitedReader)
  * [func (r *RecyclableBuffer) ResetFromReader(reader io.Reader)](#RecyclableBuffer.ResetFromReader)
  * [func (r *RecyclableBuffer) String() string](#RecyclableBuffer.String)
  * [func (r *RecyclableBuffer) Write(p []byte) (n int, err error)](#RecyclableBuffer.Write)


#### <a name="pkg-files">Package files</a>
[recyclablebuffer.go](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go)



## <a name="pkg-variables">Variables</a>
``` go
var ErrTooLarge = errors.New("read byte count too large")
```
ErrTooLarge is returned when ResetFromLimitedReader is used and the supplied Reader writes too much




## <a name="RecyclableBuffer">type</a> [RecyclableBuffer](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=596:660#L17)
``` go
type RecyclableBuffer struct {
    bytes.Reader
    // contains filtered or unexported fields
}

```
RecyclableBuffer is an io.Reader, io.ReadCloser, io.Writer, and more that comes best from a sync.Pool.
Its Close method puts itself back in the Pool it came from







### <a name="NewRecyclableBuffer">func</a> [NewRecyclableBuffer](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=731:804#L24)
``` go
func NewRecyclableBuffer(home *sync.Pool, bytes []byte) *RecyclableBuffer
```
NewRecyclableBuffer returns a RecyclableBuffer with a proper home





### <a name="RecyclableBuffer.Bytes">func</a> (\*RecyclableBuffer) [Bytes](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=1976:2017#L59)
``` go
func (r *RecyclableBuffer) Bytes() []byte
```
Bytes returns the contents of the buffer, and sets the seek pointer back to the beginning




### <a name="RecyclableBuffer.Close">func</a> (\*RecyclableBuffer) [Close](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=1076:1116#L33)
``` go
func (r *RecyclableBuffer) Close() error
```
Close puts itself back in the Pool it came from. This should absolutely **never** be
called more than once per RecyclableBuffer life.
Implements “io.Closer“ (also “io.ReadCloser“ and “io.WriteCloser“)




### <a name="RecyclableBuffer.Error">func</a> (\*RecyclableBuffer) [Error](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=2415:2456#L73)
``` go
func (r *RecyclableBuffer) Error() string
```
Error returns the contents of the buffer as a string. Implements “error“




### <a name="RecyclableBuffer.ResetFromLimitedReader">func</a> (\*RecyclableBuffer) [ResetFromLimitedReader](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=1638:1722#L47)
``` go
func (r *RecyclableBuffer) ResetFromLimitedReader(reader io.Reader, max int64) error
```
ResetFromLimitedReader performs a Reset() using the contents of the supplied Reader as the new content,
up to at most max bytes, returning ErrTooLarge if it's over. The error is not terminal, and the buffer
may continue to be used, understanding the contents will be limited




### <a name="RecyclableBuffer.ResetFromReader">func</a> (\*RecyclableBuffer) [ResetFromReader](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=1248:1308#L39)
``` go
func (r *RecyclableBuffer) ResetFromReader(reader io.Reader)
```
ResetFromReader performs a Reset() using the contents of the supplied Reader as the new content




### <a name="RecyclableBuffer.String">func</a> (\*RecyclableBuffer) [String](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=2204:2246#L66)
``` go
func (r *RecyclableBuffer) String() string
```
String returns the contents of the buffer as a string, and sets the seek pointer back to the beginning




### <a name="RecyclableBuffer.Write">func</a> (\*RecyclableBuffer) [Write](https://github.com/cognusion/go-utils/tree/master/recyclablebuffer/recyclablebuffer.go?s=2560:2621#L78)
``` go
func (r *RecyclableBuffer) Write(p []byte) (n int, err error)
```
Writer adds the bytes the written to the buffer. Implements “io.Writer“








- - -
Generated by [godoc2md](http://godoc.org/github.com/cognusion/godoc2md)
