// Обратная польская запись (Reverse Polish Notation): Реализуйте вычисление арифметического выражения в обратной польской записи с использованием стека.

func polishAritchmeticWrite(input string) int {
    stack := list.New()
    for _, char := range strings.Fields(input) {
        if char == "+" || char == "-" || char == "*" || char == "/" {
            if stack.Len() < 2 {
                fmt.Println("Stack is empty")
                return 0
            }
            number_two, err := strconv.Atoi(stack.Back().Value.(string))
            if err != nil {
                return 0
            }
            stack.Remove(stack.Back())
            number_one, err := strconv.Atoi(stack.Back().Value.(string))
            if err != nil {
                return 0
            }
            stack.Remove(stack.Back())
            switch char {
            case "+":
                result := number_one + number_two
                stack.PushBack(strconv.Itoa(result))
            case "-":
                result := number_one - number_two
                stack.PushBack(strconv.Itoa(result))
            case "*":
                result := number_one * number_two
                stack.PushBack(strconv.Itoa(result))
            case "/":
                if number_two == 0 {
                    fmt.Println("Division by zero")
                    return 0
                }
                result := number_one / number_two
                stack.PushBack(strconv.Itoa(result))
            }
        } else {
            stack.PushBack(char)
        }
    }
    result, _ := strconv.Atoi(stack.Front().Value.(string))
    return result
}
