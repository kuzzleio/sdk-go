%module kcore

%define _Complex
%enddef

%include "stl.i"
%include "exception.i"
%include "kuzzle.h"
%include "kuzzlesdk.h"
%include "core.cpp"

%ignore kuzzle;
%rename(QueryOptions) query_options;

%{
#define SWIG_FILE_WITH_INIT
#include "./templates/python/core.cpp"
using namespace BEV;
%}
%inline %{
  typedef struct {
    kuzzle* _kuzzle;
  } Kuzzle;
%}

%runtime %{
  PyObject *_bev_exception;
%}

%extend Kuzzle {
    Kuzzle(char* host, options *opts) {
        kuzzle *k = (kuzzle*)malloc(sizeof(kuzzle));
        kuzzle_new_kuzzle(k, host, (char*)"websocket", opts);

        Kuzzle *K = (Kuzzle*)calloc(1, sizeof(Kuzzle));
        K->_kuzzle = k;
        return K;
    }

    Kuzzle(char* host) {
        kuzzle *k = (kuzzle*)malloc(sizeof(kuzzle));
        kuzzle_new_kuzzle(k, host, (char*)"websocket", NULL);

        Kuzzle *K = (Kuzzle*)calloc(1, sizeof(Kuzzle));
        K->_kuzzle = k;
        return K;
    }

    ~Kuzzle() {
        unregisterKuzzle($self->_kuzzle);
        free($self->_kuzzle);
        free($self);
    }

    long long now(query_options* options=NULL) {
      int_result *result = kuzzle_now($self->_kuzzle, options);

      PyObject *o = SWIG_NewPointerObj(new BEV::MyException, SWIGTYPE_p_BEV__MyException, SWIG_POINTER_OWN);
      SWIG_Python_Raise(o, (char *)"MyException", SWIGTYPE_p_BEV__MyException);
      PyErr_SetString(PyExc_ValueError, "test");
      return -1;

      if (result->error != NULL) {
        PyErr_SetString(PyExc_ValueError, result->error);
        return -1;
      }

      return result->result;
    }
}


