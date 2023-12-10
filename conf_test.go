//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var confData = map[string]map[string][]testVarInfo{
	confGsonFile: {
		"bool_1_1": []testVarInfo{
			{"bool_1_1", true, &bool_1_1, reflect.Bool, true},
			{"bool_1_1_CDF", true, true, reflect.Bool, false},
			{"bool_1_1_CNDF", true, nil, reflect.Bool, false},
		},
		"bool_1_2": []testVarInfo{
			{"bool_1_2", false, &bool_1_2, reflect.Bool, true},
			{"bool_1_2_CDF", false, false, reflect.Bool, false},
			{"bool_1_2_CNDF", false, nil, reflect.Bool, false},
		},
		// // "bool_2_1": []testVarInfo{
		// // 	{"bool_2_1", false, &bool_2_1, reflect.Bool, true},
		// // },
		// // "bool_2_2": []testVarInfo{
		// // 	{"bool_2_2", true, &bool_2_2, reflect.Bool, true},
		// // },
		"float32_1": []testVarInfo{
			{"float32_1", float32(59.999), &float32_1, reflect.Float32, true},
			{"float32_1_CDF", float32(59.999), float32(59.999), reflect.Float32, false},
			{"float32_1_CNDF", float32(59.999), nil, reflect.Float32, false},
		},
		// // "float32_2": []testVarInfo{
		// // 	{"float32_2", float32(89.77), &float32_2, reflect.Float32, true},
		// // },
		"float64_1": []testVarInfo{
			{"float64_1", 340000.0, &float64_1, reflect.Float64, true},
			{"float64_1_CDF", 340000.0, 340000.0, reflect.Float64, false},
			{"float64_1_CNDF", 340000.0, nil, reflect.Float64, false},
		},
		// // "float64_2": []testVarInfo{
		// // 	{"float64_2", 98.7, &float64_2, reflect.Float64, true},
		// // },
		// "int8_1": []testVarInfo{
		// 	{"int8_1", int8(-33), &int8_1, reflect.Int8, true},
		// 	{"int8_1_CDF", int8(-33), int8(-33), reflect.Int8, false},
		// 	{"int8_1_CNDF", int8(-33), nil, reflect.Int8, false},
		// },
		// // "int8_2": []testVarInfo{
		// // 	{"int8_2", int8(58), &int8_2, reflect.Int8, true},
		// // },
		// "int16_1": []testVarInfo{
		// 	{"int16_1", int16(456), &int16_1, reflect.Int16, true},
		// 	{"int16_1_CDF", int16(456), int16(456), reflect.Int16, false},
		// 	{"int16_1_CNDF", int16(456), nil, reflect.Int16, false},
		// },
		// // "int16_2": []testVarInfo{
		// // 	{"int16_2", int16(-567), &int16_2, reflect.Int16, true},
		// // },
		// "int32_1": []testVarInfo{
		// 	{"int32_1", int32(703522), &int32_1, reflect.Int32, true},
		// 	{"int32_1_CDF", int32(703522), int32(703522), reflect.Int32, false},
		// 	{"int32_1_CNDF", int32(703522), nil, reflect.Int32, false},
		// },
		// // "int32_2": []testVarInfo{
		// // 	{"int32_2", int32(-89076), &int32_2, reflect.Int32, true},
		// // },
		// "int64_1": []testVarInfo{
		// 	{"int64_1", int64(325235298967), &int64_1, reflect.Int64, true},
		// 	{"int64_1_CDF", int64(325235298967), int64(325235298967), reflect.Int64, false},
		// 	{"int64_1_CNDF", int64(325235298967), nil, reflect.Int64, false},
		// },
		// // "int64_2": []testVarInfo{
		// // 	{"int64_2", int64(-98799889678), &int64_2, reflect.Int64, true},
		// // },
		// "int_1": []testVarInfo{
		// 	{"int_1", 39870935321, &int_1, reflect.Int, true},
		// 	{"int_1_CDF", 39870935321, 39870935321, reflect.Int, false},
		// 	{"int_1_CNDF", 39870935321, nil, reflect.Int, false},
		// },
		// // "int_2": []testVarInfo{
		// // 	{"int_2", 23458987898, &int_2, reflect.Int, true},
		// // },
		// "uint8_1": []testVarInfo{
		// 	{"uint8_1", uint8(27), &uint8_1, reflect.Uint8, true},
		// 	{"uint8_1_CDF", uint8(27), uint8(27), reflect.Uint8, false},
		// 	{"uint8_1_CNDF", uint8(27), nil, reflect.Uint8, false},
		// },
		// // "uint8_2": []testVarInfo{
		// // 	{"uint8_2", uint8(35), &uint8_2, reflect.Uint8, true},
		// // },
		// "uint16_1": []testVarInfo{
		// 	{"uint16_1", uint16(356), &uint16_1, reflect.Uint16, true},
		// 	{"uint16_1_CDF", uint16(356), uint16(356), reflect.Uint16, false},
		// 	{"uint16_1_CNDF", uint16(356), nil, reflect.Uint16, false},
		// },
		// // "uint16_2": []testVarInfo{
		// // 	{"uint16_2", uint16(598), &uint16_2, reflect.Uint16, true},
		// // },
		// "uint32_1": []testVarInfo{
		// 	{"uint32_1", uint32(588522), &uint32_1, reflect.Uint32, true},
		// 	{"uint32_1_CDF", uint32(588522), uint32(588522), reflect.Uint32, false},
		// 	{"uint32_1_CNDF", uint32(588522), nil, reflect.Uint32, false},
		// },
		// // "uint32_2": []testVarInfo{
		// // 	{"uint32_2", uint32(82376), &uint32_2, reflect.Uint32, true},
		// // },
		// "uint64_1": []testVarInfo{
		// 	{"uint64_1", uint64(325906298967), &uint64_1, reflect.Uint64, true},
		// 	{"uint64_1_CDF", uint64(325906298967), uint64(325906298967), reflect.Uint64, false},
		// 	{"uint64_1_CNDF", uint64(325906298967), nil, reflect.Uint64, false},
		// },
		// // "uint64_2": []testVarInfo{
		// // 	{"uint64_2", uint64(98757889678), &uint64_2, reflect.Uint64, true},
		// // },
		// "uint_1": []testVarInfo{
		// 	{"uint_1", uint(39870967321), &uint_1, reflect.Uint, true},
		// 	{"uint_1_CDF", uint(39870967321), uint(39870967321), reflect.Uint, false},
		// 	{"uint_1_CNDF", uint(39870967321), nil, reflect.Uint, false},
		// },
		// // "uint_2": []testVarInfo{
		// // 	{"uint_2", uint(23453487898), &uint_2, reflect.Uint, true},
		// // },
		// "string_1": []testVarInfo{
		// 	{"string_1", "baiqiqi5", &string_1, reflect.String, true},
		// 	{"string_1_CDF", "baiqiqi5", "baiqiqi5", reflect.String, false},
		// 	{"string_1_CNDF", "baiqiqi5", nil, reflect.String, false},
		// },
		// // "string_2": []testVarInfo{
		// // 	{"string_2", "baiqiqi3", &string_2, reflect.String, true},
		// // },
		"testdata_1_3:TestStr": []testVarInfo{
			{"testdata_1_3:TestStr", "baiqiqi3", &testdata_1_3.TestStr, reflect.String, true},
			{"testdata_1_3:TestStr_CDF", "baiqiqi3", "baiqiqi3", reflect.String, false},
			{"testdata_1_3:TestStr_CNDF", "baiqiqi3", nil, reflect.String, false},
		},
		// "testdata_1_3:TestInt": []testVarInfo{
		// 	{"testdata_1_3:TestInt", -3599, &testdata_1_3.TestInt, reflect.Int, true},
		// 	{"testdata_1_3:TestInt_CDF", -3599, -3599, reflect.Int, false},
		// 	{"testdata_1_3:TestInt_CNDF", -3599, nil, reflect.Int, false},
		// },
		// "testdata_1_3:TestFloat": []testVarInfo{
		// 	{"testdata_1_3:TestFloat", 99.7, &testdata_1_3.TestFloat, reflect.Float64, true},
		// 	{"testdata_1_3:TestFloat_CDF", 99.7, 99.7, reflect.Float64, false},
		// 	{"testdata_1_3:TestFloat_CNDF", 99.7, nil, reflect.Float64, false},
		// },
		"testdata_1_3:TestBool": []testVarInfo{
			{"testdata_1_3:TestBool", true, &testdata_1_3.TestBool, reflect.Bool, true},
			{"testdata_1_3:TestBool_CDF", true, true, reflect.Bool, false},
			{"testdata_1_3:TestBool_CNDF", true, nil, reflect.Bool, false},
		},
		"testdata_1_3:TestUint": []testVarInfo{
			{"testdata_1_3:TestUint", uint(235), &testdata_1_3.TestUint, reflect.Uint, true},
			{"testdata_1_3:TestUint_CDF", uint(235), uint(235), reflect.Uint, false},
			{"testdata_1_3:TestUint_CNDF", uint(235), nil, reflect.Uint, false},
		},

		// // "testdata_1_3": []testVarInfo{
		// // 	{"testdata_1_3:TestStr", "baiqiqi3", &testdata_1_3.TestStr, reflect.String, true},
		// // 	{"testdata_1_3:TestInt", -3599, &testdata_1_3.TestInt, reflect.Int, true},
		// // 	{"testdata_1_3:TestFloat", 99.7, &testdata_1_3.TestFloat, reflect.Float64, true},
		// // 	{"testdata_1_3:TestBool", true, &testdata_1_3.TestBool, reflect.Bool, true},
		// // 	{"testdata_1_3:TestUint", uint(235), &testdata_1_3.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_2_3": []testVarInfo{
		// // 	{"testdata_2_3:test_str", "baiqiqi4", &testdata_2_3.TestStr, reflect.String, true},
		// // 	{"testdata_2_3:test_int", -9088, &testdata_2_3.TestInt, reflect.Int, true},
		// // 	{"testdata_2_3:test_float", 99.53, &testdata_2_3.TestFloat, reflect.Float64, true},
		// // 	{"testdata_2_3:test_bool", false, &testdata_2_3.TestBool, reflect.Bool, true},
		// // 	{"testdata_2_3:test_uint", uint(234), &testdata_2_3.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_2_4": []testVarInfo{
		// // 	{"testdata_2_4:TestStr", "baiqiqi5", &testdata_2_4.TestStr, reflect.String, true},
		// // 	{"testdata_2_4:TestInt", -19088, &testdata_2_4.TestInt, reflect.Int, true},
		// // 	{"testdata_2_4:TestFloat", 991.53, &testdata_2_4.TestFloat, reflect.Float64, true},
		// // 	{"testdata_2_4:TestBool", false, &testdata_2_4.TestBool, reflect.Bool, true},
		// // 	{"testdata_2_4:TestUint", uint(2234), &testdata_2_4.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_2_5": []testVarInfo{
		// // 	{"testdata_2_5:TestStr", "baiqiqi6", &testdata_2_5.TestStr, reflect.String, true},
		// // 	{"testdata_2_5:test_int", -188, &testdata_2_5.TestInt, reflect.Int, true},
		// // 	{"testdata_2_5:test_float", -991.53, &testdata_2_5.TestFloat, reflect.Float64, true},
		// // 	{"testdata_2_5:TestBool", false, &testdata_2_5.TestBool, reflect.Bool, true},
		// // 	{"testdata_2_5:TestUint", uint(1134), &testdata_2_5.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_3_3": []testVarInfo{
		// // 	{"testdata_3_3:test_str", "baiqiqi3_3", &testdata_3_3.TestStr, reflect.String, true},
		// // 	{"testdata_3_3:test_int", -988, &testdata_3_3.TestInt, reflect.Int, true},
		// // 	{"testdata_3_3:TestFloat", 99.3, &testdata_3_3.TestFloat, reflect.Float64, true},
		// // 	{"testdata_3_3:test_bool", true, &testdata_3_3.TestBool, reflect.Bool, true},
		// // 	{"testdata_3_3:TestUint", uint(34), &testdata_3_3.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_3_4": []testVarInfo{
		// // 	{"testdata_3_4:TestStr", "baiqiqi3_4", &testdata_3_4.TestStr, reflect.String, true},
		// // 	{"testdata_3_4:TestInt", -198, &testdata_3_4.TestInt, reflect.Int, true},
		// // 	{"testdata_3_4:TestFloat", 91.53, &testdata_3_4.TestFloat, reflect.Float64, true},
		// // 	{"testdata_3_4:TestBool", false, &testdata_3_4.TestBool, reflect.Bool, true},
		// // 	{"testdata_3_4:TestUint", uint(224), &testdata_3_4.TestUint, reflect.Uint, true},
		// // },
		// // "testdata_3_5": []testVarInfo{
		// // 	{"testdata_3_5:TestStr", "baiqiqi3_5", &testdata_3_5.TestStr, reflect.String, true},
		// // 	{"testdata_3_5:test_int", -1288, &testdata_3_5.TestInt, reflect.Int, true},
		// // 	{"testdata_3_5:TestFloat", 891.53, &testdata_3_5.TestFloat, reflect.Float64, true},
		// // 	{"testdata_3_5:TestBool", true, &testdata_3_5.TestBool, reflect.Bool, true},
		// // 	{"testdata_3_5:TestUint", uint(114), &testdata_3_5.TestUint, reflect.Uint, true},
		// // },
	},
}

