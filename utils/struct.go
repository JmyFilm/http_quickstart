package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MapToStructByTag Map赋值给结构体 key为结构体的tagName标签 跳过Tag为-的字段
func MapToStructByTag[T any](m map[string]string, s T, tagName string) T {
	v := reflect.ValueOf(&s).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get(tagName)
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = t.Field(i).Name
		}

		if val, ok := m[tag]; ok && val != "" {
			if field.CanSet() {
				switch field.Kind() {
				case reflect.String:
					field.SetString(val)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
						field.SetInt(intVal)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if uintVal, err := strconv.ParseUint(val, 10, 64); err == nil {
						field.SetUint(uintVal)
					}
				case reflect.Float32, reflect.Float64:
					if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
						field.SetFloat(floatVal)
					}
				case reflect.Bool:
					if boolVal, err := strconv.ParseBool(val); err == nil {
						field.SetBool(boolVal)
					}
				default:
					field.SetZero()
				}
			}
		}
	}
	return s
}

// StructToMapByTag 结构体转为Map key为结构体的tagName标签 跳过Tag为-的字段
func StructToMapByTag(s any, tagName string) map[string]string {
	result := make(map[string]string)

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get(tagName)
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = field.Name
		} else {
			parts := strings.Split(tag, ",")
			if parts[0] != "" {
				tag = parts[0]
			} else {
				tag = field.Name
			}
		}

		result[tag] = fmt.Sprintf("%v", v.Field(i).Interface())
	}

	return result
}

// AssignStruct 赋值结构体 匹配tagName标签 跳过Tag为-的字段
func AssignStruct[T any](source any, target T, tagName string) T {
	srcVal := reflect.ValueOf(source)
	srcType := reflect.TypeOf(source)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
		srcType = srcType.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return target
	}

	srcFieldMap := make(map[string]reflect.Value)
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		if field.PkgPath != "" {
			continue
		}

		var key string
		if tagName != "" {
			if tagVal := field.Tag.Get(tagName); tagVal != "" {
				if tagVal == "-" {
					continue
				} else {
					parts := strings.Split(tagVal, ",")
					if parts[0] != "" {
						key = parts[0]
					}
				}
			} else {
				key = field.Name
			}
		}
		srcFieldMap[key] = srcVal.Field(i)
	}

	tgtVal := reflect.ValueOf(&target).Elem()
	tgtType := tgtVal.Type()

	for i := 0; i < tgtType.NumField(); i++ {
		field := tgtType.Field(i)
		if field.PkgPath != "" {
			continue
		}

		var key string
		if tagName != "" {
			if tagVal := field.Tag.Get(tagName); tagVal != "" {
				if tagVal == "-" {
					continue
				} else {
					parts := strings.Split(tagVal, ",")
					if parts[0] != "" {
						key = parts[0]
					}
				}
			} else {
				key = field.Name
			}
		}

		srcField, ok := srcFieldMap[key]
		if !ok || srcField.IsZero() {
			continue
		}

		tgtField := tgtVal.Field(i)
		if !tgtField.CanSet() {
			continue
		}

		if srcField.Type().AssignableTo(tgtField.Type()) {
			tgtField.Set(srcField)
		} else if srcField.Type().ConvertibleTo(tgtField.Type()) {
			tgtField.Set(srcField.Convert(tgtField.Type()))
		}
	}
	return target
}
