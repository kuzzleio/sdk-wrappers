CC = gcc
CFLAGS = -fPIC -I/home/kblondel/Downloads/android-studio/jre/include/linux -I/home/kblondel/Downloads/android-studio/jre/include/
LDFLAGS = -L./
LIBS = -lgokcore
SRCS = kcore_wrap.c
OBJS = $(SRCS:.c=.o)
TARGET = libkcore.so

GOCC = /usr/lib/go-1.8/bin/go
GOFLAGS = -buildmode=c-shared
GOSRC = ./go/kuzzle/Kuzzle.go
GOTARGET = libgokcore.so

SWIG = swig
SWIGTEMPLATE = kcore.i

all: java

kcore_wrap.o: kcore_wrap.c
	$(CC) -c $< -o $@ $(CFLAGS) $(LDFLAGS) $(LIBS)

core:
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)

wrapper: $(OBJS)

object:
	gcc -shared kcore_wrap.o -o libkcore.so $(LDFLAGS) $(LIBS)

swigjava:
	$(SWIG) -java $(SWIGTEMPLATE)

java: 	core swigjava wrapper object

clean:
	rm -rf build *.class *.o *.h *.so *.java *.c *~ *.go

.PHONY: all node php java wrapper swigjava clean

.DEFAULT_GOAL := all
