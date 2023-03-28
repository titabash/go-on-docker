package utility

import (
	. "app/utility/logger"
	"encoding/json"
	r "reflect"
)

func StructToMap(val interface{}) (mapVal map[string]interface{}, ok bool) {

	structVal := r.Indirect(r.ValueOf(val))
	typ := structVal.Type()

	mapVal = make(map[string]interface{})

	for i := 0; i < typ.NumField(); i++ {
		field := structVal.Field(i)

		if field.CanSet() {
			mapVal[typ.Field(i).Name] = field.Interface()
		}
	}

	return
}

func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
	byte, err := json.Marshal(mapVal)
	if err != nil {
		Log.Errorf("JSON marshal error: ", err.Error())
		return
	}

	if err := json.Unmarshal(byte, val); err != nil {
		Log.Errorf("JSON unmarshal error: ", err.Error())
		return
	}

	return
}
