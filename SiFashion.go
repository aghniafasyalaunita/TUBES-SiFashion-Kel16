package main

import "fmt"

// ==============================
// KONSTANTA DAN TIPE DATA
// ==============================

const NMAX int = 999

type Produk struct {
	id           int
	nama         string
	kategori     string // misal: "Atasan", "Bawahan", "Dress", "Aksesoris"
	ukuran       string // misal: "S", "M", "L", "XL", "XXL", "Allsize"
	warna        string // misal: "Merah", "Biru", "Hitam"
	harga        float64
	stok         int
	totalTerjual int
}

type DaftarProduk [NMAX]Produk

// ==============================
// SUBPROGRAM BANTU
// ==============================

func cetakGaris() {
	fmt.Println("============================================================")
}

// buat nampilin header menu dengan judul tertentu
func cetakHeader(judul string) {
	cetakGaris()
	fmt.Println("  ", judul)
	cetakGaris()
}

// nampilin SATU produk aja
func cetakSatuProduk(p Produk) {
	fmt.Printf("  %-21s: %d\n", "ID", p.id)
	fmt.Printf("  %-21s: %s\n", "Nama", p.nama)
	fmt.Printf("  %-21s: %s\n", "Kategori", p.kategori)
	fmt.Printf("  %-21s: %s\n", "Ukuran", p.ukuran)
	fmt.Printf("  %-21s: %s\n", "Warna", p.warna)
	fmt.Printf("  %-21s: Rp %.0f\n", "Harga", p.harga)
	fmt.Printf("  %-21s: %d pcs\n", "Stok", p.stok)
	fmt.Printf("  %-21s: %d pcs\n", "Total Terjual", p.totalTerjual)
	fmt.Println("  ----------------------------------------------------------")
}

// nampilin SEMUA isi daftar produk
func cetakSemuaProduk(tab DaftarProduk, jumlah int) {
	var i int
	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk yang terdaftar.")
	} else {
		i = 0
		for i < jumlah {
			cetakSatuProduk(tab[i])
			i++
		}
	}
}

// ngecek apakah id udah kepake atau belum
func idSudahAda(tab DaftarProduk, jumlah int, idCari int) bool {
	var found bool
	var i int = 0
	found = false
	for i < jumlah && !found {
		found = tab[i].id == idCari
		i = i + 1
	}
	return found
}

// ==============================
// FITUR A — TAMBAH, UBAH, HAPUS PRODUK
// ==============================

// prosedur buat nambahin produk baru ke daftar
func tambahProduk(tab *DaftarProduk, jumlah *int) {
	cetakHeader("TAMBAH PRODUK BARU")
	var p Produk
	if *jumlah >= NMAX {
		fmt.Println("  [!] Daftar produk sudah penuh! Kapasitas maksimal:", NMAX)
	} else {

		fmt.Print("  Masukkan ID produk (angka): ")
		fmt.Scanln(&p.id)

		// cek dulu kalau id nya udah kepake
		for idSudahAda(*tab, *jumlah, p.id) {
			fmt.Println("  [!] ID sudah dipakai, coba ID lain.")
			fmt.Printf("  %-26s: ", "Masukkan ID produk:")
			fmt.Scanln(&p.id)
		}

		fmt.Printf("  %-26s: ", "Nama produk")
		fmt.Scanln(&p.nama)
		fmt.Printf("  %-26s: ", "Kategori")
		fmt.Scanln(&p.kategori)
		fmt.Printf("  %-26s: ", "Ukuran")
		fmt.Scanln(&p.ukuran)
		fmt.Printf("  %-26s: ", "Warna")
		fmt.Scanln(&p.warna)
		fmt.Printf("  %-25s : ", "Harga (Rp)")
		fmt.Scanln(&p.harga)
		fmt.Printf("  %-26s: ", "Stok (Pcs)")
		fmt.Scanln(&p.stok)
		fmt.Printf("  %-26s: ", "Total terjual")
		fmt.Scanln(&p.totalTerjual)
		tab[*jumlah] = p
		*jumlah++

		fmt.Println("  [OK] Produk berhasil ditambahkan!")
	}

}

