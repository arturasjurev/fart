BIN_NAME="fart"
BUILD_OUTPUT = "build/$(BIN_NAME)"
BUILD_VERSION = "v0.0.1"
BUILD_DATETIME = $(shell date +%Y%m%d-%H%M%S)
BUILD_TARGET = $(shell git rev-parse --short HEAD)



WINDOWS=$(BUILD_OUTPUT)_windows_amd64.exe
LINUX=$(BUILD_OUTPUT)_linux_amd64
DARWIN=$(BUILD_OUTPUT)_darwin_amd64

BUILD_FLAGS=-ldflags "-X main.Version=${BUILD_VERSION} -X main.BuildTarget=${BUILD_TARGET} -X main.BuildDate=${BUILD_DATETIME}"

build:
	@GOOS=linux GOARCH=386 go build -o ${BUILD_OUTPUT} \
		main.go

	@echo "Build results: ${BUILD_OUTPUT}"

 
build_all: windows linux darwin


windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) $(BUILD_FLAGS) main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) $(BUILD_FLAGS)  main.go