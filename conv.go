package sailor

import (
	"bytes"
	"strconv"
	"strings"
	//   "encoding/binary"
)

//byte转16进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	// buf := buffer.String()
	buf := strings.ToUpper(buffer.String())
	return buf
}

//16进制字符串转[]byte
func HexToByte(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}

func ByteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

/*
func Int64ToBytes(i int64) []byte {
  var buf = make([]byte, 8)
  binary.BigEndian.PutUint64(buf, uint64(i))
  return buf
}

func IntToBytes(i int) []byte {
  var buf = make([]byte, 2)
  binary.BigEndian.PutUint16(buf, uint16(i))
  return buf
}

func BytesToInt64(buf []byte) int64 {
  return int64(binary.BigEndian.Uint64(buf))
}

*/
