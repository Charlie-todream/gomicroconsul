#!/bin/bash

cd Models/protos && protoc --micro_out=../ --go_out=../ *.proto && cd -