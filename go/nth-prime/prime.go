package prime

func Nth(n int) (int, bool) {
	x:=0
if n<=0{
	return 0,false
}
for i:=2;i<=104743;i++{
	if isprime(i){
		x+=1
		if x==n{
			return i,true
		}
	}

}
return 0,false
}
 func isprime(n int)bool{
for i:=2;i<=n;i++{
	if i==n{
		continue
	}
	if n%i==0{
	return false
}
}
return true
 }
