package utils

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	log "github.com/Sirupsen/logrus"
	"time"
	"ui/thrift_interface"
)

var (
	weatherClient  *thrift_interface.CityWeatherInfoServiceClient
	securityClient *thrift_interface.SecureServiceClient
)

func GetWeatherClient() (
	*thrift_interface.CityWeatherInfoServiceClient, error) {
	var (
		err       error
		res       bool
		transport thrift.TTransport
	)
	if weatherClient != nil {
		res, err = weatherClient.Ping(RequestHeader())
		if err != nil && res {
			return weatherClient, nil
		}
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err = thrift.NewTSocket(IniConf.String("weather_server"))
	if err != nil {
		log.Errorf("failed to open socket to weather server %s", err)
		return nil, err
	}
	transport = transportFactory.GetTransport(transport)
	// defer transport.Close()
	if err := transport.Open(); err != nil {
		log.Errorf("faied to open socket to weather server %s", err)
		return nil, err
	}
	weatherClient = thrift_interface.NewCityWeatherInfoServiceClientFactory(
		transport, protocolFactory)
	return weatherClient, nil
}

func GetSecurityClient() (
	*thrift_interface.SecureServiceClient, error) {
	var (
		err       error
		res       bool
		transport thrift.TTransport
	)
	if securityClient != nil {
		res, err = securityClient.Ping(RequestHeader())
		if err != nil && res {
			return securityClient, nil
		}
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err = thrift.NewTSocket(IniConf.String("security_server"))
	if err != nil {
		log.Errorf("failed to open socket to secruity server %s", err)
		return nil, err
	}
	transport = transportFactory.GetTransport(transport)
	// defer transport.Close()
	if err := transport.Open(); err != nil {
		log.Errorf("faied to open socket to security server %s", err)
		return nil, err
	}
	securityClient = thrift_interface.NewSecureServiceClientFactory(transport, protocolFactory)
	return securityClient, nil
}

func RequestHeader() *thrift_interface.CommonRequest {
	t := int32(time.Now().Unix())
	return &thrift_interface.CommonRequest{
		Requester: "ui",
		RequestId: t}
}
