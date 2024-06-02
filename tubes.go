package main

import "fmt"

const NMAX int = 1000

type DataBuku struct {
	id          int
	judul       string
	penulis     string
	genre       string
	tahunterbit int
	countF      int
	countP      float64
}

type Peminjam struct {
	idpeminjam   int
	idBdipinjam  int
	namapeminjam string
	bukudipinjam string
	haripinjam   int
	bulanpinjam  int
	tahunpinjam  int
}

type Buku [NMAX]DataBuku
type DataPinjam [NMAX]Peminjam

var Book Buku
var Pinjam DataPinjam
var nBuku int = 0
var nPinjam int = 0
var nPeminjaman int = 0

func main() {
	var PilihanMenu int = 0
	MenuUtama(&PilihanMenu)
}

func MenuUtama(PilihanMenu *int) {
	for *PilihanMenu != 4 {
		fmt.Println()
		fmt.Println("   ===== Selamat Datang =====")
		fmt.Println(" Aplikasi Pendataan Perpustakaan")
		fmt.Println("---------------------------------")
		fmt.Println("1. Edit Pendataan Buku")
		fmt.Println("2. Data Peminjaman Buku")
		fmt.Println("3. Daftar Buku")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------------------")
		fmt.Println("Pilihan Anda (1/2/3/4)?")
		fmt.Scan(PilihanMenu)
		if *PilihanMenu == 1 {
			MenuEditBuku(PilihanMenu)
		} else if *PilihanMenu == 2 {
			MenuPeminjamanBuku(PilihanMenu)
		} else if *PilihanMenu == 3 {
			MenuDaftarBuku(Book, nBuku)
		}

	}
}

func MenuEditBuku(PilihanMenu *int) {
	var PilihanMenuEdit int = 0
	for PilihanMenuEdit != 4 {
		fmt.Println()
		fmt.Println("---------------------------------")
		fmt.Println("   Apa yang ingin anda lakukan?")
		fmt.Println("---------------------------------")
		fmt.Println("1. Penambahan Data Buku")
		fmt.Println("2. Perubahan Data Buku")
		fmt.Println("3. Penghapusan Data Buku")
		fmt.Println("4. Kembali")
		fmt.Println("---------------------------------")
		fmt.Println("Pilihan Anda (1/2/3/4)?")
		fmt.Scan(&PilihanMenuEdit)
		if PilihanMenuEdit == 4 {
			*PilihanMenu = 0
		} else if PilihanMenuEdit == 1 {
			TambahDataBuku(&Book, &nBuku)
		} else if PilihanMenuEdit == 2 {
			UbahDataBuku(&Book, &nBuku)
		} else if PilihanMenuEdit == 3 {
			HapusDataBuku(&Book, &nBuku)
		}
	}
}

