#/bin/bash
go build -o $lyricsMaker -v
install_name_tool -add_rpath ./lib lyricsMaker
./lyricsMaker
