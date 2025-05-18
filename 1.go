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
	if  pilihan < 1 || pilihan > 8 {
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

	fmt.Print("\nPilih nomor Keahlian (1-7): ")
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

	fmt.Println("\n=== Edit Keahlian ===")
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

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Printf("Minat '%s' berhasil dihapus\n", minatTerhapus)
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
type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}
func tampilkanRekomendasi(daftar []Karir) {
	fmt.Println("KARIR REKOMENDASI       | KECOCOKAN | INDUSTRI     | GAJI")
	fmt.Println("------------------------+-----------+--------------+------------")
	for _, k := range daftar {
		fmt.Printf("%-24s | %-9d | %-12s | %s\n", k.Nama, k.Kecocokan, k.Industri, k.Gaji)
	}
}

func karirRekomendasi() {
	if len(minatList) == 0 || len(keahlianList) == 0 {
		fmt.Println("Silakan tambahkan minat dan keahlian terlebih dahulu.")
		return
	}


	minat := minatList[len(minatList)-1]
	keahlian := keahlianList[len(keahlianList)-1]

	fmt.Println("\n=== Karir Rekomendasi Berdasarkan: ===")
	fmt.Printf("Minat   : %s\n", minat)
	fmt.Printf("Keahlian: %s\n", keahlian)

	var rekomendasi []Karir

	if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
		rekomendasi = []Karir{
			{"UI/UX Designer", 95, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	} else if minat == "Membaca" && keahlian == "Coding"{ 
		rekomendasi = []Karir{
			{"UI/UX Designer", 70 , "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}

	}else if minat == "Menganalisis" || keahlian == "Menghitung" || keahlian == "Analisis"{
		rekomendasi = []Karir{
			{"UI/UX Designer", 70, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 95, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Fotografi" && keahlian == "Editing Video"{
		rekomendasi = []Karir{
			{"UI/UX Designer", 75, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 97, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis")  {
		rekomendasi = []Karir{
			{"UI/UX Designer", 75, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 98, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Menganalisis") {
		rekomendasi = []Karir{
			{"UI/UX Designer", 60, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 99, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
		rekomendasi = []Karir{
			{"UI/UX Designer", 75, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 97, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Menulis" && (keahlian == "Menganalisis" || keahlian == "Coding"){
		rekomendasi = []Karir{
			{"UI/UX Designer", 65, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 97, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Melukis" && keahlian == "Design Grafis"{
		rekomendasi = []Karir{
			{"UI/UX Designer", 75, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 98, "Game", "6-10 juta"},
			{"Content Strategist", 65, "Kreatif", "Bervariasi"},
		}
	}else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar"){
		rekomendasi = []Karir{
			{"UI/UX Designer", 85, "Kreatif", "6-10 juta"},
			{"Software Engineer", 90, "Teknologi", "7-12 juta"},
			{"Data Analyst", 85, "Kreatif", "4-8 juta"},
			{"Video Editor", 80, "Kreatif", "4-7 juta"},
			{"Front-End Developer", 78, "Media", "5-8 juta"},
			{"Technical Writer", 75, "Kreatif", "8-14 juta"},
			{"Music Content Editor", 73, "Pemasaran", "6-10 juta"},
			{"SEO Specialist", 70, "Media", "4-6 juta"},
			{"Illustrator Digital", 68, "Game", "6-10 juta"},
			{"Content Strategist", 95, "Kreatif", "Bervariasi"},
		}

	}else {
		fmt.Println("Belum ada rekomendasi untuk kombinasi ini.")
		return
	}

	tampilkanRekomendasi(rekomendasi)
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


func cariKarirManual(nama string, karirList, industriList, gajiList []string) {
    ditemukan := false
    fmt.Println("\n=== Hasil Pencarian Karier ===")
    for i := 0; i < len(karirList); i++ {
        if nama == karirList[i] {
            fmt.Println("Karier:", karirList[i])
            fmt.Println("Industri:", industriList[i])
            fmt.Println("Gaji:", gajiList[i])
            ditemukan = true
            break
        }
    }
    if !ditemukan {
        fmt.Println("Karier tidak ditemukan.")
    }
}

func cariIndustriManual(nama string, karirList, industriList, gajiList []string) {
    ditemukan := false
    fmt.Println("\n=== Hasil Pencarian Industri ===")
    for i := 0; i < len(industriList); i++ {
        if nama == industriList[i] {
            fmt.Println("Karier:", karirList[i])
            fmt.Println("Industri:", industriList[i])
            fmt.Println("Gaji:", gajiList[i])
            ditemukan = true
        }
    }
    if !ditemukan {
        fmt.Println("Industri tidak ditemukan.")
    }
}

