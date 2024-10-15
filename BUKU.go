package main

import (
	"fmt"
)

// Struktur untuk buku
type Buku struct {
	Judul  string
	Jumlah int
}

// Struktur untuk histori transaksi
type History struct {
	JudulBuku      string
	Jumlah         int
	JenisTransaksi string // "Pinjam" atau "Kembali"
}

var (
	historyPeminjaman []History // Menyimpan histori peminjaman
	bukuPemrograman   = []Buku{
		{"Tips Belajar Go Language", 10},
	}
	bukuFilm = []Buku{
		{"Once Upon a Time in Hollywood", 5},
	}
	bukuPrinting = []Buku{
		{"How to 3D Print", 20},
	}
)

func main() {
	fmt.Println("SELAMAT DATANG DI PERPUSTAKAAN \n VOKASI")

	// Sistem login
	var username, password string
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	// Validasi login
	if username == "Humam" && password == "2406424221" {
		fmt.Println("LOGIN BERHASIL >_<")
		MainMenu()
	} else {
		fmt.Println("LOGIN GAGAL")
	}
}

func MainMenu() {
	// Informasi pengguna
	username := "Humam"
	namaLengkap := "Humawan Ilham Marzuki"
	NPM := "2406424221"
	jenisKelamin := "Laki-laki"
	TTL := "Cilacap, 26 Juni 2005"
	umur := "19"
	MBTI := "INFP"
	makananFavorit := "Nasi Padang Ati Ampela"
	minumanFavorit := "CHATIME TARO PAKE GRASS JELLY"
	Bukufavorit := "The Song of Achilles"
	FilmFavorit := "The Florida Project"

	var pilihAngka int

	for {
		fmt.Println("\n---Menu Peminjaman Buku---")
		fmt.Println("1. Informasi Pengguna")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Menu (1-6): ")
		fmt.Scan(&pilihAngka)

		switch pilihAngka {
		case 1:
			fmt.Println("\n---INFORMASI PENGGUNA---")
			fmt.Printf("Username: %s\n", username)
			fmt.Printf("Nama Lengkap: %s\n", namaLengkap)
			fmt.Printf("NPM: %s\n,", NPM)
			fmt.Printf("Jenis kelamin: %s\n", jenisKelamin)
			fmt.Printf("TTL: %s\n", TTL)
			fmt.Printf("Umur: %s\n", umur)
			fmt.Printf("MBTI: %s\n", MBTI)
			fmt.Printf("Makanan Favorit: %s\n", makananFavorit)
			fmt.Printf("Minuman Favorit: %s\n", minumanFavorit)
			fmt.Printf("Buku Favorit: %s\n", Bukufavorit)
			fmt.Printf("Film Favorit: %s\n", FilmFavorit)

		case 2:
			fmt.Println("---DAFTAR BUKU PERPUSTAKAAN---")
			fmt.Println("1. Buku Pemrograman:")
			for i, buku := range bukuPemrograman {
				fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
			}
			fmt.Println("\n2. Buku Film:")
			for i, buku := range bukuFilm {
				fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
			}
			fmt.Println("\n3. Buku Printing:")
			for i, buku := range bukuPrinting {
				fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
			}

		case 3:
			//Menambahkan Daftar Buku = Mengembalikan Buku
			var kategori, jumlahBuku, indexBuku int
			fmt.Println("\nPilih Kategori Buku untuk Kembali:")
			fmt.Println("1. Buku Pemrograman")
			fmt.Println("2. Buku Film")
			fmt.Println("3. Buku Printing")
			fmt.Print("Pilih kategori (1-3): ")
			fmt.Scan(&kategori)

			switch kategori {
			case 1:
				fmt.Println("---DAFTAR BUKU PEMROGRAMAN---")
				for i, buku := range bukuPemrograman {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			case 2:
				fmt.Println("---DAFTAR BUKU FILM---")
				for i, buku := range bukuFilm {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			case 3:
				fmt.Println("---DAFTAR BUKU PRINTING---")
				for i, buku := range bukuPrinting {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			default:
				fmt.Println("Kategori tidak valid.")
				continue
			}

			fmt.Print("Pilih buku (nomor): ")
			fmt.Scan(&indexBuku)
			fmt.Print("Masukkan jumlah: ")
			fmt.Scan(&jumlahBuku)

			// Proses pengembalian buku
			var bukuKembali *Buku
			switch kategori {
			case 1:
				if indexBuku > 0 && indexBuku <= len(bukuPemrograman) {
					bukuKembali = &bukuPemrograman[indexBuku-1]
				}
			case 2:
				if indexBuku > 0 && indexBuku <= len(bukuFilm) {
					bukuKembali = &bukuFilm[indexBuku-1]
				}
			case 3:
				if indexBuku > 0 && indexBuku <= len(bukuPrinting) {
					bukuKembali = &bukuPrinting[indexBuku-1]
				}
			}

			if bukuKembali != nil {
				bukuKembali.Jumlah += jumlahBuku
				historyPeminjaman = append(historyPeminjaman, History{bukuKembali.Judul, jumlahBuku, "Kembali"})
				fmt.Println("TERIMA KASIH TELAH MENGEMBALIKAN BUKU.")
			} else {
				fmt.Println("Buku tidak tersedia.")
			}

		case 4:
			// Meminjam buku = mengurangi Buku
			var kategori, jumlahBuku, indexBuku int
			fmt.Println("\nPilih Kategori Buku:")
			fmt.Println("1. Buku Pemrograman")
			fmt.Println("2. Buku Film")
			fmt.Println("3. Buku Printing")
			fmt.Print("Pilih kategori (1-3): ")
			fmt.Scan(&kategori)

			switch kategori {
			case 1:
				fmt.Println("---DAFTAR BUKU PEMROGRAMAN---")
				for i, buku := range bukuPemrograman {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			case 2:
				fmt.Println("---DAFTAR BUKU FILM---")
				for i, buku := range bukuFilm {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			case 3:
				fmt.Println("---DAFTAR BUKU PRINTING---")
				for i, buku := range bukuPrinting {
					fmt.Printf("%d. Judul: %s, Jumlah: %d\n", i+1, buku.Judul, buku.Jumlah)
				}
			default:
				fmt.Println("Kategori tidak valid.")
				continue
			}

			fmt.Print("Pilih buku (nomor): ")
			fmt.Scan(&indexBuku)
			fmt.Print("Masukkan jumlah: ")
			fmt.Scan(&jumlahBuku)

			// Proses peminjaman buku
			var bukuPinjam *Buku
			switch kategori {
			case 1:
				if indexBuku > 0 && indexBuku <= len(bukuPemrograman) {
					bukuPinjam = &bukuPemrograman[indexBuku-1]
				}
			case 2:
				if indexBuku > 0 && indexBuku <= len(bukuFilm) {
					bukuPinjam = &bukuFilm[indexBuku-1]
				}
			case 3:
				if indexBuku > 0 && indexBuku <= len(bukuPrinting) {
					bukuPinjam = &bukuPrinting[indexBuku-1]
				}
			}

			if bukuPinjam != nil {
				if jumlahBuku <= bukuPinjam.Jumlah {
					bukuPinjam.Jumlah -= jumlahBuku
					historyPeminjaman = append(historyPeminjaman, History{bukuPinjam.Judul, jumlahBuku, "Pinjam"})
					fmt.Println("BUKU BERHASIL DIPINJAM. SELAMAT MEMBACA >_< ")
				} else {
					fmt.Println("Yah, tidak bisa pinjam. Jumlah buku tidak cukup.")
				}
			} else {
				fmt.Println("Buku tidak tersedia.")
			}

		case 5:
			// Tampilkan histori peminjaman
			fmt.Println("---HISTORI PEMINJAMAN---")
			for _, history := range historyPeminjaman {
				fmt.Printf("Buku: %s, Jumlah: %d, Tipe: %s\n", history.JudulBuku, history.Jumlah, history.JenisTransaksi)
			}

		case 6:
			fmt.Println("Terima kasih telah menggunakan sistem ini.")
			return // Keluar dari program

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
