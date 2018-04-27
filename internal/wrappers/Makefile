ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOTOUTDIR = $(ROOT_DIR)build
GOROOT ?= /usr/local/go
GOCC = $(GOROOT)/bin/go
JAVA_HOME ?= /usr/local
PYTHON_INCLUDE_DIR ?= /usr/include/python3.5
GOFLAGS = -buildmode=c-archive
GOSRC = ./cgo/kuzzle/
GOTARGETDIR = $(ROOTOUTDIR)/c
GOTARGET = $(GOTARGETDIR)/libkuzzlesdk.a
GOTARGETSO = $(GOTARGETDIR)/libkuzzlesdk.so
CPPDIR = $(ROOT_DIR)cpp
HEADERSDIR = $(ROOT_DIR)headers

CPP = $(CXX)
CPPFLAGS = -g -fPIC -std=c++11
INCS = -I$(HEADERSDIR) -I$(CPPDIR) -I$(ROOT_DIR)templates -I$(GOTARGETDIR)
LDFLAGS = -L$(GOTARGETDIR)
LIBS = -lkuzzlesdk
SRCS = kcore_wrap.cxx
OBJS = $(SRCS:.cxx=.o)

CPPSRCS = 	kuzzle.cpp \
					document.cpp \
					collection.cpp \
					room.cpp \
					kcore_wrap.cxx

CPPSDK = $(CPPSRCS:.cpp=.o)

SWIG = swig

all: java

kcore_wrap.o: kcore_wrap.cxx
	$(CPP) -ggdb -c $< -o $(OUTDIR)/$@ $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

%.o: cpp/%.cpp
	$(CPP) -fPIC -c $< -o $(OUTDIR)/$@ -I./headers -I./build/c/ -L./build/c -lkuzzlesdk

core:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	mkdir -p $(GOTARGETDIR)
	cd ../../ && go get ./... ; cd -
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	$(GOCC) build -o $(GOTARGETSO) $(GOFLAGS) $(GOSRC)
	mv -f $(GOTARGETDIR)/libkuzzlesdk.h $(GOTARGETDIR)/kuzzle.h

wrapper: $(OBJS)

object:
	$(CPP) -ggdb -shared $(OUTDIR)/kcore_wrap.o -o $(OUTDIR)/libkuzzle-wrapper-$(LANGUAGE).so $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

posttask:
	cp -p $(GOTARGET) $(OUTDIR)
	cp -p $(GOTARGETSO) $(OUTDIR)
	#rm -f $(OUTDIR)/kcore_wrap.o
	#rm -f kcore_wrap.cxx

swig:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -c++ -$(LANGUAGE) $(SWIGOPTS) -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(LANGINCLUDE) templates/$(LANGUAGE)/core.i

makedir:
	mkdir -p $(OUTDIR)

java: OUTDIR = $(ROOTOUTDIR)/java/io/kuzzle/sdk
java: LANGINCLUDE = -I$(JAVA_HOME)/include -I$(JAVA_HOME)/include/linux
java: LANGUAGE = java
java: SWIGOPTS = -package io.kuzzle.sdk
java: core swig wrapper object posttask
	rm -rf $(OUTDIR)/kcore_wrap.o $(OUTDIR)/libkuzzlesdk.a $(OUTDIR)/libkuzzlesdk.so
	taskset -c 1 $(JAVA_HOME)/bin/javac $(OUTDIR)/*.java
	cd build/java
	mkdir -p $(ROOTOUTDIR)/java/src/main/java
	ln -sf $(ROOTOUTDIR)/java/io/kuzzle/sdk/* $(ROOTOUTDIR)/java/src/main/java/
	cd build/java && taskset -c 1 ./gradlew sourcesJar javadocJar jar && cd -
	#rm -f $(OUTDIR)/*.java $(OUTDIR)/*.class

csharp: OUTDIR = $(ROOTOUTDIR)/csharp
csharp: LANGUAGE = csharp
csharp: SWIGOPTS = -dllimport kuzzlesdk
csharp: core swig wrapper object posttask makedir
	mcs -t:library $(OUTDIR)/*.cs -out:$(OUTDIR)/libkuzzle-csharp.so
	rm -f $(OUTDIR)/*.cs

python: OUTDIR = $(ROOTOUTDIR)/python
python: LANGINCLUDE = -I$(PYTHON_INCLUDE_DIR) -I$(ROOT_DIR)templates/$(LANGUAGE)
python: LANGUAGE = python
python: SWIGOPTS = -py3
python: CC = g++
python: CFLAGS = -fPIC
python: core swig wrapper makedir
	cp $(ROOT_DIR)templates/$(LANGUAGE)/*.py $(OUTDIR)
	cp $(ROOT_DIR)$(SRCS) $(OUTDIR)/
	cp $(ROOT_DIR)kcore_wrap.h $(OUTDIR)/
	python3 $(OUTDIR)/setup.py build_ext -I $(HEADERSDIR):$(CPPDIR):$(OUTDIR):$(GOTARGETDIR):$(OUTDIR)/templates/$(LANGUAGE) -L $(OUTDIR):$(GOTARGETDIR) -R $(GOTARGETDIR) -l kuzzlesdk -b $(OUTDIR) -t $(OUTDIR)/tmp

cpp: OUTDIR = $(ROOTOUTDIR)/cpp
cpp: makedir $(CPPSDK)
		ar rvs $(OUTDIR)/libcpp.a $(OUTDIR)/*.o
		g++ -shared -fPIC -o $(OUTDIR)/libcpp.so -Wl,--whole-archive $(OUTDIR)/libcpp.a build/c/libkuzzlesdk.a -Wl,--no-whole-archive

clean:
	rm -rf build/c build/java/src build/java/io build/java/build

.PHONY: all java csharp python wrapper swig clean object core


.DEFAULT_GOAL := all
