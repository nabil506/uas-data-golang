package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"uts_data/model"
)

func MenuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Tampilkan Produk")
		fmt.Println("3. Update Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Search Produk")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Anda memilih: Tambah Produk")
			Insert()
		case 2:
			fmt.Println("Anda memilih: Tampilkan Produk")
			Views()
		case 3:
			fmt.Println("Anda memilih: Update Produk")
			Update()
		case 4:
			fmt.Println("Anda memilih: Hapus Produk")
			Delete()
		case 5:
			fmt.Println("Anda memilih: Search Produk")
			Search()
		case 6:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func Insert() {
	var nama_produk, kategori string
	var id, harga int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Produk: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Produk: ")
	nama_produk, _ = reader.ReadString('\n')
	nama_produk = strings.TrimSpace(nama_produk)

	fmt.Print("Masukkan Kategori Produk: ")
	kategori, _ = reader.ReadString('\n')
	kategori = strings.TrimSpace(kategori)

	fmt.Print("Masukkan Harga Produk: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))

	model.Create(nama_produk, kategori, id, harga)
}

func Views() {
	fmt.Println("Daftar Produk")
	for _, emp := range model.ReadProduk() {
		fmt.Println("--- Produk Wartawarti", " ---")
		fmt.Println("ID Produk\t\t  : ", emp.ID_Produk)
		fmt.Println("Nama Produk\t\t  : ", emp.Nama_Produk)
		fmt.Println("Kategori Produk\t\t  : ", emp.Kategori)
		fmt.Println("Harga Produk\t\t  : ", emp.Harga_Produk)
		fmt.Print()
	}
}

func Update() {
	var nama_produk, kategori string
	var id, harga int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Produk: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Produk: ")
	nama_produk, _ = reader.ReadString('\n')
	nama_produk = strings.TrimSpace(nama_produk)

	fmt.Print("Masukkan Kategori Produk: ")
	kategori, _ = reader.ReadString('\n')
	kategori = strings.TrimSpace(kategori)

	fmt.Print("Masukkan Harga Produk: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))

	model.Update(nama_produk, kategori, id, harga)

}

func Delete() {
	var id int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Produk yang akan dihapus: ")
	fmt.Scan(&id)
	reader.ReadString('\n')

	model.Delete(id)
}

func Search() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	kategori, _ := reader.ReadString('\n')
	kategori = strings.TrimSpace(kategori)

	hasil := model.SearchProduk(kategori)

	if len(hasil) == 0 {
		fmt.Println("Tidak ada produk dengan kategori tersebut.")
		return
	}

	for _, emp := range hasil {
		fmt.Println("--- Produk Wartawarti", " ---")
		fmt.Println("ID Produk\t\t  : ", emp.ID_Produk)
		fmt.Println("Nama Produk\t\t  : ", emp.Nama_Produk)
		fmt.Println("Harga Produk\t\t  : ", emp.Harga_Produk)
		fmt.Println("Kategori Produk\t\t  : ", emp.Kategori)
	}
}
