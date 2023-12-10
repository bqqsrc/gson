//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"fmt"
	"strings"
)

func GetKeyWithSep(data []byte, keys, sep string, value interface{}) error {
	keyList := strings.Split(keys, sep)
	return GetKey(data, keyList, value)
}

func GetKey(data []byte, keys []string, value interface{}) error {
	hasSet := false
	if keys == nil && len(keys) == 0 {
		return UnmarshalAny(data, value)
	} else {
		currentKey := ""
		var err error
		keyFoundCallBack := func(d *Decoder, l *Lexer, isFound bool) bool {
			key := l.String()
			if keys[0] == key && isFound {
				currentKey = key
				num := len(keys)
				if num > 1 {
					keys = keys[1:num]
				} else {
					if err = d.SetAnyTarget(value, true); err != nil {
						return false
					}
					hasSet = true
				}
			} else if currentKey == key && !isFound {
				return false
			}
			return true
		}
		if err != nil {
			return err
		}
		err = DecodeData(data, keyFoundCallBack, nil, nil)
		if !hasSet {
			err = fmt.Errorf("not found key")
		}
		return err
	}
}

func GetString(data []byte, keys []string) (error, string) {
	var value string
	err := GetKey(data, keys, &value)
	return err, value
}

func GetStringD(data []byte, keys []string, defaultValue string) string {
	if err, value := GetString(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetStringWithSep(data []byte, keys, sep string) (error, string) {
	keyList := strings.Split(keys, sep)
	return GetString(data, keyList)
}

func GetStringDWithSep(data []byte, keys, sep, defaultValue string) string {
	keyList := strings.Split(keys, sep)
	return GetStringD(data, keyList, defaultValue)
}

func GetBool(data []byte, keys []string) (error, bool) {
	var value bool
	err := GetKey(data, keys, &value)
	return err, value
}

func GetBoolD(data []byte, keys []string, defaultValue bool) bool {
	if err, value := GetBool(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetBoolWithSep(data []byte, keys, sep string) (error, bool) {
	keyList := strings.Split(keys, sep)
	return GetBool(data, keyList)
}

func GetBoolDWithSep(data []byte, keys, sep string, defaultValue bool) bool {
	keyList := strings.Split(keys, sep)
	return GetBoolD(data, keyList, defaultValue)
}

func GetInt(data []byte, keys []string) (error, int) {
	var value int
	err := GetKey(data, keys, &value)
	return err, value
}

func GetIntD(data []byte, keys []string, defaultValue int) int {
	if err, value := GetInt(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetIntWithSep(data []byte, keys, sep string) (error, int) {
	keyList := strings.Split(keys, sep)
	return GetInt(data, keyList)
}

func GetIntDWithSep(data []byte, keys, sep string, defaultValue int) int {
	keyList := strings.Split(keys, sep)
	return GetIntD(data, keyList, defaultValue)
}

func GetInt8(data []byte, keys []string) (error, int8) {
	var value int8
	err := GetKey(data, keys, &value)
	return err, value
}

func GetInt8D(data []byte, keys []string, defaultValue int8) int8 {
	if err, value := GetInt8(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetInt8WithSep(data []byte, keys, sep string) (error, int8) {
	keyList := strings.Split(keys, sep)
	return GetInt8(data, keyList)
}

func GetInt8DWithSep(data []byte, keys, sep string, defaultValue int8) int8 {
	keyList := strings.Split(keys, sep)
	return GetInt8D(data, keyList, defaultValue)
}

func GetInt16(data []byte, keys []string) (error, int16) {
	var value int16
	err := GetKey(data, keys, &value)
	return err, value
}

func GetInt16D(data []byte, keys []string, defaultValue int16) int16 {
	if err, value := GetInt16(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetInt16WithSep(data []byte, keys, sep string) (error, int16) {
	keyList := strings.Split(keys, sep)
	return GetInt16(data, keyList)
}

func GetInt16DWithSep(data []byte, keys, sep string, defaultValue int16) int16 {
	keyList := strings.Split(keys, sep)
	return GetInt16D(data, keyList, defaultValue)
}

func GetInt32(data []byte, keys []string) (error, int32) {
	var value int32
	err := GetKey(data, keys, &value)
	return err, value
}

func GetInt32D(data []byte, keys []string, defaultValue int32) int32 {
	if err, value := GetInt32(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetInt32WithSep(data []byte, keys, sep string) (error, int32) {
	keyList := strings.Split(keys, sep)
	return GetInt32(data, keyList)
}

func GetInt32DWithSep(data []byte, keys, sep string, defaultValue int32) int32 {
	keyList := strings.Split(keys, sep)
	return GetInt32D(data, keyList, defaultValue)
}

func GetInt64(data []byte, keys []string) (error, int64) {
	var value int64
	err := GetKey(data, keys, &value)
	return err, value
}

func GetInt64D(data []byte, keys []string, defaultValue int64) int64 {
	if err, value := GetInt64(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetInt64WithSep(data []byte, keys, sep string) (error, int64) {
	keyList := strings.Split(keys, sep)
	return GetInt64(data, keyList)
}

func GetInt64DWithSep(data []byte, keys, sep string, defaultValue int64) int64 {
	keyList := strings.Split(keys, sep)
	return GetInt64D(data, keyList, defaultValue)
}

func GetUint(data []byte, keys []string) (error, uint) {
	var value uint
	err := GetKey(data, keys, &value)
	return err, value
}

func GetUintD(data []byte, keys []string, defaultValue uint) uint {
	if err, value := GetUint(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetUintWithSep(data []byte, keys, sep string) (error, uint) {
	keyList := strings.Split(keys, sep)
	return GetUint(data, keyList)
}

func GetUintDWithSep(data []byte, keys, sep string, defaultValue uint) uint {
	keyList := strings.Split(keys, sep)
	return GetUintD(data, keyList, defaultValue)
}

func GetUint8(data []byte, keys []string) (error, uint8) {
	var value uint8
	err := GetKey(data, keys, &value)
	return err, value
}

func GetUint8D(data []byte, keys []string, defaultValue uint8) uint8 {
	if err, value := GetUint8(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetUint8WithSep(data []byte, keys, sep string) (error, uint8) {
	keyList := strings.Split(keys, sep)
	return GetUint8(data, keyList)
}

func GetUint8DWithSep(data []byte, keys, sep string, defaultValue uint8) uint8 {
	keyList := strings.Split(keys, sep)
	return GetUint8D(data, keyList, defaultValue)
}

func GetUint16(data []byte, keys []string) (error, uint16) {
	var value uint16
	err := GetKey(data, keys, &value)
	return err, value
}

func GetUint16D(data []byte, keys []string, defaultValue uint16) uint16 {
	if err, value := GetUint16(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetUint16WithSep(data []byte, keys, sep string) (error, uint16) {
	keyList := strings.Split(keys, sep)
	return GetUint16(data, keyList)
}

func GetUint16DWithSep(data []byte, keys, sep string, defaultValue uint16) uint16 {
	keyList := strings.Split(keys, sep)
	return GetUint16D(data, keyList, defaultValue)
}

func GetUint32(data []byte, keys []string) (error, uint32) {
	var value uint32
	err := GetKey(data, keys, &value)
	return err, value
}

func GetUint32D(data []byte, keys []string, defaultValue uint32) uint32 {
	if err, value := GetUint32(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetUint32WithSep(data []byte, keys, sep string) (error, uint32) {
	keyList := strings.Split(keys, sep)
	return GetUint32(data, keyList)
}

func GetUint32DWithSep(data []byte, keys, sep string, defaultValue uint32) uint32 {
	keyList := strings.Split(keys, sep)
	return GetUint32D(data, keyList, defaultValue)
}

func GetUint64(data []byte, keys []string) (error, uint64) {
	var value uint64
	err := GetKey(data, keys, &value)
	return err, value
}

func GetUint64D(data []byte, keys []string, defaultValue uint64) uint64 {
	if err, value := GetUint64(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetUint64WithSep(data []byte, keys, sep string) (error, uint64) {
	keyList := strings.Split(keys, sep)
	return GetUint64(data, keyList)
}

func GetUint64DWithSep(data []byte, keys, sep string, defaultValue uint64) uint64 {
	keyList := strings.Split(keys, sep)
	return GetUint64D(data, keyList, defaultValue)
}

func GetFloat32(data []byte, keys []string) (error, float32) {
	var value float32
	err := GetKey(data, keys, &value)
	return err, value
}

func GetFloat32D(data []byte, keys []string, defaultValue float32) float32 {
	if err, value := GetFloat32(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetFloat32WithSep(data []byte, keys, sep string) (error, float32) {
	keyList := strings.Split(keys, sep)
	return GetFloat32(data, keyList)
}

func GetFloat32DWithSep(data []byte, keys, sep string, defaultValue float32) float32 {
	keyList := strings.Split(keys, sep)
	return GetFloat32D(data, keyList, defaultValue)
}

func GetFloat64(data []byte, keys []string) (error, float64) {
	var value float64
	err := GetKey(data, keys, &value)
	return err, value
}

func GetFloat64D(data []byte, keys []string, defaultValue float64) float64 {
	if err, value := GetFloat64(data, keys); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetFloat64WithSep(data []byte, keys, sep string) (error, float64) {
	keyList := strings.Split(keys, sep)
	return GetFloat64(data, keyList)
}

func GetFloat64DWithSep(data []byte, keys, sep string, defaultValue float64) float64 {
	keyList := strings.Split(keys, sep)
	return GetFloat64D(data, keyList, defaultValue)
}
