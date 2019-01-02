package routing_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"big-life-backend/trip/domain"
	"big-life-backend/trip/mock_service"
	"big-life-backend/trip/routing"
)

func setup(t *testing.T) (*routing.Server, *mock_service.MockTripServiceInt, *httptest.ResponseRecorder) {
	mockcontroller := gomock.NewController(t)
	defer mockcontroller.Finish()
	mockTripService := mock_service.NewMockTripServiceInt(mockcontroller)
	server := routing.NewServer(mockTripService, mux.NewRouter())
	responseRecorder := httptest.NewRecorder()
	server.Routes()

	return server, mockTripService, responseRecorder
}

func TestGetTrip(t *testing.T) {
	server, mockTripService, responseRecorder := setup(t)

	req, _ := http.NewRequest("GET", "/trip/1", nil)
	mockTripService.EXPECT().GetTrip(1).Times(1)

	server.Router.ServeHTTP(responseRecorder, req)

}

func TestCreateTrip(t *testing.T) {
	server, mockTripService, responseRecorder := setup(t)
	trip := `{
	"origin" : "India",
	"destination": "New Jersey"
	}`
	req, _ := http.NewRequest("POST", "/trip/3", strings.NewReader(trip))

	mockTripService.EXPECT().CreateTrip(domain.Trip{Id: 3, Origin: "India", Destination: "New Jersey"}).Times(1)

	server.Router.ServeHTTP(responseRecorder, req)

}

func TestServer_UpdateTripHandler(t *testing.T) {
	server, mockTripService, responseRecorder := setup(t)
	trip := `{
	"origin" : "India",
	"destination": "New Jersey"
	}`
	req, _ := http.NewRequest("PUT", "/trip/2", strings.NewReader(trip))

	mockTripService.EXPECT().UpdateTrip(2, domain.Trip{Origin: "India", Destination: "New Jersey"}).Times(1)

	server.Router.ServeHTTP(responseRecorder, req)

}

func TestServer_DeleteHandler(t *testing.T) {

	server, mockTripService, responseRecorder := setup(t)
	req, _ := http.NewRequest("DELETE", "/trip/1", nil)

	mockTripService.EXPECT().DeleteTrip(1).Times(1)

	server.Router.ServeHTTP(responseRecorder, req)

}
