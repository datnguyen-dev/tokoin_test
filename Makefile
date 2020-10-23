# Application name, version, build number
APP=tokoin_test
VERSION=0.0.1
# Using git commit hash to identify the build number
BUILD=`git rev-parse HEAD` 
BUILDDATE=`date +%FT%T%z`
REGISTRY_URL=github.com/datnguyen-dev
IMAGE_NAME=${REGISTRY_URL}/tokoin_test
IMAGE_FILE=${APP}_${VERSION}_${BUILD}

# Build flags
LDFLAGS=-ldflags "-X main.AppName=${APP} -X main.Version=${VERSION} -X main.BuildNum=${BUILD} -X main.BuildDate=${BUILDDATE}"

# Default target
.DEFAULT_GOAL: ${APP}

# Build application for linux arch
linux: clean
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/${APP}

# Build application for windows arch
windows: clean
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o dist/${APP}.exe

# Clean application
clean:
	go clean
	if [ -f dist/${APP} ]; then rm dist/${APP}; fi
	if [ -f dist/${APP}.exe ]; then rm dist/${APP}.exe; fi

.PHONY: clean install
