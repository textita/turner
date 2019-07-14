package main

import(
    "fmt"
  "log"
  "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

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

  collection := client.Database("dev-challenge").Collection("titles")
  
  cur, err := collection.Find(context.TODO(), bson.D{{}})

if err != nil {
    log.Fatal(err)
}

type otherName struct {
    TitleNameLanguage string
    TitleNameType string
    TitleNameSortable string 
    TitleName  string
  }
  
type storyLine struct {
    Description string 
    Language  string 
    Type  string
    }

  type participant struct {
   IsKey bool
   RoleType string
   IsOnScreen bool
   ParticipantType  string
   Name string
   ParticipantId int
   StoryLines []storyline
   TitleId  int
   TitleName string
   TitleNameSortable  string
  }
  
 type award struct {
 AwardWon bool
 AwardYear  string 
 Participants  []string
 Award  string
 AwardCompany  string 
}
  

  type title struct {
  Id  string
  Awards  []award
  Genres  []string
  OtherNames []otherName
  Participants []participant
  ReleaseYear  string

  }


  func main() {

  fmt.Println("hello")



  

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem title
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }
}
  
}
 