func TambahDataBuku(Book *Buku, nBuku *int) {
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("               Silahkan isi data buku yang ingin ditambahkan")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("PETUNJUK:")
	fmt.Println("Isilah dengan format: ID Buku, Judul, Penulis, Genre, dan Tahun Terbit.")
	fmt.Println("Untuk spasi saat menulis judul atau penulis, gunakan tanda '_'.")
	fmt.Println("Jika ingin berhenti, Silahkan isi 0 di bagian ID Buku.")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Data buku yang ingin ditambah:")
	var title, author, Genre string
	var ID, year int
	var duplikat bool
	fmt.Scan(&ID)
	for ID != 0 && *nBuku < NMAX {
		fmt.Scanln(&title, &author, &Genre, &year)
		duplikat = cekduplikatidbuku(Book, nBuku, ID)
		if duplikat == false {
			Book[*nBuku].id = ID
			Book[*nBuku].judul = title
			Book[*nBuku].penulis = author
			Book[*nBuku].genre = Genre
			Book[*nBuku].tahunterbit = year
			*nBuku += 1
			fmt.Scan(&ID)
		} else if duplikat == true {
			fmt.Println()
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Println("             ID Buku Sudah Terambil, Silahkan Gunakan ID Lain")
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Isilah dengan format: ID Buku, Judul, Penulis, Genre, dan Tahun Terbit.")
			fmt.Println("Untuk spasi saat menulis judul atau penulis, gunakan tanda '_'.")
			fmt.Println("Jika ingin berhenti, Silahkan isi 0 di bagian ID Buku.")
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Println("Data buku yang ingin ditambah:")
			fmt.Scan(&ID)
			duplikat = cekduplikatidbuku(Book, nBuku, ID)
		}
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("             Anda telah berhenti")
	fmt.Println("-----------------------------------------------")
}

func UbahDataBuku(Book *Buku, nBuku *int) {
	var ada, idBaru, tahunterbitBaru int
	var PilihanEditBuku int
	var x, lanjut int
	lanjut = 1
	x = -1
	ada = -1
	var judulBaru, penulisBaru, genreBaru string
	for x != 0 && ada == -1 {
		PilihanEditBuku = 0
		fmt.Println()
		fmt.Println("------------------------------------------------")
		fmt.Println("  Silahkan isi ID dari buku yang ingin diedit")
		fmt.Println("------------------------------------------------")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------")
		fmt.Println("Masukkan ID Buku:")
		fmt.Scan(&x)
		ada = seqsearchid(Book, nBuku, x)
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("            Buku Tidak Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("0. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Harap masukkan lagi ID buku yang valid:")
			fmt.Scan(&x)
			ada = seqsearchid(Book, nBuku, x)
		}
		for ada != -1 && PilihanEditBuku != 6 && lanjut == 1 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                Buku Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Data buku tersebut:")
			fmt.Println(Book[ada].id, Book[ada].judul, Book[ada].penulis, Book[ada].genre, Book[ada].tahunterbit)
			fmt.Println("-----------------------------------------------")
			fmt.Println("  Silahkan pilih data apa yang ingin diedit")
			fmt.Println("-----------------------------------------------")
			fmt.Println("1. ID Buku")
			fmt.Println("2. Judul Buku")
			fmt.Println("3. Penulis Buku")
			fmt.Println("4. Genre Buku")
			fmt.Println("5. Tahun Terbit Buku")
			fmt.Println("6. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Pilihan Anda (1/2/3/4/5/6)?")
			fmt.Scan(&PilihanEditBuku)
			if PilihanEditBuku == 1 {
				fmt.Println("Masukkan ID Baru Buku:")
				fmt.Scan(&idBaru)
				Book[ada].id = idBaru
			} else if PilihanEditBuku == 2 {
				fmt.Println("Masukkan Judul Baru Buku:")
				fmt.Scan(&judulBaru)
				Book[ada].judul = judulBaru
			} else if PilihanEditBuku == 3 {
				fmt.Println("Masukkan Nama Penulis Baru Buku:")
				fmt.Scan(&penulisBaru)
				Book[ada].penulis = penulisBaru
			} else if PilihanEditBuku == 4 {
				fmt.Println("Masukkan Genre Baru Buku:")
				fmt.Scan(&genreBaru)
				Book[ada].genre = genreBaru
			} else if PilihanEditBuku == 5 {
				fmt.Println("Masukkan Tahun Terbit Baru Buku:")
				fmt.Scan(&tahunterbitBaru)
				Book[ada].tahunterbit = tahunterbitBaru
			} else if PilihanEditBuku == 6 {
				ada = -1
				x = -1
			}
			if PilihanEditBuku < 6 {
				fmt.Println()
				fmt.Println("-----------------------------------------------")
				fmt.Println("                Edit selesai")
				fmt.Println("-----------------------------------------------")
				fmt.Println("Data baru buku tersebut:")
				fmt.Println(Book[ada].id, Book[ada].judul, Book[ada].penulis, Book[ada].genre, Book[ada].tahunterbit)
				fmt.Println("-----------------------------------------------")
				fmt.Println("     Apakah anda ingin lanjut mengedit?")
				fmt.Println("-----------------------------------------------")
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")
				fmt.Println("-----------------------------------------------")
				fmt.Println("Pilihan Anda (1/2)?")
				fmt.Scan(&lanjut)
				fmt.Println()
				if lanjut == 1 {
					ada = -1
					x = -1
				}
			}
		}
	}
}

