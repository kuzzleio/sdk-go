/* File : kcore.i */

%module(directors="1") kcore
%{
#include "exceptions.hpp"
#include "kuzzle.hpp"
#include "collection.hpp"
#include "document.hpp"
#include <assert.h>
#include <ffi.h>
%}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "exceptions.hpp"
%include "kuzzle.hpp"
%include "collection.hpp"
%include "document.hpp"