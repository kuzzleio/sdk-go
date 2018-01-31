/* File : kcore.i */

%module(directors="1") kcore
%{
#include "kuzzle.hpp"
#include <assert.h>
#include <ffi.h>
%}


%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"