#include <stdio.h>
#include <stdlib.h>
#include "inc/fmod.h"
#include "entry.h"

int testGoC() {
    printf("Test Go-C connection.\n");
    return 0;
}

char* testGoCString() {
    int length = 8;
    printf("%d\n", length);
    char* test = (char*)malloc(length * sizeof(char));
    sprintf(test, "%s", "testStr");
    return test;
}

int testFmod() {
    return 0;
}
