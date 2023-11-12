package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/burgercking/contributions-farmer/internal/config"
	"github.com/burgercking/contributions-farmer/internal/farmer"
	"github.com/burgercking/contributions-farmer/internal/vcs/github"
)

func initializeGithubFarmer(ctx context.Context) (farmer.Farmer, error) {
	cfg, err := config.New("./config/")
	if err != nil {
		return nil, err
	}

	vcs := github.New(
		github.WithAccessToken(cfg.AccessToken),
		github.WithClassicCredentials(cfg.ClassicToken, cfg.UserName, cfg.UserEmail),
		github.WithPath(cfg.ReposPath, cfg.TargetRepo, cfg.FileName, cfg.RepositoryPrefix),
	)

	return farmer.New(
		farmer.WithCommonOptions(vcs, cfg.TargetRepo, 16, cfg.StartDate, cfg.EndDate),
	)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	f, err := initializeGithubFarmer(ctx)
	if err != nil {
		log.Fatalf("%v: cannot initialize github farmer\n", err)
		os.Exit(1)
	}

	fmt.Println("starting app...")
	err = f.Run(ctx)
	if err != nil {
		log.Fatal("failed to start farmer")
	}
}
