package auth

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

const (
	TokenType = "TOKEN"
)

func AuthFlow(authtype string) (token string, err error) {
	if authtype == TokenType {
		err = survey.AskOne(&survey.Input{
			Message: "Insert your Square Cloud API token",
		}, &token)

		return
	} else {
		return "", errors.New("invalid auth type")
	}
}

// func createHttpServer() error {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
// 		http.Redirect(w, r, "https://squarecloud.app/auth/cli/success", http.StatusTemporaryRedirect)
// 	})

// 	server := &http.Server{
// 		Addr:    ":3579",
// 		Handler: mux,
// 	}

// 	return server.ListenAndServe()
// }
