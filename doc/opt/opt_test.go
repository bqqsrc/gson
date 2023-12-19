//  Copyright (C) 晓白齐齐,版权所有.

package opt

import (
	"github.com/bqqsrc/gson"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	//return
	var boolgv Bool
	err := gson.Unmarshal([]byte("false"), &boolgv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if boolgv {
		t.Fatalf("boolgv expected nil false, got true")
	}

	var intgv Int
	err = gson.Unmarshal([]byte("-564567865678"), &intgv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if intgv != -564567865678 {
		t.Fatalf("intgv expected -564567865678, got %d", intgv)
	}

	var int8gv Int8
	err = gson.Unmarshal([]byte("56"), &int8gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if int8gv != 56 {
		t.Fatalf("int8gv expected 56, got %d", int8gv)
	}

	var int16gv Int16
	err = gson.Unmarshal([]byte("6785"), &int16gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if int16gv != 6785 {
		t.Fatalf("int16gv expected 6785, got %d", int16gv)
	}

	var int32gv Int32
	err = gson.Unmarshal([]byte("38273849"), &int32gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if int32gv != 38273849 {
		t.Fatalf("int32gv expected 38273849, got %d", int32gv)
	}

	var int64gv Int64
	err = gson.Unmarshal([]byte("-338746798757"), &int64gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if int64gv != -338746798757 {
		t.Fatalf("int64gv expected -338746798757, got %d", int64gv)
	}

	var uintgv Uint
	err = gson.Unmarshal([]byte("349876493829"), &uintgv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if uintgv != 349876493829 {
		t.Fatalf("uintgv expected 349876493829, got %d", uintgv)
	}

	var uint8gv Uint8
	err = gson.Unmarshal([]byte("239"), &uint8gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if uint8gv != 239 {
		t.Fatalf("uint8gv expected 239, got %d", uint8gv)
	}

	var uint16gv Uint16
	err = gson.Unmarshal([]byte("34534"), &uint16gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if uint16gv != 34534 {
		t.Fatalf("uint16gv expected 34534, got %d", uint16gv)
	}

	var uint32gv Uint32
	err = gson.Unmarshal([]byte("373837482"), &uint32gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if uint32gv != 373837482 {
		t.Fatalf("uint32gv expected 373837482, got %d", uint32gv)
	}

	var uint64gv Uint64
	err = gson.Unmarshal([]byte("989348493849"), &uint64gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if uint64gv != 989348493849 {
		t.Fatalf("uint64gv expected 989348493849, got %d", uint64gv)
	}

	var float32gv Float32
	err = gson.Unmarshal([]byte("2.567e2"), &float32gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if float32gv != 256.7 {
		t.Fatalf("float32gv expected 256.7, got %f", float32gv)
	}

	var float64gv Float64
	err = gson.Unmarshal([]byte("592.3e-32"), &float64gv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if float64gv != 592.3e-32 {
		t.Fatalf("float64gv expected 592.3e-32, got %f", float64gv)
	}

	var stringgv String
	err = gson.Unmarshal([]byte("592.3Ee-32"), &stringgv)
	if err != nil {
		t.Fatalf("err expected nil err, got %s", err)
	}
	if stringgv != "592.3Ee-32" {
		t.Fatalf("stringgv expected 592.3Ee-32, got %s", stringgv)
	}
}
