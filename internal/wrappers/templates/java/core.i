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

%ignore *::error;
%ignore *::status;
%ignore *::stack;

%{
#include "kuzzle.cpp"
%}

%include "exceptions.i"
%include "std_string.i"
%include "typemap.i"
%include "javadoc.i"
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
      System.err.println("Native code library failed to load. \n" + e);
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

struct json_object { };

%ignore json_object::ptr;
%extend json_object {
    json_object() {
        json_object *j = json_object_new_object();
        return j;
    }

    ~json_object() {
        kuzzle_free_json_object($self);
    }

    json_object* put(char* key, char* content) {
        kuzzle_json_put($self, key, content, 0);
        return $self;
    }

    json_object* put(char* key, int content) {
        kuzzle_json_put($self, key, &content, 1);
        return $self;
    }

    json_object* put(char* key, double content) {
        kuzzle_json_put($self, key, &content, 2);
        return $self;
    }

    json_object* put(char* key, bool content) {
        kuzzle_json_put($self, key, &content, 3);
        return $self;
    }

    json_object* put(char* key, json_object* content) {
        kuzzle_json_put($self, key, content, 4);
        return $self;
    }

    char* getString(char* key) {
        return kuzzle_json_get_string($self, key);
    }

    int getInt(char* key) {
        return kuzzle_json_get_int($self, key);
    }

    double getDouble(char* key) {
        return kuzzle_json_get_double($self, key);
    }

    bool getBoolean(char* key) {
        return kuzzle_json_get_bool($self, key);
    }

    json_object* getJsonObject(char* key) {
        return kuzzle_json_get_json_object($self, key);
    }
}

%include "kuzzle.cpp"