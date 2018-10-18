package main

/*
#include <stdio.h>

void hi() {
	printf("aaa\n");
}
*/
import "C"

func main() {
	C.hi()
}