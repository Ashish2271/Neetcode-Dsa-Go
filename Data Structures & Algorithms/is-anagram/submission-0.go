func isAnagram(s string, t string) bool {
if len(s) != len(t){
	return false
}

chars1 := []rune(s)
chars2 := []rune(t)

l1 := make(map[rune]int)
l2 := make(map[rune]int)
for _,v := range chars1 {
	l1[v]++
} 
for _,v := range chars2 {
	l2[v]++
} 

for k,v := range l1{
	if v != l2[k]{
		return false
	} 
}
return true 

}
