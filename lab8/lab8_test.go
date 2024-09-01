package main

import (
	"os"
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	got, err := A("a\tb c\nd")
	if err != nil {
		t.Error(err)
	}
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("a() = %v, want %v", got, want)
	}
}

func TestA_regex(t *testing.T) {
	got, err := A("a\tb c\nd", " ")
	if err != nil {
		t.Error(err)
	}
	want := []string{"a\tb", "c\nd"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("a() = %v, want %v", got, want)
	}
}

func TestA_returns_error_when_too_few_args(t *testing.T) {
	got, err := A()
	if len(got) > 0 {
		t.Errorf("a() = %v, want {}", got)
	}
	if err == nil {
		t.Error(err)
	}
}

func TestA_returns_error_when_too_many_args(t *testing.T) {
	got, err := A("one", "two", "three")
	if len(got) > 0 {
		t.Errorf("a() = %v, want {}", got)
	}
	if err == nil {
		t.Error(err)
	}
}

func TestB(t *testing.T) {
	got, err := B("c\td a\nb")
	if err != nil {
		t.Error(err)
	}
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("b() = %v, want %v", got, want)
	}
}

func TestC(t *testing.T) {
	got := C("a\n\nb\n\n\n\n\n\n")
	want := 2
	if got != want {
		t.Errorf("c() = %v, want %v", got, want)
	}
}

func TestD(t *testing.T) {
	got := D("abb\tb c\nde")
	want := 4
	if got != want {
		t.Errorf("d() = %v, want %v", got, want)
	}
}

func TestE(t *testing.T) {
	got := E("a\tb c\ndefgh")
	want := 8
	if got != want {
		t.Errorf("e() = %v, want %v", got, want)
	}
}

func TestF(t *testing.T) {
	got := F("a\tb c\ndefgh a a a c defgh")
	want := map[string]int{
		"a":     4,
		"b":     1,
		"c":     2,
		"defgh": 2,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("f() = %v, want %v", got, want)
	}
}

func TestG(t *testing.T) {
	got := G("a ab dad abc aba gg abba")
	want := []string{"a", "aba", "abba", "dad", "gg"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("g() = %v, want %v", got, want)
	}
}

func BenchmarkB(b *testing.B) {
	// read txt files to string
	file_names := []string{"Latin-Lipsum_5.txt", "Latin-Lipsum_13.txt", "Latin-Lipsum_20.txt"}
	files := make([]string, len(file_names))
	for i, file_name := range file_names {
		file, _ := os.ReadFile(file_name)
		files[i] = string(file)
	}

	// benchmark
	// go test -bench=B -benchmem
	for i, input := range files {
		b.Run(file_names[i], func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				B(input)
			}
		})
	}
}

// go test
// go test -v -cover -coverprofile=c.out
// go tool cover -html=c.out
// go test -bench=B -benchmem
