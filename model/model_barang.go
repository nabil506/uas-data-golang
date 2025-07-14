package model

import (
	"strings"
	"uts_data/node"
)

var DaftarBarang node.Kasir

func CreateProduk(emp node.Wartawarti) bool {
	Nodebaru := &node.Kasir{
		Data:   emp,
		Lanjut: nil,
	}
	if DaftarBarang.Lanjut == nil {
		DaftarBarang.Lanjut = Nodebaru
		return true
	} else {
		temp := &DaftarBarang
		for temp.Lanjut != nil {
			temp = temp.Lanjut
		}
		temp.Lanjut = Nodebaru
		return true
	}
	return false
}

func ReadProduk() []node.Wartawarti {
	daftarBarang := []node.Wartawarti{}
	Nodebaru := &DaftarBarang
	for Nodebaru.Lanjut != nil {
		daftarBarang = append(daftarBarang, Nodebaru.Lanjut.Data)
		Nodebaru = Nodebaru.Lanjut
	}
	return daftarBarang
}

func UpdateProduk(emp node.Wartawarti, id int) bool {
	temp := DaftarBarang.Lanjut
	for temp != nil {
		if temp.Data.ID_Produk == id {
			temp.Data = emp
			return true
		}
		temp = temp.Lanjut
	}
	return false
}

func GetBarangById(id int) (*node.Wartawarti, bool) {
	temp := DaftarBarang.Lanjut
	for temp != nil {
		if temp.Data.ID_Produk == id {
			return &temp.Data, true
		}
		temp = temp.Lanjut
	}
	return nil, false
}

func DeleteProduk(id int) bool {
	if DaftarBarang.Lanjut == nil {
		return false
	}

	temp := &DaftarBarang
	for temp.Lanjut != nil {
		if temp.Lanjut.Data.ID_Produk == id {
			temp.Lanjut = temp.Lanjut.Lanjut
			return true
		}
		temp = temp.Lanjut
	}
	return false
}

func SearchProduk(kategori string) []node.Wartawarti {
	Produk := []node.Wartawarti{}
	temp := DaftarBarang.Lanjut
	for temp != nil {
		if strings.Contains(strings.ToLower(temp.Data.Kategori), strings.ToLower(kategori)) {
			Produk = append(Produk, temp.Data)
		}
		temp = temp.Lanjut
	}
	return Produk
}
