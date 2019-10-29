package godper

import (
	"reflect"
	"strconv"
	"time"
)

//Timetranfer 时间转换
func Timetranfer(ptr interface{}) {

	v := reflect.ValueOf(ptr).Elem()

	if v.Kind() != reflect.Struct {
		return
	}
	trans(v)
}

func trans(v reflect.Value) {

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			settrans(v.Type().Field(i), v.Field(i))
		}
		if v.Field(i).Kind() == reflect.Struct {
			trans(v.Field(i))
		}
	}

}

func settrans(f reflect.StructField, v reflect.Value) {
	tag := f.Tag
	timeFormate := tag.Get("godper")
	if timeFormate == "" {
		return
	}
	timestamp, _ := strconv.ParseInt(v.String(), 10, 0)
	if timestamp == 0 {
		return
	}
	res := time.Unix(timestamp, 0).Format(timeFormate)
	v.Set(reflect.ValueOf(res))
}
