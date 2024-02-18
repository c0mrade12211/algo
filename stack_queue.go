//Tasks on stack and queue 

//It's simple palindrome which was realized in stack 

func isPalindrome(input string) bool {
	stack := list.New()
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	for i := 0; i < len(input); i++ {
		stack.PushBack(input[i])
		fmt.Println("It's value will add in stack: ", input[i])
	}
	for stack.Len() > 1 {
		firstElement := stack.Front().Value
		lastElement := stack.Back().Value
		if firstElement != lastElement {
			return false
		}
		stack.Remove(stack.Front())
		stack.Remove(stack.Back())
	}
	return true
}
