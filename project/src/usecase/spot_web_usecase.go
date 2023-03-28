package usecase

import (
	"app/domain/entity"
	"app/domain/repository"
	"app/infra/googlemap"
	"app/usecase/port"
	"app/utility"
	. "app/utility/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SpotWebUsecase struct {
	spotRepository      repository.SpotRepository
	googlemapRepository repository.GoogleMapRepository
	spotOutputPort      port.SpotOutputPort
}

func NewSpotWebUsecase(fr repository.SpotRepository, gr repository.GoogleMapRepository, fo port.SpotOutputPort) *SpotWebUsecase {
	return &SpotWebUsecase{
		spotRepository:      fr,
		googlemapRepository: gr,
		spotOutputPort:      fo,
	}
}

func (fu *SpotWebUsecase) GetAllSpots(ctx context.Context) (*port.BaseJsonMultipleResponse, error) {
	lastSpotId := ""
	if ctx.Value("lastSpotId") != "" {
		lastSpotId = ctx.Value("lastSpotId").(string)
	}
	spots, err := fu.spotRepository.FindAll(ctx, lastSpotId)
	if err != nil {
		return nil, err
	}
	return fu.spotOutputPort.JsonMultipleResponse(spots, http.StatusText(200)), nil
}

func (fu *SpotWebUsecase) GetSpotDataFromGoogleMap(ctx context.Context, searchText string) error {
	resultList := []googlemap.TextSearchResult{}
	pageToken := ""
	pageTokenPointer := &pageToken
	pageNum := 1
	for {
		if pageNum > 5 {
			break
		}
		textSearchRes, _ := fu.googlemapRepository.GetPlacesByText(ctx, pageToken, searchText)
		results := textSearchRes
		for _, v := range results.Results {
			resultList = append(resultList, v)
		}
		if results.NextPageToken != "" {
			*pageTokenPointer = results.NextPageToken
		} else {
			Log.Infof("Page End")
			break
		}
		Log.Infof("Page number: %v", pageNum)
		pageNum++
		time.Sleep(time.Millisecond * 2000)
	}
	if len(resultList) > 0 {
		Log.Debugf("Search Result length: %v", len(resultList))
		for _, v := range resultList {
			getByPlaceIdRes, _ := fu.spotRepository.FindByGooglePlaceID(ctx, v.PlaceID)
			docId := fu.spotRepository.NewSpotId()
			if len(getByPlaceIdRes) == 0 {
				var reviewList []entity.Review
				var photoList []entity.Photo
				placeDetailRes, _ := fu.googlemapRepository.GetPlaceDetailsByPlaceId(ctx, v.PlaceID)
				for _, w := range placeDetailRes.Result.Reviews {
					saveTime := time.Unix(int64(w.Time), 0)
					review := entity.Review{
						Rating:    &w.Rating,
						Comment:   &w.Text,
						CreatedAt: &saveTime,
						UpdatedAt: &saveTime,
					}
					reviewList = append(reviewList, review)
				}

				Log.Infof("Name: %s, Place ID: %s, type: %s, Rating: %f, Location: {%f, %f}",
					v.Name, v.PlaceID, v.Types, placeDetailRes.Result.Rating, v.Geometry.Location.Lat, v.Geometry.Location.Lng)
				spotProvider := "operator"
				spot := entity.Spot{
					Id:       docId,
					Name:     &v.Name,
					Provider: &spotProvider,
					Location: &entity.Location{
						Lat: v.Geometry.Location.Lat,
						Lng: v.Geometry.Location.Lng,
					},
					GoogleMap: &entity.GoogleMap{
						PlaceId: &v.PlaceID,
						Types:   placeDetailRes.Result.Types,
						Reviews: reviewList,
					},
					Address:   utility.FormatAddressFromGoogleMap(placeDetailRes.Result.AddressComponents),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				err := fu.spotRepository.InsertOne(ctx, &spot)
				if err != nil {
					Log.Error(err.Error())
					return err
				}

				for _, w := range placeDetailRes.Result.Photos {
					googlePhotoUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?photo_reference=%s&key=%s&maxwidth=1000", w.PhotoReference, os.Getenv("GOOGLE_MAPS_API_KEY"))
					imageData, err := utility.GetImageFromWeb(googlePhotoUrl)
					if err != nil {
						Log.Error(err.Error())
						return err
					}
					photoProvider := "googlemap"
					storagePath, photoUrl, err := fu.spotRepository.UploadImage(ctx, imageData, photoProvider, spot.Id, w.PhotoReference)
					if err != nil {
						Log.Error(err.Error())
						return err
					}
					photo := entity.Photo{
						OriginUrl: &googlePhotoUrl,
						Url:       &photoUrl,
						Path:      &storagePath,
						Provider:  &photoProvider,
					}
					photoList = append(photoList, photo)
				}
				spot.GoogleMap.Photos = photoList
				err = fu.spotRepository.Update(ctx, spot.Id, &spot)
				if err != nil {
					Log.Error(err.Error())
					return err
				}
			} else {
				spotPlaceId := getByPlaceIdRes[0].GoogleMap.PlaceId
				Log.Infof("GoogleMap place id %s has already existed.", spotPlaceId)
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
	defer fu.spotRepository.Closer()
	return nil
}
