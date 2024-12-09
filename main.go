package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RateLimiter struct {
	timestamps map[int]uint64
	limits     [][2]uint64
}

func newRateLimiter(limits ...[2]uint64) RateLimiter {
	return RateLimiter{
		timestamps: make(map[int]uint64),
		limits:     limits,
	}
}

func (r *RateLimiter) insert(timestamp int) {
	r.timestamps[timestamp]++
}

func (r *RateLimiter) wouldExceedLimit(timestamp int, maxPackets uint64, seconds uint64) bool {
	secs := int(seconds)

	for i := 0; i < secs; i++ {
		var packetCount uint64 = 0
		for j := 0; j < secs; j++ {
			packetCount += r.timestamps[timestamp+i+j-secs+1]
		}

		if packetCount >= maxPackets {
			return true
		}
	}

	return false
}

func (r *RateLimiter) isAllowed(timestamp int) bool {
	for _, limit := range r.limits {
		if r.wouldExceedLimit(timestamp, limit[0], limit[1]) {
			return false
		}
	}

	r.insert(timestamp)
	return true
}

func withPacketsPerSec(maxPackets uint64, seconds uint64) [2]uint64 {
	return [2]uint64{maxPackets, seconds}
}

func normalizeWhitespace(s string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
}

func main() {
	r := newRateLimiter(
		withPacketsPerSec(3, 1),
		withPacketsPerSec(10, 5),
	)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := normalizeWhitespace(scanner.Text())
		parts := strings.Split(text, " ")

		var timestamps []int
		for _, part := range parts {
			timestamp, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			timestamps = append(timestamps, timestamp)
		}

		for _, timestamp := range timestamps {
			if r.isAllowed(timestamp) {
				os.Stdout.Write([]byte("a"))
			} else {
				os.Stdout.Write([]byte("d"))
			}
		}
	}
}
