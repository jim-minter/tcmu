package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/coreos/go-tcmu"
)

type config struct {
	accountName   string
	accountKey    string
	useHTTPS      bool
	containerName string
	blobName      string
	size          int64
}

var cfg = config{
	accountName:   "",
	accountKey:    "",
	useHTTPS:      true,
	containerName: "",
	blobName:      "",
	size:          1 << 30,
}

func run() error {
	// logrus.SetLevel(logrus.DebugLevel)

	err := exec.Command("modprobe", "target_core_user").Run()
	if err != nil {
		return err
	}

	b, err := newBlob()
	if err != nil {
		return err
	}
	defer b.Close()

	handler := tcmu.BasicSCSIHandler(b)
	handler.DataSizes.VolumeSize = cfg.size
	d, err := tcmu.OpenTCMUDevice("/dev/tcmufile", handler)
	if err != nil {
		return err
	}
	defer d.Close()

	fmt.Println("attached")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