func Test_Conf(t *testing.T) {
	if confGsonFile == "" {
		confGsonFile = "conf.gson"
	}
	gsF := "./testdata/" + confGsonFile
	data, _ := ioutil.ReadFile(gsF)
	resetAll()
	testConf(t, confKey, confData, data)
}

func testConf(t *testing.T, decodeAct string, decodeData map[string]map[string][]testVarInfo, data []byte) {
	byteWhiteList, ok := whiteList[decodeAct]
	if !allTest && (!ok || byteWhiteList == nil || len(byteWhiteList) == 0) {
		return
	}

	for key, value := range decodeData {

		var elemWhiteList map[string][]string
		ok := false
		if !allTest {
			elemWhiteList, ok = byteWhiteList[key]
			if !ok || elemWhiteList == nil || len(elemWhiteList) == 0 {
				continue
			}
		}

		for elemKey, elemVal := range value {
			var varWhiteList []string
			if !allTest {
				varWhiteList, ok = elemWhiteList[elemKey]
				if !ok || varWhiteList == nil || len(varWhiteList) == 0 {
					continue
				}
			}
			for _, varValue := range elemVal {
				varKey := varValue.key
				if !allTest {
					if !isInArr(varKey, varWhiteList) {
						continue
					}
				}
				keyName := decodeAct + ".go-" + key + "-" + elemKey + "-" + varKey
				testGetConf(t, varValue, elemKey, keyName, data)
			}
		}
	}
}

