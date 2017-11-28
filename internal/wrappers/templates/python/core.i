%{
#define SWIG_FILE_WITH_INIT
#include <iostream>
#include <sstream>

#include "kuzzle.hpp"
%}

%exception kuzzleio::Kuzzle::now {
  try {
    $action
  } catch (...) {
    std::cout << "BBlah";
    PyObject *o = SWIG_NewPointerObj(new kuzzleio::ForbiddenError, SWIGTYPE_p_kuzzleio__ForbiddenError, SWIG_POINTER_OWN);
    SWIG_Python_Raise(o, (char *)"MyException", SWIGTYPE_p_kuzzleio__ForbiddenError);
    SWIG_fail;
  }
}

%include "stl.i"
%include "kcore.i"

%{
#include "core.cpp"
%}

namespace kuzzleio {
  %exceptionclass KuzzleError;
  %extend KuzzleError {
    std::string __str__() const {
      std::ostringstream s;
      s << "[" << $self->status << "] " << $self->what();
      if (!$self->stack.empty()) {
        s << "\n" << $self->stack;
      }
      return s.str();
    }
  }
  %exceptionclass BadRequestError;
  %exceptionclass ForbiddenError;

  %typemap(throws) BadRequestError "(void)$1; throw;";
}

/*
%exception kuzzleio::Kuzzle::now {
  try {
    $action
  } catch (...) {
    std::cout << "Blah";
    PyObject *o = SWIG_NewPointerObj(new kuzzleio::ForbiddenError, SWIGTYPE_p_kuzzleio__ForbiddenError, SWIG_POINTER_OWN);
    SWIG_Python_Raise(o, (char *)"MyException", SWIGTYPE_p_kuzzleio__ForbiddenError);
    SWIG_fail;
  }
}
*/

%include "core.cpp"

//%ignore kuzzle;
%rename(QueryOptions) query_options;


