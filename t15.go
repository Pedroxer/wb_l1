package main

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Строки в go неизменяемые и когда делается слайс от строки создаётся новая строка с ссылкой на подстроку.
// Так как в someFunc() может создаться очень большая строка, это может привести к большому потреблению памяти.
// Ведь, при создании слайса его ёмкость будет равна длине строки в данном случае.

// Решением может быть создание новой строки и копирования в неё строки v с помощью copy.
// Таким образом, мы можем ограничить кол-во занимаемой памяти.
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v[:100])
	justString = string(tmp)
}

func main() {
	someFunc()
}
