/*****************************************************************************
*
*	File			: main.go
*
* 	Created			: 20 Aug 2024
*
*	Description		: Small Demo app used for docker multi stage docker blog.
*
*	Modified		: 20 Aug 2024
*
*	Git				:
*
*	Author			: George Leonard
*
*	Copyright © 2021: George Leonard georgelza@gmail.com aka georgelza on Discord and Mongo Community Forum
*
*
*****************************************************************************/

package main

import (
	"os"
	"time"

	glog "google.golang.org/grpc/grpclog"
)

var (
	grpcLog glog.LoggerV2
)

func init() {

	// Keeping it very simple
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func runApp() {

	grpcLog.Infoln("###############################################################")
	grpcLog.Infoln("#")
	grpcLog.Infoln("#   Project   : Docker multi stage blog")
	grpcLog.Infoln("#")
	grpcLog.Infoln("#   Comment   : Small Demo app used for docker multistage build blog demo.")
	grpcLog.Infoln("#")
	grpcLog.Infoln("#   By        : George Leonard (georgelza@gmail.com)")
	grpcLog.Infoln("#")
	grpcLog.Infoln("#   Date/Time :", time.Now().Format("2006-01-02 15:04:05"))
	grpcLog.Infoln("#")
	grpcLog.Infoln("###############################################################")
	grpcLog.Infoln("")
	grpcLog.Infoln("# 	Username:", os.Getenv("PROJECT_USERNAME"))
	grpcLog.Infoln("# 	Username:", os.Getenv("PROJECT_PASSWORD"))

}

func main() {

	grpcLog.Infoln("****** Starting           *****")

	runApp()

	grpcLog.Infoln("****** Completed          *****")

}