// prosedur buat ngubah data produk yang sudah ada
// dicari dulu by id, trs kl ketemu baru diubah
func ubahProduk(tab *DaftarProduk, jumlah int) {
	cetakHeader("UBAH DATA PRODUK")
	var i, cariID, idx int
	i = 0
	idx = -1
	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")
	} else {
		fmt.Print("  Masukkan ID produk yang mau diubah: ")
		fmt.Scanln(&cariID)

		// sequential search cari id
		for i < jumlah && idx == -1 {
			if tab[i].id == cariID {
				idx = i
			}
			i++
		}
		if idx == -1 {
			fmt.Println("  [!] Produk dengan ID tersebut tidak ditemukan.")
		} else {

			fmt.Println("\n  Data lama:")
			cetakSatuProduk(tab[idx])

			fmt.Printf("  %-26s: \n", "Masukkan data baru:")
			fmt.Printf("  %-26s: ", "Nama produk")
			fmt.Scanln(&tab[idx].nama)
			fmt.Printf("  %-26s: ", "Kategori")
			fmt.Scanln(&tab[idx].kategori)
			fmt.Printf("  %-26s: ", "Ukuran")
			fmt.Scanln(&tab[idx].ukuran)
			fmt.Printf("  %-26s: ", "Warna")
			fmt.Scanln(&tab[idx].warna)
			fmt.Printf("  %-25s : ", "Harga (Rp)")
			fmt.Scanln(&tab[idx].harga)
			fmt.Printf("  %-26s: ", "Stok (Pcs)")
			fmt.Scanln(&tab[idx].stok)
			fmt.Printf("  %-26s: ", "Total Terjual")
			fmt.Scanln(&tab[idx].totalTerjual)
			fmt.Println("  [OK] Data produk berhasil diubah!")
		}
	}

}

// prosedur buat hapus produk dari daftar kemudian array ditimpa dg isi data array selanjutny
func hapusProduk(tab *DaftarProduk, jumlah *int) {
	cetakHeader("HAPUS PRODUK")
	var i, j, idx, cariID int
	var konfirmasi string
	i = 0
	idx = -1
	if *jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")
	} else {
		fmt.Print("  Masukkan ID produk yang mau dihapus: ")
		fmt.Scanln(&cariID)

		// cari dulu indexnya pake sequential search
		for i < *jumlah && idx == -1 {
			if tab[i].id == cariID {
				idx = i
			}
			i++
		}

		if idx == -1 {
			fmt.Println("  [!] Produk dengan ID tersebut tidak ditemukan.")
		} else {
			fmt.Println("\n  Produk yang akan dihapus:")
			cetakSatuProduk(tab[idx])

			fmt.Print("  Yakin mau hapus?: ")
			fmt.Scanln(&konfirmasi)

			if konfirmasi != "Ya" {
				fmt.Println("  [!] Hapus dibatalkan.")
			} else {
				j = idx
				for i < *jumlah {
					tab[j] = tab[j+1]
					j++
				}
				*jumlah--

				fmt.Println("  [OK] Produk berhasil dihapus!")
			}
		}
	}
}

// ==============================
// FITUR B — CATAT STOK
// ==============================

// prosedur buat update stok suatu produk dg suatu kombinasi warna ukuran
// bisa tambah stok atau kurangi stok
func catatStok(tab *DaftarProduk, jumlah int) {
	cetakHeader("CATAT DETAIL STOK PRODUK")
	var i, idx, pilihan, jumlahUbahStok int
	var cariUkuran, cariWarna string
	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")
	} else {
		fmt.Printf("  %-18s: ", "Masukkan ukuran")
		fmt.Scanln(&cariUkuran)
		fmt.Printf("  %-18s: ", "Masukkan warna")
		fmt.Scanln(&cariWarna)
		// seq search buat cari produk kombinasi uk warna tertentu
		i = 0
		idx = -1
		for i < jumlah && idx == -1 {
			if (*tab)[i].ukuran == cariUkuran && (*tab)[i].warna == cariWarna {
				idx = i
			}
			i++
		}

		if idx == -1 {
			fmt.Println("  [!] Produk tidak ditemukan.")
		} else {
			fmt.Println("\n  Produk ditemukan:")
			cetakSatuProduk(tab[idx])

			fmt.Println("  Pilih aksi stok:")
			fmt.Println("  1. Tambah stok (barang masuk)")
			fmt.Println("  2. Kurangi stok (barang keluar)")
			fmt.Print("  Pilihan: ")
			fmt.Scanln(&pilihan)

			if pilihan == 1 {
				fmt.Print("  Jumlah stok: ")
				fmt.Scanln(&jumlahUbahStok)
				tab[idx].stok = tab[idx].stok + jumlahUbahStok
				fmt.Printf("  [OK] Stok bertambah jadi %d pcs\n", tab[idx].stok)
			} else if pilihan == 2 {
				fmt.Print("  Jumlah stok: ")
				fmt.Scanln(&jumlahUbahStok)
				if jumlahUbahStok > tab[idx].stok {
					fmt.Println("  [!] Stok tidak cukup!")
				} else {
					tab[idx].stok = tab[idx].stok - jumlahUbahStok
					fmt.Printf("  [OK] Stok berkurang jadi %d pcs\n", tab[idx].stok)
				}
			} else {
				fmt.Println("  [!] Pilihan tidak valid.")
			}

		}
	}

}

