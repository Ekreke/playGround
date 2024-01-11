package main

import "go.uber.org/zap"

func main() {
	mylog := zap.NewExample(zap.Fields(
		zap.Int("userId", 10),
		zap.String("username", "ekreke"),
	))
	mylog.Info("this is mylog")
	mylog.Debug("this is mylog2")

}
