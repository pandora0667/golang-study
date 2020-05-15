package main

import f "fmt"

type subject struct {
	operatingSystem float32
	database        float32
	network         float32
	java            float32
}

func main() {

	student := subject{operatingSystem: 50, database: 30, network: 60, java: 50}
	f.Println(average(&student))

}

func average(sub *subject) float32 {

	result := (sub.operatingSystem + sub.database + sub.network + sub.java) / 4
	return result

}
