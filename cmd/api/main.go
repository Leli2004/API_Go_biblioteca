package main

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/config"
	"github.com/Leli2004/API_Go_biblioteca/db"
	"github.com/labstack/echo"

	authorHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/author/delivery/http"
	authorRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/author/repository"
	authorUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/author/usecase"
	genreHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/delivery/http"
	genreRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/repository"
	genreUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/usecase"
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

	e.Start(fmt.Sprintf(":%s", config.GetServerPort()))
}
