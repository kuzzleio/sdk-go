ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOTOUTDIR = $(ROOT_DIR)build
GOROOT ?= /usr/local
GOCC = $(GOROOT)/bin/go
JAVA_HOME ?= /usr/local
PYTHON_INCLUDE_DIR ?= /usr/include/python3.5
GOFLAGS = -buildmode=c-shared
GOSRC = ./cgo/kuzzle/
GOTARGETDIR = $(ROOTOUTDIR)/c
GOTARGET = $(GOTARGETDIR)/libkuzzlesdk.so
CPPDIR = $(ROOT_DIR)cpp
HEADERSDIR = $(ROOT_DIR)headers

CPP = g++
CPPFLAGS = -g -fPIC
INCS = -I$(HEADERSDIR) -I$(CPPDIR) -I$(ROOT_DIR)templates -I$(GOTARGETDIR)
LDFLAGS = -L$(GOTARGETDIR)
LIBS = -lkuzzlesdk -ljson-c
SRCS = kcore_wrap.cxx
OBJS = $(SRCS:.cxx=.o)

SWIG = swig

all: java python

kcore_wrap.o: kcore_wrap.cxx
	$(CPP) -ggdb -c $< -o $(OUTDIR)/$@ $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

core:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	mkdir -p $(GOTARGETDIR)
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	mv -f $(GOTARGETDIR)/libkuzzlesdk.h $(GOTARGETDIR)/kuzzle.h

wrapper: $(OBJS)

object:
	$(CPP) -ggdb -shared $(OUTDIR)/kcore_wrap.o -o $(OUTDIR)/libkuzzle-wrapper-$(LANGUAGE).so $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

posttask:
	cp -p $(GOTARGET) $(OUTDIR)
	rm -f $(OUTDIR)/kcore_wrap.o
	rm -f kcore_wrap.cxx

swig:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -c++ -$(LANGUAGE) $(SWIGOPTS) -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(LANGINCLUDE) templates/$(LANGUAGE)/core.i

java: OUTDIR = $(ROOTOUTDIR)/java/io/kuzzle/sdk
java: LANGINCLUDE = -I$(JAVA_HOME)/include -I$(JAVA_HOME)/include/linux
java: LANGUAGE = java
java: SWIGOPTS = -package io.kuzzle.sdk
java: core swig wrapper object posttask
	$(JAVA_HOME)/bin/javac $(OUTDIR)/*.java
	$(JAVA_HOME)/bin/jar cf $(OUTDIR)/kuzzlesdk.jar $(OUTDIR)/*.class
	rm -f $(OUTDIR)/*.java $(OUTDIR)/*.class

csharp: OUTDIR = $(ROOTOUTDIR)/csharp
csharp: LANGUAGE = csharp
csharp: SWIGOPTS = -dllimport kuzzlesdk
csharp: core swig wrapper object posttask
	mkdir -p $(OUTDIR)
	mcs -t:library $(OUTDIR)/*.cs -out:$(OUTDIR)/libkuzzle-csharp.so
	rm -f $(OUTDIR)/*.cs

python: OUTDIR = $(ROOTOUTDIR)/python
python: LANGINCLUDE = -I$(PYTHON_INCLUDE_DIR) -I$(ROOT_DIR)templates/$(LANGUAGE)
python: LANGUAGE = python
python: SWIGOPTS = -py3
python: CC = g++
python: CFLAGS = -fPIC
python: core swig wrapper
	mkdir -p $(OUTDIR)
	cp $(ROOT_DIR)templates/$(LANGUAGE)/*.py $(OUTDIR)
	cp $(ROOT_DIR)$(SRCS) $(OUTDIR)/
	cp $(ROOT_DIR)kcore_wrap.h $(OUTDIR)/
	python3 $(OUTDIR)/setup.py build_ext -I $(HEADERSDIR):$(CPPDIR):$(OUTDIR):$(GOTARGETDIR):$(OUTDIR)/templates/$(LANGUAGE) -L $(OUTDIR):$(GOTARGETDIR) -R $(GOTARGETDIR) -l kuzzlesdk -b $(OUTDIR) -t $(OUTDIR)/tmp


clean:
	rm -rf build

.PHONY: all java csharp python wrapper swig clean object core


.DEFAULT_GOAL := all
