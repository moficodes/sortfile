package main

import (
	"io"
	"os"
	"sort"
	"testing"
)

func BenchmarkReadInts(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		_, err := readInt(f)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadfile(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		_, err := readData(f)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteData(b *testing.B) {
	discard := io.Discard
	for i := 0; i < b.N; i++ {
		err := writeData(discard, []string{"01994844cb89783e", "fc423c81a94f4713"})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteInts(b *testing.B) {
	discard := io.Discard
	for i := 0; i < b.N; i++ {
		err := writeInts(discard, []uint64{18177157573609735955, 115202725784418366})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadSortWriteData(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	discard := io.Discard
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		data, err := readData(f)
		if err != nil {
			b.Fatal(err)
		}
		sort.Strings(data)
		err = writeData(discard, data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadSortWriteInts(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	discard := io.Discard
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		data, err := readInt(f)
		if err != nil {
			b.Fatal(err)
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		err = writeInts(discard, data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
