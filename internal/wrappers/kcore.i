/* File : kcore.i */

%module kcore
%{
#include "kuzzle.hpp"

struct callback_data {
  JNIEnv *env;
  jobject obj;
};

void java_callback(int arg, struct callback_data *ptr) {
  struct callback_data *data = ptr;
  const jclass callbackInterfaceClass = (*data->env).FindClass("Callback");
  const jmethodID meth = (*data->env).GetMethodID(callbackInterfaceClass, "run", "(I)V");
  (*data->env).CallVoidMethod(data->obj, meth, (json_object*)arg);
}

%}

// 3:
%typemap(jstype) kuzzle_event_listener cb "Callback";
%typemap(jtype) kuzzle_event_listener cb "Callback";
%typemap(jni) kuzzle_event_listener cb "jobject";
%typemap(javain) kuzzle_event_listener cb "$javainput";
 // 4:
%typemap(in) (kuzzle_event_listener cb, void *userdata) {
  struct callback_data *data = malloc(sizeof *data);
  data->env = jenv;
  data->obj = JCALL1(NewGlobalRef, jenv, $input);
  JCALL1(DeleteLocalRef, jenv, $input);
  $1 = java_callback;
  $2 = data;
}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"
%include "callback.hpp"