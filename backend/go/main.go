package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Flight struct {
    ID     string `json:"id,omitempty" bson:"_id,omitempty"`
    Number string `json:"number,omitempty" bson:"number,omitempty"`
    Status string `json:"status,omitempty" bson:"status,omitempty"`
    Gate   string `json:"gate,omitempty" bson:"gate,omitempty"`
}

type Notification struct {
    Email string `json:"email"`
    Phone string `json:"phone"`
}

func CreateFlightEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    var flight Flight
    json.NewDecoder(request.Body).Decode(&flight)
    collection := client.Database("flight_status").Collection("flights")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    collection.InsertOne(ctx, flight)
    json.NewEncoder(response).Encode(flight)

    
    go sendFlightUpdate(flight)
}

func GetFlightsEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    var flights []Flight
    collection := client.Database("flight_status").Collection("flights")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var flight Flight
        cursor.Decode(&flight)
        flights = append(flights, flight)
    }
    if err := cursor.Err(); err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    json.NewEncoder(response).Encode(flights)
}

func CreateNotificationEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    var notification Notification
    json.NewDecoder(request.Body).Decode(&notification)
    
    json.NewEncoder(response).Encode(map[string]string{"message": "Settings saved"})
}

func main() {
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ = mongo.Connect(ctx, clientOptions)
    router := mux.NewRouter()
    router.HandleFunc("/api/flights", CreateFlightEndpoint).Methods("POST")
    router.HandleFunc("/api/flights", GetFlightsEndpoint).Methods("GET")
    router.HandleFunc("/api/notifications", CreateNotificationEndpoint).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}
