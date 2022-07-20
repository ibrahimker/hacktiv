package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/ibrahimker/latihan-register/handler"
	"github.com/stretchr/testify/require"
)

func TestGenerateToJson(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		go handler.GenerateToJson()
		time.Sleep(500 * time.Millisecond)
		// read from json file and write to webData
		path := "../static/weather.json"
		require.FileExists(t, path)

		file, err := ioutil.ReadFile(path)
		require.NoError(t, err)
		require.NotNil(t, file)
		webData := entity.WebData{}
		json.Unmarshal(file, &webData)
		require.NotEmpty(t, webData.Status)
		t.Log(webData)
	})
}

//func TestGenerateToJson1(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			handler.GenerateToJson()
//		})
//	}
//}
