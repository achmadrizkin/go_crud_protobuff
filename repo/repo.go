package repo

import (
	"database/sql"
	"log"

	"github.com/achmadrizkin/database"
	"github.com/achmadrizkin/model"
	"github.com/google/uuid"
)

func AddMovie(movieData model.Movie) model.Movie {
	// membuat id menggunakan uuid
	var uuid string = uuid.New().String()

	// memasukkan data buku ke dalam database
	_, err := database.DB.Query("INSERT INTO movies (id, name, descc) VALUES (?, ?, ?)",
		uuid,
		movieData.Name,
		movieData.Desc,
	)

	// jika terdapat error, tampilkan pesan error
	if err != nil {
		log.Fatalf("Insert data failed: %v", err)
		return model.Movie{}
	}

	// mengembalikan data buku yang dimasukkan
	movieData.Id = uuid

	return movieData
}

// GetBook untuk mendapatkan data buku berdasarkan id
func GetMovie(movieId string) (int, model.Movie) {

	// membuat variabel book
	// untuk menyimpan data buku berdasarkan id
	var movie model.Movie = model.Movie{}

	// mendapatkan data buku berdasarkan id
	row, err := database.DB.Query("SELECT * FROM movies WHERE id = ?", movieId)

	// menampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return 0, model.Movie{}
	}

	// Close() akan dipanggil
	// jika data dari database sudah didapatkan
	defer row.Close()

	// untuk setiap baris data
	for row.Next() {
		// masukkan berbagai atribut data buku seperti judul dll.
		// ke dalam variabel book
		switch err := row.Scan(&movie.Id, &movie.Name, &movie.Desc); err {

		// jika data tidak ditemukan
		// tampilkan pesan error
		case sql.ErrNoRows:
			log.Printf("Data not found: %v", err)
			return 0, model.Movie{}

		// jika tidak ada error
		// tampilkan data buku
		case nil:
			log.Println(movie)

		// tampilkan error untuk kondisi default
		default:
			log.Printf("Data cannot be retrieved: %v", err)
			return 0, model.Movie{}
		}
	}

	// mengembalikan angka 1 sebagai tanda data ditemukan
	// dan data buku yang ditemukan
	return 1, movie
}

// GetBooks untuk mendapatkan seluruh data buku
func GetMovies() []model.Movie {

	// mendapatkan seluruh data buku
	rows, err := database.DB.Query("SELECT * FROM movies")

	// tampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return []model.Movie{}
	}

	// close dipanggil jika
	// seluruh data berhasil diambil
	defer rows.Close()

	// membuat variabel books
	// untuk menampung berbagai data buku
	var movies []model.Movie = []model.Movie{}

	// untuk setiap data
	for rows.Next() {
		// membuat variabel book
		// untuk menyimpan sebuah data buku
		var movie model.Movie = model.Movie{}

		// masukkan berbagai atribut data buku seperti judul dll.
		// ke dalam variabel book
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Desc)

		// jika terdapat error, tampilkan error
		if err != nil {
			log.Printf("Data cannot be retrieved: %v", err)
			return []model.Movie{}
		}

		// masukkan data buku ke dalam books
		movies = append(movies, movie)
	}

	// jika jumlah data di dalam books
	// sama dengan 0, maka data kosong
	if len(movies) == 0 {
		log.Println("Movies data not found")
	}

	// mengembalikan sekumpulan data buku
	return movies
}

// UpdateBook untuk mengedit data buku
func UpdateMovie(movieData model.Movie, id string) model.Movie {

	// mengubah data buku berdasarkan id
	_, err := database.DB.Query("UPDATE movies SET name=?, descc=? WHERE id=?",
		movieData.Name,
		movieData.Desc,
		movieData.Id,
	)

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Update data failed: %v", err)
		return model.Movie{}
	}

	// mengembalikan data buku yang telah diubah
	return movieData
}

// DeleteBook untuk menghapus data buku
func DeleteMovie(id string) bool {

	// menghapus data buku berdasarkan id
	_, err := database.DB.Query("DELETE FROM movies WHERE id=?", id)

	// jika terdapat error, tampilkan pesan error
	// nilai false dikembalikan
	if err != nil {
		log.Fatalf("Delete data failed: %v", err)
		return false
	}

	// nilai true dikembalikan jika data berhasil dihapus
	return true
}
