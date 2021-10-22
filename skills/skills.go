package main

import (
	"context"
	"fmt"
	"github.com/mongo-scripts/skills/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func main() {

	config.Connect()

	skillsImport := config.Client.Database("testing").Collection("skills_import")

	cur, err := skillsImport.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err = cur.Close(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
	}()

	var m []primitive.M

	for cur.Next(context.Background()) {

		var result bson.D

		err = cur.Decode(&result)

		if err != nil {
			log.Fatalln(err)
		}

		m = append(m, result.Map())
	}

	skillsImportFlat := config.Client.Database("testing").Collection("skills_import_flat")

	var skills []interface{}

	for _, record := range m {

		categories := getCategories(record)

		newObject := bson.D{{"Name", record["Name"]}, {"Description", record["Description"]},
			{"Categories", categories}}

		skills = append(skills, newObject)

		fmt.Println(newObject)

	}

	res, err := skillsImportFlat.InsertMany(context.Background(), skills)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.InsertedIDs)

}

func getCategories(m primitive.M) []string {

	var categories []string

	for idx := 1; idx < 5; idx++ {
		if val, ok := m[fmt.Sprintf("Parent category %d", idx)]; ok {
			if val.(string) != "" {
				categories = append(categories, val.(string))
			}
		}
	}
	return categories
}
