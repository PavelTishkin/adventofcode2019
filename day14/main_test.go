package main

import "testing"

func TestCalcFuelForOre(t *testing.T) {
	input := []string{
		"157 ORE => 5 NZVS",
		"165 ORE => 6 DCFZ",
		"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
		"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
		"179 ORE => 7 PSHF",
		"177 ORE => 5 HKGWZ",
		"7 DCFZ, 7 PSHF => 2 XJWVT",
		"165 ORE => 2 GPVTF",
		"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	}
	fuelProduced := calcFuelForOre(input)
	fuelExpected := int64(82892753)

	if fuelProduced != fuelExpected {
		t.Errorf("calcFuelForOre = %d, expected %d", fuelProduced, fuelExpected)
	}

	input = []string{
		"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
		"17 NVRVD, 3 JNWZP => 8 VPVL",
		"53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
		"22 VJHF, 37 MNCFX => 5 FWMGM",
		"139 ORE => 4 NVRVD",
		"144 ORE => 7 JNWZP",
		"5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
		"5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
		"145 ORE => 6 MNCFX",
		"1 NVRVD => 8 CXFTF",
		"1 VJHF, 6 MNCFX => 4 RFSQX",
		"176 ORE => 6 VJHF",
	}
	fuelProduced = calcFuelForOre(input)
	fuelExpected = 5586022

	if fuelProduced != fuelExpected {
		t.Errorf("calcFuelForOre = %d, expected %d", fuelProduced, fuelExpected)
	}

	input = []string{
		"171 ORE => 8 CNZTR",
		"7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL",
		"114 ORE => 4 BHXH",
		"14 VRPVC => 6 BMBT",
		"6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL",
		"6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT",
		"15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW",
		"13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW",
		"5 BMBT => 4 WPTQ",
		"189 ORE => 9 KTJDG",
		"1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
		"12 VRPVC, 27 CNZTR => 2 XDBXC",
		"15 KTJDG, 12 BHXH => 5 XCVML",
		"3 BHXH, 2 VRPVC => 7 MZWV",
		"121 ORE => 7 VRPVC",
		"7 XCVML => 6 RJRHP",
		"5 BHXH, 4 VRPVC => 5 LTCX",
	}
	fuelProduced = calcFuelForOre(input)
	fuelExpected = 460664

	if fuelProduced != fuelExpected {
		t.Errorf("calcFuelForOre = %d, expected %d", fuelProduced, fuelExpected)
	}
}

func TestCalcOreForFuel(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}
	oreUsed := calcOreForFuel(input)
	oreExpected := int64(31)

	if oreUsed != oreExpected {
		t.Errorf("calcOreForFuel = %d, expected %d", oreUsed, oreExpected)
	}

	input = []string{
		"9 ORE => 2 A",
		"8 ORE => 3 B",
		"7 ORE => 5 C",
		"3 A, 4 B => 1 AB",
		"5 B, 7 C => 1 BC",
		"4 C, 1 A => 1 CA",
		"2 AB, 3 BC, 4 CA => 1 FUEL",
	}
	oreUsed = calcOreForFuel(input)
	oreExpected = 165

	if oreUsed != oreExpected {
		t.Errorf("calcOreForFuel = %d, expected %d", oreUsed, oreExpected)
	}

	input = []string{
		"157 ORE => 5 NZVS",
		"165 ORE => 6 DCFZ",
		"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
		"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
		"179 ORE => 7 PSHF",
		"177 ORE => 5 HKGWZ",
		"7 DCFZ, 7 PSHF => 2 XJWVT",
		"165 ORE => 2 GPVTF",
		"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	}
	oreUsed = calcOreForFuel(input)
	oreExpected = 13312

	if oreUsed != oreExpected {
		t.Errorf("calcOreForFuel = %d, expected %d", oreUsed, oreExpected)
	}

	input = []string{
		"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
		"17 NVRVD, 3 JNWZP => 8 VPVL",
		"53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
		"22 VJHF, 37 MNCFX => 5 FWMGM",
		"139 ORE => 4 NVRVD",
		"144 ORE => 7 JNWZP",
		"5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
		"5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
		"145 ORE => 6 MNCFX",
		"1 NVRVD => 8 CXFTF",
		"1 VJHF, 6 MNCFX => 4 RFSQX",
		"176 ORE => 6 VJHF",
	}
	oreUsed = calcOreForFuel(input)
	oreExpected = 180697

	if oreUsed != oreExpected {
		t.Errorf("calcOreForFuel = %d, expected %d", oreUsed, oreExpected)
	}

	input = []string{
		"171 ORE => 8 CNZTR",
		"7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL",
		"114 ORE => 4 BHXH",
		"14 VRPVC => 6 BMBT",
		"6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL",
		"6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT",
		"15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW",
		"13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW",
		"5 BMBT => 4 WPTQ",
		"189 ORE => 9 KTJDG",
		"1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
		"12 VRPVC, 27 CNZTR => 2 XDBXC",
		"15 KTJDG, 12 BHXH => 5 XCVML",
		"3 BHXH, 2 VRPVC => 7 MZWV",
		"121 ORE => 7 VRPVC",
		"7 XCVML => 6 RJRHP",
		"5 BHXH, 4 VRPVC => 5 LTCX",
	}
	oreUsed = calcOreForFuel(input)
	oreExpected = 2210736

	if oreUsed != oreExpected {
		t.Errorf("calcOreForFuel = %d, expected %d", oreUsed, oreExpected)
	}
}

func TestParseReactions(t *testing.T) {
	input := []string{
		"10 ORE => 5 A",
		"3 A, 1 ORE => 2 B",
	}
	reactions := parseReactions(input)

	if reactions[0].input[0].element != "ORE" {
		t.Errorf("reactions[0].input[0].element = %s, expected ORE", reactions[0].input[0].element)
	}
	if reactions[0].input[0].weight != 10 {
		t.Errorf("reactions[0].input[0].weight = %d, expected 10", reactions[0].input[0].weight)
	}
	if reactions[1].input[0].element != "A" {
		t.Errorf("reactions[1].input[1].element = %s, expected A", reactions[1].input[1].element)
	}
	if reactions[1].input[0].weight != 3 {
		t.Errorf("reactions[1].input[1].weight = %d, expected 3", reactions[1].input[1].weight)
	}
	if reactions[1].output.element != "B" {
		t.Errorf("reactions[1].output.element = %s, expected B", reactions[1].output.element)
	}
	if reactions[1].output.weight != 2 {
		t.Errorf("reactions[1].output.weight = %d, expected 2", reactions[1].output.weight)
	}

}

func TestParseElement(t *testing.T) {
	got := parseElement("5 ORE")
	if got.element != "ORE" {
		t.Errorf("element = %s, expected ORE", got.element)
	}
	if got.weight != 5 {
		t.Errorf("weight = %d, expected 5", got.weight)
	}
}
