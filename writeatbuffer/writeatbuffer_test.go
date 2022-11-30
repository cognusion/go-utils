package writeatbuffer

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func Test_WriteAtSimple(t *testing.T) {
	hw := []byte("Hello World")
	m := []byte("Matt!")

	Convey("When using WriteAtBuffer with known values, at known offsets, the bytes are consistent", t, func() {
		buff := NewBuffer(make([]byte, len(hw)))
		n, err := buff.WriteAt(hw, 0)
		So(err, ShouldBeNil)
		So(n, ShouldEqual, len(hw))
		So(buff.Bytes(), ShouldResemble, hw)
		So(len(buff.Bytes()), ShouldEqual, len(hw))

		n, err = buff.WriteAt(hw, int64(len(hw)))
		So(err, ShouldBeNil)
		So(n, ShouldEqual, len(hw))
		So(buff.Bytes(), ShouldResemble, append(hw, hw...))
		So(len(buff.Bytes()), ShouldEqual, len(append(hw, hw...)))

		n, err = buff.WriteAt(m, 6)
		So(err, ShouldBeNil)
		So(n, ShouldEqual, len(m))
		So(len(buff.Bytes()), ShouldEqual, len(append(hw, hw...)))
		So(buff.Bytes(), ShouldResemble, append([]byte("Hello Matt!"), hw...))

	})
}
