// package main
package dbOperator

//负责从API将数据转移到数据库中
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/peterhellberg/swapi"
)

type Node struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Format		string	`json:"format"`
}

type Schema struct {
	Required    []string        `json:"required"`
	Title       string          `json:"title"`
	Properties  map[string]Node `json:"properties"`
	Description string          `json:"description"`
	SSchema     string          `json:"$schema"`
	Type        string          `json:"type"`
}

func personDownLoad(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Person
	person, err := c.Person(id)
	if err != nil {
		return err
	} else if person.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	personBuckets, err := tx.CreateBucketIfNotExists([]byte("Person"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(person)
	if err != nil {
		return err
	}

	err = personBuckets.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func filmDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Film
	films, err := c.Film(id)
	if err != nil {
		return err
	} else if films.Title == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	filesBucket, err := tx.CreateBucketIfNotExists([]byte("Film"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(films)
	if err != nil {
		return err
	}

	err = filesBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func starshipsDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Starship
	starships, err := c.Starship(id)
	if err != nil {
		return err
	} else if starships.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	starshipsBucket, err := tx.CreateBucketIfNotExists([]byte("Starship"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(starships)
	if err != nil {
		return err
	}

	err = starshipsBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func vehicleDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Vehicle
	vehicle, err := c.Vehicle(id)
	if err != nil {
		return err
	} else if vehicle.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	vehicleBucket, err := tx.CreateBucketIfNotExists([]byte("Vehicle"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	err = vehicleBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func planetDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Planet
	planet, err := c.Planet(id)
	if err != nil {
		return err
	} else if planet.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	planetBucket, err := tx.CreateBucketIfNotExists([]byte("Planet"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(planet)
	if err != nil {
		return err
	}

	err = planetBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func speciesDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Species
	species, err := c.Species(id)
	if err != nil {
		return err
	} else if species.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	speciesBucket, err := tx.CreateBucketIfNotExists([]byte("Species"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(species)
	if err != nil {
		return err
	}

	err = speciesBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func CrawlData() {

	db, err := bolt.Open("my.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 88; i++ {
		fmt.Println("People:", i)
		if err := personDownLoad(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 1; i <= 100; i++ {
		fmt.Println("Planet:", i)
		if err := planetDownload(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 2; i <= 100; i++ {
		fmt.Println("Starships:", i)
		if err := starshipsDownload(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("Film:", i)
		if err := filmDownload(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 1; i <= 40; i++ {
		fmt.Println("Species:", i)
		if err := speciesDownload(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 40; i <= 80; i++ {
		fmt.Println("Vehicles:", i)
		if err := vehicleDownload(db, i); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func CrawlSchema(addr string) (Schema, error) {
	c := swapi.DefaultClient
	req, err := c.NewRequest(fmt.Sprintf(addr))
	if err != nil {
		return Schema{}, err
	}
	var schema Schema
	if _, err = c.Do(req, &schema); err != nil {
		return Schema{}, err
	}
	return schema, nil
}

func CrawlAllSchema(db *bolt.DB){
	if peopleSchema, err := CrawlSchema("people/schema"); err == nil {
		jsonCode, _ := json.Marshal(peopleSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Person"), jsonCode)
			return nil
		})
	}
	if filmSchema, err := CrawlSchema("films/schema"); err == nil {
		jsonCode, _ := json.Marshal(filmSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Film"), jsonCode)
			return nil
		})
	}
	if starshipSchema, err := CrawlSchema("starships/schema"); err == nil {
		jsonCode, _ := json.Marshal(starshipSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Starship"), jsonCode)
			return nil
		})
	}
	if vehicleSchema, err := CrawlSchema("vehicles/schema"); err == nil {
		jsonCode, _ := json.Marshal(vehicleSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Vehicle"), jsonCode)
			return nil
		})
	}
	if planetSchema, err := CrawlSchema("planets/schema"); err == nil {
		jsonCode, _ := json.Marshal(planetSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Planet"), jsonCode)
			return nil
		})
	}
	if speciesSchema, err := CrawlSchema("species/schema"); err == nil {
		jsonCode, _ := json.Marshal(speciesSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Species"), jsonCode)
			return nil
		})
	}
}