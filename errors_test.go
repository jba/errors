package errors

import (
	"io"
	"testing"

	"google.golang.org/grpc/codes"
)

func TestErr(t *testing.T) {
	err := Err{
		Space:   GRPCSpace,
		Code:    int(codes.NotFound),
		Source:  io.EOF,
		Details: "details",
	}.Printf("%s %d", "a", 1)
	if got, want := err.Message, "a 1"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := err.ErrorSource(), io.EOF; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := err.ErrorDetails(), "details"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	space, code := err.ErrorCode()
	if got, want := space, GRPCSpace; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := code, int(codes.NotFound); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
