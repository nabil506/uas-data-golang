package model

import(
	"uts_data/node"
)

func Create(nama_produk,kategori string , id, harga int){
	datanya := node.Wartawarti{
		ID_Produk: id,
		Nama_Produk: nama_produk,
		Kategori: kategori,
		Harga_Produk: harga,
	}
	 CreateProduk(datanya)
}