package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hardstylez72/cry-pay/api"
	"github.com/hardstylez72/cry-pay/internal/account"
	"github.com/hardstylez72/cry-pay/internal/arbitrum"
	"github.com/hardstylez72/cry-pay/internal/config"
	task "github.com/hardstylez72/cry-pay/internal/importers"
	"github.com/hardstylez72/cry-pay/internal/order"
	"github.com/hardstylez72/cry-pay/internal/pg"
	"github.com/hardstylez72/cry-pay/processor"
	"github.com/hardstylez72/cry-pay/processor/history"
	v1 "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	pgConn, err := pg.NewPGConnection(cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pg.WrapPgConnWithSqlx(pgConn)
	if err != nil {
		log.Fatal(err)
	}

	if err := pg.RunMigrations(pgConn, cfg.MigrationsDir, false); err != nil {
		log.Fatal(err)
	}

	accountsRepo := account.New(conn)
	ordersRepo := order.NewRepository(conn)
	historyRepo := history.New(conn)
	offsetsRepo := task.New(conn)

	procService, err := processor.New(&processor.InitConfig{
		HistoryRepo:  historyRepo,
		OrdersRepo:   ordersRepo,
		AccountsRepo: accountsRepo,
	})

	ordersService, err := order.NewService(&order.InitConfig{
		OrdersRepo:   ordersRepo,
		AccountsRepo: accountsRepo,
	})

	if !cfg.Standalone {
		arbiscanClient := arbitrum.New(&arbitrum.Config{
			Token: cfg.ArbiscanToken,
		})

		arbiscanImporterTask := task.NewArbiscanImporter(arbiscanClient, procService, offsetsRepo)
		go arbiscanImporterTask.RunImportTask(ctx, time.Second*10)

		processingTask := processor.NewProcessingTask(procService)
		go processingTask.RunTask(ctx, time.Second*10)
	}

	server := api.NewServer(&api.InitConfig{
		GrpcPort: cfg.GRPCAddr,
	}, &api.Callbacks{
		GrpcRegister: func(s grpc.ServiceRegistrar) {
			v1.RegisterFundsServiceServer(s, account.NewControllerAdapter(accountsRepo))
			v1.RegisterOrdersServiceServer(s, order.NewControllerAdapter(ordersService))
		},
	})
	server.Listen()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
}
