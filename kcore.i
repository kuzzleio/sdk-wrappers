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

%javaexception("java.lang.Exception") Kuzzle {
  $action
  if (!result) {
    jclass clazz = (*jenv)->FindClass(jenv, "java/lang/Exception");
    (*jenv)->ThrowNew(jenv, clazz, "Cannot connect");
    return $null;
  }
}

%define _Complex

%enddef
%include "libgokcore.h"
