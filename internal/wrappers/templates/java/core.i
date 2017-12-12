%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options, match="class") options;
%rename(QueryOptions) query_options;
//%rename(Kuzzle, match="class") kuzzle;

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

%{
#include "core.cpp"
%}

%include "exceptions.i"
%include "std_string.i"
%include "typemap.i"
%include "javadoc.i"
%include "../../kcore.i"

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


%typemap(javaimports) kuzzle "
/* The type Kuzzle. */"

//%extend kuzzle {
//    // ctors && dtor
//    kuzzle(char* host, options *opts) {
//        kuzzle *k = (kuzzle *)calloc(1, sizeof(kuzzle));
//        kuzzle_new_kuzzle(k, host, (char *)"websocket", opts);
//        return k;
//    }
//    kuzzle(char* host) {
//        kuzzle *k;
//        k = (kuzzle *)calloc(1, sizeof(kuzzle));
//        kuzzle_new_kuzzle(k, host, (char *)"websocket", NULL);
//        return k;
//    }
//    ~kuzzle() {
//        unregisterKuzzle($self);
//        free($self);
//    }
//
//    // checkToken
//    token_validity* checkToken(char* token) {
//        return kuzzle_check_token($self, token);
//    }
//
//    // connect
//    char* connect();
//
//    // createIndex
//    bool_result* createIndex(char* index, query_options* options) {
//        return kuzzle_create_index($self, index, options);
//    }
//    bool_result* createIndex(char* index) {
//        return kuzzle_create_index($self, index, NULL);
//    }
//
//
//    // createMyCredentials
//    json_result* createMyCredentials(char* strategy, json_object* credentials, query_options* options) {
//        return kuzzle_create_my_credentials($self, strategy, credentials, options);
//    }
//    json_result* createMyCredentials(char* strategy, json_object* credentials) {
//        return kuzzle_create_my_credentials($self, strategy, credentials, NULL);
//    }
//
//    // deleteMyCredentials
//    bool_result* deleteMyCredentials(char* strategy, query_options *options) {
//        return kuzzle_delete_my_credentials($self, strategy, options);
//    }
//    bool_result* deleteMyCredentials(char* strategy) {
//        return kuzzle_delete_my_credentials($self, strategy, NULL);
//    }
//
//    // getMyCredentials
//    json_result* getMyCredentials(char *strategy, query_options *options) {
//        return kuzzle_get_my_credentials($self, strategy, options);
//    }
//    json_result* getMyCredentials(char *strategy) {
//        return kuzzle_get_my_credentials($self, strategy, NULL);
//    }
//
//    // updateMyCredentials
//    json_result* updateMyCredentials(char *strategy, json_object* credentials, query_options *options) {
//        return kuzzle_update_my_credentials($self, strategy, credentials, options);
//    }
//    json_result* updateMyCredentials(char *strategy, json_object* credentials) {
//        return kuzzle_update_my_credentials($self, strategy, credentials, NULL);
//    }
//
//    // validateMyCredentials
//    bool_result* validateMyCredentials(char *strategy, json_object* credentials, query_options* options) {
//        return kuzzle_validate_my_credentials($self, strategy, credentials, options);
//    }
//    bool_result* validateMyCredentials(char *strategy, json_object* credentials) {
//        return kuzzle_validate_my_credentials($self, strategy, credentials, NULL);
//    }
//
//    // login
//    string_result* login(char* strategy, json_object* credentials, int expires_in) {
//        return kuzzle_login($self, strategy, credentials, &expires_in);
//    }
//    string_result* login(char* strategy, json_object* credentials) {
//        return kuzzle_login($self, strategy, credentials, NULL);
//    }
//    string_result* login(char* strategy) {
//        return kuzzle_login($self, strategy, NULL, NULL);
//    }
//
//    // getAllStatistics
//    all_statistics_result* getAllStatistics(query_options* options) {
//        return kuzzle_get_all_statistics($self, options);
//    }
//    all_statistics_result* getAllStatistics() {
//        return kuzzle_get_all_statistics($self, NULL);
//    }
//
//    // getStatistics
//    statistics_result* getStatistics(time_t time, query_options* options) {
//        return kuzzle_get_statistics($self, time, options);
//    }
//    statistics_result* getStatistics(time_t time) {
//        return kuzzle_get_statistics($self, time, NULL);
//    }
//
//    // getAutoRefresh
//    bool_result* getAutoRefresh(char* index, query_options* options) {
//        return kuzzle_get_auto_refresh($self, index, options);
//    }
//    bool_result* getAutoRefresh(char* index) {
//        return kuzzle_get_auto_refresh($self, index, NULL);
//    }
//
//    // getJwt
//    char* getJwt() {
//        return kuzzle_get_jwt($self);
//    }
//
//    // getMyRights
//    json_result* getMyRights(query_options* options) {
//        return kuzzle_get_my_rights($self, options);
//    }
//    json_result* getMyRights() {
//        return kuzzle_get_my_rights($self, NULL);
//    }
//
//    // getServerInfo
//    json_result* getServerInfo(query_options* options) {
//        return kuzzle_get_server_info($self, options);
//    }
//    json_result* getServerInfo() {
//        return kuzzle_get_server_info($self, NULL);
//    }
//
//    // listCollections
//    collection_entry_result* listCollections(char *index, query_options* options) {
//        return kuzzle_list_collections($self, index, options);
//    }
//    collection_entry_result* listCollections(char *index) {
//        return kuzzle_list_collections($self, index, NULL);
//    }
//    collection_entry_result* listCollections() {
//        return kuzzle_list_collections($self, NULL, NULL);
//    }
//
//    // listIndexes
//    string_array_result* listIndexes(query_options* options) {
//        return kuzzle_list_indexes($self, options);
//    }
//    string_array_result* listIndexes() {
//        return kuzzle_list_indexes($self, NULL);
//    }
//
//    // disconnect
//    void disconnect();
//
//    // logout
//    void logout();
//
//    // query
//    kuzzle_response* query(kuzzle_request* request, query_options* options) {
//        return kuzzle_query($self, request, options);
//    }
//    kuzzle_response* query(kuzzle_request* request) {
//        return kuzzle_query($self, request, NULL);
//    }
//
//    // refreshIndex
//    shards_result* refreshIndex(char* index, query_options* options) {
//        return kuzzle_refresh_index($self, index, options);
//    }
//    shards_result* refreshIndex(char* index) {
//        return kuzzle_refresh_index($self, index, NULL);
//    }
//
//    // addListener
//    kuzzle* addListener(enum Event ev, Callback_t cb) {
//      printf("%p\n", (void*)&cb);
//      cb(42);
//      //active(42);
//
//        //kuzzle_add_listener($self, ev, c);
////        event_callback_list[ev].push_back(c);
//
////        auto pf = [](int e, json_object* o) {
//            /*for (auto& cb : event_callback_list[e]) {
//                printf("-- %p\n", cb);
//                cb->run(static_cast<Event>(e), o);
//            }*/
////        };
//
//        //kuzzle_add_listener($self, static_cast<int>(ev), NULL);
//        return $self;
//    }
//
//    // now
//    date_result* now(query_options* options) {
//        return kuzzle_now($self, options);
//    }
//    date_result* now() {
//        return kuzzle_now($self, NULL);
//    }
//
//    // removeListener
//    kuzzle* removeListener(enum Event ev) {
//        kuzzle_remove_listener($self, (int)ev, NULL);
//        return $self;
//    }
//
//    // removeAllListener
//    kuzzle* removeAllListener() {
//    }
//}

%include "core.cpp"