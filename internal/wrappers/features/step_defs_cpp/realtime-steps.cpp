#include "steps.hpp"

GIVEN("^I subscribe to 'test-collection'$") {
  REGEX_PARAM(std::string, collection_id);

  ScenarioScope<KuzzleCtx> ctx;

  try {
    ctx->kuzzle->realtime->subscribe(ctx->index, ctx->collection, "{}", NULL, NULL);
  } catch (KuzzleException e) {
    BOOST_FAIL(e.getMessage());
  }
}