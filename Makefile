CODE_DIR=restServer
BIN_DIR=./bin
LDFLAGS += "-s -w"
BUILD_OPTIONS := -ldflags=$(LDFLAGS)

all: dep build

$(GOBIN)glide:
	@echo "Gathering GLIDE..."
	curl https://glide.sh/get | sh

dep: $(GOBIN)glide
	@echo "Gathering dependencies..."
	glide install

build:
	@echo "Building..."
	cd $(CODE_DIR) && CGO_ENABLED=0 GOOS=linux go build $(BUILD_OPTIONS) -o $(BIN_DIR)/linux/restServer -v
	cd $(CODE_DIR) && CGO_ENABLED=0 GOOS=darwin go build $(BUILD_OPTIONS) -o $(BIN_DIR)/osx/restServer -v
	cd $(CODE_DIR) && CGO_ENABLED=0 GOOS=windows go build $(BUILD_OPTIONS) -o $(BIN_DIR)/win/restServer.exe -v

clean:
	@echo "Deleting Binaries"
	cd $(CODE_DIR) && rm -rf $(BIN_DIR)/*