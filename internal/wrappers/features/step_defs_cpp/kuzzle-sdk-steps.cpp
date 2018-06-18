#include "steps.hpp"

// Anonymous namespace to handle a compile error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {
  BEFORE() { kuz_log_sep(); }

  GIVEN("^I create a user 'useradmin' with password 'testpwd' with id 'useradmin-id'$") { pending(); }

  WHEN("^I try to create a document with id '(my-document)'$")
  {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->document->create(ctx->index, ctx->collection, document_id, "{\"a\":\"document\"}");
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  GIVEN("^I update my user custom data with the pair ([\\w]+):(.+)$")
  {
    REGEX_PARAM(std::string, fieldname);
    REGEX_PARAM(std::string, fieldvalue);

    ScenarioScope<KuzzleCtx> ctx;

    string data = "{\"" + fieldname + "\":" + fieldvalue + "}";

    K_LOG_D("Updating user data with : %s", data.c_str());

    try {
      ctx->kuzzle->auth->updateSelf(data);
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
    }
  }

  THEN("^the response 'content' field contains the pair ([\\w]+):(.+)$")
  {
    K_LOG_I("Checking user content field");
    REGEX_PARAM(std::string, fieldname);
    REGEX_PARAM(std::string, expected_fieldvalue);

    ScenarioScope<KuzzleCtx> ctx;

    json_spirit::Value userContentValue;
    std::string        userContent = ctx->currentUser->content;
    json_spirit::read(userContent, userContentValue);

    json_spirit::write_formatted(userContentValue);

    json_spirit::Value fieldvalue = json_spirit::find_value(userContentValue.get_obj(), fieldname);
    switch (fieldvalue.type()) {
      case json_spirit::str_type: {
        std::string s = fieldvalue.get_str();
        K_LOG_D("Field value: \"%s\" of type string", s.c_str());
        BOOST_CHECK("\"" + s + "\"" == expected_fieldvalue);
        break;
      }
      case json_spirit::bool_type: {
        auto   b = fieldvalue.get_bool();
        string s = b ? "true" : "false";
        K_LOG_D("Field value: \"%s\" of type bool", b ? "true" : "false");
        BOOST_CHECK(s == expected_fieldvalue);
        break;
      }
      case json_spirit::int_type: {
        auto i = fieldvalue.get_int();
        K_LOG_D("Field value: %d of type int", i);
        string s = std::to_string(i);
        BOOST_CHECK(s == expected_fieldvalue);
        break;
      }
      case json_spirit::real_type: {
        float f = fieldvalue.get_real();
        K_LOG_D("Field value: %f of type real", f);
        float e = std::stof(expected_fieldvalue);
        K_LOG_D("Expected value: %f", e);
        BOOST_CHECK(f == std::stof(expected_fieldvalue));
        break;
      }
        // TODO: Add obj test case...
    }

    ctx->customUserDataType = fieldvalue.type();
  }

  THEN("^is a string$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->customUserDataType == json_spirit::str_type);
  }

  THEN("^is a number$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->customUserDataType == json_spirit::int_type || ctx->customUserDataType == json_spirit::real_type);
  }

  THEN("^is a bool$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->customUserDataType == json_spirit::bool_type);
  }

  GIVEN("^Kuzzle Server is running$")
  {
    K_LOG_I("Connecting to Kuzzle on 'localhost:7512'");
    ScenarioScope<KuzzleCtx> ctx;
    ctx->kuzzle_options         = KUZZLE_OPTIONS_DEFAULT;
    ctx->kuzzle_options.connect = MANUAL;
    try {
      ctx->kuzzle = new Kuzzle("localhost", &ctx->kuzzle_options);
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
    }
    char* error = ctx->kuzzle->connect();
    BOOST_CHECK(error == NULL);
  }

  GIVEN("^there is an index '(test-index)'$")
  {
    REGEX_PARAM(std::string, index_name);
    ScenarioScope<KuzzleCtx> ctx;
    ctx->index = index_name;

    if (!ctx->kuzzle->index->exists(index_name)) {
      K_LOG_D("Creating index: %s", index_name.c_str());
      try {
        ctx->kuzzle->index->create(index_name);
      } catch (KuzzleException e) {
        K_LOG_E(e.getMessage().c_str());
        BOOST_FAIL(e.getMessage());
      }
    } else {
      K_LOG_D("Using existing index: %s", index_name.c_str());
    }
  }

  GIVEN("^it has a collection '(test-collection)'$")
  {
    REGEX_PARAM(std::string, collection_name);
    ScenarioScope<KuzzleCtx> ctx;
    ctx->collection = collection_name;

    K_LOG_D("Creating collection: %s", collection_name.c_str());
    try {
      ctx->kuzzle->collection->create(ctx->index, ctx->collection);
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
      BOOST_FAIL(e.getMessage());
    }
  }

  GIVEN("^the collection has a document with id '(my-document-id)'$")
  {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->document->create(ctx->index, ctx->collection, document_id, "{\"a\":\"document\"}");
    } catch (KuzzleException e) {
      e.getMessage();
    }
  }

  GIVEN("^the collection doesn't have a document with id '(my-document-id)'$")
  {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->document->delete_(ctx->index, ctx->collection, document_id);
    } catch (KuzzleException e) {
    }
  }

  WHEN("^I try to create a new document with id 'my-document-id'$")
  {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;
    ctx->success = true;
    try {
      ctx->kuzzle->document->create(ctx->index, ctx->collection, document_id, "{\"a\":\"document\"}");
    } catch (KuzzleException e) {
      ctx->success = true;
    }
  }

  THEN("^I get an error with message 'document alread exists'$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->success == false);
  }

  GIVEN("^'test-index'/'test-collection' does't have a document with id "
        "'my-document-id'$")
  {
    pending();
  }

  THEN("^the document is successfully created$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->success == false);
  }

  GIVEN("^there is an user with id '([\\w\\-]+)'$")
  {
    REGEX_PARAM(std::string, user_id);
    ScenarioScope<KuzzleCtx> ctx;
    ctx->user_id = user_id;
  }

  GIVEN("^the user has 'local' credentials with name '([\\w\\-]+)' and password "
        "'([\\w\\-]+)'$")
  {
    REGEX_PARAM(std::string, username);
    REGEX_PARAM(std::string, password);
    ScenarioScope<KuzzleCtx> ctx;

    kuzzle_user_create(ctx->kuzzle, ctx->user_id, username, password);
  }

  WHEN("^I log in as '([\\w\\-]+)':'([\\w\\-]+)'$")
  {
    REGEX_PARAM(std::string, username);
    REGEX_PARAM(std::string, password);
    ScenarioScope<KuzzleCtx> ctx;

    string jwt;
    try {
      jwt = ctx->kuzzle->auth->login("local", get_login_creds(username, password));
      K_LOG_D("Logged in as '%s'", username.c_str());
      K_LOG_D("JWT is: %s", jwt.c_str());
    } catch (KuzzleException e) {
      K_LOG_W(e.getMessage().c_str());
    }
    ctx->jwt = jwt;
  }

  THEN("^the retrieved JWT is valid$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    token_validity*          v = ctx->kuzzle->auth->checkToken(ctx->jwt);
    BOOST_CHECK(v->valid);
  }

  THEN("^the retrieved JWT is invalid$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    token_validity*          v = ctx->kuzzle->auth->checkToken(ctx->jwt);
    BOOST_CHECK(!v->valid);
  }
  WHEN("^I logout$") {
      ScenarioScope<KuzzleCtx> ctx;
      ctx->kuzzle->auth->logout();
  }

  WHEN("^I get my user info$")
  {
    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->currentUser = ctx->kuzzle->auth->getCurrentUser();
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
    }

    K_LOG_D("current user = 0x%p", ctx->currentUser);
    K_LOG_D("Current user content: %s", ctx->currentUser->content);

    BOOST_CHECK_MESSAGE(ctx->currentUser != NULL, "Failed to retrieve current user");
  }

}