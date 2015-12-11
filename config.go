package main

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/migore/paypal"
)

var (
	ProductionAppID = "api-getunseen"
	StagingAppID    = "staging-api-getunseen"
)


func ProductionConfig() Config {
	config := Config {
		BaseURL: "https://apps.facebook.com/helpmehelp",
		PaypalClientId: "ATRub8NK5m1iZV1EFPcs2Ad_lcKx6A7yasQaRSj6wdjKEDPBpzZ1UZBUr4qQtxg45fG-zO8OlZ85fJx4",
		PaypalSecret: "EF4fNq7M9l_VztubdFCLsTsUnqGAoSj12WTnGWuguyQKisAC2aneCVNuXDAusmwE5EjDit67YYTMev3z",
		PaypalBase: paypal.APIBaseLive,
		ExperienceProfileId: "XP-H5BE-78MM-5XMU-LZDX",
	}
	return config
}

func StagingConfig() Config {
	config := Config {
		BaseURL: "https://apps.facebook.com/bearchcanvas",
		PaypalClientId: "AUGtRDBDZek5V-TWQZ4GCALZNfRTbObh5UjxVthXScB90X9W3iDrez2VEVZSFG4qFKDfMsnqPmx7tBze",
		PaypalSecret: "EKLTvvNjEHZHvcrH2vmdMjNBHg4BO_8S4YBr2MFMSCfFFy9rz-TdFvk9lMe595Xd-y1UMJErjudYhiRP",
		PaypalBase: paypal.APIBaseSandBox,
		ExperienceProfileId: "XP-3L6B-V2T3-RGFL-2JBZ",
	}
	return config
}

func DevConfig() Config {
	config := Config {
		BaseURL: "http://localhost:8080",
		PaypalClientId: "AUGtRDBDZek5V-TWQZ4GCALZNfRTbObh5UjxVthXScB90X9W3iDrez2VEVZSFG4qFKDfMsnqPmx7tBze",
		PaypalSecret: "EKLTvvNjEHZHvcrH2vmdMjNBHg4BO_8S4YBr2MFMSCfFFy9rz-TdFvk9lMe595Xd-y1UMJErjudYhiRP",
		PaypalBase: paypal.APIBaseSandBox,
		ExperienceProfileId: "XP-3L6B-V2T3-RGFL-2JBZ",

	}

	return config
}

type Config struct {
	BaseURL             string
	PaypalClientId      string
	PaypalSecret        string
	PaypalBase          string
	ExperienceProfileId string

}

func NewConfig(context context.Context) Config {
	appID := appengine.AppID(context)
	log.Infof(context, "AppID: %v", appID)

	if appengine.IsDevAppServer() {
		return DevConfig()
	}


	if appID == ProductionAppID {
		return ProductionConfig()
	}

	if appID == StagingAppID {
		return StagingConfig()
	}

	panic("Could not resolve environment configuration")
}
