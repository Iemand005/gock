package main

import "fmt"

/*
#include <stdio.h>

int hello() {
    printf("hello from C\n");
    return 42;
}
*/
import "C"

func hi() bool {
	return false
}

func main() {
	fmt.Print("hey!")
}
