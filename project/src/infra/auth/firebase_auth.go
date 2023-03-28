package auth

import (
	"context"

	. "app/utility/logger"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	Client *auth.Client
}

func NewFirebaseAuth(ctx context.Context) *FirebaseAuth {
	opt := option.WithCredentialsFile("secretkey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		Log.Errorf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		Log.Errorf("error getting Auth client: %v\n", err)
	}
	return &FirebaseAuth{
		Client: client,
	}
}

func (fa *FirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := fa.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
