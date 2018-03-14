package handler

import (
	"github.com/aj-jaswanth/oauth/oauth"
	"log"
	"net/http"
	"strings"
)

var page = `
<html>
	<body>
		<a href="https://github.com/login/oauth/authorize?client_id=CLIENT_ID&scope=user">authorize</a>
	</body>
</html>
`

func Home(w http.ResponseWriter, req *http.Request) {
	log.Println("home page is requested")
	processedPage := strings.Replace(page, "CLIENT_ID", oauth.ClientId, 1)
	w.Write([]byte(processedPage))
}
