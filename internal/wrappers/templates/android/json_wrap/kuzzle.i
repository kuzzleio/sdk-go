%rename (_emitEvent) kuzzleio::Server::getAllStats(Event event, const std::string& body);

%typemap(javacode) kuzzleio::Kuzzle %{
  public void emitEvent(Event event, org.json.JSONObject body) throws org.json.JSONException {
    emitEvent(event, body.toString());
  }
%}
