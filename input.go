package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func init() {
	UnsetIFlag(syscall.IXON)
	UnsetLFlag(syscall.IEXTEN)
	UnsetLFlag(syscall.ECHO)
	UnsetLFlag(syscall.ISIG)
	UnsetLFlag(syscall.ICANON)
	UnsetIFlag(syscall.ICRNL)
	UnsetIFlag(syscall.IXON)
}

func nextByte(stdin *bufio.Reader) byte {
	b, err := stdin.ReadByte()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading %v\n", err)
		return 0
	}

	return b
}
func TermIOs() syscall.Termios {
	var termios syscall.Termios
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&termios)),
	)
	return termios
}
func SetIFlag(flag uint32) {
	termios := TermIOs()
	termios.Iflag |= flag
	SetTermIOs(termios)
}
func SetTermIOs(termios syscall.Termios) {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&termios)),
	)
}
func UnsetLFlag(flag uint32) {
	termios := TermIOs()
	termios.Lflag &^= flag
	SetTermIOs(termios)
}
func UnsetIFlag(flag uint32) {
	termios := TermIOs()
	termios.Iflag &^= flag
	SetTermIOs(termios)
}
