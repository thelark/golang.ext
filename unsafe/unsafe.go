package unsafe

import (
	"strconv"
)

/**
 * 不安全的转换
 */

/**
 * TODO: string => int
 * @param s int的字符串形式
 * @return int
 */
func AtoI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

/**
 * TODO: string => int8
 * @param s int8的字符串形式
 * @param base int8字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return int8
 */
func ParseInt8(s string, base int) int8 {
	i, _ := strconv.ParseInt(s, base, 8)
	return int8(i)
}

/**
 * TODO: string => int16
 * @param s int16的字符串形式
 * @param base int16字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return int16
 */
func ParseInt16(s string, base int) int16 {
	i, _ := strconv.ParseInt(s, base, 16)
	return int16(i)
}

/**
 * TODO: string => int32
 * @param s int32的字符串形式
 * @param base int32字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return int32
 */
func ParseInt32(s string, base int) int32 {
	i, _ := strconv.ParseInt(s, base, 32)
	return int32(i)
}

/**
 * TODO: string => int64
 * @param s int64的字符串形式
 * @param base int64字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return int64
 */
func ParseInt64(s string, base int) int64 {
	i, _ := strconv.ParseInt(s, base, 64)
	return i
}

/**
 * TODO: string => uint8
 * @param s uint8的字符串形式
 * @param base uint8字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return uint8
 */
func ParseUint8(s string, base int) uint8 {
	i, _ := strconv.ParseUint(s, base, 8)
	return uint8(i)
}

/**
 * TODO: string => uint16
 * @param s uint16的字符串形式
 * @param base uint16字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return uint16
 */
func ParseUint16(s string, base int) uint16 {
	i, _ := strconv.ParseUint(s, base, 16)
	return uint16(i)
}

/**
 * TODO: string => uint32
 * @param s uint32的字符串形式
 * @param base uint32字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return uint32
 */
func ParseUint32(s string, base int) uint32 {
	i, _ := strconv.ParseUint(s, base, 32)
	return uint32(i)
}

/**
 * TODO: string => uint64
 * @param s uint64的字符串形式
 * @param base uint64字符串的进制 比如二进制 八进制 十进制 十六进制
 * @return uint64
 */
func ParseUint64(s string, base int) uint64 {
	i, _ := strconv.ParseUint(s, base, 64)
	return i
}

/**
 * TODO: string => float64
 * @param s float64的字符串形式
 * @return float64
 */
func ParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

/**
 * TODO: string => float32
 * @param s float32的字符串形式
 * @return float32
 */
func ParseFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

/**
 * TODO: string => bool
 * @param s bool的字符串形式
 * @return bool
 */
func ParseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
