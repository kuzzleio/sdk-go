#!/usr/bin/env python

import kcore
from kcore import Kuzzle

kuzzle = Kuzzle("localhost")

print(kuzzle)
print(kuzzle.server.now())
