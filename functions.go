package main

import "math/rand"

func Array() {
	// Array mit 20 Zufallszahlen
	var zufallszahlen [20]int
	for i := 0; i < 20; i++ {
		zufallszahlen[i] = rand.Intn(100)

		// Ausgabe der Zufallszahlen
		println(zufallszahlen[i])

	}

	var summe int
	for i := 0; i < 20; i++ {
		summe += zufallszahlen[i]

	}
	println("Summe aller Zufallszahlen: ", summe)

	var durchschnitt int
	durchschnitt = int(summe) / int(20)
	println("Durchschnitt aller Zufallszahlen: ", durchschnitt)

	// Slice aus Array
	var zufallszahlen2 []int
	for i := 0; i < 5; i++ {
		zufallszahlen2 = append(zufallszahlen2, zufallszahlen[i])
		// Ausgabe der Zufallszahlen
		println(zufallszahlen2[i])

	}

}
