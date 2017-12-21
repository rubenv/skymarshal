package routes

import "github.com/tedsuo/rata"

const (
	OAuthBegin      = "OAuthBegin"
	OAuthCallback   = "OAuthCallback"
	OAuthV1Begin    = "OAuthV1Begin"
	OAuthV1Callback = "OAuthV1Callback"
	LogOut          = "LogOut"
)

var OAuthRoutes = rata.Routes{
	{Path: "/oauth/logout", Method: "GET", Name: LogOut},
	{Path: "/oauth/:provider/callback", Method: "GET", Name: OAuthCallback},
	{Path: "/oauth/:provider", Method: "GET", Name: OAuthBegin},
}

var OAuthV1Routes = rata.Routes{
	{Path: "/oauth/v1/:provider/callback", Method: "GET", Name: OAuthV1Callback},
	{Path: "/oauth/v1/:provider", Method: "GET", Name: OAuthV1Begin},
}
