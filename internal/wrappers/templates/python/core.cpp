#include <exception>
#include <stdexcept>

extern "C" {
  #define _Complex
  #include "kuzzle.h"
  #include "kuzzlesdk.h"
  #include "swig.h"

  #include <stdio.h>
}

namespace BEV {
  struct Exception: std::exception
  { };

  struct MyException: Exception
  { };
}
