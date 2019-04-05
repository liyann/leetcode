/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (56.00%)
 * Total Accepted:    92.4K
 * Total Submissions: 165K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */
package main

func isPalindrome(x int) bool {
	// 转成字符串处理
	// str := strconv.Itoa(x)
	// strArr := strings.Split(str, "")
	// for i, j := 0, len(strArr)-1; i <= len(strArr)/2; i, j = i+1, j-1 {
	// 	if strArr[i] == strArr[j] {
	// 		continue
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true

	// 当做数字处理
	// if x < 0 {
	// 	return false
	// }
	// var s = []int{}
	// for x >= 1 {
	// 	y := x % 10
	// 	s = append(s, y)
	// 	x = (x - y) / 10
	// }
	// for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
	// 	if s[i] == s[j] {
	// 		continue
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true

	// 题解
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var y int
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10
}
