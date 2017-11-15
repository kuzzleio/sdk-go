#ifndef SWIG_H_
#define SWIG_H_

//Json
extern void kuzzle_wrapper_json_put(json_object*, char*, void*, int);
extern char* kuzzle_wrapper_json_get_string(json_object*, char*);
extern int kuzzle_wrapper_json_get_int(json_object*, char*);
extern double kuzzle_wrapper_json_get_double(json_object*, char*);
extern json_bool kuzzle_wrapper_json_get_bool(json_object*, char*);
extern json_object* kuzzle_wrapper_json_get_json_object(json_object*, char*);
extern void kuzzle_wrapper_json_new(json_object*);

#endif