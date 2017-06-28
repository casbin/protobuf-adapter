#!/usr/bin/env bash

protoc -I . policy.proto --go_out=.