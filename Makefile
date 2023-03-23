BUILD_DIR = $(PWD)/build
APP_NAME = apiserver

clean:
	@echo "Removing older builders"
	@rm -rf ./build
all:
	@chmod +x $(BUILD_DIR)/$(APP_NAME)
test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./test
	go tool cover -func=cover.out
build: clean test
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BUILD_DIR}/${APP_NAME} main.go
run: build all
	$(BUILD_DIR)/$(APP_NAME)