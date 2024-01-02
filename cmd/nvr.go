/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fincalc/pkg/service"
	"fmt"

	"github.com/spf13/cobra"
)

var irr float64

// nvrCmd represents the nvr command
var nvrCmd = &cobra.Command{
	Use:   "nvr",
	Short: "计算NVR（Net Present Value，净现值）",
	Long: `净现值（Net Present Value，简称 NVP）是一种用于评估投资项目或业务提案的财务指标。
它通过将未来的现金流通过某个贴现率折算到现值的总和来衡量一个项目的盈利能力。
NVP的计算涉及考虑了现金流的时间价值，即现金流越早发生，其现值越高。特点如下:
	1. 正值和负值： 如果NVP为正值，则意味着项目的现值超过了投资成本，即项目预计会产生盈利。如果NVP为负值，则项目可能不太具有吸引力，因为其现值低于投资成本。
	2. 决策标准： 在投资决策中，通常的标准是，如果NVP大于零，则项目可能是有吸引力的，因为它可以为投资者创造附加价值。
	3.与贴现率的关系： NVP的计算涉及使用贴现率。当贴现率低于项目的内部收益率（IRR）时，NVP通常会更加积极，因为更低的贴现率使得未来现金流的现值更高。
	4. 比较项目： NVP允许比较不同项目的盈利能力。在比较中，具有更高NVP的项目通常被视为更有吸引力，因为它们提供更高的净现值。
	5.不同时间跨度的考虑： NVP考虑了未来现金流的时间价值，因此它对现金流的时间分布敏感。较早发生的现金流对NVP的贡献更大。
NVR计算器使用方法：
	1. 使用完整现金流计算：
	fincalc nvr -f -300,10,20,30 -i 0.08 
	2. 使用单期投资金额、投资期数、单期收益金额、收益期数计算：
	如：每期投资30万，投资3期，每期收益4万，总期数15，第1期开始收益，贴现率0.08，最后一期返还本金
	fincalc nvr -a 30 -p 3 -C 4 -P 15 -S 1 -i 0.08 -r
`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nvr calculating ...")
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

		nvr := service.CalcNvr(cashFlows, irr)
		fmt.Printf("NVR: %.4f\n", nvr)
	},
}

func init() {
	rootCmd.AddCommand(nvrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	nvrCmd.PersistentFlags().Float64VarP(&irr, "irr", "i", 0.05, "内部收益率(不考虑外在因素), 默认为: 0.05")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nvrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
