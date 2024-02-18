// Реализация очереди через массив с функциями использования, удаления элементов и вывод размера очереди
// Практика с работой с очередью
type Queue struct {
	items []int
}

func (queue *Queue) AddElement(element int) {
	queue.items = append(queue.items, element)
	fmt.Println("Added element: ", element)
}
func (queue *Queue) UseElement() int {
	if len(queue.items) == 0 {
		fmt.Println("Queue is empty")
	}
	element := queue.items[0]
	queue.items = queue.items[1:]
	fmt.Println("Used element: ", element)

	return element

}
func (queue *Queue) LenQueue() int {
	fmt.Println("Len queue -> ", len(queue.items))
	return 1
}

func main() {
	queue := &Queue{}
	queue.AddElement(1)
	queue.AddElement(2)
	queue.AddElement(3)
	queue.AddElement(4)
	queue.AddElement(5)
	fmt.Println(queue.UseElement())
	fmt.Println(queue.UseElement())
	queue.LenQueue()
}
