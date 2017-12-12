/* File : kcore.i */

%module(directors="1") kcore
%{
#include "kuzzle.hpp"
#include <assert.h>
#include <ffi.h>
%}

%feature("director") Callback;

%javamethodmodifiers Callback::call "public abstract"
%typemap(javaout) void Callback::call ";"
%typemap(javaclassmodifiers) Callback "public abstract class"
%typemap(jstype) Callback_t "Callback";
%typemap(jtype) Callback_t "Callback";
%typemap(jni) Callback_t "jobject";
%typemap(javain) Callback_t "$javainput";
%typemap(in) Callback_t {
  $1 = (Callback_t)$input;
}

%inline %{
struct Callback {
public:
  virtual void call(const int) = 0;
  virtual ~Callback() {
    if (bound_fn) ffi_closure_free(closure);
  }

  jlong prepare_fp() {
//    if (!bound_fn) {
//      int ret;
//      args[0] = &ffi_type_uint;
//      args[1] = &ffi_type_pointer;
//      closure = static_cast<decltype(closure)>(ffi_closure_alloc(sizeof(ffi_closure), &bound_fn));
//      assert(closure);
//      ret = ffi_prep_cif(&cif, FFI_DEFAULT_ABI, 2, &ffi_type_void, args);
//      assert(ret == FFI_OK);
//      ret = ffi_prep_closure_loc(closure, &cif, java_callback, this, bound_fn);
//      assert(ret == FFI_OK);
//    }
//    return *((jlong*)&bound_fn);
    return (jlong)0;
  }
private:
  ffi_closure *closure;
  ffi_cif cif;
  ffi_type *args[1];
  void *bound_fn;

  static void java_callback(ffi_cif *cif, void *ret, void *args[], void *userdata) {
//    (void)cif;
//    (void)ret;
//    Callback *cb = static_cast<Callback*>(userdata);
//    Log l = (Log)*(unsigned*)args[0];
//    const char *s = *(const char**)args[1];
//    cb->call(l, s);
  }
};
%}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"