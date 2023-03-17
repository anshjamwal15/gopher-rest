BUILD_DIR = $(PWD)/build
APP_NAME = apiserver

build:
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BUILD_DIR}/${APP_NAME} main.go
all:
	@chmod +x $(BUILD_DIR)/$(APP_NAME)
run: build all
	$(BUILD_DIR)/$(APP_NAME)
remove:
	@echo "Removing older builders"
	@rm -rf build/
rebuild: remove run
	@$(BUILD_DIR)/$(APP_NAME)
	