%javamethodmodifiers kuzzleio::Collection::create(const std::string& index, const std::string& collection, const std::string* body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::create(const std::string& index, const std::string& collection, const std::string* body) "private";
%rename (_list) kuzzleio::Collection::list(const std::string& index, query_options *options);
%rename (_list) kuzzleio::Collection::list(const std::string& index);
%javamethodmodifiers kuzzleio::Collection::list(const std::string& index, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::list(const std::string& index) "private";
%rename (_getMapping) kuzzleio::Collection::getMapping(const std::string& index, const std::string& collection, query_options *options);
%rename (_getMapping) kuzzleio::Collection::getMapping(const std::string& index, const std::string& collection);
%javamethodmodifiers kuzzleio::Collection::getMapping(const std::string& index, const std::string& collection, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::getMapping(const std::string& index, const std::string& collection) "private";
%javamethodmodifiers kuzzleio::Collection::updateMapping(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::updateMapping(const std::string& index, const std::string& collection, const std::string& body) "private";
%rename (_getSpecifications) kuzzleio::Collection::getSpecifications(const std::string& index, const std::string& collection, query_options *options);
%rename (_getSpecifications) kuzzleio::Collection::getSpecifications(const std::string& index, const std::string& collection);
%javamethodmodifiers kuzzleio::Collection::getSpecifications(const std::string& index, const std::string& collection, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::getSpecifications(const std::string& index, const std::string& collection) "private";
%rename (_updateSpecifications) kuzzleio::Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body, query_options *options);
%rename (_updateSpecifications) kuzzleio::Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body);
%javamethodmodifiers kuzzleio::Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Collection::validateSpecifications(const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Collection::validateSpecifications(const std::string& body) "private";

i
  public void create(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    create(index, collection, body.toString(), options);
  }

  public void create(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    create(index, collection, body, null);
  }

  public org.json.JSONObject list(String index, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _list(index, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject list(String index) throws org.json.JSONException, KuzzleException {
    return list(index, null);
  }

  public org.json.JSONObject getMapping(String index, String collection, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res =_getMapping(index, collection, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getMapping(String index, String collection) throws org.json.JSONException, KuzzleException {
    return getMapping(index, collection, null);
  }

  public void updateMapping(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    updateMapping(index, collection, body.toString(), options);
  }

  public void updateMapping(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    updateMapping(index, collection, body, null);
  }

  public org.json.JSONObject getSpecifications(String index, String collection, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getSpecifications(index, collection, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getSpecifications(String index, String collection) throws org.json.JSONException, KuzzleException {
    return getSpecifications(index, collection, null);
  }

  public org.json.JSONObject updateSpecifications(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _updateSpecifications(index, collection, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject updateSpecifications(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return updateSpecifications(index, collection, body, null);
  }

  public boolean validateSpecifications(org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return validateSpecifications(body.toString(), options);
  }

  public boolean validateSpecifications(org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return validateSpecifications(body, null);
  }

%}
