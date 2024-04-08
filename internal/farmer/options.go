package farmer

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/liseloge/contributions-farmer/internal/common"
	"github.com/liseloge/contributions-farmer/internal/vcs"
)

const maxConcurrency = 16

var defaultConcurrency = runtime.NumCPU()

type Config struct {
	vcs.VCSProvider
	Repo        string
	Concurrency int
	FirstDay    time.Time
	LastDay     time.Time
}

type OptionFn = common.Options[Config]

func WithCommonOptions(vcs vcs.VCSProvider, repo string, concurrency int, firstDay, lastDay string) OptionFn {
	start, err := time.Parse(time.RFC3339, firstDay)
	if err != nil {
		log.Fatalf("Cannot convert START_DATE: %v", err)
		os.Exit(1)
	}

	end, err := time.Parse(time.RFC3339, lastDay)
	if err != nil {
		log.Fatalf("Cannot convert END_DATE: %v", err)
		os.Exit(1)
	}

	log.Println(end.String())

	return func(o *Config) {
		o.VCSProvider = vcs
		o.Repo = repo
		o.Concurrency = concurrency
		o.FirstDay = start
		o.LastDay = end
	}
}
