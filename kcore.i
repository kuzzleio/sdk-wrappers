/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "libgokcore.h"
#include <stdio.h>
%}

%define _Complex

%enddef
%include "headers/kuzzle.h"
%include "libgokcore.h"
%include "templates/kuzzle.i"
%include "templates/java/core.i"
