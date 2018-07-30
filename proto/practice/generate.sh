#!/bin/sh

protoc -I proto/ --go_out=proto/ proto/simple/simple.proto
protoc -I proto/ --go_out=proto/ proto/enum/enum.proto
protoc -I proto/ --go_out=proto/ proto/complex/complex.proto
