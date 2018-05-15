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

#ifndef _KUZZLESDK_H_
#define _KUZZLESDK_H_

#include <time.h>
#include <errno.h>
#include <stdbool.h>

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
    const long start;
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
  kuzzle* kuzzle;
} realtime;

typedef struct {
    char *type_;
    int  from;
    int  size;
    char *scroll;
} search_options;

typedef struct auth {
  void *instance;
  kuzzle *kuzzle;
} auth;

typedef struct {
  void *instance;
  kuzzle *kuzzle;
} kuzzle_index;

typedef struct {
  void *instance;
  kuzzle* kuzzle;
} server;

typedef struct {
  char *room;
  char *channel;
  int status;
  char *error;
  char *stack;
} subscribe_result;

//options passed to room constructor
typedef struct {
    char *scope;
    char *state;
    char *user;
    bool subscribe_to_self;
    char *volatiles;
} room_options;

typedef struct {
    void *instance;
    char *filters;
    room_options *options;
} room;

typedef struct {
  room *result;
  int status;
  char *error;
  char *stack;
} room_result;

typedef void (callback)(char* notification);

//options passed to query()
typedef struct {
    bool queuable;
    bool withdist;
    bool withcoord;
    long from;
    long size;
    char *scroll;
    char *scroll_id;
    char *refresh;
    char *if_exist;
    int retry_on_conflict;
    char *volatiles;
} query_options;

enum Mode {AUTO, MANUAL};
//options passed to the Kuzzle() fct

#define KUZZLE_OPTIONS_DEFAULT { \
    .queue_ttl = 120000, \
    .queue_max_size = 500, \
    .offline_mode = 0,  \
    .auto_queue = false,  \
    .auto_reconnect = true,  \
    .auto_replay = false, \
    .auto_resubscribe = true, \
    .reconnection_delay = 1000, \
    .replay_interval = 10, \
    .connect = AUTO, \
    .refresh = NULL, \
    .default_index = NULL  \
}
typedef struct {
    unsigned queue_ttl;
    unsigned long queue_max_size;
    unsigned char offline_mode;
    bool auto_queue;
    bool auto_reconnect;
    bool auto_replay;
    bool auto_resubscribe;
    unsigned long reconnection_delay;
    unsigned long replay_interval;
    enum Mode connect;
    char *refresh;
    char *default_index;
} options;

//meta of a document
typedef struct {
    char *author;
    unsigned long long created_at;
    unsigned long long updated_at;
    char *updater;
    bool active;
    unsigned long long deleted_at;
} meta;

/* === Security === */

typedef char *controllers;

typedef struct  {
    char *index;
    char **collections;
    size_t collections_length;
} policy_restriction;

typedef struct {
    char *role_id;
    policy_restriction *restricted_to;
    size_t restricted_to_length;
} policy;

typedef struct {
    char *id;
    policy *policies;
    size_t policies_length;
    kuzzle *kuzzle;
} profile;

typedef struct {
    char *id;
    char *controllers;
    kuzzle *kuzzle;
} role;

//kuzzle user
typedef struct {
    char *id;
    char *content;
    char **profile_ids;
    size_t profile_ids_length;
    kuzzle *kuzzle;
} user;

// user content passed to user constructor
typedef struct {
    char *content;
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
    kuzzle *kuzzle;
} collection;

typedef struct {
    void *instance;
    kuzzle *kuzzle;
} document;

typedef struct {
    char *id;
    meta *meta;
    char *content;
    int count;
} notification_content;

typedef struct notification_result {
    char *request_id;
    notification_content *result;
    char *volatiles;
    char *index;
    char *collection;
    char *controller;
    char *action;
    char *protocol;
    char *scope;
    char *state;
    char *user;
    char *n_type;
    char *room_id;
    unsigned long long timestamp;
    int status;
    char *error;
    char *stack;
} notification_result;

typedef struct profile_result {
    profile *profile;
    int status;
    char *error;
    char *stack;
} profile_result;

typedef struct profiles_result {
    profile *profiles;
    size_t profiles_length;
    int status;
    char *error;
    char *stack;
} profiles_result;

typedef struct role_result {
    role *role;
    int status;
    char *error;
    char *stack;
} role_result;

typedef struct roles_result {
    role *roles;
    size_t roles_length;
    int status;
    char *error;
    char *stack;
} roles_result;

typedef struct {
    char *controller;
    char *action;
    char *index;
    char *collection;
    char *value;
} user_right;

typedef struct user_rights_result {
    user_right *result;
    size_t user_rights_length;
    int status;
    char *error;
    char *stack;
} user_rights_result;

typedef struct user_result {
    user *result;
    int status;
    char *error;
    char *stack;
} user_result;

enum is_action_allowed {
    ALLOWED=0,
    CONDITIONNAL=1,
    DENIED=2
};

//statistics
typedef struct {
    char* completed_requests;
    char* connections;
    char* failed_requests;
    char* ongoing_requests;
    unsigned long long timestamp;
} statistics;

