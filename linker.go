package yocto

// #cgo LDFLAGS: -L${SRCDIR}/so/ -Wl,-rpath=${SRCDIR}/so/ -lyapi -lyocto
import "C"