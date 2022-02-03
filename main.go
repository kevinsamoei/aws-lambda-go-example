package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
)

var (
	logruslogger  = logrus.New()
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	flag.Set("v", "0")
	flag.Set("logtostderr", "true")
	flag.Parse()

}
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	fmt.Println("fmt.Println log line")
	log.Println("log.Println log line")
	logruslogger.Trace("logruslogger logs in  trace")
	logruslogger.Debug("logruslogger logs in debug ")
	logruslogger.Info("logruslogger logs in  info")
	logruslogger.Warn("logruslogger logs in warn")
	logruslogger.Error("logruslogger logs in  error")

	InfoLogger.Println("InfoLogger text about info logs")
	WarningLogger.Println("InfoLogger warn logs")
	ErrorLogger.Println("InfoLogger error log")

	glog.Info("glog logs in  info")
	glog.Warning("glog logs in warn")
	glog.Error("glog logs in  error")
	glog.Flush()

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/plain"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              "Hello from Golang",
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
