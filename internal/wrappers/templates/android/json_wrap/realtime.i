%rename (_list) kuzzleio::Realtime::list(const std::string& index, const std::string& collection, query_options *options);
%rename (_list) kuzzleio::Realtime::list(const std::string& index, const std::string& collection);
%javamethodmodifiers kuzzleio::Realtime::list(const std::string& index, const std::string& collection, query_options *options) "private";
%javamethodmodifiers kuzzleio::Realtime::list(const std::string& index, const std::string& collection) "private";
%javamethodmodifiers kuzzleio::Realtime::publish(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Realtime::publish(const std::string& index, const std::string& collection, const std::string& body) "private";
%rename (_subscribe) kuzzleio::Realtime::subscribe(const std::string& index, const std::string& collection, const std::string& body, NotificationListener* cb, room_options* options);
%rename (_subscribe) kuzzleio::Realtime::subscribe(const std::string& index, const std::string& collection, const std::string& body, NotificationListener* cb);
%javamethodmodifiers kuzzleio::Realtime::subscribe(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Realtime::subscribe(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Realtime::validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Realtime::validate(const std::string& index, const std::string& collection, const std::string& body) "private";

%typemap(javacode) kuzzleio::Realtime %{

  public org.json.JSONObject list(String index, String collection, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _list(index, collection, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject list(String index, String collection) throws org.json.JSONException, KuzzleException {
    return list(index, collection, null);
  }

  public void publish(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    publish(index, collection, body.toString(), options);
  }

  public void publish(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    publish(index, collection, body, null);    
  }

  public org.json.JSONObject subscribe(String index, String collection, org.json.JSONObject body, NotificationListener cb, RoomOptions options) throws org.json.JSONException, KuzzleException {
    String res = _subscribe(index, collection, body.toString(), cb, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject subscribe(String index, String collection, org.json.JSONObject body, NotificationListener cb) throws org.json.JSONException, KuzzleException {
    return subscribe(index, collection, body, cb, null);
  }

  public boolean validate(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return validate(index, collection, body.toString(), options);
  }

  public boolean validate(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return validate(index, collection, body, null);    
  }
%}