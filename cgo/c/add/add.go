package add

/*
#include <stdio.h>

#include "add.h"
*/
import "C"

func Add(a, b int) int {
	c := C.add(C.int(a), C.int(b))
	return int(c)
}
