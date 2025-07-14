package main

import (
	"net/http"
	"strconv"
	"text/template"
	"uts_data/model"
	"uts_data/node"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("templates/index.html")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("search")
	var dataBarang []node.Wartawarti

	if query != "" {
		dataBarang = model.SearchProduk(query)
	} else {
		dataBarang = model.ReadProduk()
	}

	templ.Execute(w, dataBarang)
}

func InsertForm(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}

func InsertProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id, _ := strconv.Atoi(r.FormValue("id"))
		harga, _ := strconv.Atoi(r.FormValue("Harga_Produk"))
		barang := node.Wartawarti{
			ID_Produk:    id,
			Nama_Produk:  r.FormValue("Nama_Produk"),
			Kategori:     r.FormValue("Kategori_Produk"),
			Harga_Produk: harga,
		}
		model.CreateProduk(barang)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	barang, ok := model.GetBarangById(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template.ParseFiles("templates/edit.html")
	tmpl.Execute(w, barang)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		harga, _ := strconv.Atoi(r.FormValue("Harga_Produk"))

		update := node.Wartawarti{
			ID_Produk:    id,
			Nama_Produk:  r.FormValue("Nama_Produk"),
			Kategori:     r.FormValue("Kategori_Produk"),
			Harga_Produk: harga,
		}
		model.UpdateProduk(update, id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	model.DeleteProduk(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	// view.MenuUtama()

	http.HandleFunc("/", Index)
	http.HandleFunc("/insert", InsertForm)
	http.HandleFunc("/insert-process", InsertProcess)
	http.HandleFunc("/update", UpdateForm)
	http.HandleFunc("/update-process", UpdateProcess)
	http.HandleFunc("/delete", DeleteBarang)

	http.ListenAndServe(":8080", nil)
}
