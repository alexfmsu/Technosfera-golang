package main

import "unsafe"

/*
   #include <stdlib.h>
   #include "slre/slre.c"

   int test_regexp(){
        return slre_match("[a-h]+", "abcdefghxxx", 11, NULL, 0, 0);
   }
*/
import "C"

func test_regexp() int {
	return regexp("[a-h]+", "abcdefghxxx", 0, false)
}

const SLRE_NO_MATCH int = -1
const SLRE_UNEXPECTED_QUANTIFIER int = -2
const SLRE_UNBALANCED_BRACKETS int = -3
const SLRE_INTERNAL_ERROR int = -4
const SLRE_INVALID_CHARACTER_SET int = -5
const SLRE_INVALID_METACHARACTER int = -6
const SLRE_CAPS_ARRAY_TOO_SMALL int = -7
const SLRE_TOO_MANY_BRANCHES int = -8
const SLRE_TOO_MANY_BRACKETS int = -9

type regexp_params struct {
	pattern     string
	expr        string
	num_caps    int
	ignore_case bool
}

func regexp(pattern string, expr string, num_caps int, ignore_case bool) int {
	Cpattern := C.CString(pattern)
	Cexpr := C.CString(expr)
	Cnum_caps := C.int(num_caps)
	Cignore_case := C.int(0)

	Clen := C.int(len(expr))

	if ignore_case {
		Cignore_case = C.int(1)
	}

	var caps []C.struct_slre_cap

	var ret C.int

	if num_caps > 0 {
		caps = make([]C.struct_slre_cap, C.int(num_caps))

		Ccaps := (*C.struct_slre_cap)(unsafe.Pointer(&caps[0]))

		ret = C.slre_match(Cpattern, Cexpr, Clen, Ccaps, Cnum_caps, Cignore_case)
	} else {
		ret = C.slre_match(Cpattern, Cexpr, Clen, nil, Cnum_caps, Cignore_case)
	}

	C.free(unsafe.Pointer(Cpattern))
	C.free(unsafe.Pointer(Cexpr))

	return int(ret)
}

func CallCGo(n int) {
	for i := 0; i < n; i++ {
		C.test_regexp()
	}
}

func CallGo(n int) {
	for i := 0; i < n; i++ {
		test_regexp()
	}
}

func main() {}
