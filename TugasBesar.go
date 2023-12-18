// Portofolio Tugas Besar Algoritma Pemrograman
package main

import "fmt"

const Size int = 100

type Mahasiswa struct {
	Nama          string
	Jawaban_Tugas string
	Jawaban_Quiz  string
	Forum_Diskusi string
	Materi        string
}

type Konten struct {
	Tugas string
	Quiz  string
	Forum string
}

type tabCourse [Size]Konten
type tabMahasiswa [Size]Mahasiswa
type tabNilai [Size]int

var jumlahSoal, jumlahJawaban, nilai int
var t tabCourse
var m tabMahasiswa
var n tabNilai

func main() {
	var pilihan string

	fmt.Println("----- SELAMAT DATANG -----")
	fmt.Println("SILAHKAN PILIH ROLE ANDA")
	fmt.Println(" ")
	fmt.Println("1. Dosen")
	fmt.Println("2. Mahasiswa")
	fmt.Println("0. Keluar")
	fmt.Print("Silahkan pilih opsi: ")

	fmt.Scanln(&pilihan)

	if pilihan == "1" {
		fmt.Println()
		menu_dosen()
	} else if pilihan == "2" {
		fmt.Println()
		menu_mahasiswa()
	} else if pilihan == "0" {
		fmt.Println("Terima Kasih telah menggunakan Aplikasi Moodle Sederhana, Have a nice day!")
		return
	} else {
		fmt.Println("Pilihan Tidak Valid")
		main()
	}
}

func menu_dosen() {
	/* IS: Terdefinisi pilihan untuk memilih opsi
	FS: Memilih opsi yang tersedia sesuai dengan index*/
	var pilihan string

	fmt.Println("Selamat Datang di e-Learning")
	fmt.Println(".............................")
	fmt.Println("1. Penambahan materi")
	fmt.Println("2. Pengeditan materi")
	fmt.Println("3. Penghapusan materi")
	fmt.Println("4. Tampilkan daftar materi")
	fmt.Println("5. Input nilai mahasiswa")
	fmt.Println("6. Data nilai mahasiswa")
	fmt.Println("0. Kembali")

	fmt.Scan(&pilihan)

	if pilihan == "1" {
		fmt.Println()
		tambahMateri(&t, &jumlahSoal)
	} else if pilihan == "2" {
		fmt.Println()
		editMateri(&t, &jumlahSoal)
	} else if pilihan == "3" {
		fmt.Println()
		hapusMateri(&t, &jumlahSoal)
	} else if pilihan == "4" {
		fmt.Println()
		tampilkanMateri(t, jumlahSoal)
	} else if pilihan == "5" {
		fmt.Println()
		nilaiMahasiswa(&n, m, jumlahJawaban, &nilai)
	} else if pilihan == "6" {
		fmt.Println()
		sortDescending(&n, &m, &nilai)
		tampilkanNilai(&n, &m, &nilai)
	} else if pilihan == "0" {
		main()
		return
	} else {
		fmt.Println("Pilihan Tidak Valid")
		menu_dosen()
	}
}

func menu_mahasiswa() {
	/* IS: Terdefinisi pilihan untuk memilih opsi
	FS: Memilih opsi yang tersedia sesuai dengan index*/

	var pilihan string

	fmt.Println("Selamat Datang di e-Learning")
	fmt.Println("----------------------------")
	fmt.Println("1. Tampilkan tugas, quiz, dan forum diskusi")
	fmt.Println("2. Jawab soal tugas, quiz, dan bertanya jawab pada forum diskusi ")
	fmt.Println("0. Kembali")

	fmt.Scan(&pilihan)

	if pilihan == "1" {
		fmt.Println()
		tampilkanTugas(t, jumlahSoal)
	} else if pilihan == "2" {
		fmt.Println()
		jawabanMateri(&m, &t, &jumlahSoal, &jumlahJawaban)
	} else if pilihan == "0" {
		main()
	} else {
		fmt.Println("Pilihan Tidak Valid")
		menu_mahasiswa()
	}
}