// ==============================
// FITUR C — PENCARIAN (SEQUENTIAL & BINARY)
// ==============================

// sequential search — cari produk berdasarkan ukuran
func seqSearchUkuran(tab DaftarProduk, jumlah int, ukuranCari string) int {
	var found, i int
	found = -1
	i = 0
	for i < jumlah && found == -1 {
		if tab[i].ukuran == ukuranCari {
			found = i
		}
		i++
	}
	return found
}

// binary search — cari produk berdasarkan warna
// karena binary brrti kita hrs urutin dlu brdasarkan warna
func selectionSortWarna(tab *DaftarProduk, jumlah int) {
	var pass, idx, i int
	var temp Produk
	pass = 1
	for pass <= jumlah-1 {
		idx = pass - 1
		i = pass
		for i < jumlah {
			if tab[idx].warna > tab[i].warna {
				idx = i
			}
			i++
		}
		temp = tab[pass-1]
		tab[pass-1] = tab[idx]
		tab[idx] = temp
		pass++
	}
}

func binarySearchWarna(tab *DaftarProduk, jumlah int, warnaCari string) int {
	selectionSortWarna(tab, jumlah)
	var found, left, right, mid int
	found = -1
	left = 0
	right = jumlah - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if warnaCari < (*tab)[mid].warna {
			right = mid - 1
		} else if warnaCari > (*tab)[mid].warna {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

// menu pencarian — user pilih mau cari by ukuran, warna, atau harga
func menuCari(tab DaftarProduk, jumlah int) {
	cetakHeader("CARI PRODUK")
	var ukuranCari, warnaCari string
	var pilihan, idxUkuran, idxWarna int

	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")

	} else {
		fmt.Println("  Cari berdasarkan apa?")
		fmt.Println("  1. Ukuran (Sequential Search)")
		fmt.Println("  2. Warna  (Binary Search)")
		fmt.Print("  Pilihan: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fmt.Print("  Masukkan ukuran yang dicari: ")
			fmt.Scanln(&ukuranCari)
			idxUkuran = seqSearchUkuran(tab, jumlah, ukuranCari)
			if idxUkuran == -1 {
				fmt.Println("  [!] Tidak ada produk dengan ukuran tersebut.")
			} else {
				fmt.Println("\n  Produk ditemukan:")
				cetakSatuProduk(tab[idxUkuran])
			}
		} else if pilihan == 2 {
			fmt.Print("  Masukkan warna yang dicari: ")
			fmt.Scanln(&warnaCari)
			idxWarna = binarySearchWarna(&tab, jumlah, warnaCari)
			if idxWarna == -1 {
				fmt.Println("  [!] Warna tidak ditemukan.")
			} else {
				fmt.Println("\n  Produk ditemukan:")
				cetakSatuProduk(tab[idxWarna])
			}
		} else {
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}

}

// ==============================
// FITUR D — PENGURUTAN (SELECTION & INSERTION SORT)
// ==============================

// selection sort ASCENDING berdasarkan HARGA
func selectionSortHarga(tab *DaftarProduk, jumlah int) {
	var pass, idx, i int
	var temp Produk

	pass = 1
	for pass <= jumlah-1 {
		idx = pass - 1
		i = pass
		for i < jumlah {
			if (*tab)[idx].harga > (*tab)[i].harga {
				idx = i
			}
			i++
		}

		temp = (*tab)[pass-1]
		(*tab)[pass-1] = (*tab)[idx]
		(*tab)[idx] = temp
		pass++
	}
}

// insertion sort DESCENDING berdasarkan STOK
func insertionSortStok(tab *DaftarProduk, jumlah int) {
	var pass, i int
	var temp Produk

	pass = 1
	for pass <= jumlah-1 {
		i = pass
		temp = (*tab)[pass]
		for i > 0 && temp.stok > (*tab)[i-1].stok {
			(*tab)[i] = (*tab)[i-1]
			i--
		}
		(*tab)[i] = temp
		pass++
	}
}

// menu pengurutan — user pilih mau sort by harga atau stok
func menuUrut(tab *DaftarProduk, jumlah int) {
	cetakHeader("URUTKAN PRODUK")
	var pilihan int
	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")
	} else {
		fmt.Println("  Urutkan berdasarkan apa?")
		fmt.Println("  1. Harga   — Dari termurah ke termahal (Selection Sort Ascending)")
		fmt.Println("  2. Stok    — Dari terbanyak ke tersedikit (Insertion Sort, Descending)")

		fmt.Print("  Pilihan: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			selectionSortHarga(tab, jumlah)
			fmt.Println("\n  [OK] Data sudah diurutkan berdasarkan harga (termurah dulu):")
			cetakSemuaProduk(*tab, jumlah)
		} else if pilihan == 2 {
			insertionSortStok(tab, jumlah)
			fmt.Println("\n  [OK] Data sudah diurutkan berdasarkan stok (terbanyak dulu):")
			cetakSemuaProduk(*tab, jumlah)
		} else {
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}

}

// ==============================
// FITUR E — STATISTIK
// ==============================

// mengurutkan berdasarkan totalterjual (populer) dulu, klo total terjual sm aj brrti urutkan berdasarkan stok
func cariTotalTerjualTerbanyak(tab DaftarProduk, jumlah int, kategoriCari string) {
	var i, idxMax int
	//sequential search
	idxMax = -1
	if jumlah == 0 {
		fmt.Println("  [!] Belum ada produk.")
	} else {
		i = 0
		for i < jumlah {
			if tab[i].kategori == kategoriCari {
				if idxMax == -1 {
					idxMax = i
				} else if tab[i].totalTerjual > tab[idxMax].totalTerjual {
					idxMax = i
				} else if tab[i].totalTerjual == tab[idxMax].totalTerjual {
					if tab[i].stok > tab[idxMax].stok {
						idxMax = i
					}
				}
			}
			i++
		}
		if idxMax == -1 {
			fmt.Println("  [!] Kategori tidak ditemukan.")
		} else {
			fmt.Println("\n  Produk paling populer: ")
			cetakSatuProduk(tab[idxMax])
		}
	}
}

func tampilStatistik(tab DaftarProduk, jumlah int) {
	var kategoriCari, konfirmasi string

	for konfirmasi != "Tidak" {
		fmt.Print("  Masukan Kategori dicari: ")
		fmt.Scanln(&kategoriCari)
		cariTotalTerjualTerbanyak(tab, jumlah, kategoriCari)

		fmt.Print("  Apakah mau mencari statistik kategori lain?")
		fmt.Scanln(&konfirmasi)
	}
	fmt.Println("  [!] Pencarian statistik selesai.")

}

// ==============================
// FUNGSI BUAT ISI DATA CONTOH
// ==============================

func isiDataContoh(tab *DaftarProduk, jumlah *int) {
	tab[0] = Produk{1, "Kaos_Polos", "Atasan", "M", "Putih", 85000, 50, 230}
	tab[1] = Produk{2, "Kaos_Polos", "Atasan", "L", "Hitam", 85000, 30, 120}
	tab[2] = Produk{3, "Kemeja_Formal", "Atasan", "M", "Biru", 250000, 20, 175}
	tab[3] = Produk{4, "Kemeja_Formal", "Atasan", "XL", "Putih", 250000, 15, 230}
	tab[4] = Produk{5, "Blue_Jeans", "Bawahan", "L", "Biru", 350000, 25, 187}
	tab[5] = Produk{6, "Celana_Chino", "Bawahan", "L", "Cream", 280000, 18, 78}
	tab[6] = Produk{7, "Topi_Baseball", "Aksesoris", "Allsize", "Hitam", 120000, 40, 65}
	tab[7] = Produk{8, "Ikat_Pinggang", "Aksesoris", "Allsize", "Coklat", 45000, 35, 89}
	tab[8] = Produk{9, "Long_Dress", "Dress", "Allsize", "Putih", 175000, 55, 99}
	tab[9] = Produk{10, "Flower_Dress", "Dress", "L", "Pink", 195000, 66, 105}
	tab[10] = Produk{11, "Korean_Dress", "Dress", "S", "Kuning", 189000, 75, 189}
	tab[11] = Produk{12, "Kalung", "Aksesoris", "Allsize", "Putih", 32000, 37, 79}
	tab[12] = Produk{13, "Blouse", "Atasan", "M", "Kuning", 185000, 56, 289}
	tab[13] = Produk{14, "Celana_Kulot", "Bawahan", "L", "Putih", 160000, 29, 177}
	tab[14] = Produk{15, "Celana_Balon", "Bawahan", "L", "Coklat", 156000, 23, 96}
	tab[15] = Produk{16, "Flare_Jeans", "Bawahan", "M", "Biru", 230000, 24, 87}
	tab[16] = Produk{17, "Bandana", "Aksesoris", "Allsize", "Coklat", 12000, 34, 45}
	tab[17] = Produk{18, "Gelang", "Aksesoris", "Allsize", "Silver", 32000, 37, 79}
	tab[18] = Produk{19, "Polkadot_Dress", "Dress", "L", "Putih", 199000, 65, 176}
	tab[19] = Produk{20, "Abaya", "Dress", "Allsize", "Hitam", 289000, 34, 80}
	*jumlah = 20
}

// ==============================
// MENU UTAMA
// ==============================

func main() {
	var daftar DaftarProduk
	var n int
	var pilihan int
	var selesai bool
	selesai = false
	isiDataContoh(&daftar, &n)
	// agar berhenti kalau user pilih 0
	for !selesai {
		cetakGaris()
		fmt.Println("  SIFASHION — Sistem Manajemen Inventaris Produk Fashion")
		cetakGaris()
		fmt.Printf("  Jumlah produk terdaftar: %d item\n", n)
		cetakGaris()
		fmt.Println("  MENU UTAMA:")
		fmt.Println("  1. Lihat semua produk")
		fmt.Println("  2. Tambah produk baru")
		fmt.Println("  3. Ubah data produk")
		fmt.Println("  4. Hapus produk")
		fmt.Println("  5. Catat stok produk")
		fmt.Println("  6. Cari produk")
		fmt.Println("  7. Urutkan produk")
		fmt.Println("  8. Statistik produk")
		fmt.Println("  0. Keluar")
		cetakGaris()
		fmt.Print("  Pilihan kamu: ")
		fmt.Scanln(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			// lihat semua produk
			cetakHeader("DAFTAR SEMUA PRODUK")
			cetakSemuaProduk(daftar, n)

		} else if pilihan == 2 {
			// tambah produk baru
			tambahProduk(&daftar, &n)

		} else if pilihan == 3 {
			// ubah data produk
			ubahProduk(&daftar, n)

		} else if pilihan == 4 {
			// hapus produk
			hapusProduk(&daftar, &n)

		} else if pilihan == 5 {
			// catat stok
			catatStok(&daftar, n)

		} else if pilihan == 6 {
			// cari produk
			menuCari(daftar, n)

		} else if pilihan == 7 {
			// urutkan produk
			menuUrut(&daftar, n)

		} else if pilihan == 8 {
			// statistik
			tampilStatistik(daftar, n)

		} else if pilihan == 0 {
			// keluar
			fmt.Println("  Terima kasih sudah memakai SiFashion. Sampai jumpa!")
			selesai = true

		} else {
			fmt.Println("  [!] Pilihan tidak valid, coba lagi.")
		}

		fmt.Println()
	}
}
