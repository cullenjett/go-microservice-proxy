package main

// Microservice holds configuration data for a microservice
type Microservice struct {
	PublicPath string
	URL        string
}

// GetMicroservices returns the list of microservices to be routed to.
// Eventually this might be a good place to fetch external configuration data...
func GetMicroservices() []Microservice {
	return []Microservice{
		{
			PublicPath: "/cars",
			URL:        "https://localhost:3000",
		},
		{
			PublicPath: "/acccount",
			URL:        "https://localhost:3001",
		},
		{
			PublicPath: "/refinance",
			URL:        "https://localhost:3001",
		},
		{
			PublicPath: "/listings",
			URL:        "https://localhost:3001",
		},
		{
			PublicPath: "/offers",
			URL:        "https://localhost:3001",
		},
		{
			PublicPath: "/prequal",
			URL:        "https://localhost:3002",
		},
		{
			PublicPath: "/message-center",
			URL:        "https://localhost:3003",
		},
	}
}
