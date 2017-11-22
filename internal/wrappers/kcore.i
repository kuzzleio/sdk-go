/* File : kcore.i */

%module kcore
%{
#define _Complex
#include "kuzzle.h"
#include "headers/kuzzlesdk.h"
#include "templates/swig.h"

#include <stdio.h>

#include <assert.h>

// 1:
struct callback_data {
  JNIEnv *env;
  jobject obj;
};

// 2:
void java_callback(int arg, void *ptr) {
  struct callback_data *data = ptr;
  const jclass callbackInterfaceClass = (*data->env)->FindClass(data->env, "Callback");
  assert(callbackInterfaceClass);
  const jmethodID meth = (*data->env)->GetMethodID(data->env, callbackInterfaceClass, "handle", "(I)V");
  assert(meth);
  (*data->env)->CallVoidMethod(data->env, data->obj, meth, (jint)arg);
}

%}
%define _Complex
%enddef

// 3:
%typemap(jstype) callback_t cb "Callback";
%typemap(jtype) callback_t cb "Callback";
%typemap(jni) callback_t cb "jobject";
%typemap(javain) callback_t cb "$javainput";
// 4:
%typemap(in,numinputs=1) (callback_t cb, void *userdata) {
  struct callback_data *data = malloc(sizeof *data);
  data->env = jenv;
  data->obj = JCALL1(NewGlobalRef, jenv, $input);
  JCALL1(DeleteLocalRef, jenv, $input);
  $1 = java_callback;
  $2 = data;
}

%include "headers/kuzzlesdk.h"
%include "kuzzle.h"
