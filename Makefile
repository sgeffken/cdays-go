PROJECT?=github.com/sgeffken/cdays-go
BUILD_PATH?=cmd/cdays

build: test
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Release=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
	-o ./bin/cdays ${PROJECT}/${BUILD_PATH}

test:
	go test -race ./...
