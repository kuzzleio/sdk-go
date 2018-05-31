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

JAVAINCLUDE = -I$(JAVA_HOME)/include -I$(JAVA_HOME)/include/linux
PYTHONINCLUDE = -I$(PYTHON_INCLUDE_DIR) -I$(ROOT_DIR)templates/python

CPP_SDK_SRCS = kuzzle.cpp \
					document.cpp \
					collection.cpp \
					realtime.cpp \
					auth.cpp \
					server.cpp \
					index.cpp

CPPSRCS = $(CPP_SDK_SRCS) \
					kcore_wrap.cxx

CPPSDK = $(CPP_SDK_SRCS:.cpp=.o)

SWIG = swig

all: python java

kcore_wrap.o: kcore_wrap.cxx
	$(CPP) -c $< -o $@ $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)

%.o: cpp/%.cpp
	$(CPP) -fPIC -c $< -o $(OUTDIR)/$@ $(CPPFLAGS) -I./headers -I./build/c/ -L./build/c -lkuzzlesdk

core_java:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	mkdir -p $(GOTARGETDIR)
	cd ../../ && go get ./... ; cd -
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	$(GOCC) build -o $(GOTARGETSO) $(GOFLAGS) $(GOSRC)
	mv -f $(GOTARGETDIR)/libkuzzlesdk.h $(GOTARGETDIR)/kuzzle.h

core_python:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	mkdir -p $(GOTARGETDIR)
	cd ../../ && go get ./... ; cd -
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
	$(GOCC) build -o $(GOTARGETSO) $(GOFLAGS) $(GOSRC)
	mv -f $(GOTARGETDIR)/libkuzzlesdk.h $(GOTARGETDIR)/kuzzle.h

wrapper_java: $(OBJS)

wrapper_python: $(OBJS)

java_object:
	$(CPP) -shared kcore_wrap.o -o $(OUTDIR)/libkuzzle-wrapper-java.so $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(JAVAINCLUDE)
	strip $(OUTDIR)/libkuzzle-wrapper-java.so

python_object:
	$(CPP) -shared kcore_wrap.o -o $(OUTDIR)/libkuzzle-wrapper-python.so $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(PYTHONINCLUDE)
	strip $(OUTDIR)/libkuzzle-wrapper-python.so

swig:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -c++ -$(LANGUAGE) $(SWIGOPTS) -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(LANGINCLUDE) templates/$(LANGUAGE)/core.i

swig_java:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -c++ -java -package io.kuzzle.sdk -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(JAVAINCLUDE) templates/java/core.i

swig_python:
	mkdir -p $(OUTDIR)
	$(SWIG) -Wall -c++ -python -py3 -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(PYTHONINCLUDE) templates/python/core.i


makedir:
	mkdir -p $(OUTDIR)

java: LANGINCLUDE = $(JAVAINCLUDE)
java: OUTDIR = $(ROOTOUTDIR)/java/io/kuzzle/sdk
java: makedir core_java swig_java wrapper_java java_object
	#rm -rf $(OUTDIR)/kcore_wrap.o $(OUTDIR)/libkuzzlesdk.a $(OUTDIR)/libkuzzlesdk.so
	taskset -c 1 $(JAVA_HOME)/bin/javac $(OUTDIR)/*.java
	mkdir -p $(ROOTOUTDIR)/io/kuzzle/sdk/resources
	mv build/java/io/kuzzle/sdk/libkuzzle-wrapper-java.so build/io/kuzzle/sdk/resources
	mkdir -p $(ROOTOUTDIR)/java/src/main/java	
	ln -sf $(ROOTOUTDIR)/java/io/kuzzle/sdk/* $(ROOTOUTDIR)/java/src/main/java/
	cd build/java && taskset -c 1 sh gradlew sourcesJar jar javadocJar && cd -
	rm -f $(OUTDIR)/*.java $(OUTDIR)/*.class
	cp -p $(GOTARGET) $(OUTDIR)
	cp -p $(GOTARGETSO) $(OUTDIR)
	rm -f kcore_wrap.o
	rm -f kcore_wrap.cxx

csharp: OUTDIR = $(ROOTOUTDIR)/csharp
csharp: LANGUAGE = csharp
csharp: SWIGOPTS = -dllimport kuzzlesdk $(OBJS)
csharp: core swig wrapper object posttask makedir
	mcs -t:library $(OUTDIR)/*.cs -out:$(OUTDIR)/libkuzzle-csharp.so
	rm -f $(OUTDIR)/*.cs

python: LANGINCLUDE = $(PYTHONINCLUDE)
python: OUTDIR = $(ROOTOUTDIR)/python
python: CC = g++
python: CFLAGS = -fPIC
python: makedir core_python swig_python wrapper_python python_object
	cp $(ROOT_DIR)templates/python/*.py $(OUTDIR)
	cp $(ROOT_DIR)$(SRCS) $(OUTDIR)/
	cp $(ROOT_DIR)kcore_wrap.h $(OUTDIR)/
	python3 $(OUTDIR)/setup.py build_ext -I $(HEADERSDIR):$(CPPDIR):$(OUTDIR):$(GOTARGETDIR):$(OUTDIR)/templates/python -L $(OUTDIR):$(GOTARGETDIR) -R $(GOTARGETDIR) -l kuzzlesdk  -b $(OUTDIR) -t $(OUTDIR)/tmp
	cp -p $(GOTARGET) $(OUTDIR)
	cp -p $(GOTARGETSO) $(OUTDIR)
	rm -f kcore_wrap.o
	rm -f kcore_wrap.cxx

cpp: OUTDIR = $(ROOTOUTDIR)/cpp
cpp: makedir $(CPPSDK)
cpp: LANGUAGE = c++
cpp: core
		ar rvs $(OUTDIR)/libcpp.a $(OUTDIR)/*.o
		g++ -shared -fPIC -o $(OUTDIR)/libcpp.so -Wl,--whole-archive $(OUTDIR)/libcpp.a build/c/libkuzzlesdk.a -Wl,--no-whole-archive

clean:
	rm -rf build/c build/java/src build/java/io build/java/build kcore_wrap.cxx kcore_wrap.h
	rm -rf build/io

.PHONY: all java csharp python wrapper swig clean object core


.DEFAULT_GOAL := all
