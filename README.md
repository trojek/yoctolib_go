# Yocto GO library

This is Yocto Go library. It contains of C++ library wrapper generated with SWIG.
Actual communication with Yocto module is implemented in C++. In order to run,
this library requires to be linked with libyocto.so and libyapi.so.

## linker.go

Contains linker flags required to use yocto shared lib.

## yocto.go

Contains generated wrapper for C++ lib.

## yoctolib_go_wrap.cxx

Contains C++ wrapper providing API callable by CGO.

## headers/

Contains header files required to compile C++ wrapper. These files were copied from yoctolib_cpp.

## so/

Contains prebuilt shared lib with yoctolib_cpp implementation.
