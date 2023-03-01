package service

import (
	"github.com/achmadrizkin/model"
	"github.com/achmadrizkin/repo"
)

// AddBook untuk menambahkan data buku
func AddMovie(bookData model.Movie) model.Movie {
	return repo.AddMovie(bookData)
}

// GetBook untuk mendapatkan data buku berdasarkan id
func GetMovie(bookId string) (int, model.Movie) {
	return repo.GetMovie(bookId)
}

// GetBooks untuk mendapatkan seluruh data buku
func GetMovies() []model.Movie {
	return repo.GetMovies()
}

// UpdateBook untuk mengedit data buku
func UpdateMovie(bookData model.Movie, id string) model.Movie {
	return repo.UpdateMovie(bookData, id)
}

// DeleteBook untuk menghapus data buku
func DeleteMovie(id string) bool {
	return repo.DeleteMovie(id)
}
