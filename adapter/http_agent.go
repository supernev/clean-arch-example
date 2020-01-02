package adapter

// HTTPAgent - HTTP implementation for agent interface
type HTTPAgent struct {
	requestTrainHandler func(ballerID uint64, playerID uint64)
}

// GetIP - Get IP of client
func (mod *HTTPAgent) GetIP() string {
	return "0.0.0.0"
}

// GetOS - Get OS of client
func (mod *HTTPAgent) GetOS() string {
	return "Android"
}

// GetVersion - Get version of client
func (mod *HTTPAgent) GetVersion() string {
	return "1.0.1"
}

// OnRequestTrain - Register a handler for the request
func (mod *HTTPAgent) OnRequestTrain(fn func(ballerID uint64, playerID uint64)) {
	mod.requestTrainHandler = fn
}

// ResponseTrain -
func (mod *HTTPAgent) ResponseTrain(ResultCode uint32, OldLevel uint32, NewLevel uint32) {
	// Write and flush
}
