BUILD_DIR = build

.PHONY: build test clean

build:
	@cmake -S . -B $(BUILD_DIR)
	@cmake --build $(BUILD_DIR)

test: build
	@cd $(BUILD_DIR) && ctest --output-on-failure

clean:
	@rm -rf $(BUILD_DIR)
