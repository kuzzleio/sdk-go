/* File : kcore.i */

%module(directors="1") kcore
%{
#include "kuzzle.hpp"
#include "callback.hpp"
%}

%define _Complex
%enddef


%feature("director") Callback;
%feature("director") CallbackWrapper;

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"
%include "callback.hpp"