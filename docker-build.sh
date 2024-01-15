#!/bin/bash

go build -ldflags "-linkmode 'external' -extldflags '-static'"
docker build . -t tabbybox/ddns