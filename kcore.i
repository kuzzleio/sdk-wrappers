/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "libgokcore.h"
#include <stdio.h>
%}

%pragma(java) jniclasscode=%{
  static {
    try {
        System.loadLibrary("kcore");
    } catch (UnsatisfiedLinkError e) {
      System.err.println("Native code library failed to load. \n" + e);
      System.exit(1);
    }
  }
%}

%define _Complex

%enddef
%include "libgokcore.h"
%include "templates/kuzzle.i"
%include "headers/kuzzle.h"
