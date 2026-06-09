package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 100

//STRUKTUR DATA

type Waktu struct {
	Hari, Bulan, Tahun int
}

type Assessment struct {
	ID          int
	UserID      int
	Nama        string
	Jawaban     [5]int //Skala Likert 1-5 untuk 5 pertanyaan
	TotalSkor   int
	Tanggal     Waktu
	Rekomendasi string
}

type DaftarAssessment [NMAX]Assessment

//FUNGSI UTAMA & MENU

func main() {
	var data DaftarAssessment
	var nData int = 0
	var pilihan int
	var status string

	inisialisasiData(&data, &nData)

	for {
		cls()
		header()
		fmt.Printf("%-35s[1] 📋 Kelola Self-Assessment\n", "")
		fmt.Printf("%-35s[2] 📊 Laporan & Statistik\n", "")
		fmt.Printf("%-35s[3] 🚪 Keluar\n", "")
		fmt.Printf("\n%-35s%s\n", "", status)
		fmt.Printf("%-35sPilih: ", "")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuKelola(&data, &nData, &status)
		} else if pilihan == 2 {
			menuLaporan(data, nData)
		} else if pilihan == 3 {
			break
		}
		status = ""
	}
}

func menuKelola(A *DaftarAssessment, n *int, status *string) {
	var pil int
	for {
		cls()
		header()
		tampilkanTabel(*A, *n)
		fmt.Printf("%-35s[1] ➕ Tambah Assessment\n", "")
		fmt.Printf("%-35s[2] 📝 Ubah Data\n", "")
		fmt.Printf("%-35s[3] 🗑️ Hapus Data\n", "")
		fmt.Printf("%-35s[4] 🔍 Cari Berdasarkan User ID\n", "")
		fmt.Printf("%-35s[5] 📶 Urutkan Data\n", "")
		fmt.Printf("%-35s[6] ⬅️ Kembali\n", "")
		fmt.Printf("\n%-35sPilih: ", "")
		fmt.Scan(&pil)

		if pil == 1 {
			tambahAssessment(A, n)
		} else if pil == 2 {
			ubahAssessment(A, *n)
		} else if pil == 3 {
			hapusAssessment(A, n)
		} else if pil == 4 {
			cariAssessment(*A, *n)
		} else if pil == 5 {
			urutkanAssessment(A, *n)
		} else if pil == 6 {
			break
		}
	}
}

//FUNGSI-FUNGSI BARU UNTUK MENGATASI ERROR

func header() {
	fmt.Printf("\n%-35s==================================================\n", "")
	fmt.Printf("%-35s              MENTAL HEALTH MANAGEMENT           \n", "")
	fmt.Printf("%-35s==================================================\n", "")
}

func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func inisialisasiData(A *DaftarAssessment, n *int) {
	// Menambahkan 2 data tiruan awal agar tabel tidak kosong
	A[0] = Assessment{ID: 1, UserID: 101, Nama: "Andi", TotalSkor: 12, Rekomendasi: "Stres Ringan", Tanggal: Waktu{1, 5, 2026}}
	A[1] = Assessment{ID: 2, UserID: 102, Nama: "Budi", TotalSkor: 8, Rekomendasi: "Kondisi Stabil", Tanggal: Waktu{5, 5, 2026}}
	*n = 2
}

func tampilkanTabel(A DaftarAssessment, n int) {
	fmt.Printf("%-35s┌────┬─────────┬──────────────┬──────┬─────────────────┐\n", "")
	fmt.Printf("%-35s│ ID │ User ID │     Nama     │ Skor │     Tanggal     │\n", "")
	fmt.Printf("%-35s├────┼─────────┼──────────────┼──────┼─────────────────┤\n", "")
	for i := 0; i < n; i++ {
		fmt.Printf("%-35s│ %-2d │ %-7d │ %-12s │ %-4d │ %02d-%02d-%-9d │\n", "", A[i].ID, A[i].UserID, A[i].Nama, A[i].TotalSkor, A[i].Tanggal.Hari, A[i].Tanggal.Bulan, A[i].Tanggal.Tahun)
	}
	fmt.Printf("%-35s└────┴─────────┴──────────────┴──────┴─────────────────┘\n", "")
}

