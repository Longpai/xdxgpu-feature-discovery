// WARNING: This file has automatically been generated on Tue, 07 May 2024 14:39:20 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package xdxml

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "xdxml.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// MAX_DEVICES as defined in xdxml/xdxml.h:11
	MAX_DEVICES = 2
	// PARAMETER_PAGE_OFFSET as defined in xdxml/xdxml.h:13
	PARAMETER_PAGE_OFFSET = 0x58000
	// FW_GLOABL_STRUCT_OFFSET as defined in xdxml/xdxml.h:14
	FW_GLOABL_STRUCT_OFFSET = 0x59000
	// CMCU_VARS_OFFSET as defined in xdxml/xdxml.h:15
	CMCU_VARS_OFFSET = 0x6FF000
	// DEVICE_NAME_BUFFER_SIZE as defined in xdxml/xdxml.h:20
	DEVICE_NAME_BUFFER_SIZE = 64
	// DEVICE_PCI_BUS_ID_BUFFER_SIZE as defined in xdxml/xdxml.h:25
	DEVICE_PCI_BUS_ID_BUFFER_SIZE = 32
	// XDXGPU_INFO_UTIL_INFO as defined in xdxml/xdxml.h:32
	XDXGPU_INFO_UTIL_INFO = 0x02
	// RGXFWIF_DM_MAX as defined in xdxml/xdxml.h:33
	RGXFWIF_DM_MAX = (uint32(9))
	// RGX_NUM_OS_SUPPORTED as defined in xdxml/xdxml.h:34
	RGX_NUM_OS_SUPPORTED = 1
	// RET_SIZE as defined in xdxml/xdxml.h:41
	RET_SIZE = 1000000
	// MAX_PATH_LEN as defined in xdxml/xdxml.h:42
	MAX_PATH_LEN = 100
	// PROCESS_NAME_LEN as defined in xdxml/xdxml.h:43
	PROCESS_NAME_LEN = 30
)

// Return as declared in xdxml/xdxml.h:76
type Return int32

// Return enumeration from xdxml/xdxml.h:76
const (
	ERROR                         Return = -1
	SUCCESS                       Return = 0
	ERROR_UNINITIALIZED           Return = 1
	ERROR_INVALID_ARGUMENT        Return = 2
	ERROR_NOT_SUPPORTED           Return = 3
	ERROR_NO_PERMISSION           Return = 4
	ERROR_ALREADY_INITIALIZED     Return = 5
	ERROR_NOT_FOUND               Return = 6
	ERROR_INSUFFICIENT_SIZE       Return = 7
	ERROR_INSUFFICIENT_POWER      Return = 8
	ERROR_DRIVER_NOT_LOADED       Return = 9
	ERROR_TIMEOUT                 Return = 10
	ERROR_IRQ_ISSUE               Return = 11
	ERROR_LIBRARY_NOT_FOUND       Return = 12
	ERROR_FUNCTION_NOT_FOUND      Return = 13
	ERROR_CORRUPTED_INFOROM       Return = 14
	ERROR_GPU_IS_LOST             Return = 15
	ERROR_RESET_REQUIRED          Return = 16
	ERROR_OPERATING_SYSTEM        Return = 17
	ERROR_LIB_RM_VERSION_MISMATCH Return = 18
	ERROR_IN_USE                  Return = 19
	ERROR_MEMORY                  Return = 20
	ERROR_NO_DATA                 Return = 21
	ERROR_VGPU_ECC_NOT_SUPPORTED  Return = 22
	ERROR_UNKNOWN                 Return = 999
)
