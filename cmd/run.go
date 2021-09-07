package cmd

import (
	"sync"

	"github.com/pg-sharding/spqr/app"
	"github.com/pg-sharding/spqr/internal"
	"github.com/pg-sharding/spqr/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
)

var (
	dataFolder string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&dataFolder, "data", "d", "", "path to data folder")
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run SPQR",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Load(dataFolder); err != nil {
			return err
		}
		spqr, err := internal.NewSpqr(dataFolder)
		if err != nil {
			return errors.Wrap(err, "SPQR creation failed")
		}

		app := app.NewApp(spqr)

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			err := app.ProcPG()
			tracelog.ErrorLogger.FatalOnError(err)
			wg.Done()
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			err := app.ServHttp()
			tracelog.ErrorLogger.FatalOnError(err)
			wg.Done()
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			err := app.ProcADM()
			tracelog.ErrorLogger.FatalOnError(err)
			wg.Done()
		}(wg)

		wg.Wait()

		return nil
	},
}
