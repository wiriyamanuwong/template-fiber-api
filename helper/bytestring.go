// Package helper is a collection of useful functions
// from https://github.com/tsingson/fasthttp-guide
package helper

import (
	"encoding/binary"
	"os"
	"strings"
	"unsafe"
)

const toLower = 'a' - 'A'

var toLowerTable = func() [256]byte {
	var a [256]byte
	for i := 0; i < 256; i++ {
		c := byte(i)
		if c >= 'A' && c <= 'Z' {
			c += toLower
		}
		a[i] = c
	}
	return a
}()

// LowercaseBytes low
func LowercaseBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		p := &b[i]
		*p = toLowerTable[*p]
	}
}

// B2S converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
// update https://blog.devops.dev/fast-string-to-byte-and-byte-to-string-conversion-in-go-1-20-85bdb859ee67
//
// Note it may break if string and/or slice header will change
// in the future go versions.
// nolint
func B2S(bs []byte) string {
	if len(bs) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

// S2B converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
// nolint
func S2B(s string) []byte {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// StrBuilder strings builder
func StrBuilder(args ...string) string {
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}

// GetFileSize get file size
func GetFileSize(fullFilename string) int64 {
	fileInfo, err := os.Stat(fullFilename)
	if err != nil {
		return 0
	}
	fileSize := fileInfo.Size() // size
	return fileSize
}

// Int64ToBytes int64 to byte, use for timestamp
// nolint
func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 byte to int64
// nolint
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
