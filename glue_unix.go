package gooz

// #cgo LDFLAGS: -loo2core
//
// #include <stdlib.h>
// int OodleLZ_Decompress(
//		const void *compBuf,
//		size_t compBufSize,
//		void *rawBuf,
// 		size_t rawLen,
// 		int fuzzSafe,
// 		int checkCRC,
// 		int verbosity,
// 		void *decBufBase,
// 		size_t decBufSize,
// 		void *fpCallback,
// 		void *callbackUserData,
// 		void *decoderMemory,
// 		size_t decoderMemorySize,
// 		int threadPhase
// );
import "C"

import (
	"fmt"
)

// Decompress behaves similarly to copy(), but passes the data through the
// OodleLZ decompressor.
//
// The size of the output buffer is significant to the decompressor.
func Decompress(in []byte, out []byte) (int, error) {
	i_buf := C.CBytes(in)
	o_buf := C.malloc(C.size_t(len(out)))
	defer C.free(i_buf)
	defer C.free(o_buf)

	r := int(C.OodleLZ_Decompress(
		i_buf, C.size_t(len(in)),
		o_buf, C.size_t(len(out)),
		1,   // fuzzSafe = OodleLZ_FuzzSafe_Yes
		1,   // checkCRC = OodleLZ_CheckCRC_Yes
		0,   // verbosity = OodleLZ_Verbosity_None
		nil, // decBufBase
		0,   // decBufSize
		nil, // fpCallback
		nil, // callbackUserData
		nil, // decoderMemory
		0,   // decoderMemorySize
		3,   // threadPhase = OodleLZ_Decode_ThreadPhaseAll,
	))

	if r != len(out) {
		return 0, fmt.Errorf("expected to decompress %d bytes but only got %d", len(out), r)
	}

	return copy(out, C.GoBytes(o_buf, C.int(len(out)))), nil
}
