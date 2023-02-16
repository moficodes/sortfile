package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var (
	file string
)

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
	flag.StringVar(&file, "file", "input.txt", "file to read")
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func MemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	message := fmt.Sprintf("alloc = %s, totalAlloc = %s, sys = %s, numGC = %v", humanReadableFilesize(int64(m.Alloc)), humanReadableFilesize(int64(m.TotalAlloc)), humanReadableFilesize(int64(m.Sys)), m.NumGC)
	return message
}

func humanReadableFilesize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "kMGTPE"[exp])
}

func readData(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, scanner.Err()
}

func readInt(r io.Reader) ([]uint64, error) {
	scanner := bufio.NewScanner(r)
	var res []uint64
	for scanner.Scan() {
		i, err := strconv.ParseUint(scanner.Text(), 16, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, scanner.Err()
}

func writeData(w io.Writer, data []string) error {
	for _, d := range data {
		_, err := w.Write([]byte(d + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeInts(w io.Writer, data []uint64) error {
	for _, d := range data {
		_, err := fmt.Fprintf(w, "%016x\n", d)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	input, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	data, err := readData(input)
	if err != nil {
		os.Exit(1)
	}

	sort.Strings(data)

	output, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	writeData(output, data)
}
