BINARY_NAME=lyricsMaker
setEnv:
	export CGO_CFLAGS_ALLOW=".*"
	export CGO_LDFLAGS_ALLOW=".*"
init:
ifeq ($(OS),Windows_NT)
else
ifeq ($(shell uname), Darwin)
	otool -L ./$(BINARY_NAME)
	install_name_tool -add_rpath $(shell pwd)/lib $(BINARY_NAME)
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
macOS:
	make prepare
	fyne package -os darwin -name $(BINARY_NAME)
	cp font.ttf $(BINARY_NAME).app/Contents/MacOS/
	mkdir -p $(BINARY_NAME).app/Contents/MacOS/lib
	cp -r ./lib/*.dylib $(BINARY_NAME).app/Contents/MacOS/lib
	install_name_tool -add_rpath ./lib $(BINARY_NAME).app/Contents/MacOS/src
	install_name_tool -add_rpath /Volumes/$(BINARY_NAME)/$(BINARY_NAME).app/Contents/MacOS/lib $(BINARY_NAME).app/Contents/MacOS/src
	install_name_tool -add_rpath /Applications/$(BINARY_NAME).app/Contents/MacOS/lib $(BINARY_NAME).app/Contents/MacOS/src
run:
	./$(BINARY_NAME)
clear:
	rm -rf ./$(BINARY_NAME)
clean:
	make clear
