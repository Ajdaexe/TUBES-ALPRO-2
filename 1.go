package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat ")
		fmt.Println("7. Tampilkan Keahlian ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
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
			karirRekomendasi()
		case 6:
			tampilminat("Daftar minat", minatList)
		case 7:
			tampilkeahlian("Daftar kealhlian", keahlianList)
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-10")
		}
	}
}

func tambahMinat() {
	var pilihan int
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
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Print("\nPilih nomor minat (1-8): ")
	fmt.Scan(&pilihan)
	if pilihan > 8 {
		fmt.Println("Input tidak valid")
		return
	}
	minatTerpilih := daftarMinat[pilihan-1]
	minatList = append(minatList, minatTerpilih)
	fmt.Printf("Minat '%s' berhasil ditambahkan\n", minatTerpilih)
}

func tambahKeahlian() {
	var pilihan int
	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Menganalisis",
		"Melukis",
	}
	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	fmt.Print("\nPilih nomor minat (1-7): ")
	fmt.Scan(&pilihan)
	if pilihan > 7 {
		fmt.Println("Input tidak valid")
		return
	}
	keahlianTerpilih := daftarKeahlian[pilihan-1]
	keahlianList = append(keahlianList, keahlianTerpilih)
	fmt.Printf("keahlian '%s' berhasil ditambahkan\n", keahlianTerpilih)
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. hapus Minat")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	}
}

func editKeahlian() {
	fmt.Println("\nMemanggil fungsi: editKeahlian()")
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. hapus keahlian")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", minatTerhapus)
}

func hapusKeahlian() {
	fmt.Println("\nMemanggil fungsi: hapusKeahlian()")
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
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
	fmt.Println("\nRekomendasi karir yang cocok:")

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				fmt.Println("- UI/UX Designer")
			} else if minat == "Membaca" && keahlian == "Coding" {
				fmt.Println("- Software Engineer")
			} else if minat == "Menganalisis" || keahlian == "Menghitung" || keahlian == "Analisis" {
				fmt.Println("- Data Analyst")
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				fmt.Println("- Video Editor")
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				fmt.Println("- Front-End Developer")
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Menganalisis") {
				fmt.Println("- Technical Writer")
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				fmt.Println("- Music Content Editor")
			} else if minat == "Menulis" && (keahlian == "Menganalisis" || keahlian == "Coding") {
				fmt.Println("- SEO Specialist")
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				fmt.Println("- Ilustrator Digital")
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				fmt.Println("- Content Strategist")
			}
		}
	}
	rekomendasi := []struct {
		No        int
		Kombinasi string
		Karir     string
		Kecocokan string
		Industri  string
		Gaji      string
	}
	var hasil []struct {
        Kombinasi    string
        Karir       string
        Kecocokan   string
        Industri    string
        Gaji        string
    }
	{	{1, "Menggambar + Melukis + Design Grafis", "UI/UX Designer", "91%", "Desain & Kreatif", "7.000.000 - 12.000.000"},
		{2, "Membaca + Coding", "Software Engineer", "88%", "Teknologi", "9.000.000 - 20.000.000"},
		{3, "Menganalisis + Menghitung + Analisis", "Data Analyst", "93%", "Teknologi, Keuangan", "8.000.000 - 15.000.000"},
		{4, "Fotografi + Editing Video", "Video Editor", "86%", "Media & Konten", "5.000.000 - 9.000.000"},
		{5, "Menggambar + Coding + Design Grafis", "Front-End Developer", "89%", "Teknologi, Web Design", "7.000.000 - 14.000.000"},
		{6, "Menulis + Menghafal + Menganalisis", "Technical Writer", "84%", "Dokumentasi, Teknologi, Pendidikan", "6.000.000 - 10.000.000"},
		{7, "Bermain alat musik + Menghafal", "Music Content Editor", "79%", "Musik Digital, Media", "4.000.000 - 9.000.000"},
		{8, "Menulis + Menganalisis + Coding", "SEO Specialist", "82%", "Digital Marketing", "6.000.000 - 10.000.000"},
		{9, "Melukis + Design Grafis", "Ilustrator Digital", "87%", "Desain, Game Art", "5.000.000 - 10.000.000"},
		{10, "Membaca + Menulis + Menggambar", "Content Strategist", "85%", "Media & Komunikasi", "6.500.000 - 10.000.000"},
	}
	fmt.Println("\n+----------------------------------------+-----------------------+-------------+------------------------------+---------------------------+")
    fmt.Println("| Kombinasi Minat + Keahlian            | Karir Potensial       | Kecocokan   | Industri yang Cocok          | Gaji (Rp)                 |")
    fmt.Println("+----------------------------------------+-----------------------+-------------+------------------------------+---------------------------+")

    if len(hasil) == 0 {
        fmt.Println("| Tidak ditemukan rekomendasi karir yang cocok                                                    |")
    } else {
        for _, h := range hasil {
            fmt.Printf("| %-38s | %-21s | %-11s | %-28s | %-25s |\n", 
                h.Kombinasi, h.Karir, h.Kecocokan, h.Industri, h.Gaji)
        }
    }

    fmt.Println("+----------------------------------------+-----------------------+-------------+------------------------------+---------------------------+")
}
func tampilminat(minat string, minatList []string) {
	var i int
	fmt.Println("\n=== Daftar Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada data minat")
		return
	}

	for i, minat = range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}
}
func tampilkeahlian(judul string, keahlianList []string) {
	fmt.Printf("\n=== %s ===\n", judul)
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada data keahlian")
		return
	}

	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
}
