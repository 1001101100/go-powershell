package powershell

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

/*

#cgo CFLAGS: -I.
#cgo LDFLAGS: ${SRCDIR}/../../native-powershell/native-powershell-bin/psh_host.dll


#include <stddef.h>
#include <string.h>
#include "powershell.h"

*/
import "C"

func makeString(str *C.wchar_t) string {
	ptr := unsafe.Pointer(str)
	count := C.wcslen(str) + 1
	arr := make([]uint16, count)
	ptrwchar := unsafe.Pointer(&arr[0])

	C.memcpy(ptrwchar, ptr, count*2)

	s := windows.UTF16ToString(arr)
	return s
}

func makeCString(str string) *C.wchar_t {
	cs, _ := windows.UTF16PtrFromString(str)
	ptrwchar := unsafe.Pointer(cs)
	return C.MallocCopy((*C.wchar_t)(ptrwchar))
}

//export logWchart
// commandWchart the C function pointer that dispatches to the Golang function for SimpleLogging
func logWchart(context uint64, str *C.wchar_t) {
	if context != 0 {
		s := makeString(str)
		contextInterface := getRunspaceContext(context)
		contextInterface.Log.Write(s)
	}
}

//export commandWchart
// commandWchart the C function pointer that dispatches to the Golang function for Send-HostCommand
func commandWchart(context uint64, cMessage *C.wchar_t, input *C.PowerShellObject, inputCount uint64, ret *C.JsonReturnValues) {

	var resultsWriter callbackResultsWriter
	if context != 0 {
		contextInterface := getRunspaceContext(context)
		inputArr := make([]Object, inputCount)
		for i := uint32(0); uint64(i) < inputCount; i++ {
			inputArr[i] = makePowerShellObjectIndexed(input, i)
		}
		message := makeString(cMessage)
		contextInterface.Callback.Callback(message, inputArr, &resultsWriter)
	}
	resultsWriter.filloutResults(ret)
}
