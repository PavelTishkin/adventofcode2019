package main

import "testing"

func TestFftFeedback(t *testing.T) {
	input := stringToNumArray("80871224585914546619083218645595")
	pattern := []int64{0, 1, 0, -1}
	got := numArrayToString(fftFeedback(input, pattern, 100))
	if got[:8] != "24176176" {
		t.Errorf("fftFeedback first 8 digits are %s, expected 24176176", got[:8])
	}

	input = stringToNumArray("19617804207202209144916044189917")
	got = numArrayToString(fftFeedback(input, pattern, 100))
	if got[:8] != "73745418" {
		t.Errorf("fftFeedback first 8 digits are %s, expected 73745418", got[:8])
	}

	input = stringToNumArray("69317163492948606335995924319873")
	got = numArrayToString(fftFeedback(input, pattern, 100))
	if got[:8] != "52432133" {
		t.Errorf("fftFeedback first 8 digits are %s, expected 52432133", got[:8])
	}
}

func TestFft(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	expandedPattern := expandPattern([]int64{0, 1, 0, -1}, len(input))
	got := fft(input, expandedPattern)
	expected := []int64{4, 8, 2, 2, 6, 1, 5, 8}
	if !int64SlicesEqual(got, expected) {
		t.Errorf("fft = %v, expected %v", got, expected)
	}
	got = fft(got, expandedPattern)
	expected = []int64{3, 4, 0, 4, 0, 4, 3, 8}
	if !int64SlicesEqual(got, expected) {
		t.Errorf("fft = %v, expected %v", got, expected)
	}
}

func TestMulInputByPattern(t *testing.T) {
	input := []int64{9, 8, 7, 6, 5}
	pattern := []int64{1, 2, 3}
	got := mulInputByPattern(input, pattern)
	expected := int64(2)
	if got != expected {
		t.Errorf("mulInputByPattern = %d, expected = %d", got, expected)
	}
}

func TestExpandPattern(t *testing.T) {
	pattern := []int64{0, 1, 0, -1}
	got := expandPattern(pattern, 12)
	if len(got) != 12 {
		t.Errorf("len(expandPattern) = %d, expected 12", len(got))
	}
	expected := []int64{1, 0, -1, 0, 1, 0, -1, 0, 1, 0, -1, 0}
	if !int64SlicesEqual(got[0], expected) {
		t.Errorf("expandPattern[0] = %v, expected %v", got[0], expected)
	}
	expected = []int64{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0}
	if !int64SlicesEqual(got[1], expected) {
		t.Errorf("expandPattern[1] = %v, expected %v", got[1], expected)
	}
	expected = []int64{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1, 0}
	if !int64SlicesEqual(got[2], expected) {
		t.Errorf("expandPattern[2] = %v, expected %v", got[2], expected)
	}
	expected = []int64{0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, -1}
	if !int64SlicesEqual(got[3], expected) {
		t.Errorf("expandPattern[3] = %v, expected %v", got[3], expected)
	}
}

func TestStringToNumArray(t *testing.T) {
	got := stringToNumArray("12345")
	expected := []int64{1, 2, 3, 4, 5}

	if !int64SlicesEqual(got, expected) {
		t.Errorf("stringToNumArray = %v, expected %v", got, expected)
	}

	got = stringToNumArray("1")
	expected = []int64{1}

	if !int64SlicesEqual(got, expected) {
		t.Errorf("stringToNumArray = %v, expected %v", got, expected)
	}

	got = stringToNumArray("")

	if len(got) != 0 {
		t.Errorf("len(stringToNumArray) = %d, expected 0", len(got))
	}
}

func int64SlicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
