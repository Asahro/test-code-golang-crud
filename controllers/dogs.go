package controllers

import (
	"crud/extparty"
	"crud/models"
	"crud/utils"
	"math"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type ResponDogs struct {
	Id   int
	Name string
	Race string
	Age  string
}

type ResponBreeds struct {
	Breeds     string
	SubBreeds  interface{}
	BreedImage []string `json:"breedImage,omitempty"`
}

type DogsController struct {
	beego.Controller
}

func (c *DogsController) ReadDogs() {
	var resp utils.ResponseSchema
	resp.Code = 400
	resp.Data = nil
	var responDogs []ResponDogs
	var responDog ResponDogs

	dogID, _ := c.GetInt("dogId")
	limit, _ := c.GetInt("limit")
	page, _ := c.GetInt("offset")
	if limit < 1 {
		limit = 10
	}
	offset := page * limit

	dogID = int(math.Abs(float64(dogID)))
	limit = int(math.Abs(float64(limit)))
	offset = int(math.Abs(float64(offset)))

	if dogID == 0 {
		data, err := models.ReadDogs(limit, offset)
		if err != nil {
			resp.Message = err.Error()
		} else {
			for i := 0; i < len(data); i++ {
				responDog.Id = data[i].Id
				responDog.Name = data[i].Name
				responDog.Race = data[i].Race
				responDog.Age = data[i].Age
				responDogs = append(responDogs, responDog)
			}
			resp.Data = responDogs
			resp.Message = "success"
			resp.Code = 200
		}
	} else {
		if !models.IsDogExist(dogID) {
			resp.Message = "dog's data not found"
		} else {
			if data, err := models.ReadDogById(dogID); err != nil {
				resp.Message = "dog's data not found"
			} else {
				resp.Message = "success"
				resp.Code = 200
				resp.Data = data
			}
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "GET, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *DogsController) CreateDog() {
	var resp utils.ResponseSchema
	resp.Code = 400
	resp.Data = nil

	name := c.GetString("name")
	race := c.GetString("race")
	age := c.GetString("age")

	valid := validation.Validation{}
	valid.Required(name, "name")
	valid.Required(race, "race")
	valid.Required(age, "age")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		var dog models.Dogs
		dog.Name = name
		dog.Race = race
		dog.Age = age

		err := models.CreateDog(dog)

		if err != nil {
			resp.Code = 500
			resp.Message = "failed"
			resp.Data = err
		} else {
			resp.Code = 200
			resp.Message = "success"
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *DogsController) UpdateDog() {
	var resp utils.ResponseSchema
	resp.Code = 400
	resp.Data = nil

	dogId, _ := c.GetInt("dogId")
	dogId = int(math.Abs(float64(dogId)))
	name := c.GetString("name")
	race := c.GetString("race")
	age := c.GetString("age")

	valid := validation.Validation{}
	valid.Required(dogId, "dog id")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		if !models.IsDogExist(dogId) {
			resp.Message = "dog's data not found"
		} else {
			dog := orm.Params{
				"id":   dogId,
				"name": name,
				"race": race,
				"age":  age,
			}

			for i, val := range dog {
				if val == "" {
					delete(dog, i)
				}
			}

			err := models.UpdateDog(dog)

			if err != nil {
				resp.Code = 500
				resp.Message = err.Error()
			} else {
				resp.Code = 200
				resp.Message = "success"
			}
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *DogsController) DeleteDog() {
	var resp utils.ResponseSchema
	resp.Code = 400
	resp.Data = nil

	dogId, _ := c.GetInt("dogId")
	dogId = int(math.Abs(float64(dogId)))

	valid := validation.Validation{}
	valid.Required(dogId, "dog id")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			resp.Message = err.Key + " " + strings.ToLower(err.Message)
		}
	} else {
		if !models.IsDogExist(dogId) {
			resp.Message = "dog's data not found"
		} else {
			if err := models.DeleteDog(dogId); err != nil {
				resp.Code = 500
				resp.Message = "failed"
			} else {
				resp.Code = 200
				resp.Message = "success"
			}
		}
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *DogsController) BreedList() {
	var resp utils.ResponseSchema
	resp.Code = 400
	var breed ResponBreeds
	var breeds []ResponBreeds

	listData, _ := extparty.GetDogBreeds()

	for _, key := range reflect.ValueOf(listData.Message).MapKeys() {
		breed.Breeds = key.String()
		breed.SubBreeds = listData.Message[key.String()]
		breeds = append(breeds, breed)
	}

	resp.Data = breeds
	resp.Message = "success"
	resp.Code = 200

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *DogsController) BreedDetail() {
	var resp utils.ResponseSchema
	resp.Code = 400
	breedName := c.Ctx.Input.Param(":id")
	var breedDetailRespon ResponBreeds
	var listSubBreed []ResponBreeds
	var subBreed ResponBreeds

	listImage, _ := extparty.GetEmagesDogBreeds(breedName)
	listData, _ := extparty.GetDogBreeds()

	for _, subBreedKey := range listData.Message[breedName] {
		subBreed.Breeds = subBreedKey
		subBreed.SubBreeds = nil
		listSubBreed = append(listSubBreed, subBreed)
	}

	breedDetailRespon.BreedImage = listImage.Message
	breedDetailRespon.Breeds = breedName
	breedDetailRespon.SubBreeds = listSubBreed

	resp.Data = breedDetailRespon
	resp.Message = "success"
	resp.Code = 200

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Method", "POST, OPTIONS")
	c.Ctx.Output.Header("Content-Type", "application/json, application/x-www-form-urlencoded")
	c.Ctx.Output.SetStatus(resp.Code)
	c.Data["json"] = resp
	c.ServeJSON()
}
