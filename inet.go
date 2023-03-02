package inet

import (
	"bytes"
	"strconv"
	"strings"
)

// Aton == > inet_aton
func Aton(ipstring string) uint32 {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt uint32 = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | uint32(tempInt)
		pos -= 8
	}
	return ipInt
}

// Ntoa == > inet_ntoa
func Ntoa(ipInt uint32) string {
	ipSegs := make([]string, 4)
	var l int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < l; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[l-i-1] = strconv.Itoa(int(tempInt))
		ipInt = ipInt >> 8
	}
	for i := 0; i < l; i++ {
		buffer.WriteString(ipSegs[i])
		if i < l-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
