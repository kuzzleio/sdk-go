/* File : kcore.i */

%module(directors="1") kcore
%{
#include "listeners.hpp"
#include "exceptions.hpp"
#include "event_emitter.hpp"
#include "kuzzle.hpp"
#include "collection.hpp"
<<<<<<< HEAD
=======
#include "room.hpp"
#include "document.hpp"
#include "index.hpp"
>>>>>>> 1.x
#include "server.hpp"
#include <assert.h>
#include <ffi.h>
%}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "listeners.hpp"
%include "exceptions.hpp"
%include "event_emitter.hpp"
%include "kuzzle.hpp"
%include "collection.hpp"
<<<<<<< HEAD
=======
%include "room.hpp"
%include "document.hpp"
%include "index.hpp"
>>>>>>> 1.x
%include "server.hpp"
