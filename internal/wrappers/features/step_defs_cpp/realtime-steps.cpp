#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {

  GIVEN("^I subscribe to (.*)$") {
    ScenarioScope<KuzzleCtx> ctx;

    CustomNotificationListener listener;

    try {
      ctx->kuzzle->realtime->subscribe(ctx->index, ctx->collection, "{}", &listener);
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  WHEN("^I create a document in (.*)$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    query_options options = {0};
    options.refresh = (char *)"wait_for";

    try {
      ctx->kuzzle->document->create(ctx->index, ctx->collection, "", "{\"foo\":\"bar\"}", &options);
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  THEN("^I receive a notification$") {
    ScenarioScope<KuzzleCtx> ctx;
    usleep(60000);
    BOOST_CHECK(ctx->notif_result != NULL);
  }
}