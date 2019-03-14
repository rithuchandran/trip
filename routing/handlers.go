package routing

import (
	"cloud.google.com/go/logging"
	"encoding/json"
	"fmt"
	"github.com/rithuchandran/trip/domain"
	"github.com/rithuchandran/trip/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//noinspection GoUnhandledErrorResult
func (s Server) GetTripHandler(writer http.ResponseWriter, request *http.Request) {
	lg := logger.NewLogger()
	lg.Log(logging.Entry{Severity: logging.Info, Payload: "Inside create trip handler"})
	parameters := mux.Vars(request)
	id, _ := strconv.Atoi(parameters["id"])
	t, e := s.tripService.GetTrip(id)
	if e != nil {
		fmt.Fprint(writer, e)
		return
	}
	json.NewEncoder(writer).Encode(t)
}

//func (s Server) LogHandler(writer http.ResponseWriter, request *http.Request) {
//	ctx := context.Background()
//
//	// Sets your Google Cloud Platform project ID.
//	projectID := "turnkey-env-234412"
//
//	// Creates a client.
//	client, err := logging.NewClient(ctx, projectID)
//	if err != nil {
//		fmt.Printf("Failed to create client: %v", err)
//		os.Exit(2)
//	}
//	err = client.Ping(ctx)
//	if err != nil {
//		fmt.Printf("Failed ping the logging service: %v", err)
//		os.Exit(2)
//	}
//
//	const name = "stderr"
//	logger2 := client.Logger(name)
//	defer logger2.Flush() // Ensure the entry is written.
//
//	infolog := logger2.StandardLogger(logging.Info)
//	infolog.Printf("infolog is a standard Go log.Logger with INFO severity.")
//}

//noinspection GoUnhandledErrorResult
func (s Server) CreateTripHandler(writer http.ResponseWriter, request *http.Request) {
	lg := logger.NewLogger()
	lg.Log(logging.Entry{Severity: logging.Info, Payload: "Inside create trip handler"})
	var trip domain.Trip
	err := json.NewDecoder(request.Body).Decode(&trip)
	trip.Id, _ = strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		fmt.Fprint(writer, "invalid input!")
		return
	}
	s.tripService.CreateTrip(trip)
	lg.Flush()
	fmt.Fprint(writer, "Your trip has been created")
}

//noinspection GoUnhandledErrorResult
func (s Server) UpdateTripHandler(writer http.ResponseWriter, request *http.Request) {
	logger.NewLogger().Log(logging.Entry{Severity: logging.Info, Payload: "Inside update trip handler"})
	var trip domain.Trip
	err := json.NewDecoder(request.Body).Decode(&trip)
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		fmt.Fprint(writer, "invalid input!")
		return
	}
	s.tripService.UpdateTrip(id, trip)
	s.logger.Flush()
	fmt.Fprint(writer, "Your trip has been updated")
}

//noinspection GoUnhandledErrorResult
func (s Server) DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	logger.NewLogger().Log(logging.Entry{Severity: logging.Info, Payload: "Inside delete trip handler"})
	parameters := mux.Vars(request)
	id, _ := strconv.Atoi(parameters["id"])
	e := s.tripService.DeleteTrip(id)
	if e != nil {
		fmt.Fprint(writer, e)
		return
	}
	s.logger.Flush()
	fmt.Fprint(writer, "Your trip has been deleted")
}
//
//func logfunc() {
//	ctx := context.Background()
//
//	// Sets your Google Cloud Platform project ID.
//	projectID := "YOUR_PROJECT_ID"
//
//	// Creates a client.
//	client, err := logging.NewClient(ctx, projectID)
//	if err != nil {
//		log.Fatalf("Failed to create client: %v", err)
//	}
//
//	// Sets the name of the log to write to.
//	logName := "appengine.googleapis.com/stderr"
//
//	// Selects the log to write to.
//	logger := client.Logger(logName)
//
//	// Sets the data to log.
//	text := "Hello, world!"
//
//	// Adds an entry to the log buffer.
//	logger.Log(logging.Entry{Payload: text})
//
//	// Closes the client and flushes the buffer to the Stackdriver Logging
//	// service.
//	if err := client.Close(); err != nil {
//		log.Fatalf("Failed to close client: %v", err)
//	}
//
//	fmt.Printf("Logged: %v\n", text)
//}
