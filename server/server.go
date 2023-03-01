package server

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/achmadrizkin/model"
	pb "github.com/achmadrizkin/proto"
	"github.com/achmadrizkin/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

// fungsi untuk memetakan data dari book model ke bookpb Book
func mapToBookpb(movie model.Movie) *pb.Movie {
	return &pb.Movie{
		Id:   movie.Id,
		Name: movie.Name,
		Desc: movie.Desc,
	}
}

// GetBook untuk mendapatkan data buku berdasarkan id
func (*Server) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {

	// mengambil inputan id dari request
	var bookId string = req.GetId()

	// memanggil GetBook untuk mendapatkan data buku berdasarkan id
	_, result := service.GetMovie(bookId)

	// jika tidak ditemukan, beri pesan error
	if result.Id != bookId {
		return &pb.GetMovieResponse{Status: false, Data: nil}, errors.New("Data not found!")
	}

	// mengolah hasil dari GetBook
	// untuk data di objek response
	var bookData *pb.Movie = &pb.Movie{
		Id:   result.Id,
		Name: result.Name,
		Desc: result.Desc,
	}

	// memberikan response berupa data buku
	return &pb.GetMovieResponse{
		Status: true,
		Data:   bookData,
	}, nil

}

// AddBook untuk menambahkan data buku
func (*Server) AddMovie(ctx context.Context, req *pb.AddMovieRequest) (*pb.AddMovieResponse, error) {

	// menerima request
	var movieRequest *pb.Movie = req.GetMovie()

	// membuat variabel bookData
	// untuk dimasukkan ke storage
	var bookData model.Movie = model.Movie{
		Id:   movieRequest.GetId(),
		Name: movieRequest.GetName(),
		Desc: movieRequest.GetDesc(),
	}

	// memanggil AddBook untuk menambahkan data buku
	var insertedMovie model.Movie = service.AddMovie(bookData)

	// mengembalikan response berupa data buku
	// yang telah ditambahkan
	return &pb.AddMovieResponse{
		Status: true,
		Data:   mapToBookpb(insertedMovie),
	}, nil
}

// UpdateBook untuk mengubah data buku
func (*Server) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {

	// menerima request
	var movieRequest *pb.Movie = req.GetMovie()

	// membuat variabel bookData
	// untuk mengubah data buku
	var movieData model.Movie = model.Movie{
		Id:   movieRequest.GetId(),
		Name: movieRequest.GetName(),
		Desc: movieRequest.GetDesc(),
	}

	// memanggil UpdateBook untuk mengubah data buku berdasarkan id
	var updatedMovie model.Movie = service.UpdateMovie(movieData, movieData.Id)

	// mengembalikan response berupa
	// data buku yang telah diubah
	return &pb.UpdateMovieResponse{
		Status: true,
		Data:   mapToBookpb(updatedMovie),
	}, nil
}

// DeleteBook untuk menghapus data buku
func (*Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {

	// mengambil inputan id dari request
	var bookId string = req.GetBookId()

	// memanggil DeleteBook untuk menghapus data buku berdasarkan id
	var result bool = service.DeleteMovie(bookId)

	// mengembalikan response
	return &pb.DeleteMovieResponse{
		Status: result,
	}, nil
}

// GetBooks untuk mendapatkan seluruh data buku
func (*Server) GetMovies(req *pb.GetMoviesRequest, stream pb.MovieService_GetMoviesServer) error {

	// memanggil GetBooks untuk mendapatkan seluruh data buku
	var books []model.Movie = service.GetMovies()

	// melakukan iterasi pada setiap data buku
	for _, book := range books {
		// setiap data buku
		// dikirimkan melalui stream
		stream.Send(&pb.GetMoviesResponse{
			Status: true,
			Data:   mapToBookpb(book),
		})
	}

	// karena tidak terdapat error
	// kembalian nilai nil / kosong
	return nil
}

// AddBatchBook untuk menambahkan sekumpulan data buku
func (*Server) AddBatchMovie(stream pb.MovieService_AddBatchMovieServer) error {
	// untuk setiap request
	for {
		// menerima request
		req, err := stream.Recv()

		// jika request tidak ada
		// tutup stream
		if err == io.EOF {
			return stream.SendAndClose(&pb.AddBatchMovieResponse{
				Status:  true,
				Message: "All book data inserted successfully!",
			})
		}

		// jika terdapat error di server
		// tampilkan pesan error
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error, insert batch failed: %v", err),
			)
		}

		// mendapatkan request data buku
		var bookData *pb.Movie = req.GetMovie()

		// membuat variabel book
		// untuk dimasukkan ke storage
		var book model.Movie = model.Movie{
			Id:   bookData.GetId(),
			Name: bookData.GetName(),
			Desc: bookData.GetDesc(),
		}

		// tambahkan data buku
		service.AddMovie(book)
	}
}
