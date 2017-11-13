/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "kuzzle.h"
#include "headers/kuzzlesdk.h"
#include "templates/swig.h"

#include <stdio.h>
%}
%define _Complex

%enddef

%include "headers/kuzzlesdk.h"
%include "kuzzle.h"
