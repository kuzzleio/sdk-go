#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json-c/json.h>
#include <time.h>
#include <errno.h>
#include <stdbool.h>

typedef struct {
    void *instance;
} kuzzle;

enum {
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
} event;

//define a request
typedef struct {
    char *request_id;
    char *controller;
    char *action;
    char *index;
    char *collection;
    json_object *body;
    char *id;
    long from;
    long size;
    char *scroll;
    char *scroll_id;
    char *strategy;
    unsigned long long expires_in;
    json_object *volatiles;
    char *scope;
    char *state;
    char *user;
    long start;
    long stop;
    long end;
    unsigned char bit;
    char *member;
    char *member1;
    char *member2;
    char **members;
    size_t members_length;
    double lon;
    double lat;
    double distance;
    char *unit;
    json_object *options;
    char **keys;
    size_t keys_length;
    long cursor;
    long offset;
    char *field;
    char **fields;
    size_t fields_length;
    char *subcommand;
    char *pattern;
    long idx;
    char *min;
    char *max;
    char *limit;
    unsigned long count;
    char *match;
} kuzzle_request;

//query object used by query()
typedef struct {
    json_object *query;
    unsigned long long timestamp;
    char   *request_id;
} query_object;

typedef struct {
    query_object **queries;
    size_t queries_length;
} offline_queue;

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
    json_object *volatiles;
} query_options;

//options passed to room constructor
typedef struct {
    char *scope;
    char *state;
    char *user;
    bool subscribe_to_self;
    json_object *volatiles;
} room_options;

enum Mode {AUTO, MANUAL};
//options passed to the Kuzzle() fct
typedef struct {
    unsigned queue_ttl;
    unsigned long queue_max_size;
    unsigned char offline_mode;
    unsigned char auto_queue;
    unsigned char auto_reconnect;
    unsigned char auto_replay;
    unsigned char auto_resubscribe;
    unsigned long reconnection_delay;
    unsigned long replay_interval;
    enum Mode connect;
    char *refresh;
    char *default_index;
    json_object *headers;
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

typedef json_object controllers;

typedef struct {
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
    json_object *controllers;
    kuzzle *kuzzle;
} role;

//kuzzle user
typedef struct {
    char *id;
    json_object *content;
    char **profile_ids;
    size_t profile_ids_length;
    kuzzle *kuzzle;
} user;

// user content passed to user constructor
typedef struct {
    json_object *content;
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
    char *index;
    char *collection;
    kuzzle *kuzzle;
} collection;

typedef struct {
    char *id;
    char *index;
    meta *meta;
    shards *shards;
    json_object *content;
    int version;
    char *result;
    bool created;
    char *collection;
    collection *_collection;
} document;

typedef struct {
    document *result;
    int status;
    char *error;
    char *stack;
} document_result;

typedef struct {
    char *id;
    meta *meta;
    json_object *content;
    int count;
} notification_content;

typedef struct {
    char *request_id;
    notification_content *result;
    json_object *volatiles;
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

typedef struct {
    profile *profile;
    int status;
    char *error;
    char *stack;
} profile_result;

typedef struct {
    profile *profiles;
    size_t profiles_length;
    int status;
    char *error;
    char *stack;
} profiles_result;

typedef struct {
    role *role;
    int status;
    char *error;
    char *stack;
} role_result;

typedef struct {
    char *controller;
    char *action;
    char *index;
    char *collection;
    char *value;
} user_right;

typedef struct {
    user_right *user_rights;
    size_t user_rights_length;
    int status;
    char *error;
    char *stack;
} user_rights_result;

typedef struct {
    user *user;
    int status;
    char *error;
    char *stack;
} user_result;

enum {
    ALLOWED=0,
    CONDITIONNAL=1,
    DENIED=2
} is_action_allowed;

//statistics
typedef struct {
    json_object* completed_requests;
    json_object* connections;
    json_object* failed_requests;
    json_object* ongoing_requests;
    unsigned long long timestamp;
} statistics;

typedef struct {
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
typedef struct {
    double (*result)[2];
    size_t result_length;
    int status;
    char *error;
    char *stack;
} geopos_result;

//check_token
typedef struct {
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
    json_object *result;
    json_object *volatiles;
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
typedef struct {
    int status;
    char *error;
    char *stack;
} void_result;

//any json result
typedef struct {
    json_object *result;
    int status;
    char *error;
    char *stack;
} json_result;

//any array of json_object result
typedef struct {
    json_object **result;
    size_t result_length;
    int status;
    char *error;
    char *stack;
} json_array_result;

//any boolean result
typedef struct {
    bool result;
    int status;
    char *error;
    char *stack;
} bool_result;

//any integer result
typedef struct {
    long long result;
    int status;
    char *error;
    char *stack;
} int_result;

//any double result
typedef struct {
    double result;
    int status;
    char *error;
    char *stack;
} double_result;

//any array of integers result
typedef struct {
    long long *result;
    size_t result_length;
    int status;
    char *error;
    char*stack;
} int_array_result;

// any string result
typedef struct {
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
    json_object* query;
    json_object* sort;
    json_object* aggregations;
    json_object* search_after;
} search_filters;

typedef struct {
    document *hits;
    size_t hits_length;
    unsigned total;
    char *scroll_id;
} document_search;

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
typedef struct {
    bool acknowledged;
    bool shards_acknowledged;
    int status;
    char *error;
    char *stack;
} ack_result;

typedef struct {
    shards *result;
    int status;
    char *error;
    char *stack;
} shards_result;

typedef struct {
    bool strict;
    json_object *fields;
    json_object *validators;
} specification;

typedef struct {
    specification *validation;
    char *index;
    char *collection;
} specification_entry;

typedef struct {
    specification *result;
    int status;
    char *error;
    char *stack;
} specification_result;

typedef struct {
    document_search *result;
    int status;
    char *error;
    char *stack;
} search_result;

typedef struct {
    profile_search *result;
    int status;
    char *error;
    char *stack;
} search_profiles_result;

typedef struct {
    role_search *result;
    int status;
    char *error;
    char *stack;
} search_roles_result;

typedef struct {
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

typedef struct {
    specification_search *result;
    int status;
    char *error;
    char *stack;
} specification_search_result;

typedef struct {
    json_object *mapping;
    collection *collection;
} mapping;

typedef struct {
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

#endif
