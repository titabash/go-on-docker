package cmd

import (
	"app/adapter/controller"
	"app/adapter/gateway"
	"app/adapter/presenter"
	"app/usecase"
	"context"

	"github.com/spf13/cobra"
)

var FromGoogleMapCmd = &cobra.Command{
	Use:   "search-googlemap",
	Short: "Get spots data from google map",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		spotRepository := gateway.NewSpotGateway(ctx)
		googlemapRepository := gateway.NewGoogleMapGateway()
		spotOutputPort := presenter.NewBaseJsonPresenter()
		uc := usecase.NewSpotWebUsecase(spotRepository, googlemapRepository, spotOutputPort)
		c := controller.NewSpotContoller(uc)
		c.GetSpotDataFromGoogleMap(ctx, "東京 杉並区 子育て")
		return
	},
}

func init() {
	rootCmd.AddCommand(FromGoogleMapCmd)
}
