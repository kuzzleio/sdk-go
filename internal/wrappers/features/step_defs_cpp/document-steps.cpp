#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {
  WHEN("^I create a document with id \'([^\"]*)\'$") {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {

    }
    catch (e) {
      
    }
  }

  THEN("^I get an error with message 'Document already exists'$") {

  }

}
