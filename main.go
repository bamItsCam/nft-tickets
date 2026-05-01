package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/bamItsCam/nft-tickets/internal/contract"
	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/bamItsCam/nft-tickets/internal/handlers"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	godotenv.Load()
	contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	if contractAddressHex == "" {
		slog.Error("missing CONTRACT_ADDRESS, please set it")
		return
	}
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		slog.Error("missing OWNER_KEY, please set it")
		return
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		slog.Error("missing DATABASE_URL, please set it")
		return
	}
	ethNodeUrl := os.Getenv("ETH_NODE_URL")
	if ethNodeUrl == "" {
		slog.Error("missing ETH_NODE_URL, please set it")
		return
	}
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		slog.Error("missing BASE_URL, please set it")
		return
	}

	contractAddress := common.HexToAddress(contractAddressHex)

	client, err := ethclient.Dial(ethNodeUrl)
	if err != nil {
		slog.Error("failed to dial eth", "error", err)
		return
	}
	defer client.Close()

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		slog.Error("failed to connect to postgres", "error", err)
		return
	}

	ticket := contract.NewTicket()
	instance := ticket.Instance(client, contractAddress)

	privateKey, err := crypto.HexToECDSA(ownerKey)
	if err != nil {
		slog.Error("failed to convert owner key", "error", err)
		return
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		slog.Error("failed to fetch the chainId", "error", err)
		return
	}

	app := handlers.App{
		Querier:  db.New(pool),
		Instance: instance,
		Client:   client,
		ChainId:  chainId,
		Opts:     bind.NewKeyedTransactor(privateKey, chainId),
		Ticket:   ticket,
		BaseUrl:  baseUrl,
	}

	slog.Info("setup successful", "owner_address", app.Opts.From)

	mux := http.NewServeMux()
	config := huma.DefaultConfig("NFT Tickets API", "1.0.0")
	api := humago.New(mux, config)

	handlers.RegisterEventRoutes(api, &app)
	handlers.RegisterUserRoutes(api, &app)
	handlers.RegisterTicketRoutes(api, &app)

	slog.Info("serving...")

	if err := http.ListenAndServe(":8080", logMiddleware(mux)); err != nil {
		slog.Error("server stopped", "error", err)
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "status", rw.status)
	})
}
