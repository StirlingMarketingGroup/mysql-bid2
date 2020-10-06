package main

// #include <string.h>
// #include <stdbool.h>
// #include <mysql.h>
// #cgo CFLAGS: -O3 -I/usr/include/mysql -fno-omit-frame-pointer
import "C"
import (
	"crypto/rand"
	"time"
	"unsafe"
)

// main function is needed even for generating shared object files
func main() {}

func msg(message *C.char, s string) {
	m := C.CString(s)
	defer C.free(unsafe.Pointer(m))

	C.strcpy(message, m)
}

//export bid2_init
func bid2_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	initid.maybe_null = C.bool(false)

	return C.bool(false)
}

//export bid2
func bid2(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	t := time.Now().UnixNano()
	b := make([]byte, 8)

	b[0] = byte((t >> 070) & 0xff)
	b[1] = byte((t >> 060) & 0xff)
	b[2] = byte((t >> 050) & 0xff)
	b[3] = byte((t >> 040) & 0xff)
	b[4] = byte((t >> 030) & 0xff)
	b[5] = byte((t >> 020) & 0xff)
	b[6] = byte((t >> 010) & 0xff)

	rand.Read(b[7:])

	*length = uint64(len(b))
	*isNull = 0
	return C.CString(string(b))
}
