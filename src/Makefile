BINARY_NAME=lyrisMaker
setEnv:
	export CGO_CFLAGS_ALLOW=".*"
	export CGO_LDFLAGS_ALLOW=".*"
init:
ifeq ($(OS),Windows_NT)
else
ifeq ($(shell uname), Darwin)
	otool -L ./$(BINARY_NAME)
	install_name_tool -add_rpath ./lib $(BINARY_NAME)
else
	PLATFORM = "Unix/Linux"
	@echo $(PLATFORM)
	echo $$(uname)
endif
endif
prepare:
	make setEnv clear
test:
	make prepare
	go build -o $(BINARY_NAME) -v
	make init
build:
	make prepare
	go build -o $(BINARY_NAME) -v -ldflags="-s -w"
	make init
run:
	./$(BINARY_NAME)
clear:
	rm -rf ./$(BINARY_NAME)
clean:
	make clear
