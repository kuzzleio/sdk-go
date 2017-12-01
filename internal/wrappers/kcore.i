/* File : kcore.i */

%module kcore
%{
#include "kuzzle.hpp"

struct callback_data {
  JNIEnv *env;
  jobject obj;
};

void java_callback(json_object* arg, void *ptr) {
  printf("one");
  struct callback_data *data = (callback_data*)ptr;
  printf("two");
  const jclass callbackInterfaceClass = (*data->env).FindClass("Callback");
  printf("three");
  const jmethodID meth = (*data->env).GetMethodID(callbackInterfaceClass, "handle", "(I)V");
  printf("four");
  (*data->env).CallVoidMethod(data->obj, meth, arg);
  printf("five");
}

%}

%typemap(jstype) kuzzle_event_listener "Callback";
%typemap(jtype) kuzzle_event_listener "Callback";
%typemap(jni) kuzzle_event_listener "jobject";
%typemap(javain) kuzzle_event_listener "$javainput";

%typemap(in) (kuzzle_event_listener cb, json_object *userdata) {
  struct callback_data *data = malloc(sizeof *data);
  data->env = jenv;
  data->obj = JCALL1(NewGlobalRef, jenv, $input);
  JCALL1(DeleteLocalRef, jenv, $input);
  $1 = java_callback;
  $2 = userdata;
}

%define _Complex
%enddef

%include "kuzzlesdk.h"
%include "kuzzle.h"
%include "kuzzle.hpp"
%include "callback.hpp"