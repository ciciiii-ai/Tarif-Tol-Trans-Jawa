package main

import "fmt"

func main() {
	// DATA UTAMA
	// Kita simpan nama kota dan posisi KM di dua daftar terpisah
	// Supaya kodenya tidak terlihat rumit
	namaKota := []string{"Ngawi", "Madiun", "Nganjuk", "Mojokerto", "Surabaya", "Malang"}
	posisiKM := []int{579, 603, 647, 712, 745, 790}

	// 1. TAMPILKAN MENU (List Menurun)
	fmt.Println("=== DAFTAR GERBANG TOL ===")
	// Kita pakai perulangan (loop) untuk mencetak daftar ke bawah
	for i := 0; i < len(namaKota); i++ {
		// %d = angka urut, %s = nama kota
		fmt.Printf("%d. %s\n", i+1, namaKota[i])
	}

	// 2. INPUT GERBANG (Satu per satu)
	var masuk, keluar int

	fmt.Print("\nNomor Gerbang Masuk : ")
	fmt.Scan(&masuk)

	fmt.Print("Nomor Gerbang Keluar: ")
	fmt.Scan(&keluar)

	// Cek apakah nomor yang dimasukkan benar
	// Angka tidak boleh kurang dari 1 atau lebih dari jumlah kota (6)
	if masuk < 1 || masuk > 6 || keluar < 1 || keluar > 6 {
		fmt.Println("Error: Nomor gerbang tidak ada!")
		return
	}

	// 3. HITUNG JARAK
	// Ambil KM dari daftar 'posisiKM'. Ingat, komputer mulai hitung dari 0, jadi kita kurangi 1.
	jarak := posisiKM[keluar-1] - posisiKM[masuk-1]

	// Jika hasilnya minus (karena balik arah), ubah jadi positif
	if jarak < 0 {
		jarak = -jarak
	}

	// 4. PILIH GOLONGAN
	fmt.Println("\n--- Jenis Kendaraan ---")
	fmt.Println("1. Mobil   (Golongan 1)")
	fmt.Println("2. Truk    (Golongan 2)")
	fmt.Println("3. Trailer (Golongan 3)")

	var golongan, hargaPerKM int
	fmt.Print("Pilih (1-3): ")
	fmt.Scan(&golongan)

	if golongan == 1 {
		hargaPerKM = 1000
	} else if golongan == 2 {
		hargaPerKM = 1500
	} else {
		hargaPerKM = 2000
	}

	totalHarga := jarak * hargaPerKM

	// 5. PEMBAYARAN
	var uangBayar int
	fmt.Printf("\nTotal Tagihan : Rp %d\n", totalHarga)
	fmt.Print("Masukkan Uang : Rp ")
	fmt.Scan(&uangBayar)

	// 6. CETAK STRUK (Jika uang cukup)
	if uangBayar >= totalHarga {
		kembalian := uangBayar - totalHarga

		fmt.Println("\n")
		fmt.Println("===============================")
		fmt.Println("      STRUK TRANSAKSI TOL      ")
		fmt.Println("===============================")
		// Mengambil nama kota dari daftar 'namaKota'
		fmt.Printf("Asal      : %s\n", namaKota[masuk-1])
		fmt.Printf("Tujuan    : %s\n", namaKota[keluar-1])
		fmt.Printf("Jarak     : %d KM\n", jarak)
		fmt.Println("-------------------------------")
		fmt.Printf("Tagihan   : Rp %d\n", totalHarga)
		fmt.Printf("Bayar     : Rp %d\n", uangBayar)
		fmt.Println("-------------------------------")
		fmt.Printf("KEMBALI   : Rp %d\n", kembalian)
		fmt.Println("===============================")
		fmt.Println("    TERIMA KASIH HATI-HATI     ")
	} else {
		fmt.Println("\nTRANSAKSI GAGAL: Uang Anda Kurang!")
	}
}
