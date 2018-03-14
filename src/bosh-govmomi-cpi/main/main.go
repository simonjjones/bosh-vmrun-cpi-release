package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/cppforlife/bosh-cpi-go/rpc"

	"bosh-govmomi-cpi/action"
	"bosh-govmomi-cpi/config"
	"bosh-govmomi-cpi/govc"
	"bosh-govmomi-cpi/stemcell"
)

var (
	configPathOpt = flag.String("configPath", "", "Path to configuration file")
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // todo MAC generation

	logger, fs, compressor := basicDeps()

	defer logger.HandlePanic("Main")

	flag.Parse()
	cpiConfig, err := config.NewConfigFromPath(*configPathOpt, fs)
	if err != nil {
		logger.Error("main", "Loading cfg %s", err.Error())
		os.Exit(1)
	}

	govcClient := govc.NewClient(govc.NewGovcConfig(cpiConfig), logger)
	stemcellClient := stemcell.NewClient(compressor, fs, logger)
	cpiFactory := action.NewFactory(govcClient, stemcellClient, cpiConfig, fs, logger)

	cli := rpc.NewFactory(logger).NewCLI(cpiFactory)

	err = cli.ServeOnce()
	if err != nil {
		logger.Error("main", "Serving once: %s", err)
		os.Exit(1)
	}
}

func basicDeps() (boshlog.Logger, boshsys.FileSystem, boshcmd.Compressor) {
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr)
	fs := boshsys.NewOsFileSystem(logger)
	cmdRunner := boshsys.NewExecCmdRunner(logger)
	compressor := boshcmd.NewTarballCompressor(cmdRunner, fs)

	return logger, fs, compressor
}