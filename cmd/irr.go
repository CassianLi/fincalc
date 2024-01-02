/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fincalc/pkg/service"
	"fmt"

	"github.com/spf13/cobra"
)

// irrCmd represents the irr command
var irrCmd = &cobra.Command{
	Use:   "irr",
	Short: "IRR（Internal Rate of Return，内部收益率）",
	Long: `IRR表示的是一个项目的内部盈利率，即使项目的净现值（NVP）等于零的贴现率。通俗地说，IRR是使得投资的成本与项目的未来现金流相抵消的贴现率。

IRR的计算涉及到项目的现金流，即投资在不同期间内产生的现金收入和支出。IRR是使得项目的现金流通过某个贴现率折算到现值的总和等于零的贴现率。
特点如下:
	1. 项目的盈利能力： IRR表示投资项目内部的盈利率，即项目自身的回报率。如果IRR较高，说明项目具有较高的盈利潜力。
	2. 决策标准： 通常，投资者会比较IRR与其资本成本或期望的回报率。如果IRR高于这些标准，项目可能是有吸引力的。
	3. 与贴现率的关系： IRR与净现值（NVP）密切相关。当IRR等于贴现率时，NVP为零。因此，IRR提供了一个与NVP相互验证的指标。
	4. 多个IRR的情况： 有时，一个项目可能存在多个IRR，特别是在现金流变化方向发生变化时。在实际应用中，需要谨慎解释和选择合适的IRR。
	5. 比较项目： IRR允许比较不同项目的内部回报率。通常，较高的IRR被视为更有吸引力的投资。

IRR计算器使用方法：
	1. 使用完整现金流计算：
	fincalc irr -c -300,10,20,30
	2. 使用单期投资金额、投资期数、单期收益金额、收益期数计算：
	如：每期投资30万，投资3期，每期收益4万，总期数15，第1期开始收益，返还本金
	fincalc irr -a 30 -p 3 -C 4 -P 15 -S 2 -r

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("irr calculating ...")
		if len(cashFlows) == 0 {
			cashFlows = calcNvrByInvest(investAmount, investPeriod, incomeAmount, period, incomeStart, isReturnPrincipal)
		}

		fmt.Println("Cash Flows: ", cashFlows)
		fmt.Println("Periods: ", period)

		// 引入通胀率，计算名义现金流
		if inflationRate > 0 {
			cashFlows = calcNominalCashFlows(cashFlows, inflationRate)
			fmt.Println("Nominal Cash Flows: ", cashFlows)
		}

		irr := service.CalcIrr(cashFlows)
		fmt.Printf("IRR: %.4f\n", irr)
	},
}

func init() {
	rootCmd.AddCommand(irrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// irrCmd.PersistentFlags().String("foo", "", "A help for foo")
	//irrCmd.PersistentFlags().Float64SliceVarP(&cashFlows, "cash-flows", "c", []float64{}, "现金流")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// irrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
