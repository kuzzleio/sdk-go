// Copyright 2015-2018 Kuzzle
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

#ifndef _KUZZLESDK_H_
#define _KUZZLESDK_H_

#include <time.h>
#include <errno.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>

enum Mode {AUTO, MANUAL};
//options passed to the Kuzzle() fct

#define KUZZLE_OPTIONS_DEFAULT { \
    .queue_ttl = 120000, \
    .queue_max_size = 500, \
    .offline_mode = MANUAL,  \
    .auto_queue = false,  \
    .auto_reconnect = true,  \
    .auto_replay = false, \
    .auto_resubscribe = true, \
    .reconnection_delay = 1000, \
    .replay_interval = 10, \
    .refresh = NULL \
}

enum Event {
    CONNECTED,
    DISCARDED,
    DISCONNECTED,
    LOGIN_ATTEMPT,
    NETWORK_ERROR,
    OFFLINE_QUEUE_POP,
    OFFLINE_QUEUE_PUSH,
    QUERY_ERROR,
    RECONNECTED,
    JWT_EXPIRED,
    ERROR
};

enum is_action_allowed {
    ALLOWED,
    CONDITIONNAL,
    DENIED
};


# ifdef __cplusplus
namespace kuzzleio {
# endif

//query object used by query()
typedef struct {
    char *query;
    unsigned long long timestamp;
    char   *request_id;
} query_object;

typedef struct {
    query_object **queries;
    size_t queries_length;
} offline_queue;

typedef void (*kuzzle_event_listener)(int, char*, void*);

//define a request
typedef struct {
    const char *request_id;
    const char *controller;
    const char *action;
    const char *index;
    const char *collection;
    const char *body;
    const char *id;
    long from;
    long size;
    const char *scroll;
    const char *scroll_id;
    const char *strategy;
    unsigned long long expires_in;
    const char *volatiles;
    const char *scope;
    const char *state;
    const char *user;
    long start;
    long stop;
    long end;
    unsigned char bit;
    const char *member;
    const char *member1;
    const char *member2;
    char **members;
    size_t members_length;
    double lon;
    double lat;
    double distance;
    const char *unit;
    const char *options;
    const char **keys;
    size_t keys_length;
    long cursor;
    long offset;
    const char *field;
    const char **fields;
    size_t fields_length;
    const char *subcommand;
    const char *pattern;
    long idx;
    const char *min;
    const char *max;
    const char *limit;
    unsigned long count;
    const char *match;
} kuzzle_request;

typedef offline_queue* (*kuzzle_offline_queue_loader)(void);
typedef bool (*kuzzle_queue_filter)(const char*);

typedef struct {
    void *instance;
    kuzzle_queue_filter filter;
    kuzzle_offline_queue_loader loader;
} kuzzle;

typedef struct {
  void *instance;
  kuzzle* k;
} realtime;

typedef struct {
    const char *type_;
    int  from;
    int  size;
    const char *scroll;
} search_options;

typedef struct auth {
  void *instance;
  kuzzle *k;
} auth;

typedef struct {
  void *instance;
  kuzzle *k;
} kuzzle_index;

typedef struct {
  void *instance;
  kuzzle* k;
} server;

typedef struct {
  const char *room;
  const char *channel;
  int status;
  const char *error;
  const char *stack;
} subscribe_result;

//options passed to room constructor
typedef struct {
    const char *scope;
    const char *state;
    const char *user;
    bool subscribe_to_self;
    const char *volatiles;
} room_options;

typedef struct {
    void *instance;
    const char *filters;
    room_options *options;
} room;

typedef struct {
  room *result;
  int status;
  const char *error;
  const char *stack;
} room_result;

typedef void (callback)(char* notification);

//options passed to query()
typedef struct {
    bool queuable;
    bool withdist;
    bool withcoord;
    long from;
    long size;
    const char *scroll;
    const char *scroll_id;
    const char *refresh;
    const char *if_exist;
    int retry_on_conflict;
    const char *volatiles;
} query_options;

typedef struct {
    unsigned queue_ttl;
    unsigned long queue_max_size;
    enum Mode offline_mode;
    bool auto_queue;
    bool auto_reconnect;
    bool auto_replay;
    bool auto_resubscribe;
    unsigned long reconnection_delay;
    unsigned long replay_interval;
    const char *refresh;
} options;

//meta of a document
typedef struct {
    const char *author;
    unsigned long long created_at;
    unsigned long long updated_at;
    const char *updater;
    bool active;
    unsigned long long deleted_at;
} meta;

/* === Security === */

typedef char *controllers;

typedef struct  {
    const char *index;
    char **collections;
    size_t collections_length;
} policy_restriction;

typedef struct {
    const char *role_id;
    policy_restriction *restricted_to;
    size_t restricted_to_length;
} policy;

typedef struct {
    const char *id;
    policy *policies;
    size_t policies_length;
    kuzzle *k;
} profile;

typedef struct {
    const char *id;
    const char *controllers;
    kuzzle *k;
} role;

// kuzzle user
typedef struct {
    const char *id;
    const char *content;
    char **profile_ids;
    size_t profile_ids_length;
    kuzzle *k;
} kuzzle_user;

// user content passed to user constructor
typedef struct {
    const char *content;
    char **profile_ids;
    size_t profile_ids_length;
} user_data;

/* === Dedicated response structures === */

typedef struct {
  int failed;
  int successful;
  int total;
} shards;

typedef struct {
    void *instance;
    kuzzle *k;
} collection;

typedef struct {
    void *instance;
    kuzzle *k;
} document;

typedef struct {
    const char *id;
    meta *m;
    const char *content;
    int count;
} notification_content;

typedef struct notification_result {
    const char *request_id;
    notification_content *result;
    const char *volatiles;
    const char *index;
    const char *collection;
    const char *controller;
    const char *action;
    const char *protocol;
    const char *scope;
    const char *state;
    const char *user;
    const char *n_type;
    const char *room_id;
    unsigned long long timestamp;
    int status;
    const char *error;
    const char *stack;
} notification_result;

typedef struct profile_result {
    profile *p;
    int status;
    const char *error;
    const char *stack;
} profile_result;

typedef struct profiles_result {
    profile *profiles;
    size_t profiles_length;
    int status;
    const char *error;
    const char *stack;
} profiles_result;

typedef struct role_result {
    role *r;
    int status;
    const char *error;
    const char *stack;
} role_result;

typedef struct roles_result {
    role *roles;
    size_t roles_length;
    int status;
    const char *error;
    const char *stack;
} roles_result;

typedef struct {
    const char *controller;
    const char *action;
    const char *index;
    const char *collection;
    const char *value;
} user_right;

typedef struct user_rights_result {
  user_right *result;
  size_t user_rights_length;
  int status;
  const char *error;
  const char *stack;
} user_rights_result;

typedef struct user_result {
    kuzzle_user *result;
    int status;
    const char *error;
    const char *stack;
} user_result;

//statistics
typedef struct {
    const char* completed_requests;
    const char* connections;
    const char* failed_requests;
    const char* ongoing_requests;
    unsigned long long timestamp;
} statistics;

typedef struct statistics_result {
    statistics* result;
    int status;
    const char *error;
    const char *stack;
} statistics_result;

typedef struct all_statistics_result {
    statistics* result;
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} all_statistics_result;

// ms.geopos
typedef struct geopos_result {
    double (*result)[2];
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} geopos_result;

// ms.geopoint
typedef struct point {
    float lat;
    float lon;
    const char *name;
} point;

// ms.msHashField
typedef struct ms_hash_field {
  const char *field;
  const char *value;
} ms_hash_field;

// ms.keyValue
typedef struct ms_key_value {
  const char *key;
  const char *value;
} ms_key_value;

// ms.sortedSet
typedef struct ms_sorted_set {
  float score;
  const char *member;
} ms_sorted_set;

//check_token
typedef struct token_validity {
    bool valid;
    const char *state;
    unsigned long long expires_at;
    int status;
    const char *error;
    const char *stack;
} token_validity;

/* === Generic response structures === */

// raw Kuzzle response
typedef struct {
    const char *request_id;
    const char *result;
    const char *volatiles;
    const char *index;
    const char *collection;
    const char *controller;
    const char *action;
    const char *room_id;
    const char *channel;
    int status;
    const char *error;
    const char *stack;
} kuzzle_response;

//any void result
typedef struct error_result {
    int status;
    const char *error;
    const char *stack;
} error_result;

//any json result
typedef struct json_result {
    const char *result;
    int status;
    const char *error;
    const char *stack;
} json_result;

//any array of char result
typedef struct json_array_result {
    char **result;
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} json_array_result;

//any boolean result
typedef struct bool_result {
    bool result;
    int status;
    const char *error;
    const char *stack;
} bool_result;

//any integer result
typedef struct int_result {
    long long result;
    int status;
    const char *error;
    const char *stack;
} int_result;

typedef struct date_result {
    long long result;
    int status;
    const char *error;
    const char *stack;
} date_result;

//any double result
typedef struct double_result {
    double result;
    int status;
    const char *error;
    const char *stack;
} double_result;

//any array of integers result
typedef struct int_array_result {
    long long *result;
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} int_array_result;

// any string result
typedef struct string_result {
    const char *result;
    int status;
    const char *error;
    const char *stack;
} string_result;

//any array of strings result
typedef struct string_array_result {
    char **result;
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} string_array_result;

typedef struct {
    profile *hits;
    size_t hits_length;
    unsigned total;
    const char *scroll_id;
} profile_search;

typedef struct {
    role *hits;
    size_t hits_length;
    unsigned total;
} role_search;

typedef struct {
    kuzzle_user *hits;
    size_t hits_length;
    unsigned total;
    const char *scroll_id;
} user_search;

//any delete* function
typedef struct ack_result {
    bool acknowledged;
    bool shards_acknowledged;
    int status;
    const char *error;
    const char *stack;
} ack_result;

typedef struct shards_result {
    shards *result;
    int status;
    const char *error;
    const char *stack;
} shards_result;

typedef struct {
    bool strict;
    const char *fields;
    const char *validators;
} specification;

typedef struct {
    specification *validation;
    const char *index;
    const char *collection;
} specification_entry;

typedef struct specification_result {
    specification *result;
    int status;
    const char *error;
    const char *stack;
} specification_result;

typedef struct search_result {
    const char *documents;
    unsigned fetched;
    unsigned total;
    const char *aggregations;
    const char *filters;
    query_options *options;
    const char *collection;
    int status;
    const char *error;
    const char *stack;
} search_result;

typedef struct search_profiles_result {
    profile_search *result;
    int status;
    const char *error;
    const char *stack;
} search_profiles_result;

typedef struct search_roles_result {
    role_search *result;
    int status;
    const char *error;
    const char *stack;
} search_roles_result;

typedef struct search_users_result {
    user_search *result;
    int status;
    const char *error;
    const char *stack;
} search_users_result;

typedef struct {
    specification_entry *hits;
    size_t hits_length;
    unsigned total;
    const char *scroll_id;
} specification_search;

typedef struct specification_search_result {
    specification_search *result;
    int status;
    const char *error;
    const char *stack;
} specification_search_result;

typedef struct  {
    bool persisted;
    const char *name;
} collection_entry;

typedef struct collection_entry_result {
    collection_entry* result;
    size_t result_length;
    int status;
    const char *error;
    const char *stack;
} collection_entry_result;

typedef void (*kuzzle_notification_listener)(notification_result*, void*);
typedef void (*kuzzle_subscribe_listener)(room_result*, void*);

# ifdef __cplusplus
}
# endif

#endif
