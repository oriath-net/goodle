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
	"unsafe"
)

// Decompress behaves similarly to copy(), but passes the data through the
// OodleLZ decompressor.
//
// The size of the output buffer is significant to the decompressor.
func Decompress(in []byte, out []byte) (int, error) {
	r := int(C.OodleLZ_Decompress(
		unsafe.Pointer(&in[0]),
		C.size_t(len(in)),
		unsafe.Pointer(&out[0]),
		C.size_t(len(out)),
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

	if r < 0 {
		return 0, fmt.Errorf("OodleLZ_Decompress error %d", r)
	}

	if r != len(out) {
		return 0, fmt.Errorf("expected to decompress %d bytes but got %d", len(out), r)
	}

	return r, nil
}
