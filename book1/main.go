package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID     primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title  string             `json:"title"   bson:"title,omitempty"`
	Author string             `json:"author"   bson:"author,omitempty"`
	Price  float64            `json:"price"    bson:"price,omitempty"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func DatabaseConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CloseConnection(c *mongo.Client) {
	c.Disconnect(context.Background())
}

func main() {
	// Initialize router

	router := mux.NewRouter()

	router.Use(corsMiddleware)
	con := DatabaseConnection()
	defer CloseConnection(con)
	// Define routes
	router.HandleFunc("/books", getBooks)
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", createBook)
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook)
	// router.HandleFunc("/book",updateBookDetails)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			// Handle preflight request
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	con := DatabaseConnection()
	defer CloseConnection(con)
	collection = con.Database("book").Collection("book")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	var books []Book
	for cur.Next(context.Background()) {
		var book Book
		cur.Decode(&book)
		books = append(books, book)

	}
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// // id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid book ID", http.StatusBadRequest)
	// 	return
	// }

	// // for _, book := range books {
	// // 	if book.ID == id {
	// // 		json.NewEncoder(w).Encode(book)
	// // 		return
	// // 	}
	// // }

	// http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	con := DatabaseConnection()
	defer CloseConnection(con)
	// fmt.Printf("%+v", r.Body)
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book1 := bson.M{
		"_id":    primitive.NewObjectID(),
		"title":  book.Title,
		"author": book.Author,
		"price":  book.Price,
	}

	coll := con.Database("book").Collection("book")
	_, err = coll.InsertOne(context.Background(), book1)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	con := DatabaseConnection()
	defer CloseConnection(con)
	params := mux.Vars(r)
	bookID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coll := con.Database("book").Collection("book")
	result, err := coll.DeleteOne(context.Background(), bson.M{"_id": bookID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Book deleted successfully")
}
