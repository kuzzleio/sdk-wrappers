/* File : kcore.i */

%module kcore
%{
#define bool uchar_bool
#define true 1
#define false 0
typedef unsigned char bool;

#define _Complex
#include "libgokcore.h"
#include <stdio.h>
%}
%define _Complex

%enddef

%include "headers/kuzzle.h"
%include "libgokcore.h"
