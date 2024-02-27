package cli

import (
	"com.tuntun.rangers/service/chaindata/src/middleware/log"
	"com.tuntun.rangers/service/chaindata/src/middleware/mysql"
	"fmt"
	"os"
	"runtime"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	version = "0.0.1"

	instanceSection = "instance"

	indexKey = "index"
)

type GX struct {
	init bool
}

func NewGX() *GX {
	return &GX{}
}

func (gx *GX) Run() {
	// control+c 信号
	ctrlC := signals()
	quitChan := make(chan bool)
	go gx.handleExit(ctrlC, quitChan)

	app := kingpin.New("chaindata", "blockchain data service")
	app.HelpFlag.Short('h')

	//版本号
	versionCmd := app.Command("version", "show Rangers Service version")

	// mine
	startCmd := app.Command("start", "service start")
	//mysqlAddr := startCmd.Flag("mysql", "the mysql addr").String()

	command, err := app.Parse(os.Args[1:])

	if err != nil {
		kingpin.Fatalf("%s, try --help", err)
	}

	switch command {
	case versionCmd.FullCommand():
		fmt.Println("chaindata Service Version:", version)
		os.Exit(0)

	case startCmd.FullCommand():
		go func() {
			runtime.SetBlockProfileRate(1)
			runtime.SetMutexProfileFraction(1)
		}()

		fmt.Println("Welcome to chaindata service")

		mysql.InitMySql()
	}

	<-quitChan
}

func (gx *GX) handleExit(ctrlC <-chan bool, quit chan<- bool) {
	<-ctrlC
	fmt.Println("exiting...")

	if gx.init {
		mysql.CloseMysql()
		log.Close()
		quit <- true
	} else {
		log.Close()
		os.Exit(0)
	}
}
