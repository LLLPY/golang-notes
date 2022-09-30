package main

/**
  * 
  * @param s string字符串 
  * @return bool布尔型
*/
func isValid( s string ) bool {
    // write code here

	stack := [10000]string
	for i:=0;i<len(s);i++{
		if s[i]=="(" || s[i]=="[" || s[i]=="{"{
			stack=append(stack,s[i])
		}
		else{

			if len(stack)==0{
				return false
			}
			tmp:=stack[]

		}



	}

}


func main(){

	s:="((())){}"
	isValid(s)



}

