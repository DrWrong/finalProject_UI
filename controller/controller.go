package controller

import (
	// "fmt"
	"github.com/DrWrong/finalProject_UI/thrift_interface"
	"github.com/DrWrong/finalProject_UI/utils"
	"github.com/Unknwon/macaron"
)

//Request plain weather info text
// it get the city name from an GET parameter "q" and make a RPC to the weather info service.
// then it return the info string
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

// Request encrypted weather info text
//  it get the city name from an GET parameter "q" and make a RPC to the weather info service.
//  then it invoke the security service to encrypt the text
//  then it return the encrypted text
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

// Request to decrypt an encrypted string
// it get the cipher text and decrypt key from GET parameters and then invoke the security service to decrypt the text
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
