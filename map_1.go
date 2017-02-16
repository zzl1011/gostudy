package main

import "fmt"

func main() {
	week := make(map[int]string)
	week[1] = "Monday"
	week[2] = "Tuesday"
	week[3] = "Wednesday"
	week[4] = "Thursday"
	week[5] = "Friday"
	week[6] = "Saturday"
	week[7] = "Sunday"

	for key, value := range week {
		fmt.Printf("key is: %d, value is %s\n", key, value)
		if value == "Tuesday" {
			fmt.Println("Tuesday is exist")
		}
		if value == "Hollyday" {
			fmt.Println("Hollyday is exist")
		}
	}


}
