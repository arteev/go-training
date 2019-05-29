#include <stdio.h>
#include "_cgo_export.h"

void call_go() {
    printf("in C \n");
    fflush(stdout);
    CallFromC();    
}