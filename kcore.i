/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "libgokcore.h"
#include "kuzzle.h"
#include "templates/swig.h"

#include <stdio.h>
%}
%define _Complex

%enddef

%include "headers/structs.h"
%include "libgokcore.h"
