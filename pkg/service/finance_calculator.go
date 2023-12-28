package service

import (
	"fmt"
	"math"
)

const (
	// 最大迭代次数
	maxIterations = 100
	// 误差容忍度
	tolerance = 0.00001
)

// CalcNvr 通过现金流和内部收益率计算净现值NVR
func CalcNvr(cashFlows []float64, irr float64) float64 {
	var nvr float64
	for i, cashFlow := range cashFlows {
		nvr += cashFlow / math.Pow(1+irr, float64(i))
	}
	return nvr
}

// CalcIrr 通过现金流计算内部收益率IRR，使用牛顿迭代法
func CalcIrr(cashFlows []float64) float64 {
	var lowerRate = -0.99
	var upperRate = 0.99

	lowerNpv := CalcNvr(cashFlows, lowerRate)
	upperNpv := CalcNvr(cashFlows, upperRate)

	if lowerNpv*upperNpv > 0 {
		return math.NaN() // 无法找到IRR
	}

	for i := 0; i < maxIterations; i++ {
		rate := (lowerRate + upperRate) / 2
		npv := CalcNvr(cashFlows, rate)

		if math.Abs(npv) < tolerance {
			fmt.Println("迭代次数: ", i+1)
			return rate
		}

		if npv*lowerNpv < 0 {
			upperRate = rate
			upperNpv = npv
		} else {
			lowerRate = rate
			lowerNpv = npv
		}
	}

	return math.NaN() // 无法收敛到IRR
}
