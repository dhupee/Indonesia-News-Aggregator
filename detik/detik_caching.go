package detik

import(
	"context"
	"encoding/json"
	"log"

	db "github.com/dhupee/Indonesia-News-Aggregator/db"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func SetDetikNewsCache(url string, news DetikNewsStruct) error {
	// make sure rdb is not nil
	if rdb == nil {
		rdb = db.GetDB()
	}

	// parse the news struct to json
	jsonNews, err := json.Marshal(news)
	if err != nil {
		return err
	}

	// set the news in the cache
	err = rdb.Set(ctx, url, jsonNews, 0).Err()
	if err != nil {
		log.Printf("Error setting news in cache: %v", err)
		return err
	}

	return nil
}

func GetDetikNewsCache(url string) (DetikNewsStruct, error) {
	// make sure rdb is not nil
	if rdb == nil {
		rdb = db.GetDB()
	}

	// get the news from the cache
	jsonNews, err := rdb.Get(ctx, url).Result()
	if err != nil {
		log.Printf("Error getting news from cache or news haven't been cached: %v", err)
		return DetikNewsStruct{}, err
	}

	// parse the json to struct
	var news DetikNewsStruct
	err = json.Unmarshal([]byte(jsonNews), &news)
	if err != nil {
		log.Printf("Error unmarshalling news from cache: %v", err)
		return DetikNewsStruct{}, err
	}

	return news, nil
}
