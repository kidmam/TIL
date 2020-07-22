package globals_and_constants

// https://qvault.io/2019/10/21/golang-constant-maps-slices/

// this is a global constant
const safeRateLimit = 10

// this is a global variable
var dangerousRateLimit = 10

const rateLimit = 10

// this is meant to be constant! Please don't mutate it!
var supportedNetworks = []string{"facebook", "twitter", "instagram"}

func getSupportedNetworks() []string {
	return []string{"facebook", "twitter", "instagram"}
}

func getSupportedNetworkRatelimits() map[string]int {
	return map[string]int{
		"facebook":  10,
		"twitter":   50,
		"instagram": 60,
	}
}
