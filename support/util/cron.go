package util

import (
	"fmt"
	"math/big"
)

// generateModularDailyCronSchedule returns a daily crontab schedule
// where, given a is input's integer representation, the minute is a % 60
// and hour is a % 24.
func GenerateModularDailyCronSchedule(input []byte) string {
	a := big.NewInt(0).SetBytes(input)
	var hi, mi big.Int
	m := mi.Mod(a, big.NewInt(60))
	h := hi.Mod(a, big.NewInt(24))
	return fmt.Sprintf("%d %d * * *", m.Int64(), h.Int64())
}