func HapusDataBuku(Book *Buku, nBuku *int) {
	var x, lanjut, ada, i, KonfirmasiHapus int
	x = -1
	ada = -1

	for x != 0 && ada == -1 {
		lanjut = 1
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println(" Silahkan isi ID dari buku yang ingin dihapus")
		fmt.Println("-----------------------------------------------")
		fmt.Println("0. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Masukkan ID Buku:")
		fmt.Scan(&x)
		ada = seqsearchid(Book, nBuku, x)
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("             Buku Tidak Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("0. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Harap masukkan lagi ID buku yang valid:")
			fmt.Scan(&x)
			ada = seqsearchid(Book, nBuku, x)
		}
		for ada != -1 && lanjut == 1 {
			fmt.Println()
			fmt.Println("---------------------------------------------------------")
			fmt.Println("                     Buku ditemukan")
			fmt.Println("---------------------------------------------------------")
			fmt.Println("Data buku tersebut:")
			fmt.Println(Book[ada].id, Book[ada].judul, Book[ada].penulis, Book[ada].genre, Book[ada].tahunterbit)
			fmt.Println("---------------------------------------------------------")
			fmt.Println("Apakah anda yakin ingin menghapus data buku tersebut?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Println("---------------------------------------------------------")
			fmt.Println("Pilihan Anda (1/2)?")
			fmt.Scan(&KonfirmasiHapus)
			if KonfirmasiHapus == 1 {
				for i = ada; i < *nBuku-1; i++ {
					Book[i] = Book[i+1]
				}
				*nBuku = *nBuku - 1
				fmt.Println()
				fmt.Println("-----------------------------------------------")
				fmt.Println("            Buku Berhasil Dihapus")
				fmt.Println("-----------------------------------------------")
				fmt.Println(" Apakah anda ingin lanjut menghapus buku lain?")
				fmt.Println("-----------------------------------------------")
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")
				fmt.Println("-----------------------------------------------")
				fmt.Println("Pilihan Anda (1/2)?")
				fmt.Scan(&lanjut)
				if lanjut == 1 {
					lanjut = -1
					ada = -1
					x = -1
				}
			} else if KonfirmasiHapus == 2 {
				ada = -1
				x = -1
			}

		}
	}
}

func MenuPeminjamanBuku(PilihanMenu *int) {
	var PilihanMenuPeminjaman int = 0
	for PilihanMenuPeminjaman != 7 {
		fmt.Println()
		fmt.Println("-------------------------------------------------------")
		fmt.Println("              Apa yang ingin anda lakukan?")
		fmt.Println("-------------------------------------------------------")
		fmt.Println("1. Penambahan Data Peminjaman Buku")
		fmt.Println("2. Perubahan Data Peminjaman Buku")
		fmt.Println("3. Penghapusan Data Peminjaman Buku")
		fmt.Println("4. Hitung Tarif Peminjaman dan Denda Keterlambatan")
		fmt.Println("5. Buku-buku Yang Sedang Dipinjam")
		fmt.Println("6. Buku-buku Terfavorit")
		fmt.Println("7. Kembali")
		fmt.Println("-------------------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3/4/5/6/7)?")
		fmt.Scan(&PilihanMenuPeminjaman)
		if PilihanMenuPeminjaman == 7 {
			*PilihanMenu = 0
		} else if PilihanMenuPeminjaman == 1 {
			TambahDataPinjamBuku(&Pinjam, &nPinjam)
		} else if PilihanMenuPeminjaman == 2 {
			UbahDataPinjamBuku(&Pinjam, &nPinjam)
		} else if PilihanMenuPeminjaman == 3 {
			HapusDataPinjamBuku(&Pinjam, &nPinjam)
		} else if PilihanMenuPeminjaman == 4 {
			TarifDanDenda(&Pinjam, &nPinjam)
		} else if PilihanMenuPeminjaman == 5 {
			StatusPinjam(Pinjam, nPinjam)
		} else if PilihanMenuPeminjaman == 6 {
			BukuFavorit(&Book, nBuku)
		}
	}
}

func TambahDataPinjamBuku(Pinjam *DataPinjam, nPinjam *int) {
	var x int = -1
	var kembali, harikembali, bulankembali, tahunkembali int
	var ada, lanjut int
	var duplikat, duplikatpeminjaman bool
	for x != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("Silahkan masukkan ID buku yang ingin ditambahkan ke dalam data peminjaman")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("PETUNJUK:")
		fmt.Println("Jika ingin kembali, Silahkan isi 0 di bagian ID Buku.")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("ID buku yang ingin dipinjam:")
		fmt.Scan(&x)
		ada = seqsearchid(&Book, &nBuku, x)
		duplikatpeminjaman = cekduplikatidbukudipinjam(*Pinjam, *nPinjam, x)
		for duplikatpeminjaman == true {
			fmt.Println()
			fmt.Println("-------------------------------------------------------")
			fmt.Println("                 Buku sudah terpinjam")
			fmt.Println("-------------------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Mohon masukkan ID buku yang lain.")
			fmt.Println("Jika ingin kembali, Silahkan isi 0 di bagian ID Buku.")
			fmt.Println("-------------------------------------------------------")
			fmt.Println("ID buku yang ingin dipinjam:")
			fmt.Scan(&x)
			duplikatpeminjaman = cekduplikatidbukudipinjam(*Pinjam, *nPinjam, x)
			ada = seqsearchid(&Book, &nBuku, x)
			if x == 0 {
				duplikatpeminjaman = false
				x = 0
			}
		}
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-------------------------------------------------------")
			fmt.Println("                 Buku tidak ditemukan")
			fmt.Println("-------------------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Mohon masukkan ID buku yang valid.")
			fmt.Println("Jika ingin kembali, Silahkan isi 0 di bagian ID Buku.")
			fmt.Println("-------------------------------------------------------")
			fmt.Println("ID buku yang ingin dipinjam:")
			fmt.Scan(&x)
			ada = seqsearchid(&Book, &nBuku, x)
		}
		for ada != -1 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("               Buku ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Berikut adalah data buku tersebut:")
			fmt.Println(Book[ada].id, Book[ada].judul, Book[ada].penulis, Book[ada].genre, Book[ada].tahunterbit)
			fmt.Println("-----------------------------------------------")
			fmt.Println("    Masukkan ID peminjam dan nama peminjam")
			fmt.Println("-----------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Gunakan '_' sebagai spasi saat memasukkan nama.")
			fmt.Println("Untuk kembali, masukkan 0 ke bagian ID buku.")
			fmt.Println("-----------------------------------------------")
			fmt.Println("ID dan nama peminjam buku:")
			var namapeminjam string
			var ID int
			fmt.Scan(&ID)
			duplikat = cekduplikatidpeminjam(*Pinjam, *nPinjam, ID)
			if ID == 0 {
				ada = -1
			}
			for duplikat == true {
				fmt.Scanln(&namapeminjam)
				fmt.Println()
				fmt.Println("-----------------------------------------------------")
				fmt.Println("ID Peminjam Sudah Terambil, Silahkan Gunakan ID Lain")
				fmt.Println("-----------------------------------------------------")
				fmt.Println("      Masukkan ID peminjam dan nama peminjam")
				fmt.Println("-----------------------------------------------------")
				fmt.Println("PETUNJUK:")
				fmt.Println("Gunakan '_' sebagai spasi saat memasukkan nama.")
				fmt.Println("Untuk kembali, masukkan 0 ke bagian ID buku.")
				fmt.Println("-----------------------------------------------------")
				fmt.Println("ID dan nama peminjam buku:")
				fmt.Scan(&ID)
				duplikat = cekduplikatidpeminjam(*Pinjam, *nPinjam, ID)
				if ID == 0 {
					duplikat = false
				}
			}
			if ID != 0 && *nPinjam < NMAX {
				fmt.Scanln(&namapeminjam)
				Pinjam[*nPinjam].idpeminjam = ID
				Pinjam[*nPinjam].namapeminjam = namapeminjam
				ada = seqsearchid(&Book, &nBuku, x)
				fmt.Println()
				fmt.Println("---------------------------------------------------------")
				fmt.Println("               Masukkan tanggal peminjaman")
				fmt.Println("---------------------------------------------------------")
				fmt.Println("PETUNJUK:")
				fmt.Println("Gunakan format hari, bulan, dan tahun.")
				fmt.Println("Gunakan spasi untuk memisahkan hari, bulan, dan tahun.")
				fmt.Println("Untuk kembali, masukkan 0 ke bagian hari.")
				fmt.Println("---------------------------------------------------------")
				fmt.Println("Hari, bulan, dan tahun peminjaman buku:")
				var hari, bulan, tahun int
				fmt.Scan(&hari)
				if hari != 0 && *nPinjam < NMAX {
					fmt.Scanln(&bulan, &tahun)
					Pinjam[*nPinjam].haripinjam = hari
					Pinjam[*nPinjam].bulanpinjam = bulan
					Pinjam[*nPinjam].tahunpinjam = tahun
					Pinjam[*nPinjam].bukudipinjam = Book[ada].judul
					Pinjam[*nPinjam].idBdipinjam = Book[ada].id
					Book[ada].countF = Book[ada].countF + 1
					*nPinjam += 1
					nPeminjaman += 1
					kembali = hitungharikembali(Pinjam, ada)
					tahunkembali = kembali / 360
					kembali = kembali % 360
					bulankembali = kembali / 30
					kembali = kembali % 30
					harikembali = kembali
					fmt.Println()
					fmt.Println("------------------------------------------------------------------")
					fmt.Println("                Data peminjaman berhasil disimpan")
					fmt.Printf("          Peminjam harus mengembalikan buku pada %d/%d/%d\n", harikembali, bulankembali, tahunkembali)
					fmt.Println("------------------------------------------------------------------")
					fmt.Println("Apakah anda ingin lanjut melakukan penambahan data peminjaman?")
					fmt.Println("1. Ya")
					fmt.Println("2. Tidak")
					fmt.Println("------------------------------------------------------------------")
					fmt.Println("Pilihan Anda (1/2)?")
					fmt.Scan(&lanjut)
					if lanjut == 1 {
						x = -1
						ada = -1
					} else if lanjut == 2 {
						x = 0
						ada = -1
					}
				}
			}

		}

	}
}

