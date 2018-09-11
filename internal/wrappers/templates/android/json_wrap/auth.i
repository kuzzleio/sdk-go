%javamethodmodifiers kuzzleio::Auth::createMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options) "private";
%javamethodmodifiers kuzzleio::Auth::createMyCredentials(const std::string& strategy, const std::string& credentials) "private";
%rename (_getMyCredentials) kuzzleio::Auth::getMyCredentials(const std::string& strategy, query_options *options);
%rename (_getMyCredentials) kuzzleio::Auth::getMyCredentials(const std::string& strategy);
%javamethodmodifiers kuzzleio::Auth::getMyCredentials(const std::string& strategy, query_options *options) "private";
%javamethodmodifiers kuzzleio::Auth::getMyCredentials(const std::string& strategy) "private";
%javamethodmodifiers kuzzleio::Auth::login(const std::string& strategy, const std::string& credentials, int expiresIn) "private";
%javamethodmodifiers kuzzleio::Auth::login(const std::string& strategy, const std::string& credentials) "private";
%rename (_updateMyCredentials) kuzzleio::Auth::updateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options);
%rename (_updateMyCredentials) kuzzleio::Auth::updateMyCredentials(const std::string& strategy, const std::string& credentials);
%javamethodmodifiers kuzzleio::Auth::updateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options) "private";
%javamethodmodifiers kuzzleio::Auth::updateMyCredentials(const std::string& strategy, const std::string& credentials) "private";
%javamethodmodifiers kuzzleio::Auth::updateSelf(const std::string& content, query_options* options) "private";
%javamethodmodifiers kuzzleio::Auth::updateSelf(const std::string& content) "private";
%javamethodmodifiers kuzzleio::Auth::validateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options) "private";
%javamethodmodifiers kuzzleio::Auth::validateMyCredentials(const std::string& strategy, const std::string& credentials) "private";

%typemap(javaimports) kuzzleio::Auth %{
import android.webkit.WebView;
import android.webkit.WebViewClient;
import java.io.IOException;
import java.net.HttpURLConnection;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.URI;
%}
%typemap(javacode) kuzzleio::Auth %{
  public org.json.JSONObject createMyCredentials(String strategy, org.json.JSONObject credentials, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = createMyCredentials(strategy, credentials.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject createMyCredentials(String strategy, org.json.JSONObject credentials) throws org.json.JSONException, KuzzleException {
    return createMyCredentials(strategy, credentials, null);
  }

  public org.json.JSONObject getMyCredentials(String strategy, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getMyCredentials(strategy, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getMyCredentials(String strategy) throws org.json.JSONException, KuzzleException {
    return getMyCredentials(strategy, null);
  }

  public String login(String strategy, org.json.JSONObject credentials, int expiresIn) throws org.json.JSONException, KuzzleException {
    return login(strategy, credentials.toString(), expiresIn);
  }

  public String login(String strategy, org.json.JSONObject credentials) throws org.json.JSONException, KuzzleException {
    return login(strategy, credentials.toString());
  }

  public org.json.JSONObject updateMyCredentials(String strategy, org.json.JSONObject credentials, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _updateMyCredentials(strategy, credentials.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject updateMyCredentials(String strategy, org.json.JSONObject credentials) throws org.json.JSONException, KuzzleException {
    return updateMyCredentials(strategy, credentials, null);
  }

  public KuzzleUser updateSelf(org.json.JSONObject content, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return updateSelf(content.toString(), options);
  }

  public KuzzleUser updateSelf(org.json.JSONObject content) throws org.json.JSONException, KuzzleException {
    return updateSelf(content, null);
  }

  public boolean validateMyCredentials(String strategy, org.json.JSONObject credentials, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return validateMyCredentials(strategy, credentials.toString(), options);
  }

  public boolean validateMyCredentials(String strategy, org.json.JSONObject credentials) throws org.json.JSONException, KuzzleException {
    return validateMyCredentials(strategy, credentials, null);
  }

  /**
   * WebViewClient to forward kuzzle's jwt token after an OAuth authentication
   */
  protected class KuzzleWebViewClient extends WebViewClient {
    @Override
    public boolean shouldOverrideUrlLoading(WebView view, final String url) {
      if (url.contains("code=")) {
        new Thread(new Runnable() {
          @Override
          public void run() {
            try {
              HttpURLConnection conn = (HttpURLConnection) URI.create(url).toURL().openConnection();
              conn.setRequestMethod("GET");
              conn.setUseCaches(false);

              BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
              StringBuilder sb = new StringBuilder();
              String line;
              while ((line = br.readLine()) != null) {
                sb.append(line);
              }
              br.close();

              org.json.JSONObject response = new org.json.JSONObject(sb.toString());
              if (response.isNull("error")) {
                Auth.this.setJwt(response.getString("jwt"));
              } else {
                Auth.this.get_kuzzle().emitEvent(Event.LOGIN_ATTEMPT, new org.json.JSONObject()
                  .put("success", false)
                  .put("error", response.getJSONObject("error")));
              }
            } catch (org.json.JSONException | IOException e) {
              e.printStackTrace();
            }
          }
        }).start();
      } else {
        view.loadUrl(url);
      }
      return true;
    }
  }

  /**
   * Gets kuzzle web view client.
   *
   * @return the kuzzle web view client
   */
  public KuzzleWebViewClient getKuzzleWebViewClient() {
    return new KuzzleWebViewClient();
  }
%}


