# UDP Packet Rate Limiter

## Problem

Take a continuous stream of timestamps on stdin

Example Input: 1074 1074 1076 1076 1076 1076

The input represents a timestamp at which a UDP packet was received from a particular customer.

For each input, respond with either a 'd' which instructs the router to DROP the packet, or an 'a' which instructs the router to ACCEPT the packet

Accept as many packets as possible without exceeding:

3 packets per 1 second

10 packets per 5 seconds

## Usage

```bash
./test.sh
```

## Source

https://leetcode.com/discuss/interview-question/1424264/cloudflare-oa
