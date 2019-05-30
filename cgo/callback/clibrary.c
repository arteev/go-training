#include <stdio.h>
#include "clibrary.h"

void c_func(callback_func callback)
{
	int arg = 1000;
	printf("C.c_func(): arg = %d\n", arg);
	int result = callback(arg);
	printf("C.c_func(): %d\n", result);
	fflush(stdout);
}