func testGetConf(t *testing.T, varValue testVarInfo, confKey, showKey string, data []byte) {
	var err error
	if strings.HasSuffix(varValue.key, "_CDF") {
		var ret interface{}
		switch varValue.basicType {
		case reflect.Bool:
			ret = GetBoolDWithSep(data, confKey, ":", varValue.got.(bool))
		case reflect.Int:
			ret = GetIntDWithSep(data, confKey, ":", varValue.got.(int))
		case reflect.Int8:
			ret = GetInt8DWithSep(data, confKey, ":", varValue.got.(int8))
		case reflect.Int16:
			ret = GetInt16DWithSep(data, confKey, ":", varValue.got.(int16))
		case reflect.Int32:
			ret = GetInt32DWithSep(data, confKey, ":", varValue.got.(int32))
		case reflect.Int64:
			ret = GetInt64DWithSep(data, confKey, ":", varValue.got.(int64))
		case reflect.Uint:
			ret = GetUintDWithSep(data, confKey, ":", varValue.got.(uint))
		case reflect.Uint8:
			ret = GetUint8DWithSep(data, confKey, ":", varValue.got.(uint8))
		case reflect.Uint16:
			ret = GetUint16DWithSep(data, confKey, ":", varValue.got.(uint16))
		case reflect.Uint32:
			ret = GetUint32DWithSep(data, confKey, ":", varValue.got.(uint32))
		case reflect.Uint64:
			ret = GetUint64DWithSep(data, confKey, ":", varValue.got.(uint64))
		case reflect.Float32:
			ret = GetFloat32DWithSep(data, confKey, ":", varValue.got.(float32))
		case reflect.Float64:
			ret = GetFloat64DWithSep(data, confKey, ":", varValue.got.(float64))
		case reflect.String:
			ret = GetStringDWithSep(data, confKey, ":", varValue.got.(string))
		}
		checkTestError(t, err, true, "")
		checkBasicResult(t, showKey, varValue.expected, ret, varValue.basicType, false)
	} else if strings.HasSuffix(varValue.key, "_CNDF") {
		var ret interface{}
		switch varValue.basicType {
		case reflect.Bool:
			err, ret = GetBoolWithSep(data, confKey, ":")
		case reflect.Uint:
			err, ret = GetUintWithSep(data, confKey, ":")
		case reflect.Uint8:
			err, ret = GetUint8WithSep(data, confKey, ":")
		case reflect.Uint16:
			err, ret = GetUint16WithSep(data, confKey, ":")
		case reflect.Uint32:
			err, ret = GetUint32WithSep(data, confKey, ":")
		case reflect.Uint64:
			err, ret = GetUint64WithSep(data, confKey, ":")
		case reflect.Int:
			err, ret = GetIntWithSep(data, confKey, ":")
		case reflect.Int8:
			err, ret = GetInt8WithSep(data, confKey, ":")
		case reflect.Int16:
			err, ret = GetInt16WithSep(data, confKey, ":")
		case reflect.Int32:
			err, ret = GetInt32WithSep(data, confKey, ":")
		case reflect.Int64:
			err, ret = GetInt64WithSep(data, confKey, ":")
		case reflect.Float32:
			err, ret = GetFloat32WithSep(data, confKey, ":")
		case reflect.Float64:
			err, ret = GetFloat64WithSep(data, confKey, ":")
		case reflect.String:
			err, ret = GetStringWithSep(data, confKey, ":")
		}
		checkTestError(t, err, true, "")
		checkBasicResult(t, showKey, varValue.expected, ret, varValue.basicType, false)
	} else {
		if varValue.got != nil && varValue.gotPointer {
			err = GetKeyWithSep(data, confKey, ":", varValue.got)
			checkTestError(t, err, true, "")
			checkBasicResult(t, showKey, varValue.expected, varValue.got, varValue.basicType, varValue.gotPointer)
		} else {
			t.Fatalf("conf_test.go:%s err: got is nil or gotPointer is false, got expected not nil, and gotPointer expected true", showKey)
		}
	}
}

//TODO 添加没有传入key的测试样例