func tambahMateri(T *tabCourse, jumlahSoal *int) {
	/* IS: terdefinisi soal tugas, quiz, dan forum
	FS: Menambah tugas, quiz, dan forum sesuai index yang diinput ke array. Masukan berhenti ketika kalimat diakhiri dengan ' .'*/
	var tugas, quiz, forum string
	var kata1, kata2, kata3 string
	fmt.Println("Masukkan Soal Tugas: ")
	for kata1 != "." {
		fmt.Scan(&kata1)
		if kata1 != "." {
			tugas += kata1 + " "
			T[*jumlahSoal].Tugas = tugas
		}
	}

	fmt.Println("Masukkan Soal Quiz: ")
	for kata2 != "." {
		fmt.Scan(&kata2)
		if kata2 != "." {
			quiz += kata2 + " "
			T[*jumlahSoal].Quiz = quiz
		}
	}
	fmt.Println("Masukkan Diskusi Forum: ")
	for kata3 != "." {
		fmt.Scan(&kata3)
		if kata3 != "." {
			forum += kata3 + " "
			T[*jumlahSoal].Forum = forum
		}
	}

	*jumlahSoal++
	fmt.Println("Terima kasih, konten berhasil ditambahkan")

	fmt.Println()
	menu_dosen()
}

func editMateri(T *tabCourse, jumlahSoal *int) {
	/* IS: terdefinisi index untuk mencari soal tugas, quiz, dan forum
	FS: Mengedit tugas, quiz, dan forum sesuai index yang diinput ke array*/
	var idx int
	var kata1, kata2, kata3 string
	var tugas, quiz, forum string

	fmt.Print("Masukkan Indeks Konten: ")
	fmt.Scan(&idx)
	index := idx - 1
	if index < 0 || index >= *jumlahSoal {
		fmt.Println("Maaf indeks konten tidak valid.")
		menu_dosen()
	}

	fmt.Println("Masukkan Soal Tugas Baru: ")
	for kata1 != "." {
		fmt.Scan(&kata1)
		if kata1 != "." {
			tugas += kata1 + " "
			T[*jumlahSoal].Tugas = tugas
		}
	}

	fmt.Println("Masukkan Soal Quiz Baru: ")
	for kata2 != "." {
		fmt.Scan(&kata2)
		if kata2 != "." {
			quiz += kata2 + " "
			T[*jumlahSoal].Quiz = quiz
		}
	}

	fmt.Println("Masukkan Diskusi Forum Baru: ")
	for kata3 != "." {
		fmt.Scan(&kata3)
		if kata3 != "." {
			forum += kata3 + " "
			T[*jumlahSoal].Forum = forum
		}
	}

	T[index] = Konten{Tugas: tugas, Quiz: quiz, Forum: forum}
	fmt.Println("Terima kasih, konten berhasil diubah")

	fmt.Println()
	menu_dosen()
}

func hapusMateri(T *tabCourse, jumlahSoal *int) {
	/* IS: terdefinisi index untuk mencari soal tugas, quiz, dan forum
	FS: Menghapus tugas, quiz, dan forum sesuai index yang diinput ke array*/
	var idx int

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	indeks := idx - 1
	if indeks < 0 || indeks >= *jumlahSoal {
		fmt.Println("Maaf, indeks tidak valid.")
		menu_dosen()
	}

	for i := indeks; i < *jumlahSoal-1; i++ {
		(T)[i] = (T)[i+1]
	}

	*jumlahSoal--
	fmt.Println("Konten berhasil dihapus.")

	fmt.Println()
	menu_dosen()
}

func tampilkanMateri(T tabCourse, jumlahSoal int) {
	/* IS: -
	FS: Menampilkan tugas, quiz, dan forum sesuai index yang diinput ke array*/

	fmt.Println("Daftar Materi:")
	for i := 0; i < jumlahSoal; i++ {
		fmt.Println()
		fmt.Printf("Materi ke-%d\n", i+1)
		fmt.Println("Tugas:", T[i].Tugas)
		fmt.Println("Quiz:", T[i].Quiz)
		fmt.Println("Forum:", T[i].Forum)
		fmt.Println()
	}

	fmt.Println()
	menu_dosen()
}

func nilaiMahasiswa(N *tabNilai, M tabMahasiswa, jumlahJawaban int, nilai *int) {
	/* IS: terdefinisi nama mahasiswa dalam index array
	FS: Menginput nilai mahasiswa sesuai dengan jawaban mahasiswa*/
	var total_nilai int

	fmt.Println("Silahkan input nilai mahasiswa")
	fmt.Println("------------------------------")

	for i := 0; i < jumlahJawaban; i++ {
		fmt.Println()
		fmt.Printf("Nama: %v \n ", M[i].Nama)
		fmt.Printf("Tugas: %v \n ", M[i].Jawaban_Tugas)
		fmt.Printf("Quiz: %v \n", M[i].Jawaban_Quiz)
		fmt.Printf("Forum: %v \n", M[i].Forum_Diskusi)
		fmt.Println()

		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&total_nilai)

		N[*nilai] = total_nilai
		*nilai++
	}

	fmt.Println()
	menu_dosen()
}

