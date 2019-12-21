/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.24%)
 * Likes:    978
 * Dislikes: 0
 * Total Accepted:    107.2K
 * Total Submissions: 273.2K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 */
 func isValid(s string) bool {
	if len(s)==0{
		return true
	}
	if string(s[0])==")" || string(s[0])=="}"||string(s[0])=="]"{
		return false
	}
	if len(s)%2!=0{
		return false
	}
	var a []string
	for _,i:=range s{
		if string(i) == "(" || string(i) == "{" || string(i) == "[" {
			a = append(a,string(i))
		} else{
			if a[len(a)-1]=="("{
				if string(i) != ")"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="{"{
				if string(i)!="}"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="["{
				if string(i)!="]"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}
		}
	}
	if len(a) != 0{
		return false
	} else{
		return true
	}
}

