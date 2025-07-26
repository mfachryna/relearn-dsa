package relearndsa

func isValid(s string) bool {
    var stack []rune

	for _, val := range s {
		if val == '(' ||  val == '[' ||  val == '{' {
            stack = append(stack,val)
            continue
        }
        if len(stack) == 0 {
            return false
        }

        lastStack := stack[len(stack)-1]
        if (val == ')' && lastStack == '(') || (val == '}' && lastStack == '{') || (val == ']' && lastStack == '[') {
            stack = stack[:len(stack)-1]
            continue
        }
        return false
    }
    if (len(stack) > 0) {
        return false
    }
    return true
}