package node

type Wartawarti struct{
	Nama_Produk,Kategori string
	Harga_Produk int	
	ID_Produk int
}

type Kasir struct{
	Data Wartawarti
	Lanjut *Kasir
}