package main

import "testing"

func TestGetFileNameWithoutExtension(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"cfg.json", "cfg"},
		{"cfg.example.json", "cfg.example"},
		{"cfg", "cfg"},
	}
	for i, v := range tests {
		actual := getFileNameWithoutExtension(v.input)
		expected := v.expected
		t.Logf("Check case %d: %s(actual) == %s(expected)", i, actual, expected)
		if actual != expected {
			t.Errorf("Error on case %d: %s(actual) != %s(expected)", i, actual, expected)
		}

	}
}