func sortDescending(N *tabNilai, M *tabMahasiswa, nilai *int) {
	/*IS: -
	FS: mengurutkan nilai mahasiswa dari nilai tertinggi secara menurun */
	for i := 0; i < *nilai-1; i++ {
		maxIndex := i
		for j := i + 1; j < *nilai; j++ {
			if N[j] > N[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			// Swap nilai
			N[i], N[maxIndex] = N[maxIndex], N[i]
			// Swap nama
			M[i].Nama, M[maxIndex].Nama = M[maxIndex].Nama, M[i].Nama
			// Swap materi
			M[i].Materi, M[maxIndex].Materi = M[maxIndex].Materi, M[i].Materi
		}
	}
}

func tampilkanNilai(N *tabNilai, M *tabMahasiswa, nilai *int) {
	/*IS: -
	FS: menampilkan nilai mahasiswa*/
	for i := 0; i < *nilai; i++ {
		fmt.Printf("Nama Mahasiswa: %s\n", M[i].Nama)
		fmt.Printf("Materi ke- %s\n", M[i].Materi)
		fmt.Printf("Nilai: %d\n", N[i])
		fmt.Println()
	}

	menu_dosen()
}

func tampilkanTugas(T tabCourse, jumlahSoal int) {
	/* IS: -
	FS: Menampilkan tugas, quiz, dan forum sesuai index yang diinput ke array*/
	fmt.Println("Daftar tugas, quiz dan forum:")
	for i := 0; i < jumlahSoal; i++ {
		fmt.Println()
		fmt.Printf("Materi ke-%d\n", i+1)
		fmt.Println("Tugas:", T[i].Tugas)
		fmt.Println("Quiz:", T[i].Quiz)
		fmt.Println("Forum:", T[i].Forum)
		fmt.Println()
	}

	fmt.Println()
	menu_mahasiswa()
}

func jawabanMateri(M *tabMahasiswa, T *tabCourse, jumlahSoal *int, jumlahJawaban *int) {
	/* IS: terdefinisi jawaban tugas, quiz, dan forum
	FS: Menjawab tugas, quiz, dan forum yang sudah diinput oleh dosen. Masukan berhenti ketika kalimat diakhiri dengan ' .'*/
	var jawaban_tugas, jawaban_quiz, forum_diskusi string
	var kata1, kata2, kata3 string
	var nama, materi string

	for i := 0; i < *jumlahSoal; i++ {
		kata1 = " "
		kata2 = " "
		kata3 = " "

		fmt.Print("Materi ke-")
		fmt.Scan(&materi)
		M[*jumlahJawaban].Materi = materi

		fmt.Print("Masukkan nama: ")
		fmt.Scan(&nama)
		M[*jumlahJawaban].Nama = nama

		fmt.Println("Tugas:", (T)[i].Tugas)
		fmt.Print("Masukkan jawaban tugas: ")

		for kata1 != "." {
			fmt.Scan(&kata1)
			if kata1 != "." {
				jawaban_tugas += kata1 + " "
			}
		}
		M[*jumlahJawaban].Jawaban_Tugas = jawaban_tugas

		fmt.Println("Quiz:", (T)[i].Quiz)
		fmt.Print("Masukkan jawaban quiz: ")

		for kata2 != "." {
			fmt.Scan(&kata2)
			if kata2 != "." {
				jawaban_quiz += kata2 + " "
			}
		}
		M[*jumlahJawaban].Jawaban_Quiz = jawaban_quiz

		fmt.Println("Forum Diskusi:", (T)[i].Forum)
		fmt.Print("Masukkan forum diskusi: ")

		for kata3 != "." {
			fmt.Scan(&kata3)
			if kata3 != "." {
				forum_diskusi += kata3 + " "
			}
		}
		M[*jumlahJawaban].Forum_Diskusi = forum_diskusi

		*jumlahJawaban++
	}

	fmt.Println("Jawaban tugas, quiz, dan forum diskusi berhasil disimpan")
	fmt.Println()

	menu_mahasiswa()
}