func UbahDataPinjamBuku(Pinjam *DataPinjam, nPinjam *int) {
	var ada, idBaru, haribaru, bulanbaru, tahunbaru int
	var judulbaru, namabaru string
	var PilihanEditBuku int
	var x, lanjut int
	lanjut = 1
	x = -1
	ada = -1
	for x != 0 && ada == -1 {
		PilihanEditBuku = 0
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("Silahkan isi ID dari peminjam yang ingin diedit")
		fmt.Println("-----------------------------------------------")
		fmt.Println("0. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Masukkan ID Peminjam:")
		fmt.Scan(&x)
		ada = seqsearchpinjam(Pinjam, nPinjam, x)
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("              ID Tidak Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("0. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Harap masukkan lagi ID peminjam yang valid:")
			fmt.Scan(&x)
			ada = seqsearchpinjam(Pinjam, nPinjam, x)
		}
		for ada != -1 && PilihanEditBuku != 6 && lanjut == 1 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("             Data Peminjam Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Data peminjam tersebut:")
			fmt.Println(Pinjam[ada].idpeminjam, Pinjam[ada].namapeminjam)
			fmt.Println("Buku yang dipinjam:")
			fmt.Println(Pinjam[ada].bukudipinjam)
			fmt.Printf("Tanggal Peminjaman: %d/%d/%d\n", Pinjam[ada].haripinjam, Pinjam[ada].bulanpinjam, Pinjam[ada].tahunpinjam)
			fmt.Println("-----------------------------------------------")
			fmt.Println("  Silahkan pilih data apa yang ingin diedit")
			fmt.Println("-----------------------------------------------")
			fmt.Println("1. ID Peminjam")
			fmt.Println("2. Nama Peminjam")
			fmt.Println("3. Judul Buku")
			fmt.Println("4. Tanggal Pinjam")
			fmt.Println("5. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Pilihan Anda (1/2/3/4/5)?")
			fmt.Scan(&PilihanEditBuku)
			if PilihanEditBuku == 1 {
				fmt.Println("Masukkan ID Baru Peminjam:")
				fmt.Scan(&idBaru)
				Pinjam[ada].idpeminjam = idBaru
			} else if PilihanEditBuku == 2 {
				fmt.Println("Masukkan Nama Baru Peminjam:")
				fmt.Scan(&namabaru)
				Pinjam[ada].namapeminjam = namabaru
			} else if PilihanEditBuku == 3 {
				fmt.Println("Masukkan Judul Baru Buku:")
				fmt.Scan(&judulbaru)
				Pinjam[ada].bukudipinjam = judulbaru
			} else if PilihanEditBuku == 4 {
				fmt.Println("Masukkan Tanggal Pinjam Baru:")
				fmt.Println("Hari:")
				fmt.Scan(&haribaru)
				fmt.Println("Bulan:")
				fmt.Scan(&bulanbaru)
				fmt.Println("Tahun:")
				fmt.Scan(&tahunbaru)
				Pinjam[ada].haripinjam = haribaru
				Pinjam[ada].bulanpinjam = bulanbaru
				Pinjam[ada].tahunpinjam = tahunbaru
			} else if PilihanEditBuku == 5 {
				ada = -1
				x = -1
			}
			if PilihanEditBuku < 5 {
				fmt.Println()
				fmt.Println("-----------------------------------------------")
				fmt.Println("                Edit selesai")
				fmt.Println("-----------------------------------------------")
				fmt.Println("Data baru peminjaman tersebut:")
				fmt.Println(Pinjam[ada].idpeminjam, Pinjam[ada].namapeminjam)
				fmt.Println("Buku yang dipinjam:")
				fmt.Println(Pinjam[ada].bukudipinjam)
				fmt.Printf("Tanggal Peminjaman: %d/%d/%d\n", Pinjam[ada].haripinjam, Pinjam[ada].bulanpinjam, Pinjam[ada].tahunpinjam)
				fmt.Println("-----------------------------------------------")
				fmt.Println("     Apakah anda ingin lanjut mengedit?")
				fmt.Println("-----------------------------------------------")
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")
				fmt.Println("-----------------------------------------------")
				fmt.Println("Pilihan Anda (1/2)?")
				fmt.Scan(&lanjut)
				if lanjut == 1 {
					ada = -1
					x = -1
				}
			}
		}
	}
}

