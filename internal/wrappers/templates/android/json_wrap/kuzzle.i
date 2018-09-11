%rename (_emitEvent) kuzzleio::Kuzzle::emitEvent(Event& event, const std::string& body);

%typemap(javacode) kuzzleio::Kuzzle %{
  public void emitEvent(Event event, org.json.JSONObject body) throws org.json.JSONException {
    _emitEvent(event, body.toString());
  }
%}