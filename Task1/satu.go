package Task1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	Name string
	ID   int
}

var (
	morningActivities = make(chan string, 10)
	workSignal        = make(chan bool, 5)
	done              = make(chan bool)
)

func MorningRoutine(worker Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	
	activities := []string{
		fmt.Sprintf("%s sedang mandi", worker.Name),
		fmt.Sprintf("%s membuat kopi", worker.Name),
		fmt.Sprintf("%s menyiapkan sarapan", worker.Name),
		fmt.Sprintf("%s merapikan kamar tidur", worker.Name),
	}
	
	for _, activity := range activities {
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		
		morningActivities <- activity
		fmt.Printf("[%s] %s \n", 
			time.Now().Format("15:04:05"), activity)
	}
}

func WorkSchedule(worker Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	
	<-workSignal
	
	fmt.Printf("[%s] %s mulai bekerja \n", 
		time.Now().Format("15:04:05"), worker.Name)
	
	workDuration := time.Duration(rand.Intn(4)+3) * time.Second
	time.Sleep(workDuration)
	
	fmt.Printf("[%s] %s selesai bekerja\n", 
		time.Now().Format("15:04:05"), worker.Name)
}

func SimulateWorkerLife(workers []Worker) {
	var wg sync.WaitGroup
	
	fmt.Printf("\nWaktu mulai: %s\n", time.Now().Format("15:04:05"))
	fmt.Println("\nRUTINITAS PAGI:")
	
	for _, worker := range workers {
		wg.Add(1)
		go MorningRoutine(worker, &wg)
	}
	
	for _, worker := range workers {
		wg.Add(1)
		go WorkSchedule(worker, &wg)
	}
	
	go func() {
		activityCount := 0
		expectedActivities := len(workers) * 4 
		
		for range morningActivities {
			activityCount++
			if activityCount == expectedActivities {
				fmt.Printf("\n[%s] Semua rutinitas pagi selesai!\n", 
					time.Now().Format("15:04:05"))
				
				fmt.Println("\nJADWAL KERJA:")
				for range workers {
					workSignal <- true
				}
				close(morningActivities)
				break
			}
		}
	}()
	
	go func() {
		wg.Wait()
		fmt.Printf("\n[%s]Berangkat kerja! Semua pekerja selesai!\n", 
			time.Now().Format("15:04:05"))
		done <- true
	}()
	
	 <-done

	
	wg.Wait()
}