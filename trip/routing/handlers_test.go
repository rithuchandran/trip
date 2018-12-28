package routing_test

import (
	"big-life-backend/trip/domain"
	"big-life-backend/trip/mock_service"
	"big-life-backend/trip/routing"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setup(t *testing.T) (*routing.Server, *mock_service.MockTripServiceInt, *httptest.ResponseRecorder) {
	mockcontroller := gomock.NewController(t)
	defer mockcontroller.Finish()
	sr := mock_service.NewMockTripServiceInt(mockcontroller)
	s := routing.NewServer(sr, mux.NewRouter())
	rr := httptest.NewRecorder()
	s.Routes()

	return s, sr, rr
}

func TestGetTrip(t *testing.T) {
	s, sr, rr := setup(t)

	req, _ := http.NewRequest("GET", "/trip/1", nil)
	sr.EXPECT().GetTrip(1).Times(1)

	s.Router.ServeHTTP(rr, req)

}

func TestCreateTrip(t *testing.T) {
	s, sr, rr := setup(t)
	trip := `{
	"origin" : "India",
	"destination": "New Jersey"
	}`
	req, _ := http.NewRequest("POST", "/trip/3", strings.NewReader(trip))

	sr.EXPECT().CreateTrip(domain.Trip{Id: 3, Origin: "India", Destination: "New Jersey"}).Times(1)

	s.Router.ServeHTTP(rr, req)

}

func TestServer_UpdateTripHandler(t *testing.T) {
	s, sr, rr := setup(t)
	trip := `{
	"origin" : "India",
	"destination": "New Jersey"
	}`
	req, _ := http.NewRequest("PUT", "/trip/2", strings.NewReader(trip))

	sr.EXPECT().UpdateTrip(2, domain.Trip{Origin: "India", Destination: "New Jersey"}).Times(1)

	s.Router.ServeHTTP(rr, req)

}

func TestServer_DeleteHandler(t *testing.T) {

	s, sr, rr := setup(t)
	req, _ := http.NewRequest("DELETE", "/trip/1", nil)

	sr.EXPECT().DeleteTrip(1).Times(1)

	s.Router.ServeHTTP(rr, req)

}
