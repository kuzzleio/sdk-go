ROOTOUTDIR = ./build
GOROOT ?= /usr/local
JAVAHOME ?= /usr/local
GOCC = $(GOROOT)/bin/go
GOFLAGS = -buildmode=c-shared
GOSRC = ./cgo/kuzzle/
GOTARGETDIR = $(ROOTOUTDIR)/c
GOTARGET = $(GOTARGETDIR)/libkuzzlesdk.so

CC = gcc
CFLAGS = -std=c99 -fPIC 
INCS = -I$(PWD)/headers -I$(GOTARGETDIR)
LDFLAGS = -L$(GOTARGETDIR)
LIBS = -lkuzzlesdk -ljson-c
SRCS = kcore_wrap.c
OBJS = $(SRCS:.c=.o)

SWIG = swig

all: java

kcore_wrap.o: kcore_wrap.c
	$(CC) -ggdb -c $< -o $(OUTDIR)/$@ $(CFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE) 

core:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	mkdir -p $(GOTARGETDIR)
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	mv -f $(GOTARGETDIR)/libkuzzlesdk.h $(GOTARGETDIR)/kuzzle.h

wrapper: $(OBJS)

object:
	gcc -ggdb -shared $(OUTDIR)/kcore_wrap.o -o $(OUTDIR)/libkuzzle-wrapper-$(LANGUAGE).so $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

posttask:
	cp -p $(GOTARGET) $(OUTDIR)
	rm -f $(OUTDIR)/kcore_wrap.o
	rm -f kcore_wrap.c

swig:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -$(LANGUAGE) $(SWIGOPTS) -outdir $(OUTDIR) -o kcore_wrap.c $(INCS) $(LANGINCLUDE) templates/$(LANGUAGE)/core.i

java: OUTDIR = $(ROOTOUTDIR)/java/io/kuzzle/sdk
java: LANGINCLUDE = -I$(JAVAHOME)/include -I$(JAVAHOME)/include/linux
java: LANGUAGE = java
java: SWIGOPTS = -package io.kuzzle.sdk
java: core swig wrapper object posttask
	$(JAVAHOME)/bin/javac $(OUTDIR)/*.java
	$(JAVAHOME)/bin/jar cf $(OUTDIR)/kuzzlesdk.jar $(OUTDIR)/*.class
	rm -f $(OUTDIR)/*.java $(OUTDIR)/*.class

csharp: OUTDIR = $(ROOTOUTDIR)/csharp
csharp: LANGUAGE = csharp
csharp: SWIGOPTS = -dllimport kuzzlesdk
csharp: core swig wrapper object posttask
	mkdir -p $(OUTDIR)
	mcs -t:library $(OUTDIR)/*.cs -out:$(OUTDIR)/libkuzzle-csharp.so
	rm -f $(OUTDIR)/*.cs

clean:
	rm -rf build

.PHONY: all java csharp wrapper swig clean object core

.DEFAULT_GOAL := all
