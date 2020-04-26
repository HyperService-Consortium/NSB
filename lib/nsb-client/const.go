package nsbcli

const (
	httpPrefix   = "http://"
	httpsPrefix  = "https://"
	maxBytesSize = 64 * 1024
)

var (
	GlobalClient = NewNSBClient("")
)
