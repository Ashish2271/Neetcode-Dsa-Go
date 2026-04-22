type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    return fmt.Sprintf("%d§%s", len(strs), strings.Join(strs, "§"))
}

func (s *Solution) Decode(encoded string) []string {
    parts := strings.SplitN(encoded, "§", 2)
    count, _ := strconv.Atoi(parts[0])
    
    if count == 0 {
        return []string{}
    }
    if parts[1] == "" {
        return make([]string, count) // returns [""] for count=1
    }
    
    return strings.Split(parts[1], "§")
}

