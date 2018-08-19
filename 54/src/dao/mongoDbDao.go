package dao

import (
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

const Host = "localhost"
const Port = 27017
var Url = Host + ":" + strconv.Itoa(Port)
const DatabaseName = "exercise54"
const UrlMappingCollectionName = "urlMappings"
const invocationsCollectionName = "invokations"

type UrlMapping struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	LongUrl string `bson:"longUrl"`
	ShortUrl string `bson:"shortUrl"`
}

type invocation struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	ShortUrl string `bson:"shortUrl"`
	NumberOfInvocations int `bson:"numberOfInvocations"`
}

func GetAllUrlMappings() (urlMappings []*UrlMapping, err error){
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	urlMappingCollection := session.DB(DatabaseName).C(UrlMappingCollectionName)

	// Fetch data from DB.
	err = urlMappingCollection.Find(bson.M{}).All(&urlMappings)
	return
}

func GetAllShortUrls() (shortUrls []string, err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	urlMappingCollection := session.DB(DatabaseName).C(UrlMappingCollectionName)

	// Fetch data from DB.
	var urlMappings []*UrlMapping
	err = urlMappingCollection.Find(bson.M{}).All(&urlMappings)

	if err != nil {
		return
	}

	for _, urlMapping := range urlMappings {
		shortUrls = append(shortUrls, urlMapping.ShortUrl)
	}

	return
}

func GetShortUrlForLongUrl(longUrl string) (shortUrl string, err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	urlMappingCollection := session.DB(DatabaseName).C(UrlMappingCollectionName)

	// Fetch data from DB.
	var urlMapping *UrlMapping
	err = urlMappingCollection.Find(bson.M{"longUrl": longUrl}).One(&urlMapping)

	if err != nil {
		return
	}

	shortUrl = urlMapping.ShortUrl
	return
}

func GetLongUrlForShortUrl(shortUrl string) (longUrl string, err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	urlMappingCollection := session.DB(DatabaseName).C(UrlMappingCollectionName)

	// Fetch data from DB.
	var urlMapping *UrlMapping
	err = urlMappingCollection.Find(bson.M{"shortUrl": shortUrl}).One(&urlMapping)

	if err != nil {
		return
	}

	longUrl = urlMapping.LongUrl
	return
}

func MappingExists(longUrl string) (found bool, err error) {
	shortUrl, err := GetShortUrlForLongUrl(longUrl)
	if err == mgo.ErrNotFound {
		err = nil
		found = false // To be overly clear...
		return
	}

	found = len(shortUrl) == 0// Should not happen...
	return
}

func AddMapping(shortUrl string, longUrl string) (err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	urlMappingCollection := session.DB(DatabaseName).C(UrlMappingCollectionName)

	// Insert document.
	err = urlMappingCollection.Insert(UrlMapping{ShortUrl: shortUrl, LongUrl: longUrl})
	return err
}

func GetNumberOfUrlInvocations(shortUrl string) (numberOfInvocations int, err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	invocationsCollection := session.DB(DatabaseName).C(invocationsCollectionName)

	var invocation invocation
	err = invocationsCollection.Find(bson.M{"shortUrl": shortUrl}).One(&invocation)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when finding invocation with shortURL: %s\n", shortUrl)
		return
	}

	numberOfInvocations = invocation.NumberOfInvocations
	return
}

func ShortUrlInvoked(shortUrl string) (err error) {
	// Fetch session.
	session := getSession()
	defer session.Close()

	// Fetch collection.
	invocationsCollection := session.DB(DatabaseName).C(invocationsCollectionName)

	var invc invocation
	err = invocationsCollection.Find(bson.M{"shortUrl": shortUrl}).One(&invc)

	if err != nil {
		// First invocation?
		if err == mgo.ErrNotFound {
			// Then insert a new document.
			err = invocationsCollection.Insert(invocation{ShortUrl: shortUrl, NumberOfInvocations: 1})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not insert incocation document for shortURL: %s\n", shortUrl)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error when finding invocation with shortURL: %s\n", shortUrl)
		}
		return
	}

	// Not the first invocation.
	fmt.Fprintf(os.Stderr, "invc with shortURL: %s already existed. Will inc it.\n", shortUrl)
	query := bson.M{"shortUrl": shortUrl}
	change := bson.M{"$inc": bson.M{"numberOfInvocations": 1}}
	err = invocationsCollection.Update(query, change)

	return
}


func getSession() (session *mgo.Session) {
	// Open session to DB.
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return
}
