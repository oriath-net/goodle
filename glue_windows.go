package gooz

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

var oodleDLL *windows.DLL

func init() {
	var dllNames []string

	envDll := os.Getenv("OODLE_DLL")
	if envDll != "" {
		dllNames = []string{envDll}
	} else {
		dllNames = []string{
			"oo2core_9_win64.dll",
			"oo2core_8_win64.dll",
			"oo2core_7_win64.dll",
			"oo2core_6_win64.dll",
			"oo2core_5_win64.dll",
			"oo2core_4_win64.dll",
			"oo2core_3_win64.dll",
		}
	}

	for _, name := range dllNames {
		dh, err := windows.LoadDLL(name)
		if err == nil {
			oodleDLL = dh
			return
		}
	}

	log.Fatal(
		"Unable to load Oodle DLL. Place oo2core_#_win64.dll alongside pogo.exe\n" +
			"or set OODLE_DLL to a path to the library.\n",
	)
}

// Decompress behaves similarly to copy(), but passes the data through the
// OodleLZ decompressor.
//
// The size of the output buffer is significant to the decompressor.
func Decompress(in []byte, out []byte) (int, error) {
	fn, err := oodleDLL.FindProc("OodleLZ_Decompress")
	if err != nil {
		panic(err)
	}

	ur, _, err := fn.Call(
		uintptr(unsafe.Pointer(&in[0])),
		uintptr(len(in)),
		uintptr(unsafe.Pointer(&out[0])),
		uintptr(len(out)),
		1, // fuzzSafe = OodleLZ_FuzzSafe_Yes
		1, // checkCRC = OodleLZ_CheckCRC_Yes
		0, // verbosity = OodleLZ_Verbosity_None
		0, // decBufBase
		0, // decBufSize
		0, // fpCallback
		0, // callbackUserData
		0, // decoderMemory
		0, // decoderMemorySize
		3, // threadPhase = OodleLZ_Decode_ThreadPhaseAll,
	)
	r := int(ur)

	if r < 0 {
		return 0, fmt.Errorf("OodleLZ_Decompress error %d", r)
	}

	if r != len(out) {
		return 0, fmt.Errorf("expected to decompress %d bytes but got %d", len(out), r)
	}

	return int(r), nil
}
