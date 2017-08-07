/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "libgokcore.h"
#include <stdio.h>
%}

%extend kuzzle {
    char* connect() {
        kuzzle_wrapper_connect();
    }
}

%define _Complex

%enddef
%include "libgokcore.h"
