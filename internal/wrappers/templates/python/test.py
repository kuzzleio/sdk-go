#!/usr/bin/env python

import kcore
from kcore import Kuzzle

kuzzle = Kuzzle("localhost:7512")

print(kuzzle)
print(kuzzle.now())
