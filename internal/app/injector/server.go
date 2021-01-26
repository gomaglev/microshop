package injector

// // InitServerConfig for injector
// func InitServerConfig() rpc.ServerConfig {
// 	return rpc.ServerConfig{
// 		Addr:                 config.C.GRPC.Addr,
// 		Port:                 config.C.GRPC.Port,
// 		CertFile:             config.C.GRPC.CertFile,
// 		KeyFile:              config.C.GRPC.KeyFile,
// 		EnableLogrus:         config.C.Interceptor.EnableLogrus,
// 		EnableRateLimit:      config.C.Interceptor.EnableRateLimit,
// 		EnableRecovery:       config.C.Interceptor.EnableRecovery,
// 		EnableAuthentication: config.C.Interceptor.EnableAuthentication,
// 		EnableAuthorization:  config.C.Interceptor.EnableAuthorization,
// 		Authentic:            config.C.BasicAuth.AuthFunc(),
// 		Authorize:            config.C.Authorizer.Enforce(),
// 	}
// }

// // InitGatewayConfig for injector
// func InitGatewayConfig() rpc.GatewayConfig {
// 	return rpc.GatewayConfig{
// 		Host:       config.C.Gateway.Host,
// 		Port:       config.C.Gateway.Port,
// 		PathPrefix: config.C.Gateway.ApiUrl,
// 	}
// }
