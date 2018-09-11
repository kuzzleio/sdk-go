/* File : kcore.i */

%module(directors="1") kuzzlesdk
%{
#include "exceptions.hpp"
#include "event_emitter.hpp"
#include "kuzzle.hpp"
#include "collection.hpp"
#include "index.hpp"
#include "server.hpp"
#include "document.hpp"
#include "realtime.hpp"
#include "auth.hpp"
#include <assert.h>
%}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "exceptions.hpp"
%include "event_emitter.hpp"
%include "kuzzle.hpp"
%include "collection.hpp"
%include "index.hpp"
%include "server.hpp"
%include "document.hpp"
%include "realtime.hpp"
%include "auth.hpp"
