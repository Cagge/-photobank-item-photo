package resize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	. "photobank-item-photo/app/pkg/downloader/models"
	. "photobank-item-photo/app/pkg/resize/models"
	"reflect"
	"testing"
)

type mockHttpClient struct {
	execCounter       int
	possibleResponses []*http.Response
}

func (m *mockHttpClient) Do(request *http.Request) (*http.Response, error) {
	defer func() {
		m.execCounter++
	}()
	execCounter := m.execCounter
	if execCounter > 1 {
		execCounter = 1
	}
	return m.possibleResponses[m.execCounter], nil
}

func TestControlResize(t *testing.T) {
	type args struct {
		a    ConfigDataPhoto
		c    ConfigResize
		dhc  httpClientI
		dhc2 httpClientI
	}
	ex, _ := os.Getwd()
	exPath := filepath.Dir(ex)
	ref140byte := make([][]byte, 1)
	cancRead140 := fmt.Sprint(exPath, "/resize/testfile/referenc140.jpg")
	cancReadSRC := fmt.Sprint(exPath, "/resize/testfile/referencSRC.jpg")
	ref140byte[0], _ = os.ReadFile(cancRead140)
	refSRC, _ := os.Open(cancReadSRC)
	f := Data{
		Count: 1,
		Photos: []Phot{
			{
				0, 6537837, "regular",
			},
		},
	}
	byteStruct, _ := json.Marshal(f)
	refSRCbyte, _ := ioutil.ReadAll(refSRC)
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "BLA",
			args: args{
				a: ConfigDataPhoto{PHotodate{"***", "***", "https://reqres.in/api/users", "http://photobank.sima-land.ru:8080/static/items/"}},
				c: ConfigResize{REsize{[]uint{140}, "/tmp/"}},
				dhc: &mockHttpClient{0,
					[]*http.Response{
						{Body: ioutil.NopCloser(bytes.NewReader(byteStruct))},
						{Body: ioutil.NopCloser(bytes.NewReader(refSRCbyte))},
					},
				},
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ControlResize(tt.args.a, tt.args.c, tt.args.dhc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ControlResize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ControlResize() got = %v, want %v", got, tt.want) //count img
			}

			if !reflect.DeepEqual(got1[0], ref140byte[0]) {
				t.Errorf("error out img !equal refer") //byte
			}

		})
	}
}
