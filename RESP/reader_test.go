package resp

import "testing"

func TestReader(t *testing.T) {
	input := "$5\r\nAhmed\r\n"
	result := reader(input)
	expected := "Ahmed"
	if result != expected {
		t.Errorf("reader(\"$5\r\nAhmed\\r\\n\") = %s; want %s", result, expected)
	}
}
