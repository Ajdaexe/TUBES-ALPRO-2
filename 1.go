package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

func main() {
	var keluar bool

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Hapus minat")
		fmt.Println("6. Hapus keahlian")
		fmt.Println("7. Karir rekomendasi")
		fmt.Println("8. Urut rentang gaji/bulan")
		fmt.Println("9. Statistik kecocokan karier")
		fmt.Println("10. Tampilkan Minat ")
		fmt.Println("11. Tampilkan Keahlian ")
		fmt.Println("12. Keluar")

		var pilihan int
		fmt.Print("\nMasukkan pilihan (1-12): ")
		_, err := fmt.Scan(&pilihan)
		if err != nil {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			hapusMinat()
		case 6:
			hapusKeahlian()
		case 7:
			karirRekomendasi()
		case 8:
			urutRentangGaji()
		case 9:
			statistikKecocokanKarir()
		case 10:
			tampilminat()
		case 11:
			tampilkeahlian()
		case 12:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-10")
		}
	}
}

func tambahMinat() {
	fmt.Println("\n=== Tambah Minat ===")

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	// Tampilkan daftar minat
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	// Input pilihan
	var pilihan int
	fmt.Print("\nPilih nomor minat (1-8): ")
	_, err := fmt.Scan(&pilihan)
	if err != nil || pilihan < 1 || pilihan > len(daftarMinat) {
		fmt.Println("Input tidak valid")
		return
	}

	// Tambahkan ke list
	minatTerpilih := daftarMinat[pilihan-1]
	minatList = append(minatList, minatTerpilih)
	fmt.Printf("Minat '%s' berhasil ditambahkan\n", minatTerpilih)
}

func tambahKeahlian() {
	fmt.Println("\n=== Tambah Keahlian ===")

	var keahlian string
	fmt.Print("Masukkan keahlian baru: ")
	fmt.Scan(&keahlian)

	keahlianList = append(keahlianList, keahlian)
	fmt.Printf("Keahlian '%s' berhasil ditambahkan\n", keahlian)
}

func editMinat() {
	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}

	// Tampilkan minat yang ada
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	// Pilih minat yang akan diedit
	var pilihan int
	fmt.Print("Pilih nomor minat yang akan diedit: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	// Input minat baru
	var minatBaru string
	fmt.Print("Masukkan minat baru: ")
	fmt.Scan(&minatBaru)

	minatList[pilihan-1] = minatBaru
	fmt.Println("Minat berhasil diupdate")
}

func editKeahlian() {
	fmt.Println("\nMemanggil fungsi: editKeahlian()")
	// Implementasi serupa dengan editMinat()
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}

	// Tampilkan minat yang ada
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	// Pilih minat yang akan dihapus
	var pilihan int
	fmt.Print("Pilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	// Hapus minat
	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Printf("Minat '%s' berhasil dihapus\n", minatTerhapus)
}

func hapusKeahlian() {
	fmt.Println("\nMemanggil fungsi: hapusKeahlian()")
	// Implementasi serupa dengan hapusMinat()
}

func karirRekomendasi() {
	fmt.Println("\n=== Rekomendasi Karir ===")
	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Anda:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)
	fmt.Println("\nRekomendasi karir yang mungkin:")

	// Contoh rekomendasi sederhana
	for _, minat := range minatList {
		switch minat {
		case "Menulis":
			fmt.Println("- Penulis/Konten Creator")
		case "Menggambar":
			fmt.Println("- Desainer Grafis")
		case "Fotografi":
			fmt.Println("- Fotografer")
			// Tambahkan case lainnya
		}
	}
}
func tampilminat() {
	fmt.Println("\nMemanggil fungsi: tampilminat")
}
func tampilkeahlian() {
	fmt.Println("\nMemanggil fungsi: tampilkeahlian")
}

func urutRentangGaji() {
	fmt.Println("\nMemanggil fungsi: urutRentangGaji()")
}

func statistikKecocokanKarir() {
	fmt.Println("\nMemanggil fungsi: statistikKecocokanKarir()")
}
