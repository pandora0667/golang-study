package main

import f "fmt"

func main() {

	a := make(map[string]int)

	a["age"] = 27
	a["height"] = 175
	f.Println(a)

	b := map[string]float64{
		"pi":    3.141592,
		"sqrt2": 1.41421356,
	}

	f.Println(b["pi"], b["sqrt2"])

	f.Println("-------------------")

	capacityUnit := make(map[string]string)

	capacityUnit["1byte"] = "1024 bit"
	capacityUnit["1MB"] = "1024 Byte"
	capacityUnit["1GB"] = "1024 MB"
	capacityUnit["1TB"] = "1024 GB"
	capacityUnit["1PB"] = "1024 TB"

	value, ok := capacityUnit["1YB"]
	f.Println(value, ok)

	if result, ok := capacityUnit["1GB"]; ok {
		f.Println(result, ok)
	}

	f.Println("-------------------")

	for key, result := range capacityUnit {
		f.Println(key, " : ", result)
	}

	f.Println("-------------------")

	for _, result := range capacityUnit {
		f.Println(result)
	}

	delete(capacityUnit, "1PB")
	f.Println(capacityUnit)
}
