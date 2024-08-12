package main

/*
#cgo LDFLAGS: -lntdll

#include <windows.h>
#include <stdio.h>
#include <stdint.h>

typedef struct {
    int Ssn;
    const char *Name;
    PVOID Address;
} SystemCall;

typedef struct {
    DWORD_PTR ImageBase;
    PIMAGE_DOS_HEADER DosHeader;
    PIMAGE_NT_HEADERS NtHeaders;
    PIMAGE_EXPORT_DIRECTORY ExportDirectory;
    PIMAGE_RUNTIME_FUNCTION_ENTRY ExceptionDirectory;
} PeFile;

PeFile ParsePEFile(void* peBase) {
    PeFile peFile;
    peFile.ImageBase = (DWORD_PTR)peBase;
    peFile.DosHeader = (PIMAGE_DOS_HEADER)peFile.ImageBase;
    peFile.NtHeaders = (PIMAGE_NT_HEADERS)(peFile.ImageBase + peFile.DosHeader->e_lfanew);
    peFile.ExportDirectory = (PIMAGE_EXPORT_DIRECTORY)(peFile.ImageBase +
                             peFile.NtHeaders->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress);
    peFile.ExceptionDirectory = (PIMAGE_RUNTIME_FUNCTION_ENTRY)(peFile.ImageBase +
                                   peFile.NtHeaders->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_EXCEPTION].VirtualAddress);
    return peFile;
}

// Extract the System Service Number (SSN) from a function address
int GetSSN(void* fnAddress) {
    unsigned char* fnAddr = (unsigned char*)fnAddress;
    int i = 0;

    while (1) {
        if (fnAddr[i] == 0x4C && fnAddr[i + 1] == 0x8B && fnAddr[i + 2] == 0xD1 && fnAddr[i + 3] == 0xB8) {
            unsigned char high = fnAddr[i + 5];
            unsigned char low = fnAddr[i + 4];
            return (high << 8) | low;
        }
        if (fnAddr[i] == 0xC3 || (fnAddr[i] == 0x0f && fnAddr[i + 1] == 0x05)) {
            break;
        }
        i++;
    }
    return -1; // Indicates no SSN found
}

// Retrieve the System Call Table (SSDT)
SystemCall* GetSSDT(PeFile peFile, int* outCount) {
    DWORD* funcNames = (DWORD*)(peFile.ImageBase + peFile.ExportDirectory->AddressOfNames);
    WORD* funcOrds = (WORD*)(peFile.ImageBase + peFile.ExportDirectory->AddressOfNameOrdinals);
    DWORD* funcAddrs = (DWORD*)(peFile.ImageBase + peFile.ExportDirectory->AddressOfFunctions);

    SystemCall* sysCalls = NULL;
    *outCount = 0;

    for (DWORD i = 0; i < peFile.ExportDirectory->NumberOfFunctions; i++) {
        const char* fnName = (const char*)(peFile.ImageBase + funcNames[i]);
        WORD fnOrdinal = funcOrds[i];
        void* fnAddr = (void*)(peFile.ImageBase + funcAddrs[fnOrdinal]);

        if (strncmp(fnName, "Zw", 2) == 0) {
            int ssn = GetSSN(fnAddr);
            if (ssn != -1) {
                sysCalls = (SystemCall*)realloc(sysCalls, sizeof(SystemCall) * (*outCount + 1));
                sysCalls[*outCount].Ssn = ssn;
                sysCalls[*outCount].Name = fnName;
                sysCalls[*outCount].Address = fnAddr;
                (*outCount)++;
            }
        }
    }

    return sysCalls;
}

*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    // Load NTDLL and parse it using the C function
    peBase := C.LoadLibraryA(C.CString("NTDLL"))
    peFile := C.ParsePEFile(unsafe.Pointer(peBase))

    var sysCallCount C.int
    sysCalls := C.GetSSDT(peFile, &sysCallCount)

    // Iterate over the system calls and print them
    for i := 0; i < int(sysCallCount); i++ {
        sysCall := (*C.SystemCall)(unsafe.Pointer(uintptr(unsafe.Pointer(sysCalls)) + uintptr(i)*unsafe.Sizeof(*sysCalls)))
        fmt.Printf("-----------------\n")
        fmt.Printf("Name: %s\n", C.GoString(sysCall.Name))
        fmt.Printf("SSN: %d (0x%04X)\n", int(sysCall.Ssn), int(sysCall.Ssn))
        fmt.Printf("Address: %p\n", sysCall.Address)
    }

    // Free the allocated memory for sysCalls (not necessary in Go, but for good practice)
    C.free(unsafe.Pointer(sysCalls))
}
