package main

import(
  "fmt"
  "log"
  "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

/*
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
   StoryLines []storyLine
   
  }
  
 type award struct {
 AwardWon bool
 AwardYear  string 
 Award  string
 AwardCompany  string 
}

*/

type title struct {
  Id  string `json:"_id"`
  


/*
  Awards  []award
  Genres  []string
  OtherNames []otherName
  Participants []participant
  ReleaseYear  string

  TitleId  int
  TitleName string
  TitleNameSortable  string

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
  
  cur, err := collection.Find(context.TODO(), bson.D{{}})

if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a tim

count := 0

for cur.Next(context.TODO()) {
    
  fmt.Println("*")

    var elem title
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }
   count++;
   fmt.Println("Id ",elem.Id," count ",count)
}
}
 