func tambahAssessment(A *DaftarAssessment, n *int) {
	if *n >= NMAX {
		fmt.Printf("%-35sPenyimpanan Penuh!\n", "")
		return
	}
	var d Assessment
	fmt.Printf("\n%-35sUser ID: ", "")
	fmt.Scan(&d.UserID)
	fmt.Printf("%-35sNama Pengguna: ", "")
	fmt.Scan(&d.Nama)
	fmt.Printf("%-35sTanggal (DD MM YYYY): ", "")
	fmt.Scan(&d.Tanggal.Hari, &d.Tanggal.Bulan, &d.Tanggal.Tahun)

	daftarPertanyaan := [5]string{
		"Merasa cemas atau panik tiba-tiba",
		"Mengalami gangguan tidur (insomnia)",
		"Kehilangan minat pada aktivitas hobi",
		"Merasa lelah sepanjang hari",
		"Sulit fokus/berkonsentrasi saat belajar",
	}

	fmt.Printf("%-35sInput 5 Jawaban (Skala 1-5):\n", "")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-37sPertanyaan %d. %-40s: ", "", i+1, daftarPertanyaan[i])
		fmt.Scan(&d.Jawaban[i])
		d.TotalSkor += d.Jawaban[i]
	}
	d.ID = *n + 1
	if d.TotalSkor <= 10 {
		d.Rekomendasi = "Kondisi Stabil"
	} else if d.TotalSkor <= 18 {
		d.Rekomendasi = "Stres Ringan"
	} else {
		d.Rekomendasi = "Stres Berat"
	}
	A[*n] = d
	*n++
}

func ubahAssessment(A *DaftarAssessment, n int) {
	var id int
	fmt.Printf("%-35sMasukkan ID Assessment yang diubah: ", "")
	fmt.Scan(&id)
	for i := 0; i < n; i++ {
		if A[i].ID == id {
			fmt.Printf("%-35sNama Baru: ", "")
			fmt.Scan(&A[i].Nama)
			fmt.Printf("%-35sTotal Skor Baru: ", "")
			fmt.Scan(&A[i].TotalSkor)
			return
		}
	}
}

func hapusAssessment(A *DaftarAssessment, n *int) {
	var id int
	fmt.Printf("%-35sMasukkan ID Assessment yang akan dihapus: ", "")
	fmt.Scan(&id)
	idx := -1
	for i := 0; i < *n; i++ {
		if A[i].ID == id {
			idx = i
			break
		}
	}
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Printf("%-35sData berhasil dihapus!\n", "")
	}
}

func cariAssessment(A DaftarAssessment, n int) {
	var target int
	fmt.Printf("%-35sMasukkan User ID yang dicari: ", "")
	fmt.Scan(&target)
	for i := 0; i < n; i++ {
		if A[i].UserID == target {
			fmt.Printf("%-35s-> ID: %d | Nama: %s | Skor: %d | %s\n", "", A[i].ID, A[i].Nama, A[i].TotalSkor, A[i].Rekomendasi)
		}
	}
	fmt.Printf("\n%-35sTekan Enter...", "")
	fmt.Scanln()
	fmt.Scanln()
}

func urutkanAssessment(A *DaftarAssessment, n int) {
	//Urutkan dengan Selection Sort berdasarkan total skor terendah
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if A[j].TotalSkor < A[minIdx].TotalSkor {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
	fmt.Printf("%-35sData berhasil diurutkan berdasarkan Skor!\n", "")
	fmt.Printf("\n%-35sTekan Enter...", "")
	fmt.Scanln()
}

func menuLaporan(A DaftarAssessment, n int) {
	cls()
	header()

	if n == 0 {
		fmt.Printf("%-35s⚠️ Belum ada data assessment yang tersimpan.\n", "")
		fmt.Printf("\n%-35sTekan Enter untuk kembali...", "")
		fmt.Scanln() // baris 1 (pembilas buffer)
		fmt.Scanln() // baris 2 (penahan layar)
		return
	}

	var total int
	for i := 0; i < n; i++ {
		total += A[i].TotalSkor
	}
	rata := float64(total) / float64(n)

	fmt.Printf("%-35s============================================\n", "")
	fmt.Printf("%-35s       LAPORAN & STATISTIK KESEHATAN MENTAL  \n", "")
	fmt.Printf("%-35s============================================\n", "")
	fmt.Printf("%-35sTotal Sesi Assessment : %d data\n", "", n)
	fmt.Printf("%-35sRata-rata Skor Global : %.2f\n", "", rata)
	fmt.Printf("%-35s============================================\n", "")

	// Menahan terminal agar tidak langsung diclear oleh fungsi main
	fmt.Printf("\n%-35sTekan [Enter] untuk kembali ke Menu Utama...", "")
	fmt.Scanln()
	fmt.Scanln()
}
