// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef SWIG_H_
#define SWIG_H_

#define __swig_build

//Json
extern void kuzzle_json_put(json_object*, char*, void*, int);
extern char* kuzzle_json_get_string(json_object*, char*);
extern int kuzzle_json_get_int(json_object*, char*);
extern double kuzzle_json_get_double(json_object*, char*);
extern json_bool kuzzle_json_get_bool(json_object*, char*);
extern json_object* kuzzle_json_get_json_object(json_object*, char*);
extern void kuzzle_json_new(json_object*);

#endif