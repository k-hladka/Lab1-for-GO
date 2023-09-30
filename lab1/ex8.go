package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	//Задание.
	//1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести
	var (
		defaultInt8  int8
		defaultInt16 int16
		defaultInt32 int32
	)
	fmt.Printf("delaultInt8 = %d\ndefaultInt16 = %d\ndefaultInt32 = %d\n", defaultInt8, defaultInt16, defaultInt32)
}
