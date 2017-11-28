#include "kuzzle.hpp"

namespace kuzzleio {

  KuzzleError::KuzzleError(int status, const std::string& error, const std::string& stack)
    : std::runtime_error(error) {
    this->status = status;
    this->stack = stack;
  }

}
