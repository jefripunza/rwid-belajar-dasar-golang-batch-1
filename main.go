package main

import (
	"fmt"
	"time"
)

func main() {

	// ----------------------------------------------------- //
	// Kondisional : if else switch select statement

	// if else
	human := false
	if human {
		fmt.Println("I'm a Human")
	} else {
		fmt.Println("I'm not a Human")
	}

	// switch
	switch {
	case human:
		fmt.Println("I'm a Human")
	case !human:
		fmt.Println("I'm not a Human")
	default:
		fmt.Println("I don't know")
	}

	// select statement
	// Buat channel untuk komunikasi
	humanChan := make(chan string)

	// Jalankan goroutine untuk mengirim pesan ke channel
	go func() {
		time.Sleep(3 * time.Second) // Simulasikan proses yang memakan waktu
		humanChan <- "I'm a Human"  // Kirim pesan ke channel
	}()

	// Select statement untuk memilih case yang siap
	select {
	case message := <-humanChan: // Tunggu pesan dari goroutine
		fmt.Println(message)
	case <-time.After(2 * time.Second): // Timeout jika channel tidak menerima pesan dalam 2 detik
		fmt.Println("No response, I don't know")
	}

	// ----------------------------------------------------- //
	// Looping : for, for range, break, continue

	// for
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("for %d\n", i)
	}

	// for range
	for _, i := range []int{1, 2, 3, 4, 5} {
		if i == 2 {
			continue
		}
		fmt.Printf("range %d\n", i)
	}

	// for {
	// 	fmt.Println("Hello World")
	// }

	// enum

	// Define a custom type for Color
	type Color string

	// Define constants for each color
	const (
		Red   Color = "RED"
		Green Color = "GREEN"
		Blue  Color = "BLUE"
	)

	// Assign the color
	var warna Color = Blue

	// Use the color
	fmt.Println("Selected color:", warna)

}
