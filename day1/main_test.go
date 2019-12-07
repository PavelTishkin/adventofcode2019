package main

import "testing"

func TestCalcMass(t *testing.T) {
	got := calcMass(12)
	if got != 2 {
		t.Errorf("calcMass = %d; want 2", got)
	}

	got = calcMass(14)
	if got != 2 {
		t.Errorf("calcMass = %d; want 2", got)
	}

	got = calcMass(1969)
	if got != 654 {
		t.Errorf("calcMass = %d; want 654", got)
	}

	got = calcMass(100756)
	if got != 33583 {
		t.Errorf("calcMass = %d; want 33583", got)
	}
}

func TestCalcMassReq(t *testing.T) {
	got := calcMassRec(14)
	if got != 2 {
		t.Errorf("calcMassRec = %d; want 2", got)
	}

	got = calcMassRec(1969)
	if got != 966 {
		t.Errorf("calcMassRec = %d; want 966", got)
	}

	got = calcMassRec(100756)
	if got != 50346 {
		t.Errorf("calcMassRec = %d; want 50346", got)
	}
}
