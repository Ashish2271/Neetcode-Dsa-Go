func groupAnagrams(strs []string) [][]string {

	bigarr := [][]string{}
	used := make([]bool, len(strs))
	
for i:=0; i<len(strs); i++{
	if used[i]{
		continue 
	}
	arr := []string{strs[i]}
	
  for j:=i+1; j<len(strs); j++{
	
	if len(strs[i]) == len(strs[j]){
		
		 if isAnagram(strs[i],strs[j]){
			used[j] = true
	arr = append(arr, strs[j])
  }
                                               
	}

   }
   bigarr = append(bigarr,arr)

}
 return bigarr 
}



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