%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options, match="class") options;
%rename(QueryOptions) query_options;
%rename(JsonObject) json_object;
%rename(JsonResult) json_result;
%rename(LoginResult) login_result;
%rename(BoolResult) bool_result;
%rename(Statistics) statistics;
%rename(AllStatisticsResult) all_statistics_result;
%rename(StatisticsResult) statistics_result;
%rename(CollectionsList) collection_entry;
%rename(CollectionsListResult) collection_entry_result;
%rename(StringArrayResult) string_array_result;
%rename(KuzzleResponse) kuzzle_response;
%rename(KuzzleRequest) kuzzle_request;
%rename(ShardsResult) shards_result;
%rename(DateResult) date_result;
%rename(UserData) user_data;
%rename(KuzzleUser, match="class") kuzzle_user;
%rename(RoomOptions) room_options;
%rename(SearchFilters) search_filters;
%rename(SearchResult) search_result;
%rename(NotificationResult) notification_result;
%rename(NotificationContent) notification_content;
%rename(SubscribeToSelf) subscribe_to_self;

%rename(delete) delete_;

%ignore *::error;
%ignore *::status;
%ignore *::stack;

%feature("director") NotificationListener;
%feature("director") EventListener;
%feature("director") SubscribeListener;

%{
#include "kuzzle.cpp"
#include "collection.cpp"
#include "auth.cpp"
#include "index.cpp"
#include "server.cpp"
#include "document.cpp"
#include "realtime.cpp"
%}

%include "../java/exceptions.i"
%include "std_string.i"
%include "json_wrap/document.i"
%include "json_wrap/server.i"
%include "json_wrap/collection.i"
%include "json_wrap/auth.i"

%include "kcore.i"

%include "std_vector.i"

typedef long long time_t;

%template(StringVector) std::vector<std::string>;

%typemap(out) const StringVector& %{
    return $1;
%}

%pragma(java) jniclasscode=%{
  static {
    try {
      System.loadLibrary("kuzzle-wrapper-android");
    } catch (Exception e) {
      System.err.println("Native code library failed to load. \n");
      e.printStackTrace();
      System.exit(1);
    }
  }
%}

%extend options {
    options() {
        options *o = kuzzle_new_options();
        return o;
    }

    ~options() {
        free($self);
    }
}

%include "kuzzle.cpp"
%include "collection.cpp"
%include "document.cpp"
%include "realtime.cpp"
%include "auth.cpp"
%include "index.cpp"
%include "server.cpp"
