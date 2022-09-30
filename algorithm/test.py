#
# 
# @param s string字符串 
# @return bool布尔型
#
class Solution:
    def isValid(self , s ):
        # write code here
        
        
        stack=[]
        for i in s:
            if i in ['(','{','[']:
                stack.append(i)
            else:
                if not stack:return False
                tmp=stack.pop()
                if (i==')' and tmp!='(') or (i==']' and tmp!='[') or (i=='}' and tmp!='{'):
                    return False
        return not stack
        
        
        
        
print(Solution().isValid("([)]"))
