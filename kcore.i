/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "libgokcore.h"
#include "kuzzle.h"
#include <json-c/json.h>
#include <stdio.h>
%}
%define _Complex

%enddef

%import "json-c/json_object_private.h"
%import "json-c/json.h"
%include "headers/kuzzle.h"
%include "libgokcore.h"
