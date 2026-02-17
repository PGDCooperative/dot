#!/bin/sh

go build -trimpath -o build/dot app/rl/*.go && cd build/ && ./dot && cd ..