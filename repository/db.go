package Repository

import (
	"context"
	"log"
	"sync"

	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//I have used below constants just to hold required database config's.

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "bms"
)

func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func getMovies(collection string, filter bson.M, options *options.FindOptions) ([]model.Movie, error) {
	var result model.Movie
	var results []model.Movie

	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	//Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	cursor, err := db.Find(context.TODO(), filter, options)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func GetMovieCount(collection string) (int64, error) {
	client, err := GetMongoClient()
	if err != nil {
		return 0, err
	}

	// Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	res, err := db.CountDocuments(context.TODO(), bson.M{})

	if err != nil {
		return 0, err
	}

	return res, nil
}

func GetMovie(collection string, filter bson.M) (model.Movie, error) {
	var result model.Movie

	client, err := GetMongoClient()
	if err != nil {
		return model.Movie{}, err
	}

	//Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	db.FindOne(context.TODO(), filter).Decode(&result)

	return result, nil
}

func getAggregationPipeline(movieName string) []primitive.M {
	pipeline := []bson.M{
		{"$sort": bson.M{"start_date": 1}},
		{"$match": bson.M{"movie_name": movieName}},
		{
			"$lookup": bson.M{
				"from":         "movie",
				"localField":   "movie_name",
				"foreignField": "movie_name",
				"as":           "movie",
			},
		},
		{"$unwind": bson.M{"path": "$movie", "preserveNullAndEmptyArrays": true}},
		{
			"$lookup": bson.M{
				"from":         "theatre",
				"localField":   "theatre_name",
				"foreignField": "theatre_name",
				"as":           "theatre",
			},
		},
		{"$unwind": bson.M{"path": "$theatre", "preserveNullAndEmptyArrays": true}},
	}

	return pipeline
}

func GetShowsOfTheMovie(collection string, movieName string) ([]model.MovieShow, error) {
	var MovieShow model.MovieShow
	var MovieShows []model.MovieShow
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	// Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	pipeline := getAggregationPipeline(movieName)
	cursor, err := db.Aggregate(context.TODO(), pipeline)

	if err != nil {
		// return err
		log.Println("found some error in database query of count movies", err)
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&MovieShow)
		if err != nil {
			return nil, err
		}
		MovieShows = append(MovieShows, MovieShow)
	}
	return MovieShows, nil
}

func GetCoupons(collection string) ([]model.Coupon, error) {
	var result model.Coupon
	var results []model.Coupon

	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	//Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	cursor, err := db.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func getTheatres(collection string, filter bson.M) ([]model.Theatre, error) {
	var result model.Theatre
	var results []model.Theatre

	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	//Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	cursor, err := db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func Login(collection string, filter bson.M) (model.LoggedInObj, error) {
	var loggedInObj model.LoggedInObj
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return model.LoggedInObj{}, err
	}

	db := client.Database(DB).Collection(collection)
	log.Println("filter", filter)
	db.FindOne(context.TODO(), filter).Decode(&loggedInObj)

	log.Println("login data", loggedInObj)
	return loggedInObj, nil
}

func GetBookings(collection string, filter bson.M) ([]model.LoggedInObj, error) {
	var booking model.LoggedInObj
	var bookings []model.LoggedInObj
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	// Create a handle to the respective collection in the database.
	db := client.Database(DB).Collection(collection)
	findOptions := options.Find()

	// Sort by `price` field descending
	findOptions.SetSort(bson.M{"my_bookings.date": -1})
	cursor, retrieveErr := db.Find(context.TODO(), filter, findOptions)
	if retrieveErr != nil {
		return nil, retrieveErr
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&booking)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func createDocument(collection string, data interface{}, filter bson.M) (interface{}, error) {
	var existingData interface{}
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(DB).Collection(collection)
	db.FindOne(context.TODO(), filter).Decode(&existingData)
	log.Println("existingData", existingData)

	if existingData == nil || existingData == "" {
		_, err = db.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	} else {
		return "Record already existing", nil
	}
	return "Record Created", nil
}

func updateDocument(collection string, data bson.M, query bson.M) (*mongo.UpdateResult, error) {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(DB).Collection(collection)
	log.Println("collection, query, data", collection, query, data)
	cursor, err := db.UpdateOne(context.TODO(), query, data)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("cursor data is:", cursor)
	return cursor, nil
}

func deleteDocument(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(DB).Collection(collection)
	res, err := db.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

/*
registerUser function to register new users
*/
// func registerUser(c *gin.Context) {
// 	decoder := json.NewDecoder(c.Request.Body)
// 	var user User

// 	error := decoder.Decode(&user)
// 	if error != nil {
// 		panic(error)
// 	}

// 	isValidEmail := validateEmail(user.Email)
// 	fmt.Println("the result of the validation", user.Email)

// 	if isValidEmail {
// 		//Get MongoDB connection using connectionhelper.
// 		client, err := db.GetMongoClient()
// 		if err != nil {
// 			fmt.Println("err in user register fun", err)
// 			c.JSON(http.StatusForbidden, err)
// 		}

// 		buf := []byte(user.Password)
// 		ciphertext := make([]byte, 64)
// 		sha3.ShakeSum256(ciphertext, buf)
// 		user.Password = string(ciphertext)
// 		session := sessions.Default(c)

// 		//Create a handle to the respective collection in the database.
// 		collection := client.Database("bms").Collection("user")

// 		var existingUserData User

// 		collection.FindOne(c, bson.M{"email": user.Email}).Decode(&existingUserData)

// 		fmt.Println("the existing user detail", existingUserData)

// 		if existingUserData.Email != "" {
// 			c.JSON(http.StatusOK, "user already exists")
// 		} else {
// 			result, insertErr := collection.InsertOne(c, user)
// 			if insertErr != nil {
// 				msg := "error found while user registration"
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}

// 			fmt.Println("result", result)
// 			session.Set("email", user.Email)
// 			session.Set("gender", user.Gender)
// 			session.Set("is_admin", user.IsAdmin)
// 			session.Set("name", user.Name)
// 			session.Save()

// 			c.JSON(http.StatusOK, bson.M{"email": user.Email, "gender": user.Gender, "is_admin": user.IsAdmin, "name": user.Name})
// 		}
// 	} else {
// 		c.JSON(http.StatusUnauthorized, "email id is not valid, please check your email id")
// 	}
// }
