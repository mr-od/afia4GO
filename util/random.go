package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{ENU, AWK, OWR, UMU, ABA, NGN, PHC, CAL, UYO}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomDescription() string {
	return RandomString(100)
}

func GenerateFileName(filename string) string {
	tim := time.Now().Format("20060102150405")
	splitsName := strings.Split(filename, ".")
	var lenfilename = len(splitsName)
	if lenfilename == 1 {
		return filename + "_" + tim
	} else {
		var newFileName = ""
		for i := 0; i < lenfilename-1; i++ {
			newFileName += splitsName[i]
		}
		newFileName += "_" + tim + "." + splitsName[lenfilename-1]
		return newFileName
	}
}

func RandomProduct() string {
	return RandomString(64)
}
