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

%include "templates/swig.h"
%include "headers/kuzzle.h"
%include "libgokcore.h"