func HapusDataPinjamBuku(Pinjam *DataPinjam, nPinjam *int) {
	var x, lanjut, ada, i, KonfirmasiHapus int
	x = -1
	ada = -1

	for x != 0 && ada == -1 {
		lanjut = 1
		fmt.Println()
		fmt.Println("------------------------------------------------")
		fmt.Println("Silahkan isi ID dari Peminjam yang ingin dihapus")
		fmt.Println("------------------------------------------------")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------")
		fmt.Println("Masukkan ID Peminjam:")
		fmt.Scan(&x)
		ada = seqsearchpinjam(Pinjam, nPinjam, x)
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("            Peminjam Tidak Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("0. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Harap masukkan lagi ID peminjam yang valid:")
			fmt.Scan(&x)
			ada = seqsearchpinjam(Pinjam, nPinjam, x)
		}
		for ada != -1 && lanjut == 1 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("             Data Peminjam Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Data peminjam tersebut:")
			fmt.Println(Pinjam[ada].idpeminjam, Pinjam[ada].namapeminjam)
			fmt.Println("Buku yang dipinjam:")
			fmt.Println(Pinjam[ada].bukudipinjam)
			fmt.Printf("Tanggal Peminjaman: %d/%d/%d\n", Pinjam[ada].haripinjam, Pinjam[ada].bulanpinjam, Pinjam[ada].tahunpinjam)
			fmt.Println("-----------------------------------------------")
			fmt.Println("Apakah anda yakin ingin menghapus data tersebut?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Pilihan Anda (1/2)?")
			fmt.Scan(&KonfirmasiHapus)
			if KonfirmasiHapus == 1 {
				for i = ada; i < *nPinjam-1; i++ {
					Pinjam[i] = Pinjam[i+1]
				}
				*nPinjam = *nPinjam - 1
				fmt.Println()
				fmt.Println("-------------------------------------------------")
				fmt.Println("             Data Berhasil Dihapus")
				fmt.Println("-------------------------------------------------")
				fmt.Println("Apakah anda ingin lanjut menghapus peminjam lain?")
				fmt.Println("-------------------------------------------------")
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")
				fmt.Println("-------------------------------------------------")
				fmt.Println("Pilihan Anda (1/2)?")
				fmt.Scan(&lanjut)
				if lanjut == 1 {
					lanjut = -1
					ada = -1
					x = -1
				}
			} else if KonfirmasiHapus == 2 {
				ada = -1
				x = -1
			}

		}
	}
}

