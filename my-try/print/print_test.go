package print_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/print"
)

func TestPrintAnythingTo_PrintsStringToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	print.PrintAnythingTo(buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintAnythingTo_PrintsIntToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	print.PrintAnythingTo(buf, 123)
	want := "123\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintAnythingTo_PrintsFloatToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	print.PrintAnythingTo(buf, 123.22)
	want := "123.22\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
