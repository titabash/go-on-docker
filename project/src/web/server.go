package web

import (
	"app/adapter/controller"
	"app/adapter/gateway"
	"app/adapter/presenter"
	"app/infra/auth"
	"app/usecase"
	. "app/utility/logger"
	openapi "app/web/go"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	jsonPresenter := presenter.NewBaseJsonPresenter()
	list := []interface{}{
		"a",
	}
	jsonPresenter.JsonMultipleResponse(list, http.StatusText(200))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		query := r.URL.Query()
		s := fmt.Sprintf("Query param → name: %s\n", query["name"])
		io.WriteString(w, s)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
			// Authorization ヘッダのフォーマットが適切でない
		}
		idToken := strings.Split(r.Header.Get("Authorization"), " ")[1]
		firebaseAuth := auth.NewFirebaseAuth(r.Context())
		userToken, err := firebaseAuth.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			// 渡されてきたトークンが適切ではない
		}
		Log.Debug(userToken.UID)
		body := r.Body
		defer body.Close()
		buf := new(bytes.Buffer)
		io.Copy(buf, body)
		var bodyMap map[string]interface{}
		json.Unmarshal(buf.Bytes(), &bodyMap)
		s := fmt.Sprintf("Body param → name: %s\n", bodyMap["name"])
		jsonPresenter := presenter.NewBaseJsonPresenter()
		jsonPresenter.JsonResponse(s, http.StatusText(200))
	}
}

func handler3(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx := context.Background()
		query := r.URL.Query()
		Log.Debugf("last_spot_id: %s", query.Get("last_spot_id"))
		ctx = context.WithValue(ctx, "lastSpotId", query.Get("last_spot_id"))
		spotRepository := gateway.NewSpotGateway(ctx)
		googlemapRepository := gateway.NewGoogleMapGateway()
		spotOutputPort := presenter.NewBaseJsonPresenter()
		uc := usecase.NewSpotWebUsecase(spotRepository, googlemapRepository, spotOutputPort)
		c := controller.NewSpotContoller(uc)
		c.GetAllSpots(ctx)
	}
}

func Server() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handler)
	// mux.HandleFunc("/foo", handler1)
	// mux.HandleFunc("/bar", handler2)
	// mux.HandleFunc("/spots", handler3)
	// c := cors.Default()
	// handler := c.Handler(mux)

	Log.Info("Golang Web server initilizing........ ")

	DefaultApiService := NewHandler()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	portNumber := 8080
	Log.Infof("Port number: %v", portNumber)
	// err := http.ListenAndServe(":"+strconv.Itoa(portNumber), handler)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	Log.Fatalf(err.Error())
}