func TarifDanDenda(Pinjam *DataPinjam, nPinjam *int) {
	var harikembali, bulankembali, tahunkembali, x, ada, kembali int
	var telat bool
	x = -1
	for x != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("            Silahkan masukkan ID dari peminjam yang ingin dicek")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("PETUNJUK:")
		fmt.Println("Jika ingin kembali, Silahkan isi 0.")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("ID peminjam buku:")
		fmt.Scan(&x)
		ada = seqsearchpinjam(Pinjam, nPinjam, x)
		for ada == -1 && x != 0 {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("            Peminjam Tidak Ditemukan")
			fmt.Println("-----------------------------------------------")
			fmt.Println("0. Kembali")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Harap masukkan lagi ID peminjam yang valid:")
			fmt.Scan(&x)
			ada = seqsearchpinjam(Pinjam, nPinjam, x)
		}
		for x != 0 && ada != -1 {
			fmt.Println()
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("                       Data Peminjam Ditemukan")
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("Data peminjam tersebut:")
			fmt.Println(Pinjam[ada].idpeminjam, Pinjam[ada].namapeminjam)
			fmt.Println("Buku yang dipinjam:")
			fmt.Println(Pinjam[ada].bukudipinjam)
			fmt.Printf("Tanggal Peminjaman: %d/%d/%d\n", Pinjam[ada].haripinjam, Pinjam[ada].bulanpinjam, Pinjam[ada].tahunpinjam)
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("               Silahkan masukkan tanggal pengembalian buku")
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Jika ingin kembali, Silahkan isi 0 di bagian hari.")
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("Tanggal pengembalian buku:")
			fmt.Println("Hari:")
			fmt.Scan(&harikembali)
			if harikembali != 0 {
				fmt.Println("Bulan:")
				fmt.Scan(&bulankembali)
				fmt.Println("Tahun:")
				fmt.Scan(&tahunkembali)
				telat = ketelatan(Pinjam, harikembali, bulankembali, tahunkembali, ada)
				for telat == true && ada != -1 {
					fmt.Println()
					fmt.Println("--------------------------------------------------------------------------")
					fmt.Println("             Masa peminjaman telah melewati 14 hari")
					fmt.Println("             Peminjam telat dalam mengembalikan buku")
					fmt.Println("--------------------------------------------------------------------------")
					fmt.Println("Peminjam kena denda sebesar Rp.", hitungdenda(Pinjam, harikembali, bulankembali, tahunkembali, ada))
					fmt.Println("--------------------------------------------------------------------------")
					fmt.Println("Masukkan 0 Jika ingin kembali.")
					fmt.Scan(&kembali)
					if kembali == 0 {
						ada = -1
					}
				}
				for telat == false && ada != -1 {
					fmt.Println()
					fmt.Println("--------------------------------------------------------------------------")
					fmt.Println("               Masa peminjaman belum melewati 14 hari")
					fmt.Printf("       Peminjam memiliki sebanyak %d hari untuk mengembalikan buku\n", hitungharidenda(Pinjam, harikembali, bulankembali, tahunkembali, ada))
					fmt.Println("--------------------------------------------------------------------------")
					fmt.Println("Masukkan 0 Jika ingin kembali.")
					fmt.Scan(&kembali)
					if kembali == 0 {
						ada = -1
					}
				}
			} else if harikembali == 0 {
				x = 0
			}

		}
	}
}

func StatusPinjam(Pinjam DataPinjam, nPinjam int) {
	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-18s %-20s\n", "ID Buku", "Judul", "ID Peminjam", "Nama Peminjam")
		fmt.Println("-------------------------------------------------------------------------------------")
		for i = 0; i < nPinjam; i++ {
			fmt.Printf("%-15d %-20s %-18d %-20s\n", Pinjam[i].idBdipinjam, Pinjam[i].bukudipinjam, Pinjam[i].idpeminjam, Pinjam[i].namapeminjam)
		}
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Println("Masukkan 0 Jika ingin kembali.")
		fmt.Scan(&kembali)
	}

}

func BukuFavorit(Book *Buku, nBuku int) {
	var i, kembali int
	sortfavorit(Book, nBuku)
	persentasebuku(Book, nBuku)
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("---------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-18s %-20s\n", "Rank", "Judul", "Banyak Peminjaman", "Persentase Peminjaman")
		fmt.Println("---------------------------------------------------------------------------------")
		for i = 0; i < 5; i++ {
			fmt.Printf("%-15d %-20s %-18d %-20.2f\n", i+1, Book[i].judul, Book[i].countF, Book[i].countP)
		}
		fmt.Println("---------------------------------------------------------------------------------")
		fmt.Println("Masukkan 0 Jika ingin kembali.")
		fmt.Scan(&kembali)
	}
}

func MenuDaftarBuku(Book Buku, nBuku int) {
	var i, kembali int
	kembali = -1
	for kembali != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Pencarian Buku")
		fmt.Println("2. Urutkan Buku")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3)?")
		fmt.Scan(&kembali)
		if kembali == 1 {
			PencarianBuku(Book, nBuku)
		} else if kembali == 2 {
			MenuSortBuku(Book, nBuku)
		}
	}
}

