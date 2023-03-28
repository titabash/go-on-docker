package gateway

import (
	"app/domain/entity"
	"app/infra/db"
	"app/infra/storage"
	"app/utility"
	. "app/utility/logger"
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
)

type SpotGateway struct {
	collection        string
	dbConnection      *db.FirestoreConnection
	storageConnection *storage.FirebaseStorageConnection
}

func NewSpotGateway(ctx context.Context) *SpotGateway {
	return &SpotGateway{
		collection:        "spots",
		dbConnection:      db.NewFirestoreConnection(ctx),
		storageConnection: storage.NewFirebaseStorageConnection(ctx),
	}
}

func (fg *SpotGateway) NewSpotId() string {
	newDocId := fg.dbConnection.Client.Collection(fg.collection).NewDoc().ID
	return newDocId
}

func (fg *SpotGateway) FindByGooglePlaceID(ctx context.Context, googlePlaceId string) ([]*entity.Spot, error) {
	collection := fg.dbConnection.Client.Collection(fg.collection)
	iter := collection.Where("google_place_id", "==", googlePlaceId).Documents(ctx)
	docSnapshots, _ := iter.GetAll()

	var spots []*entity.Spot

	for _, v := range docSnapshots {
		var spot *entity.Spot
		utility.MapToStruct(v.Data(), &spot)
		spots = append(spots, spot)
	}

	return spots, nil
}

func (fg *SpotGateway) InsertOne(ctx context.Context, spot *entity.Spot) error {
	err := fg.dbConnection.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		query := fg.dbConnection.Client.Collection(fg.collection).Where("google_map.place_id", "==", spot.GoogleMap.PlaceId).Limit(1)
		docs, err := tx.Documents(query).GetAll()
		if err != nil {
			return err
		}
		if len(docs) > 0 {
			Log.Error("Error: unique field already exists")
			return errors.New("unique field already exists")
		}

		docRef := fg.dbConnection.Client.Collection(fg.collection).Doc(spot.Id)
		spotMap, _ := utility.StructToMap(spot)
		err = tx.Create(docRef, spotMap)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (fg *SpotGateway) FindByDocumentID(ctx context.Context, id string) (*entity.Spot, error) {
	collection := fg.dbConnection.Client.Collection(fg.collection)
	docRef := collection.Doc(id)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}
	var spot *entity.Spot
	utility.MapToStruct(docSnapshot.Data(), spot)
	return spot, nil
}

func (fg *SpotGateway) Update(ctx context.Context, id string, spot *entity.Spot) error {
	collection := fg.dbConnection.Client.Collection(fg.collection)
	docRef := collection.Doc(id)
	_, err := docRef.Set(ctx, spot)
	if err != nil {
		return err
	}
	return nil
}

func (fg *SpotGateway) FindAll(ctx context.Context, lastSpotId string) ([]*entity.Spot, error) {
	collection := fg.dbConnection.Client.Collection(fg.collection).OrderBy("created_at", firestore.Asc).Limit(20)
	if lastSpotId != "" {
		lastDoc, err := fg.dbConnection.Client.Collection(fg.collection).Doc(lastSpotId).Get(ctx)
		if err != nil {
			return nil, err
		}
		collection = collection.StartAfter(lastDoc)
	}
	iter := collection.Documents(ctx)
	docSnapshots, _ := iter.GetAll()

	var spots []*entity.Spot

	for _, v := range docSnapshots {
		var spot *entity.Spot
		utility.MapToStruct(v.Data(), &spot)
		spots = append(spots, spot)
	}

	return spots, nil
}

func (fg *SpotGateway) UploadImage(ctx context.Context, file []byte, provider string, spotId string, name string) (string, string, error) {
	connection := fg.storageConnection
	fileName := fmt.Sprintf("/spots/images/%s/%s/%s.png", provider, spotId, name)
	err := connection.UploadFile(ctx, file, fileName)
	if err != nil {
		Log.Error(err.Error())
		return "", "", err
	}
	mediaUrl, err := connection.GetMedialink(ctx, fileName)
	if err != nil {
		Log.Error(err.Error())
		return "", "", err
	}

	return fileName, mediaUrl, nil
}

func (fg *SpotGateway) FindBetweenLatLng(ctx context.Context, northEast *entity.Location, southWest *entity.Location) ([]*entity.Spot, error) {
	collection := fg.dbConnection.Client.Collection(fg.collection)
	collection.Where("location.lat", "<", northEast.Lat)
	collection.Where("location.lat", ">", southWest.Lat)
	collection.Where("location.lng", "<", northEast.Lng)
	collection.Where("location.lng", ">", southWest.Lng)
	iter := collection.Documents(ctx)
	docSnapshots, _ := iter.GetAll()

	var spots []*entity.Spot

	for _, v := range docSnapshots {
		var spot *entity.Spot
		utility.MapToStruct(v.Data(), &spot)
		spots = append(spots, spot)
	}

	return spots, nil
}

func (fg *SpotGateway) Closer() error {
	fg.dbConnection.Client.Close()
	return nil
}
