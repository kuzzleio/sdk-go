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
%rename(User, match="class") user;
%rename(RoomOptions) room_options;
%rename(SearchFilters) search_filters;
%rename(SearchResult) search_result;
%rename(NotificationResult) notification_result;
%rename(NotificationContent) notification_content;
%rename(SubscribeToSelf) subscribe_to_self;

// struct options
%rename(queueMaxSize) queue_max_size;
%rename(offlineMode) offline_mode;
%rename(autoQueue) auto_queue;
%rename(autoReconnect) auto_reconnect;
%rename(autoReplay) auto_replay;
%rename(autoReplay) auto_replay;
%rename(autoResubscribe) auto_resubscribe;
%rename(reconnectionDelay) reconnection_delay;
%rename(replayInterval) replay_interval;
%rename(defaultIndex) default_index;

// struct query_options
%rename(scrollId) scroll_id;
%rename(ifExist) if_exist;
%rename(retryOnConflict) retry_on_conflict;

// struct kuzzle_request
%rename(membersLength) members_length;
%rename(keysLength) keys_length;
%rename(fieldsLength) fields_length;

// struct role
%rename(profileIds) profile_ids;
%rename(profileIdsLength) profile_ids_length;

// struct notification_result
%rename(nType) n_type;
%rename(roomId) room_id;

// struct statistics
%rename(completedRequests) completed_requests;
%rename(failedRequests) failed_requests;
%rename(ongoingRequests) ongoing_requests;

// struct token_validity
%rename(expiresAt) expires_at;

// struct kuzzle_response
%rename(requestId) request_id;
%rename(roomId) room_id;

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

%include "exceptions.i"
%include "std_string.i"
%include "typemap.i"
%include "javadoc/kuzzle.i"
%include "javadoc/document.i"
%include "javadoc/room.i"
%include "javadoc/collection.i"
%include "../../kcore.i"

%include "std_vector.i"

%template(StringVector) std::vector<std::string>;

%typemap(out) const StringVector& %{
    return $1;
%}

%pragma(java) jniclasscode=%{
  static {
    try {
      System.loadLibrary("kuzzle-wrapper-java");
    } catch (UnsatisfiedLinkError e) {
      try {
        java.io.InputStream inputStream = kuzzlesdk.class.getResourceAsStream("/libkuzzle-wrapper-java.so");
        java.nio.file.Path path = java.nio.file.FileSystems.getDefault().getPath("").toAbsolutePath();
        String sharedObject = path.toString() + "/libs/libkuzzle-wrapper-java.so";

        try {
          java.io.File folder = new java.io.File(path.toString() + "/libs/");
          folder.mkdir();
        } catch(Exception ee) {}

        java.io.OutputStream outputStream = new java.io.FileOutputStream(new java.io.File(sharedObject));

        int read = 0;
        byte[] bytes = new byte[1024];

        while ((read = inputStream.read(bytes)) != -1) {
          outputStream.write(bytes, 0, read);
        }

        System.load(path.toString() + "/libs/libkuzzle-wrapper-java.so");
      } catch (Exception ex) {
        System.err.println("Native code library failed to load. \n");
        ex.printStackTrace();
        System.exit(1);
      }
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
