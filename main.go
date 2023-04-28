package main

import (
	"github.com/TykTechnologies/tyk/gateway"
)

// @Version 5.0.1
// @Title Tyk Gateway API
// @Description Rest API For Tyk Gateway!
// @ContactName Person Example
// @ContactEmail person@example.com
// @ContactURL http://www.example.com
// @TermsOfServiceUrl http://www.example.com
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Server http://localhost:5000 Local
// @Server http://www.example-dev.com Development
// @Server http://www.example.com Production
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader http bearer Input your token
func main() {

	gateway.Start()
}
