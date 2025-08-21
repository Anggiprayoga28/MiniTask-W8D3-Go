package Task2

import (
	"fmt"
	"time"
)

type Message struct {
	Author  string
	Content string
	Time    string
}

func KirimPesanKeChannel(ch chan Message, author, content string) {
	waktu := time.Now().Format("15:04")
	pesan := Message{
		Author:  author,
		Content: content,
		Time:    waktu,
	}
	ch <- pesan
	fmt.Printf("Pesan dari %s berhasil dikirim\n", author)
}

func BuatPesan(author, content string) Message {
	waktu := time.Now().Format("15:04")
	return Message{
		Author:  author,
		Content: content,
		Time:    waktu,
	}
}

func KirimPesan(ch chan Message, author, content string) {
	pesan := BuatPesan(author, content)
	ch <- pesan
	fmt.Printf("Pesan dari %s berhasil dikirim\n", author)
}

func PapanTulis(ch chan Message, selesai chan bool) {
	
	for {
		pesan, masihTerbuka := <-ch
		if !masihTerbuka {
			fmt.Println("Channel tertutup, papan tulis berhenti")
			selesai <- true
			return
		}
		fmt.Printf("[%s] %s: %s\n", pesan.Time, pesan.Author, pesan.Content)
	}
}

func KirimBanyakPesan(ch chan Message) {
	fmt.Println("\nMengirim pesan dari berbagai anggota keluarga...")
	
	daftarPesan := []struct {
		pengirim string
		isi      string
	}{
		{"Ayah", "Jangan lupa beli susu"},
		{"Ibu", "Makan siang sudah siap"},
		{"Kakak", "Mau nonton film?"},
		{"Adik", "Besok ada ulangan"},
		{"Ayah", "Pulang agak telat"},
		{"Ibu", "Cuci piring ya"},
	}
	
	for i, data := range daftarPesan {
		KirimPesanKeChannel(ch, data.pengirim, data.isi)
		time.Sleep(1 * time.Second)
		fmt.Printf("Progress: %d/%d\n", i+1, len(daftarPesan))
	}
	
	fmt.Println("Semua pesan berhasil dikirim!")
}