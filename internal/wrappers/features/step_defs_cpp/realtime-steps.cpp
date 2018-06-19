#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {

  GIVEN("^I subscribe to 'test-collection'$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->realtime->subscribe(ctx->index, ctx->collection, "{}", NULL);
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  WHEN("^I create a document in \"test-collection\"$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->document->create(ctx->index, ctx->collection, "", "{\"foo\":\"bar\"}");
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }


}