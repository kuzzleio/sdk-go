#ifndef __SDK_WRAPPERS_INTERNAL
#define __SDK_WRAPPERS_INTERNAL

typedef char *char_ptr;
typedef policy *policy_ptr;
typedef policy_restriction *policy_restriction_ptr;
typedef query_object *query_object_ptr;

// used by memory_storage.geopos
typedef double geopos_arr[2];

static void set_errno(int err) {
  errno = err;
}

static void kuzzle_notify(kuzzle_notification_listener f, notification_result* res, void* data) {
    f(res, data);
}

static void kuzzle_trigger_event(int event, kuzzle_event_listener f, char* res, void* data) {
    f(event, res, data);
}

static void room_on_subscribe(kuzzle_subscribe_listener f, room_result* res, void* data) {
    f(res, data);
}

static bool kuzzle_filter_query(kuzzle_queue_filter f, const char *rq) {
  return f(rq);
}

#endif
