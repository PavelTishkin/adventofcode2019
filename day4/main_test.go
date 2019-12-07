package main

import "testing"

func TestHasTrueDouble(t *testing.T) {
	strArr := numToStrArray(112233)
	got := hasTrueDouble(strArr)
	if !got {
		t.Error("hasTrueDouble = false; want true")
	}

	strArr = numToStrArray(123444)
	got = hasTrueDouble(strArr)
	if got {
		t.Error("hasTrueDouble = true; want false")
	}

	strArr = numToStrArray(111122)
	got = hasTrueDouble(strArr)
	if !got {
		t.Error("hasTrueDouble = false; want true")
	}

	strArr = numToStrArray(11123)
	got = hasTrueDouble(strArr)
	if got {
		t.Error("hasTrueDouble = true; want false")
	}

	strArr = numToStrArray(12333)
	got = hasTrueDouble(strArr)
	if got {
		t.Error("hasTrueDouble = true; want false")
	}

	strArr = numToStrArray(1112333)
	got = hasTrueDouble(strArr)
	if got {
		t.Error("hasTrueDouble = true; want false")
	}

	strArr = numToStrArray(11122333)
	got = hasTrueDouble(strArr)
	if !got {
		t.Error("hasTrueDouble = false; want true")
	}
}

func TestHasDouble(t *testing.T) {
	strArr := numToStrArray(123)
	got := hasDouble(strArr)
	if got {
		t.Error("hasDouble = true; want false")
	}

	strArr = numToStrArray(1233)
	got = hasDouble(strArr)
	if !got {
		t.Error("hasDouble = false; want true")
	}

	strArr = numToStrArray(222)
	got = hasDouble(strArr)
	if !got {
		t.Error("hasDouble = false; want true")
	}
}

func TestIsIncreasing(t *testing.T) {
	strArr := numToStrArray(123)
	got := isIncreasing(strArr)
	if !got {
		t.Error("isIncreasing = false; want true")
	}

	strArr = numToStrArray(1233)
	got = isIncreasing(strArr)
	if !got {
		t.Error("isIncreasing = false; want true")
	}

	strArr = numToStrArray(1231)
	got = isIncreasing(strArr)
	if got {
		t.Error("isIncreasing = true; want false")
	}
}

func TestNumToStrArr(t *testing.T) {
	got := numToStrArray(1)
	if got[0] != "1" {
		t.Errorf("numToStrArray = %v; want {\"1\"}", got)
	}

	got = numToStrArray(152)
	if got[0] != "1" || got[1] != "5" || got[2] != "2" {
		t.Errorf("numToStrArray = %v; want {\"1\", \"5\", \"2\"}", got)
	}
}
