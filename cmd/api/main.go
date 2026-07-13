package main

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/config"
	"github.com/Leli2004/API_Go_biblioteca/db"
	"github.com/labstack/echo"

	authorHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/author/delivery/http"
	authorRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/author/repository"
	authorUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/author/usecase"
	
	bookHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/book/delivery/http"
	bookRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/book/repository"
	bookUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/book/usecase"
	
	genreHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/delivery/http"
	genreRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/repository"
	genreUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/usecase"
	
	publisherHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/publisher/delivery/http"
	publisherRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/publisher/repository"
	publisherUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/publisher/usecase"
	
	userHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/user/delivery/http"
	userRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/user/repository"
	userUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/user/usecase"

	bookCopieHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie/delivery/http"
	bookCopieRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie/repository"
	bookCopieUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie/usecase"
)

func main() {
	if err := config.Load(); err != nil {
		panic(fmt.Errorf("ERROR: erro ao carregar configurações: %w", err))
	}

	dbSqlx, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}
	defer dbSqlx.Close()

	e := echo.New()

	// Author
	authorRepo := authorRepository.NewRepository(dbSqlx)
	authorUC := authorUseCase.NewUseCase(authorRepo)
	authorhandler := authorHttp.NewHandler(authorUC)
	authorHttp.MapRoutes(e, authorhandler)

	// Genre
	genreRepo := genreRepository.NewRepository(dbSqlx)
	genreUC := genreUseCase.NewUseCase(genreRepo)
	genreHandler := genreHttp.NewHandler(genreUC)
	genreHttp.MapRoutes(e, genreHandler)

	// Publisher
	publisherRepo := publisherRepository.NewRepository(dbSqlx)
	publisherUC := publisherUseCase.NewUseCase(publisherRepo)
	publisherHandler := publisherHttp.NewHandler(publisherUC)
	publisherHttp.MapRoutes(e, publisherHandler)

	// Book
	bookRepo := bookRepository.NewRepository(dbSqlx)
	bookUC := bookUseCase.NewUseCase(bookRepo)
	bookHandler := bookHttp.NewHandler(bookUC)
	bookHttp.MapRoutes(e, bookHandler)

	// Book Copie
	bookCopieRepo := bookCopieRepository.NewRepository(dbSqlx)
	bookCopieUC := bookCopieUseCase.NewUseCase(bookCopieRepo)
	bookCopieHandler := bookCopieHttp.NewHandler(bookCopieUC)
	bookCopieHttp.MapRoutes(e, bookCopieHandler)

	// User
	userRepo := userRepository.NewRepository(dbSqlx)
	userUC := userUseCase.NewUseCase(userRepo)
	userHandler := userHttp.NewHandler(userUC)
	userHttp.MapRoutes(e, userHandler)

	e.Start(fmt.Sprintf(":%s", config.GetServerPort()))
}
