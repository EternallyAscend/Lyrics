#include <stdio.h>
#include <stdlib.h>
#include "inc/fmod.h"
#include "entry.h"

char* testGoCString() {
    int length = 16;
    char* test = (char*)malloc(length * sizeof(char));
    sprintf(test, "%s", "CGO: Connected.");
    return test;
}