func PencarianBuku(Book Buku, nBuku int) {
	var x string
	var ada int
	x = "hai"
	ada = 0
	for x != "Back" {
		fmt.Println()
		fmt.Println("------------------------------------------------------")
		fmt.Println(" Silahkan masukkan judul dari buku yang ingin dicari")
		fmt.Println("------------------------------------------------------")
		fmt.Println("PETUNJUK:")
		fmt.Println("Masukkan 'Back' Untuk Kembali.")
		fmt.Println("Gunakan '_' sebagai spasi saat memasukkan nama.")
		fmt.Println("------------------------------------------------------")
		fmt.Println("Judul buku:")
		fmt.Scan(&x)
		if x != "Back" {
			ada = seqsearchjudul(Book, nBuku, x)
		}
		for ada == -1 {
			fmt.Println()
			fmt.Println("------------------------------------------------------")
			fmt.Println(" Buku tidak ditemukan, masukkan judul buku yang valid")
			fmt.Println("------------------------------------------------------")
			fmt.Println("PETUNJUK:")
			fmt.Println("Masukkan 'Back' Untuk Kembali.")
			fmt.Println("Gunakan '_' sebagai spasi saat memasukkan nama.")
			fmt.Println("------------------------------------------------------")
			fmt.Println("Judul buku:")
			fmt.Scan(&x)
			if x == "Back" {
				ada = 0
				x = "Back"
			} else {
				ada = seqsearchjudul(Book, nBuku, x)
			}
		}
		for ada != -1 && x != "Back" {
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println("             Buku ditemukan")
			fmt.Println("Data buku tersebut:")
			fmt.Println(Book[ada].id, Book[ada].judul, Book[ada].penulis, Book[ada].genre, Book[ada].tahunterbit)
			fmt.Println("-----------------------------------------------")
			fmt.Println("Masukkan 'Back' Untuk kembali.")
			fmt.Scan(&x)
		}
	}
}

func MenuSortBuku(Book Buku, nBuku int) {
	var pilihsort int
	pilihsort = -1
	for pilihsort != 6 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Kategori Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. ID Buku")
		fmt.Println("2. Tahun Terbit Buku")
		fmt.Println("3. Judul")
		fmt.Println("4. Penulis")
		fmt.Println("5. Genre")
		fmt.Println("6. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3/4/5/6)?")
		fmt.Scan(&pilihsort)
		if pilihsort == 1 {
			sortID(Book, nBuku)
		} else if pilihsort == 2 {
			sortYear(Book, nBuku)
		} else if pilihsort == 3 {
			sortJudul(Book, nBuku)
		} else if pilihsort == 4 {
			sortPenulis(Book, nBuku)
		} else if pilihsort == 5 {
			sortGenre(Book, nBuku)
		}
	}
}

func sortGenre(Book Buku, nBuku int) {
	var pilihtipesort int
	pilihtipesort = -1
	for pilihtipesort != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Tipe Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. A to Z")
		fmt.Println("2. Z to A")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3)?")
		fmt.Scan(&pilihtipesort)
		if pilihtipesort == 1 {
			genre_alfabetascend(Book, nBuku)
		} else if pilihtipesort == 2 {
			genre_alfabetdescend(Book, nBuku)
		}
	}
}

func sortPenulis(Book Buku, nBuku int) {
	var pilihtipesort int
	pilihtipesort = -1
	for pilihtipesort != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Tipe Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. A to Z")
		fmt.Println("2. Z to A")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3)?")
		fmt.Scan(&pilihtipesort)
		if pilihtipesort == 1 {
			penulis_alfabetascend(Book, nBuku)
		} else if pilihtipesort == 2 {
			penulis_alfabetdescend(Book, nBuku)
		}
	}
}

func sortJudul(Book Buku, nBuku int) {
	var pilihtipesort int
	pilihtipesort = -1
	for pilihtipesort != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Tipe Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. A to Z")
		fmt.Println("2. Z to A")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3)?")
		fmt.Scan(&pilihtipesort)
		if pilihtipesort == 1 {
			judul_alfabetascend(Book, nBuku)
		} else if pilihtipesort == 2 {
			judul_alfabetdescend(Book, nBuku)
		}
	}
}

func sortID(Book Buku, nBuku int) {
	var pilihtipesort int
	pilihtipesort = -1
	for pilihtipesort != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Tipe Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. High-to-Low")
		fmt.Println("2. Low-to-High")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilihan Anda (1/2/3)?")
		fmt.Scan(&pilihtipesort)
		if pilihtipesort == 1 {
			sortIDdescend(Book, nBuku)
		} else if pilihtipesort == 2 {
			sortIDascend(Book, nBuku)
		}
	}
}

func sortYear(Book Buku, nBuku int) {
	var pilihtipesort int
	pilihtipesort = -1
	for pilihtipesort != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("       Silahkan Pilih Tipe Pengurutan")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. High-to-Low")
		fmt.Println("2. Low-to-High")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Tipe yang dipilih:")
		fmt.Scan(&pilihtipesort)
		if pilihtipesort == 1 {
			sortyeardescend(Book, nBuku)
		} else if pilihtipesort == 2 {
			sortyearascend(Book, nBuku)
		}
	}
}

