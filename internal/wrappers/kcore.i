/* File : kcore.i */

%module(directors="1") kcore
%{
#include "kuzzle.hpp"
%}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"
