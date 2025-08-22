package Task3

import (
	"fmt"
	"sync"
	"time"
)

type Microwave struct {
	mu     sync.Mutex
	inUse  bool
	userID string
}

func (m *Microwave) Use(userID string, duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	fmt.Printf("[%s] %s mulai menggunakan microwave\n", 
		time.Now().Format("15:04:05"), userID)
	
	m.inUse = true
	m.userID = userID

	time.Sleep(duration)

	fmt.Printf("[%s] %s selesai menggunakan microwave\n", 
		time.Now().Format("15:04:05"), userID)
	
	m.inUse = false
	m.userID = ""
}

func (m *Microwave) TryUse(userID string, duration time.Duration) bool {
	if m.mu.TryLock() {
		defer m.mu.Unlock()
		
		fmt.Printf("[%s] %s berhasil menggunakan microwave\n", 
			time.Now().Format("15:04:05"), userID)
		
		m.inUse = true
		m.userID = userID
		
		time.Sleep(duration)
		
		fmt.Printf("[%s] %s selesai menggunakan microwave\n", 
			time.Now().Format("15:04:05"), userID)
		
		m.inUse = false
		m.userID = ""
		return true
	}
	
	fmt.Printf("[%s] %s tidak bisa menggunakan microwave (sedang dipakai)\n", 
		time.Now().Format("15:04:05"), userID)
	return false
}

