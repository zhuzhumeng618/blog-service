package convert

import (
	"strconv"
)

// StrTo 自定义类型，用于 string 转换
type StrTo string

// String s 转换为 string 类型
func (s StrTo) String() string {
	return string(s)
}

// Int s 转换为 Int 类型
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt s 转换为 Int 类型，不做错误处理
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// UInt32 s 转换为 UInt32 类型
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUInt32 s 转换为 UInt32 类型，不做错误处理
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
