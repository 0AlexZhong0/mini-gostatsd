package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Metric struct {
	path      string
	value     int64
	timestamp int64
}

// conform to the Graphite data path format, <path.subpath.more_sub> <key_value> <timestamp>
func (m *Metric) toText() string {
	return m.path + " " + fmt.Sprint(m.value) + " " + fmt.Sprint(m.timestamp)
}

type Stats struct {
	metrics []Metric
}

// add a metric to metrics in Stats
func (s *Stats) add(path string, value int64, timestamp int64) {
	s.metrics = append(s.metrics, Metric{path, value, timestamp})
}

// convert metrics to text
func (s *Stats) toText() string {
	result := ""

	for _, m := range s.metrics {
		result = result + m.toText() + "\n"
	}

	return result
}

const (
	graphiteHost    = "localhost"
	graphitePort    = "2003"
	// udp connection raises a write error
	connType        = "tcp"
	flushInterval   = 2 * time.Second
	globalNameSpace = "mini_gostatsd"
)

var (
	graphiteStats = make(map[string]int64)
)

// send stat in interval
func flushStat(conn net.Conn, interval time.Duration, stats *Stats) {
	for range time.Tick(interval) {
		log.Println("Sending metrics...")
		postStat(conn, stats)
	}
}

// write to Graphite
func postStat(conn net.Conn, stats *Stats) {
	currTimestamp := time.Now().Unix()

	var lastFlush int64

	if _, exists := graphiteStats["last_flush"]; !exists {
		lastFlush = 0
	} else {
		lastFlush = graphiteStats["last_flush"]
	}

	stats.add(globalNameSpace+".graphiteStats.last_flush", lastFlush, currTimestamp)
	graphiteStats["last_flush"] = time.Now().Unix()

	// send text data to graphite
	_, err := conn.Write([]byte(stats.toText()))

	if err != nil {
		log.Fatalln("Failed to write to backend", err)
	}
}

func main() {
	// establish the connection
	conn, err := net.Dial(connType, graphiteHost+":"+graphitePort)

	defer conn.Close()

	if err != nil {
		log.Fatalln("Connection failed", err)
	}

	stats := Stats{make([]Metric, 0)}

	for {
		// send aggregated metrics
		flushStat(conn, flushInterval, &stats)
	}
}
