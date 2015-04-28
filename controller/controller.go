package controller

import (
	// "fmt"
	"github.com/Unknwon/macaron"
	"ui/thrift_interface"
	"ui/utils"
)

func Plain(ctx *macaron.Context) string {
	cityname := ctx.Query("q")
	weatherclient, err := utils.GetWeatherClient()
	if err != nil {
		return "The Weather info Service seems down"
	}
	request := &thrift_interface.CityWeatherInfoRequest{cityname}
	info, _ := weatherclient.GetCityWeatherInfo(
		utils.RequestHeader(), request)
	return info
}

func Encrypted(ctx *macaron.Context) string {
	var securityClient *thrift_interface.SecureServiceClient
	cityname := ctx.Query("q")
	weatherclient, err := utils.GetWeatherClient()
	if err != nil {
		return "The Weather info Service seems down"
	}
	request := &thrift_interface.CityWeatherInfoRequest{cityname}
	info, _ := weatherclient.GetCityWeatherInfo(
		utils.RequestHeader(), request)

	securityClient, err = utils.GetSecurityClient()
	if err != nil {
		return "The security service seems down"
	}
	info, _ = securityClient.Encrypted(
		utils.RequestHeader(), info)
	return info
}

func Decrypted(ctx *macaron.Context) string {
	cipher := ctx.Query("cipher")
	key := ctx.Query("key")
	securityClient, err := utils.GetSecurityClient()
	if err != nil {
		return "The security service seems down"
	}
	info, err := securityClient.Decrypted(
		utils.RequestHeader(), cipher, key)
	if err != nil {
		return err.Error()
	}
	return info
}

// func Decrypted(ctx)
