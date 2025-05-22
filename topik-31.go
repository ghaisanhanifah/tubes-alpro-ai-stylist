package main

import (
	"fmt"
)

// Struktur Data
const MaksData = 100

const (
	KategoriKasual   = "Kasual"
	KategoriFormal   = "Formal"
	KategoriOlahraga = "Olahraga"

	MusimPanas  = "Panas"
	MusimHujan  = "Hujan"
	MusimDingin = "Dingin"

	AcaraKantor   = "Kantor"
	AcaraPesta    = "Pesta"
	AcaraKasual   = "Kasual"
	AcaraFormal   = "Formal"
	AcaraOlahraga = "Olahraga"
)

type Pakaian struct {
	ID              int
	Nama            string
	Warna           string
	Kategori        string
	TingkatFormal   int
	TerakhirDipakai int
	Musim           string
	Acara           string
}

var daftarPakaian [MaksData]Pakaian
var jumlahPakaian int = 0

func input(prompt string) string {
	fmt.Print(prompt)
	var a, b string
	n, _ := fmt.Scanln(&a, &b)
	if n == 2 {
		return a + " " + b
	}
	return a
}

func inputInt(prompt string) int {
	fmt.Print(prompt)
	var val int
	fmt.Scanln(&val)
	return val
}

func tambahPakaian() {
	if jumlahPakaian >= MaksData {
		fmt.Println("Data pakaian sudah penuh.")
		return
	}
	var p Pakaian
	p.ID = jumlahPakaian + 1
	p.Nama = input("Nama Pakaian: ")
	p.Warna = input("Warna: ")
	p.Kategori = input("Kategori (Kasual/Formal/Olahraga): ")
	p.TingkatFormal = inputInt("Tingkat Formalitas (1-5): ")
	p.TerakhirDipakai = inputInt("Tanggal Terakhir Dipakai (YYYYMMDD): ")
	p.Musim = input("Cocok untuk musim apa? (Panas/Hujan/Dingin): ")
	p.Acara = input("Cocok untuk acara apa? (Kasual/Formal/Olahraga/Kantor/Pesta): ")

	daftarPakaian[jumlahPakaian] = p
	jumlahPakaian++
	fmt.Println("Pakaian berhasil ditambahkan.")
}

func tampilkanSemuaPakaian() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada data pakaian.")
		return
	}
	for i := 0; i < jumlahPakaian; i++ {
		tampilkanDetail(daftarPakaian[i])
	}
}

func tampilkanDetail(p Pakaian) {
	fmt.Println("--------------------------")
	fmt.Println("ID:", p.ID)
	fmt.Println("Nama:", p.Nama)
	fmt.Println("Warna:", p.Warna)
	fmt.Println("Kategori:", p.Kategori)
	fmt.Println("Tingkat Formalitas:", p.TingkatFormal)
	fmt.Println("Terakhir Dipakai:", p.TerakhirDipakai)
	fmt.Println("Musim:", p.Musim)
	fmt.Println("Acara:", p.Acara)
}

