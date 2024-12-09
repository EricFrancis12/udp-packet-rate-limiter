#!/bin/bash

go build -o udp-packet-rate-limiter

input_stream="1074 1074 1076 1076 1076 1076 1090"

# Echo the input stream and pipe it into the compiled Go program
echo $input_stream | ./udp-packet-rate-limiter
