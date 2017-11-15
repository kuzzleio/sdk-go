CC = gcc
CFLAGS = -std=c99 -fPIC -I$(PWD)/headers -I${JAVA_HOME}/include -I${JAVA_HOME}/include/linux
LDFLAGS = -L./
LIBS = -lkuzzlesdk -ljson-c
SRCS = kcore_wrap.c
OBJS = $(SRCS:.c=.o)
TARGET = libkcore.so

GOROOT ?= /usr/local
GOCC = $(GOROOT)/bin/go
GOFLAGS = -buildmode=c-shared
GOSRC = ./cgo/kuzzle/
GOTARGET = libkuzzlesdk.so

SWIG = swig

all: java

kcore_wrap.o: kcore_wrap.c
	$(CC) -ggdb -c $< -o $@ $(CFLAGS) $(LDFLAGS) $(LIBS)

core:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	mv -f libkuzzlesdk.h kuzzle.h

wrapper: $(OBJS)

object:
	gcc -ggdb -shared kcore_wrap.o -o libkuzzle.so $(LDFLAGS) $(LIBS)

swigjava:
	$(SWIG) -Wall -java -package io.kuzzle.sdk -outdir ./io/kuzzle/sdk -o kcore_wrap.c -I/usr/local/include templates/java/core.i

java: 	core swigjava wrapper object

clean:
	rm -rf build *.class *.o *.h *.so io/kuzzle/sdk/*.java io/kuzzle/sdk/*.class *.c *~ *.go

.PHONY: all java wrapper swigjava clean object core

.DEFAULT_GOAL := all
