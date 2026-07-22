package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	loanHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/loan/delivery/http"
	loanRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/loan/repository"
	loanUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/loan/usecase"

	reservationHttp "github.com/Leli2004/API_Go_biblioteca/internal/api/reservation/delivery/http"
	reservationRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/reservation/repository"
	reservationUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/reservation/usecase"

	fineRepository "github.com/Leli2004/API_Go_biblioteca/internal/api/fine/repository"
	fineUseCase "github.com/Leli2004/API_Go_biblioteca/internal/api/fine/usecase"
	"github.com/Leli2004/API_Go_biblioteca/internal/worker"
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
	authorRepo := authorRepository.NewRepository()
	authorUC := authorUseCase.NewUseCase(dbSqlx, authorRepo)
	authorhandler := authorHttp.NewHandler(authorUC)
	authorHttp.MapRoutes(e, authorhandler)

	// Genre
	genreRepo := genreRepository.NewRepository()
	genreUC := genreUseCase.NewUseCase(dbSqlx, genreRepo)
	genreHandler := genreHttp.NewHandler(genreUC)
	genreHttp.MapRoutes(e, genreHandler)

	// Publisher
	publisherRepo := publisherRepository.NewRepository()
	publisherUC := publisherUseCase.NewUseCase(dbSqlx, publisherRepo)
	publisherHandler := publisherHttp.NewHandler(publisherUC)
	publisherHttp.MapRoutes(e, publisherHandler)

	// Book
	bookRepo := bookRepository.NewRepository()
	bookUC := bookUseCase.NewUseCase(dbSqlx, bookRepo)
	bookHandler := bookHttp.NewHandler(bookUC)
	bookHttp.MapRoutes(e, bookHandler)

	// Book Copie
	bookCopieRepo := bookCopieRepository.NewRepository()
	bookCopieUC := bookCopieUseCase.NewUseCase(dbSqlx, bookCopieRepo)
	bookCopieHandler := bookCopieHttp.NewHandler(bookCopieUC)
	bookCopieHttp.MapRoutes(e, bookCopieHandler)

	// User
	userRepo := userRepository.NewRepository()
	userUC := userUseCase.NewUseCase(dbSqlx, userRepo)
	userHandler := userHttp.NewHandler(userUC)
	userHttp.MapRoutes(e, userHandler)

	// Loan
	loanRepo := loanRepository.NewRepository()
	loanUC := loanUseCase.NewUseCase(dbSqlx, loanRepo)
	loanHandler := loanHttp.NewHandler(loanUC)
	loanHttp.MapRoutes(e, loanHandler)

	// Reservation
	reservationRepo := reservationRepository.NewRepository()
	reservationUC := reservationUseCase.NewUseCase(dbSqlx, reservationRepo)
	reservationHandler := reservationHttp.NewHandler(reservationUC)
	reservationHttp.MapRoutes(e, reservationHandler)

	// Fine checker worker
	fineRepo := fineRepository.NewRepository()
	fineUC := fineUseCase.NewUseCase(dbSqlx, fineRepo)
	fineChecker := worker.NewFineChecker(fineUC)
	
	workerCtx, cancelWorker := context.WithCancel(context.Background())
	defer cancelWorker()
	go fineChecker.Run(workerCtx)

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- e.Start(fmt.Sprintf(":%s", config.GetServerPort()))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case <-quit:
	case err := <-serverErr:
		if err != nil {
			e.Logger.Error(err)
		}
	}

	cancelWorker()
	if err := e.Shutdown(context.Background()); err != nil {
		e.Logger.Error(err)
	}
}