typedef struct statistics_result {
    statistics* result;
    int status;
    char *error;
    char *stack;
} statistics_result;

typedef struct all_statistics_result {
    statistics* result;
    size_t result_length;
    int status;
    char *error;
    char *stack;
} all_statistics_result;

// ms.geopos
typedef struct geopos_result {
    double (*result)[2];
    size_t result_length;
    int status;
    char *error;
    char *stack;
} geopos_result;

// ms.geopoint
typedef struct point {
    float lat;
    float lon;
    char *name;
} point;

// ms.msHashField
typedef struct ms_hash_field {
  char *field;
  char *value;
} ms_hash_field;

// ms.keyValue
typedef struct ms_key_value {
  char *key;
  char *value;
} ms_key_value;

// ms.sortedSet
typedef struct ms_sorted_set {
  float score;
  char *member;
} ms_sorted_set;

//check_token
typedef struct token_validity {
    bool valid;
    char *state;
    unsigned long long expires_at;
    int status;
    char *error;
    char *stack;
} token_validity;

/* === Generic response structures === */

// raw Kuzzle response
typedef struct {
    char *request_id;
    char *result;
    char *volatiles;
    char *index;
    char *collection;
    char *controller;
    char *action;
    char *room_id;
    char *channel;
    int status;
    char *error;
    char *stack;
} kuzzle_response;

//any void result
typedef struct void_result {
    int status;
    char *error;
    char *stack;
} void_result;

//any json result
typedef struct json_result {
    char *result;
    int status;
    char *error;
    char *stack;
} json_result;

//any array of char result
typedef struct json_array_result {
    char **result;
    size_t result_length;
    int status;
    char *error;
    char *stack;
} json_array_result;

//any boolean result
typedef struct bool_result {
    bool result;
    int status;
    char *error;
    char *stack;
} bool_result;

//any integer result
typedef struct int_result {
    long long result;
    int status;
    char *error;
    char *stack;
} int_result;

typedef struct date_result {
    long long result;
    int status;
    char *error;
    char *stack;
} date_result;

//any double result
typedef struct double_result {
    double result;
    int status;
    char *error;
    char *stack;
} double_result;

//any array of integers result
typedef struct int_array_result {
    long long *result;
    size_t result_length;
    int status;
    char *error;
    char*stack;
} int_array_result;

// any string result
typedef struct string_result {
    char *result;
    int status;
    char *error;
    char *stack;
} string_result;

//any array of strings result
typedef struct string_array_result {
    char **result;
    size_t result_length;
    int status;
    char *error;
    char *stack;
} string_array_result;

typedef struct {
    char* query;
    char* sort;
    char* aggregations;
    char* search_after;
} search_filters;

typedef struct {
    profile *hits;
    size_t hits_length;
    unsigned total;
    char *scroll_id;
} profile_search;

typedef struct {
    role *hits;
    size_t hits_length;
    unsigned total;
} role_search;

typedef struct {
    user *hits;
    size_t hits_length;
    unsigned total;
    char *scroll_id;
} user_search;

//any delete* function
typedef struct ack_result {
    bool acknowledged;
    bool shards_acknowledged;
    int status;
    char *error;
    char *stack;
} ack_result;

typedef struct shards_result {
    shards *result;
    int status;
    char *error;
    char *stack;
} shards_result;

typedef struct {
    bool strict;
    char *fields;
    char *validators;
} specification;

typedef struct {
    specification *validation;
    char *index;
    char *collection;
} specification_entry;

typedef struct specification_result {
    specification *result;
    int status;
    char *error;
    char *stack;
} specification_result;

typedef struct search_result {
    char *documents;
    unsigned fetched;
    unsigned total;
    char *aggregations;
    search_filters *filters;
    query_options *options;
    char *collection;
    int status;
    char *error;
    char *stack;
} search_result;

typedef struct search_profiles_result {
    profile_search *result;
    int status;
    char *error;
    char *stack;
} search_profiles_result;

typedef struct search_roles_result {
    role_search *result;
    int status;
    char *error;
    char *stack;
} search_roles_result;

typedef struct search_users_result {
    user_search *result;
    int status;
    char *error;
    char *stack;
} search_users_result;

typedef struct {
    specification_entry *hits;
    size_t hits_length;
    unsigned total;
    char *scroll_id;
} specification_search;

typedef struct specification_search_result {
    specification_search *result;
    int status;
    char *error;
    char *stack;
} specification_search_result;

typedef struct {
    char *mapping;
    collection *collection;
} mapping;

typedef struct mapping_result {
    mapping *result;
    int status;
    char *error;
    char *stack;
} mapping_result;

typedef struct  {
    bool persisted;
    char* name;
} collection_entry;

typedef struct collection_entry_result {
    collection_entry* result;
    size_t result_length;
    int status;
    char* error;
    char* stack;
} collection_entry_result;

typedef void (*kuzzle_notification_listener)(notification_result*, void*);
typedef void (*kuzzle_subscribe_listener)(room_result*, void*);

#endif
