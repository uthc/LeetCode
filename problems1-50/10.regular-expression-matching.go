// 10. 正则表达式匹配

// 给定一个字符串 (s) 和一个字符模式 (p)。实现支持 '.' 和 '*' 的正则表达式匹配。

// '.' 匹配任意单个字符。
// '*' 匹配零个或多个前面的元素。
// 匹配应该覆盖整个字符串 (s) ，而不是部分字符串。

// 说明:

// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
// 示例 1:

// 输入:
// s = "aa"
// p = "a"
// 输出: false
// 解释: "a" 无法匹配 "aa" 整个字符串。
// 示例 2:

// 输入:
// s = "aa"
// p = "a*"
// 输出: true
// 解释: '*' 代表可匹配零个或多个前面的元素, 即可以匹配 'a' 。因此, 重复 'a' 一次, 字符串可变为 "aa"。
// 示例 3:

// 输入:
// s = "ab"
// p = ".*"
// 输出: true
// 解释: ".*" 表示可匹配零个或多个('*')任意字符('.')。
// 示例 4:

// 输入:
// s = "aab"
// p = "c*a*b"
// 输出: true
// 解释: 'c' 可以不被重复, 'a' 可以被重复一次。因此可以匹配字符串 "aab"。
// 示例 5:

// 输入:
// s = "mississippi"
// p = "mis*is*p*."
// 输出: false

func isMatch(s string, p string) bool {
    if len(s)==0&&len(p)==0{
        return true
    }
    if len(p)==1{
        if p==s{
            return true
        }
		if p=="."&&len(s)==1{
            return true
        }        
        return false
    }
    judge:=true
    si,pi:=0,0
    for judge{
        if si>=len(s)||pi>=len(p){
            break
        }
        if pi<len(p)-1&&string(p[pi+1])=="*"{            
            //匹配0次到n次    
			if pi+2==len(p){
				for si<len(s){
					if !isMatch(string(s[si]),string(p[pi])){
						return false
					}
					si++
				}
				return true
			}
            n:=0
			for ;si+n<len(s)&&isMatch(string(s[si+n]), string(p[pi]));n++{}			
            for x:=0;x<=n;x++{				
				if isMatch(string(s[si+x:]), string(p[pi+2:])){
                    return true
                }				
            }
			return false
            //pi=pi+2
        }else{
            //匹配1次
            judge=isMatch(string(s[si]), string(p[pi]))
            si++
            pi++
        }
    }
    if judge==false{
        return false
    }
	if si<len(s){
		return false
	}
    if si==len(s)&&pi==len(p){
        return true
	}			
	if pi==len(p)-1||p[len(p)-1]!='*'{
		return false
	}
	pi++
	for ;pi<len(p);pi=pi+2{
		if p[pi]!='*'{
			return false
		}
	}
    return judge
}