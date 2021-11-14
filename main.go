package main

import (
	"github.com/leishi1313/Shifter/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	cmd.Execute()
}
