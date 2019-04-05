/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.20%)
 * Total Accepted:    106.3K
 * Total Submissions: 330.3K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */
package main

import (
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	value := 0
	var y int
	if x > 0 {
		y = x
	} else {
		y = -x
	}

	str := strings.Split(strconv.Itoa(y), "")
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	m := ""
	for _, s := range str {
		if s != "0" || m != "" {
			m = m + s
		}
	}
	value, _ = strconv.Atoi(m)

	if x < 0 {
		value = -value
	}
	if float64(value) > math.Pow(2, 31)-1 || float64(value) < -math.Pow(2, 31) {
		value = 0
	}
	return value

}

func reverse2(x int) int {
	var rev int
	intMax := math.Pow(2, 31)
	intMin := -math.Pow(2, 31)
	for x != 0 {
		pop := x % 10
		x = x / 10
		if float64(rev) > intMax/10 || (float64(rev) == intMax/10 && pop > 7) {
			return 0
		}
		if float64(rev) < intMin/10 || (float64(rev) == intMin/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
