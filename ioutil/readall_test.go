package ioutil

import (
	"sync"

	. "github.com/smartystreets/goconvey/convey"

	"bytes"
	"io/ioutil"
	"testing"
)

type TestReader struct {
	stuff *bytes.Buffer
}

func (tr *TestReader) Read(p []byte) (n int, err error) {
	n, err = tr.stuff.Read(p)
	return
}

func (tr *TestReader) Write(p []byte) (n int, err error) {
	n, err = tr.stuff.Write(p)
	return
}

func Test_ReadAllSimple(t *testing.T) {
	hw := []byte("Hello World")

	Convey("When using ReadAll on an io.Reader with a known value, the bytes are consistent", t, func() {
		buff := bytes.NewBuffer(hw)
		val, err := ReadAll(buff)
		So(err, ShouldBeNil)
		So(val, ShouldResemble, hw)

	})
}

func Test_ReadAllGoros(t *testing.T) {
	Convey("When using ReadAll on io.Readers from many goros with known values, the bytes are consistent", t, func(c C) {

		var wg sync.WaitGroup

		for i := 0; i < 1500; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				hw := []byte("This is not the !!! Typical Hello World! WEEEEEEEEEE How fun! This is not the !!! Typical Hello World! WEEEEEEEEEE How fun!")
				hwbuf := bytes.NewReader(hw)

				val, err := ReadAll(hwbuf)
				c.So(err, ShouldBeNil)
				c.So(val, ShouldResemble, hw)
			}()
		}

		wg.Wait()
	})
}

func BenchmarkIoUtilReadAll(b *testing.B) {
	var (
		val  []byte
		err  error
		hw   = []byte("Hello World")
		shw  = string(hw)
		buff = bytes.NewReader(hw)
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err = ioutil.ReadAll(buff)
		if err != nil {
			panic(err)
		} else if string(val) != shw {
			panic("WTF " + string(val))
		}
		buff.Seek(0, 0)
	}
}

func BenchmarkReadAll(b *testing.B) {
	var (
		val  []byte
		err  error
		hw   = []byte("Hello World")
		shw  = string(hw)
		buff = bytes.NewReader(hw)
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, err = ReadAll(buff)
		if err != nil {
			panic(err)
		} else if string(val) != shw {
			panic("WTF " + string(val))
		}
		buff.Seek(0, 0)
	}
}
