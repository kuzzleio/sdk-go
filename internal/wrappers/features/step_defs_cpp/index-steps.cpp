#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {

  GIVEN("^there is no index called \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, index_name);
    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->index->delete_(index_name);
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
      BOOST_FAIL(e.getMessage());
    }
  }

  GIVEN("^there is an index \'([^\"]*)\'$")
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

  GIVEN("^there is the indexes \'([^\"]*)\' and \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, index_name1);
    REGEX_PARAM(std::string, index_name2);
    ScenarioScope<KuzzleCtx> ctx;

    if (! ctx->kuzzle->index->exists(index_name1)) {
      try {
        ctx->kuzzle->index->create(index_name1);
      } catch(KuzzleException e) {
        K_LOG_E(e.getMessage().c_str());
        BOOST_FAIL(e.getMessage());
      }
    } else {
      K_LOG_D("Using existing index: %s", index_name1.c_str());
    }

    if (! ctx->kuzzle->index->exists(index_name2)) {
      try {
        ctx->kuzzle->index->create(index_name2);
      } catch(KuzzleException e) {
        K_LOG_E(e.getMessage().c_str());
        BOOST_FAIL(e.getMessage());
      }
    } else {
      K_LOG_D("Using existing index: %s", index_name2.c_str());
    }
  }

  WHEN("^I delete the indexes \'([^\"]*)\' and \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, index_name1);
    REGEX_PARAM(std::string, index_name2);
    ScenarioScope<KuzzleCtx> ctx;

    std::vector<std::string> v;

    v.push_back(index_name1);
    v.push_back(index_name2);

    try {
      ctx->kuzzle->index->mDelete(v);
    } catch(KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
      BOOST_FAIL(e.getMessage());
    }
  }

  WHEN("^I create an index called \'([^\"]*)\'$")
  {

    REGEX_PARAM(std::string, index_name);
    ScenarioScope<KuzzleCtx> ctx;

    K_LOG_D("Creating index: %s", index_name.c_str());
    try {
      ctx->kuzzle->index->create(index_name);
      ctx->index = index_name;
    } catch (KuzzleException e) {
      ctx->success = false;
      K_LOG_E(e.getMessage().c_str());
    }
  }

  THEN("^the index should exist$")
  {
    ScenarioScope<KuzzleCtx> ctx;
    BOOST_CHECK(ctx->kuzzle->index->exists(ctx->index));
  }

  THEN("^indexes \'([^\"]*)\' and \'([^\"]*)\' don't exist$")
  {
    REGEX_PARAM(std::string, index_name1);
    REGEX_PARAM(std::string, index_name2);
    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(
      !ctx->kuzzle->index->exists(index_name1) &&
      !ctx->kuzzle->index->exists(index_name2)
    );
  }

  WHEN("^I list indexes$")
  {
    ScenarioScope<KuzzleCtx> ctx;

    ctx->string_array = ctx->kuzzle->index->list();
  }

  THEN("^I get \'([^\"]*)\' and \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, index_name1);
    REGEX_PARAM(std::string, index_name2);

    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(std::find(ctx->string_array.begin(), ctx->string_array.end(), index_name1) != ctx->string_array.end());
    BOOST_CHECK(std::find(ctx->string_array.begin(), ctx->string_array.end(), index_name2) != ctx->string_array.end());
  }

}
