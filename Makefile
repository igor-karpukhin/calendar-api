SERVICE := calendar-api
PACKAGE := github.com/igor-karpukhin/${SERVICE}
COMMITMSG := ${shell git log -1 --pretty=%B}
VERSION := ${shell git describe --tags --always}
COMMIT := ${shell git rev-parse HEAD}
BUILDTIME := ${shell date -u '+%Y-%m-%d_%H:%M:%S'}
LDFLAGS := -s -w -X '${PACKAGE}/version.Version=${VERSION}' \
					-X '${PACKAGE}/version.BuildTime=${BUILDTIME}' \
					-X '${PACKAGE}/version.Commit=${COMMIT}' \
					-X '${PACKAGE}/version.CommitMsg=${COMMITMSG}'
ifdef OSX
	TARGET_OS=darwin
else
	TARGET_OS=linux
endif

.PHONY: clean build all

all: clean build

build:
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=$(TARGET_OS) go build -ldflags "${LDFLAGS}" -a -o bin/${SERVICE}

clean:
	rm -rf bin/

