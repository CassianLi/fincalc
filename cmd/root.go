/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math"
	"os"

	"github.com/spf13/cobra"
)

// 现金流
var cashFlows []float64

// 创建变量，单期投资金额，投资期数，单期收益金额，投资期数，第几期开始收益
var investAmount float64
var investPeriod int
var incomeAmount float64
var period int
var incomeStart int

// 是否返还本金
var isReturnPrincipal bool

// 通胀率
var inflationRate float64

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fincalc",
	Short: "金融计算器",
	Long:  `可用于计算一些常见的金融数据，如：NVR（Net Present Value，净现值）、IRR（Internal Rate of Return，内部收益率）等。例如:`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fincalc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().Float64SliceVarP(&cashFlows, "cash-flows", "f", []float64{}, "现金流")
	rootCmd.PersistentFlags().Float64VarP(&investAmount, "invest-amount", "a", 0, "每期投资额，默认为: 0")
	rootCmd.PersistentFlags().IntVarP(&investPeriod, "invest-period", "p", 1, "投资期数，默认为: 1")
	rootCmd.PersistentFlags().Float64VarP(&incomeAmount, "income-amount", "C", 0, "每期收益金额，默认为: 0")
	rootCmd.PersistentFlags().IntVarP(&period, "income-period", "P", 1, "总计期数，默认为: 1")
	rootCmd.PersistentFlags().IntVarP(&incomeStart, "income-start", "S", 1, "第几期开始收益，默认为: 1")
	rootCmd.PersistentFlags().BoolVarP(&isReturnPrincipal, "is-return-principal", "r", true, "是否返还本金，默认为: true")
	rootCmd.PersistentFlags().Float64VarP(&inflationRate, "inflation-rate", "I", 0, "通胀率，允许输入通胀率计算名义现金流（=实际现金流*(1+i%)^t）, 默认为: 0")

}

// calcNvrByInvest 通过单期投资金额、投资期数、单期收益金额、总期数、第几期开始收益，计算NVR计算现金流
func calcNvrByInvest(investAmount float64, investPeriod int, incomeAmount float64, period int, incomeStart int, isReturnPrincipal bool) []float64 {
	var cashFlows []float64
	for i := 0; i < investPeriod; i++ {
		if i < incomeStart-1 {
			cashFlows = append(cashFlows, -investAmount)
		}
		if i >= incomeStart-1 && i < investPeriod {
			cashFlows = append(cashFlows, incomeAmount-investAmount)
		}
	}
	for i := investPeriod; i < period; i++ {
		cashFlows = append(cashFlows, incomeAmount)
	}

	if isReturnPrincipal {
		// cashFlow最后一个元素加上所有投资额
		cashFlows[len(cashFlows)-1] += float64(investPeriod) * investAmount
	}
	return cashFlows
}

// 计算名义现金流
func calcNominalCashFlows(cashFlows []float64, inflationRate float64) []float64 {
	var nominalCashFlows []float64
	for i, cashFlow := range cashFlows {
		nominalCashFlows = append(nominalCashFlows, cashFlow*math.Pow(1+inflationRate, float64(i)))
	}
	return nominalCashFlows
}
