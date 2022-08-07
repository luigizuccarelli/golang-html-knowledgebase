#!/bin/bash

SERVER_PORT=3020
NAME=golang-html-knowledgebase
VERSION=1.0.1
LOG_LEVEL=trace

export SERVER_PORT NAME VERSION LOG_LEVEL

./build/microservice

