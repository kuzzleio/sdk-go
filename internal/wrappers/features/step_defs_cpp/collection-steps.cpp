#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {
  WHEN("^I create a collection \'([^\"]*)\'$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->collection->create(ctx->index, collection_id);
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  THEN("^the collection \'([^\"]*)\' should exists$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->kuzzle->collection->exists(ctx->index, collection_id) == true);
  }

  WHEN("^I check if the collection \'([^\"]*)\' exists$") {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    ctx->success = ctx->kuzzle->collection->exists(ctx->index, collection_id);
  }

  THEN("^the collection should exist$") {
    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->success == true);
  }

  GIVEN("^it has a collection \'([^\"]*)\'$")
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

}
