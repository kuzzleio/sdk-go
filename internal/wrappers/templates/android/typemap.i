%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%typemap(javacode) kuzzleio::Document %{

  public org.json.JSONObject create(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws Exception {
    String res = create(index, collection, id, body.toString());

    return new org.json.JSONObject(res);
  }

%}