func ubahPakaian() {
	id := inputInt("Masukkan ID pakaian yang ingin diubah: ")
	idx := cariIndexByID(id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	p := &daftarPakaian[idx]
	p.Nama = input("Nama baru: ")
	p.Warna = input("Warna baru: ")
	p.Kategori = input("Kategori baru: ")
	p.TingkatFormal = inputInt("Tingkat Formalitas baru (1-5): ")
	p.TerakhirDipakai = inputInt("Tanggal Terakhir Dipakai baru (YYYYMMDD): ")
	p.Musim = input("Musim baru (Panas/Hujan/Dingin): ")
	p.Acara = input("Acara baru (Kasual/Formal/Olahraga/Kantor/Pesta): ")
	fmt.Println("Pakaian berhasil diubah.")
}

func hapusPakaian() {
	id := inputInt("Masukkan ID pakaian yang ingin dihapus: ")
	idx := cariIndexByID(id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahPakaian-1; i++ {
		daftarPakaian[i] = daftarPakaian[i+1]
	}
	jumlahPakaian--
	fmt.Println("Pakaian berhasil dihapus.")
}

func cariPakaian() {
	keyword := input("Masukkan warna atau kategori yang dicari: ")
	ketemu := false
	for i := 0; i < jumlahPakaian; i++ {
		if daftarPakaian[i].Warna == keyword || daftarPakaian[i].Kategori == keyword {
			tampilkanDetail(daftarPakaian[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ditemukan pakaian dengan keyword tersebut.")
	}
}

func urutkanPakaian() {
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Tingkat Formalitas")
	fmt.Println("2. Terakhir Dipakai")
	pilih := inputInt("Pilih kategori (1/2): ")
	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	urut := inputInt("Pilih urutan (1/2): ")

	for i := 1; i < jumlahPakaian; i++ {
		key := daftarPakaian[i]
		j := i - 1
		// Pastikan `)` penutup kondisi dan `{` pembuka blok loop berada di baris yang sama
		for j >= 0 &&
			(pilih == 1 && ((urut == 1 && daftarPakaian[j].TingkatFormal > key.TingkatFormal) || (urut == 2 && daftarPakaian[j].TingkatFormal < key.TingkatFormal))) ||
			(pilih == 2 && ((urut == 1 && daftarPakaian[j].TerakhirDipakai > key.TerakhirDipakai) || (urut == 2 && daftarPakaian[j].TerakhirDipakai < key.TerakhirDipakai))) {
			i
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = key
	}
	fmt.Println("Data pakaian berhasil diurutkan.")
	tampilkanSemuaPakaian()
}

func rekomendasiKondisi() {
	cuaca := input("Masukkan kondisi cuaca saat ini (Panas/Hujan/Dingin): ")
	acara := input("Masukkan jenis acara (Kasual/Formal/Olahraga/Kantor/Pesta): ")
	terbaik := -1
	maxSkor := -1

	for i := 0; i < jumlahPakaian; i++ {
		skor := 0
		if daftarPakaian[i].Musim == cuaca {
			skor++
		}
		if daftarPakaian[i].Acara == acara {
			skor++
		}
		if daftarPakaian[i].TingkatFormal >= 3 {
			skor++
		}
		if skor > maxSkor || (skor == maxSkor && daftarPakaian[i].TerakhirDipakai < daftarPakaian[terbaik].TerakhirDipakai) {
			maxSkor = skor
			terbaik = i
		}
	}
	if terbaik != -1 {
		fmt.Println("Rekomendasi terbaik untuk cuaca dan acara saat ini:")
		tampilkanDetail(daftarPakaian[terbaik])
	} else {
		fmt.Println("Tidak ditemukan pakaian yang sesuai kondisi tersebut.")
	}
}

func cariIndexByID(id int) int {
	for i := 0; i < jumlahPakaian; i++ {
		if daftarPakaian[i].ID == id {
			return i
		}
	}
	return -1
}

func main() {
	for {
		fmt.Println("\n==== DIGITAL AI STYLIST ====")
		fmt.Println("1. Tambah Pakaian")
		fmt.Println("2. Ubah Pakaian")
		fmt.Println("3. Hapus Pakaian")
		fmt.Println("4. Cari Pakaian")
		fmt.Println("5. Urutkan Pakaian")
		fmt.Println("6. Rekomendasi Berdasarkan Cuaca dan Acara")
		fmt.Println("7. Tampilkan Semua Pakaian")
		fmt.Println("8. Keluar")
		pilihan := inputInt("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahPakaian()
		case 2:
			ubahPakaian()
		case 3:
			hapusPakaian()
		case 4:
			cariPakaian()
		case 5:
			urutkanPakaian()
		case 6:
			rekomendasiKondisi()
		case 7:
			tampilkanSemuaPakaian()
		case 8:
			fmt.Println("Terima kasih telah menggunakan AI Stylist!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
