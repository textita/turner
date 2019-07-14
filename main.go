package main

import(
  "fmt"
  "log"
  "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/bson/primitive"
 )

/*

type OtherName struct {
    TitleNameLanguage string
    TitleNameType string
    TitleNameSortable string 
    TitleName  string
  }
  
type StoryLine struct {
    Description string 
    Language  string 
    Type  string
    }

  type Participant struct {
   IsKey bool
   RoleType string
   IsOnScreen bool
   ParticipantType  string
   Name string
   ParticipantId int
   StoryLines []StoryLine
   
  }
  
 type Award struct {
 AwardWon bool
 AwardYear  string 
 Award  string
 AwardCompany  string 
}

*/

type Title struct {
//ID primitive.ObjectID  `json:"_id" bson:"_id"`

  TitleId  int `json:"title_id" bson:"title_id"`
  TitleName string `json:"title_name" bson:"title_name"`
  TitleNameSortable  string `json:"tile_name_sortable" bson:"title_name_sortable"`

/*
  Awards  []Award
  Genres  []string
  OtherNames []OtherName
  Participants []Participant
  ReleaseYear  string
*/
  
}

func main() {
  
connStr:= "mongodb://readonly:turner@ds043348.mongolab.com:43348/dev-challenge"
  
// Set client options
clientOptions := options.Client().ApplyURI(connStr)

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")

  collection := client.Database("dev-challenge").Collection("Titles")
  
// Pass these options to the Find method
findOptions := options.Find()

// Here's an array in which you can store the decoded documents
var results []*Title

cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem Title
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }
results = append(results, &elem)
    }

if err := cur.Err(); err != nil {
    log.Fatal(err)
}

// Close the cursor once finished
cur.Close(context.TODO())

//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

for _, r := range results {
      fmt.Printf("%+v\n",r)
}

}
