package conf

// CommonConfig is set of common configurations
type CommonConfig struct { // nolint
	APIEndpoint    *string
	APIKey         *string
	AppVersion     string
	ExtendedOutput *bool
	IsDebugMode    bool
}

// StartConfig is set of configurations for starting the process
type StartConfig struct {
	Common       *CommonConfig
	Platform     *string
	Asynchronous *bool
}

// SuccessConfig is set of configurations for ending the process with success state
type SuccessConfig struct {
	Common  *CommonConfig
	Message *string
}