func seqsearchpinjam(Pinjam *DataPinjam, nPinjam *int, x int) int {
	var ketemu int = -1
	var i int = 0
	for i < *nPinjam && ketemu == -1 {
		if Pinjam[i].idpeminjam == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func seqsearchid(Book *Buku, nBuku *int, x int) int {
	var ketemu int = -1
	var i int = 0
	for i < *nBuku && ketemu == -1 {
		if Book[i].id == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func seqsearchjudul(Book Buku, nBuku int, x string) int {
	var ketemu int = -1
	var i int = 0
	for i < nBuku && ketemu == -1 {
		if Book[i].judul == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func seqsearchpenulis(Book Buku, nBuku int, x string) int {
	var ketemu int = -1
	var i int = 0
	for i < nBuku && ketemu == -1 {
		if Book[i].penulis == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func seqsearchgenre(Book Buku, nBuku int, x string) int {
	var ketemu int = -1
	var i int = 0
	for i < nBuku && ketemu == -1 {
		if Book[i].genre == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func ketelatan(Pinjam *DataPinjam, harikembali, bulankembali, tahunkembali, x int) bool {
	var totalharikembali, totalharipinjam int
	totalharikembali = harikembali + bulankembali*30 + tahunkembali*360
	totalharipinjam = Pinjam[x].haripinjam + Pinjam[x].bulanpinjam*30 + Pinjam[x].tahunpinjam*360
	if totalharikembali-totalharipinjam > 14 {
		return true
	} else {
		return false
	}
}

func hitungdenda(Pinjam *DataPinjam, harikembali, bulankembali, tahunkembali, x int) int {
	var totalharikembali, totalharipinjam, denda int
	totalharikembali = harikembali + bulankembali*30 + tahunkembali*360
	totalharipinjam = Pinjam[x].haripinjam + Pinjam[x].bulanpinjam*30 + Pinjam[x].tahunpinjam*360
	denda = ((totalharikembali - totalharipinjam) - 14) * 10000
	return denda
}

func hitungharikembali(Pinjam *DataPinjam, x int) int {
	var totalharikembali, totalharipinjam int
	totalharipinjam = Pinjam[x].haripinjam + Pinjam[x].bulanpinjam*30 + Pinjam[x].tahunpinjam*360
	totalharikembali = totalharipinjam + 14
	return totalharikembali
}

func hitungharidenda(Pinjam *DataPinjam, harikembali, bulankembali, tahunkembali, x int) int {
	var totalharikembali, totalharipinjam, denda int
	totalharikembali = harikembali + bulankembali*30 + tahunkembali*360
	totalharipinjam = Pinjam[x].haripinjam + Pinjam[x].bulanpinjam*30 + Pinjam[x].tahunpinjam*360
	denda = (14 - (totalharikembali - totalharipinjam))
	return denda
}

func sortfavorit(Book *Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].countF < Book[j].countF {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}
}

func sortIDdescend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].id < Book[j].id {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")

		fmt.Scan(&kembali)
	}
}

func sortIDascend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].id > Book[j].id {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func sortyearascend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].tahunterbit > Book[j].tahunterbit {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func sortyeardescend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].tahunterbit < Book[j].tahunterbit {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func judul_alfabetascend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].judul > Book[j].judul {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func judul_alfabetdescend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].judul < Book[j].judul {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")

		fmt.Scan(&kembali)
	}
}

func penulis_alfabetascend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].penulis > Book[j].penulis {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func penulis_alfabetdescend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].penulis < Book[j].penulis {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")

		fmt.Scan(&kembali)
	}
}

func genre_alfabetascend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].genre > Book[j].genre {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")
		fmt.Scan(&kembali)
	}
}

func genre_alfabetdescend(Book Buku, nBuku int) {
	var pass, j, idx int
	var tabTemp Buku
	for pass = 1; pass < nBuku; pass++ {
		idx = pass - 1
		for j = pass; j < nBuku; j++ {
			if Book[idx].genre < Book[j].genre {
				idx = j
			}
		}
		tabTemp[0] = Book[idx]
		Book[idx] = Book[pass-1]
		Book[pass-1] = tabTemp[0]
	}

	var i, kembali int
	kembali = -1
	for kembali != 0 {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("%-15s %-20s %-20s %-18s %-18s\n", "ID", "Judul", "Penulis", "Genre", "Tahun Terbit")
		fmt.Println("--------------------------------------------------------------------------------------------")
		for i = 0; i < nBuku; i++ {
			fmt.Printf("%-15d %-20s %-20s %-18s %-18d\n", Book[i].id, Book[i].judul, Book[i].penulis, Book[i].genre, Book[i].tahunterbit)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Masukkan angka 0 untuk kembali.")

		fmt.Scan(&kembali)
	}
}

func persentasebuku(Book *Buku, nBuku int) {
	var i int
	for i = 0; i < nBuku; i++ {
		Book[i].countP = ((float64(Book[i].countF) / float64(nPeminjaman)) * 100 / 100) * 100
	}
}

func cekduplikatidbuku(Book *Buku, nBuku *int, x int) bool {
	var ketemu bool = false
	var i int = 0
	for i < *nBuku && !ketemu {
		if Book[i].id == x {
			ketemu = true
		}
		i++
	}
	return ketemu
}

func cekduplikatidpeminjam(Pinjam DataPinjam, nPinjam int, x int) bool {
	var ketemu bool = false
	var i int = 0
	for i < nPinjam && !ketemu {
		if Pinjam[i].idpeminjam == x {
			ketemu = true
		}
		i++
	}
	return ketemu
}

func cekduplikatidbukudipinjam(Pinjam DataPinjam, nPinjam int, x int) bool {
	var ketemu bool = false
	var i int = 0
	for i < nPinjam && !ketemu {
		if Pinjam[i].idBdipinjam == x {
			ketemu = true
		}
		i++
	}
	return ketemu
}
