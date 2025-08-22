package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Anggiprayoga28/Task1"
	"github.com/Anggiprayoga28/Task2"
	"github.com/Anggiprayoga28/Task3"
)

func main() {
// Nomor 1
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("SIMULASI KEHIDUPAN PEKERJA KANTORAN")
	
	workers := []Task1.Worker{
		{Name: "Sidik", ID: 1},
		{Name: "Bobi", ID: 2},
		{Name: "Alex", ID: 3},
	}
	
	Task1.SimulateWorkerLife(workers)
	
	
	fmt.Println("\nProgram selesai!")


// Nomor 2

	channelPesan := make(chan Task2.Message, 5)
	channelSelesai := make(chan bool)
	
	fmt.Println("PAPAN TULIS VIRTUAL KELUARGA")
	go Task2.PapanTulis(channelPesan, channelSelesai)
	time.Sleep(500 * time.Millisecond)
	Task2.KirimBanyakPesan(channelPesan)
	
	
	fmt.Println("\nPesan tambahan:")
	Task2.KirimPesan(channelPesan, "Nenek", "Kapan main ke rumah nenek?")
	time.Sleep(1 * time.Second)
	
	Task2.KirimPesan(channelPesan, "Kakek", "Jangan lupa siram tanaman")
	time.Sleep(1 * time.Second)
	
	time.Sleep(2 * time.Second)
	
	fmt.Println("\nMenutup papan tulis...")
	close(channelPesan)
	
	<-channelSelesai
	
	fmt.Println("Program selesai!")

//Nomor 3
	fmt.Println("=== Simulasi Berbagi Microwave di Kontrakan ===")
	fmt.Println()

	microwave := &Task3.Microwave{}

	housemates := []string{"Anggi", "Selamet", "Titus", "Farid", "Sidik"}

	var wg sync.WaitGroup


	for i, name := range housemates {
		wg.Add(1)
		go func(person string, delay int) {
			defer wg.Done()
			
			time.Sleep(time.Duration(delay) * time.Millisecond * 100)
			
			heatDuration := time.Duration(rand.Intn(3)+1) * time.Second
			
			fmt.Printf("[%s] %s ingin menggunakan microwave\n", 
				time.Now().Format("15:04:05"), person)
			
			microwave.Use(person, heatDuration)
		}(name, i)
	}

	wg.Wait()

	for i, name := range housemates {
		wg.Add(1)
		go func(person string, delay int) {
			defer wg.Done()
			
			time.Sleep(time.Duration(delay) * time.Millisecond * 200)
			
			heatDuration := time.Duration(rand.Intn(2)+1) * time.Second
			
			fmt.Printf("[%s] %s mencoba menggunakan microwave\n", 
				time.Now().Format("15:04:05"), person)
			
			if !microwave.TryUse(person, heatDuration) {
				time.Sleep(1 * time.Second)
				fmt.Printf("[%s] %s mencoba lagi...\n", 
					time.Now().Format("15:04:05"), person)
				microwave.Use(person, heatDuration)
			}
		}(name, i)
	}

	wg.Wait()

	fmt.Println()
	fmt.Println("Simulasi Selesai")
}