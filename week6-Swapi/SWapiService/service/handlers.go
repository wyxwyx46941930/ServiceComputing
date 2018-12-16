package service

import (
    "net/http"
    "github.com/unrolled/render"
    "encoding/json"
	"fmt"
	"strconv"
	"errors"
	"github.com/gorilla/mux"
	"github.com/SYSUServiceOnComputingCloud2018/SwapiService/dbOperator"
	"github.com/peterhellberg/swapi"
	"github.com/boltdb/bolt"

)
const (
	ErrorResponseCode   = "404" // 错误响应code
	SuccessResponseCode = "200"    // 正确响应code
)

// type ResponseMessage struct {
// 	Code    string         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

type Peoples struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []interface{} `json:"results"`
}

func rootHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {

		formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
		formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
		formatter.JSON(w,http.StatusOK,struct{
			Films string `json:"films"`
			Peoples string `json:"people"`
			Planets string `json:"planets"`
			Species string `json:"species"`
			Starships string `json:"starships"`
			Vehicles string `json:"vehicles"`
		}{Films:"http://localhost:3000/api/films/",
		Peoples:"http://localhost:3000/api/people/",
		Planets:"http://localhost:3000/api/planets/",
		Species:"http://localhost:3000/api/species/",
		Starships:"http://localhost:3000/api/starships/",
		Vehicles:"http://localhost:3000/api/vehicles/"})
    }
}

func peopleHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		vars := req.URL.Query();
		search, search_ok:= vars["search"]
		page_param , page_ok := vars["page"]
		if search_ok{
			fmt.Printf("param 'search' string is [%s]\n", search[0])
			v , err := dbOperator.GetElementsBySearchField(db,"Person",search[0])
			users := Peoples{Count:len(v)}
			if err == nil {
				for i := 0; i < len(v); i++ {
					var user swapi.Person
					err = json.Unmarshal(v[i], &user)
					if err != nil {
						fmt.Println(err)
					} else {
						// fmt.Println(user)
						users.Results = append(users.Results,user)
					}
				}
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,users)
			} else{
				fmt.Println(err)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
			}

		} else{
			// 得到query变量
			page := 0
			if page_ok {
				fmt.Printf("param 'page' string is [%s]\n", page_param)
				page1,err :=strconv.Atoi(page_param[0])
				if err != nil{
					fmt.Println(err)
				}
				page = page1
			} else {
				fmt.Printf("query param 'page' does not exist\n")
				fmt.Printf("The default page index is 1\n")
				page = 1
			}
			// 读取数据库
			v , err := dbOperator.GetAllResources(db,"Person")
			users := Peoples{}

			// 判断页码逻辑
			if page <= 0 {
				err := errors.New("Page index <= 0.")
				fmt.Println(err)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
				return 
			}
			if (page-1)*10 >= len(v){
				err := errors.New("Page index >= max_#page.")
				fmt.Println(err)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
				return 
			}
			if page*10 > len(v){
				users.Count = len(v) % 10
			} else {
				users.Count = 10
			}

			
			if err == nil {
				for i := (page-1)*10; i < (page-1)*10+users.Count; i++ {
					var user swapi.Person
					err = json.Unmarshal(v[i], &user)
					if err != nil {
						fmt.Println(err)
					} else {
						// fmt.Println(user)
						users.Results = append(users.Results,user)
					}
				}
				if users.Count == 10{
					users.Next = "localhost:3000/api/people/?page="+strconv.Itoa(page+1)
				}
				if page != 1{
					users.Previous = "localhost:3000/api/people/?page="+strconv.Itoa(page-1)
				}
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,users)
			} else{
				fmt.Println(err)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
			}
		}
	}
}

func peopleSchemaHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){

		// 输出schema
		jsonData, _ := dbOperator.GetSchemaByBucket(db, "Person")
		var schema dbOperator.Schema //定义在crawler.go中
		err := json.Unmarshal(jsonData, &schema)
		if err == nil {
			fmt.Println(schema)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
			formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
			formatter.JSON(w,http.StatusOK,schema)
		} else {
			fmt.Println(err)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		}
	}
}
	
func peopleIdHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
	
	return func(w http.ResponseWriter, req *http.Request){
		
		vars := mux.Vars(req)
		// 获取id
		id := vars["id"]
		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Person", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var user swapi.Person
			err = json.Unmarshal(v, &user)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(user.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", user)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,user)
			}
		}
	}
}

func planetsHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {

		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Planet", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var planet swapi.Planet
			err = json.Unmarshal(v, &planet)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(planet.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", planet)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,planet)
			}
		}
	}
}

func filmsHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {

		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Film", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var film swapi.Film
			err = json.Unmarshal(v, &film)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(film.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", film)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,film)
			}
		}
	}
}

func speciesHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		
		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Species", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var species swapi.Species
			err = json.Unmarshal(v, &species)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(species.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", species)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,species)
			}
		}
	}
}

func vehiclesHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {

		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Vehicle", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var vehicle swapi.Vehicle
			err = json.Unmarshal(v, &vehicle)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(vehicle.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", vehicle)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,vehicle)
			}
		}
	}
}

func starshipsHandler(formatter *render.Render,db *bolt.DB) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {

		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Starship", id)
		if err != nil {
			fmt.Println(err)
			// WriteResponse(w, ErrorResponseCode, "failed", nil)
			formatter.Text(w, http.StatusOK, "HTTP/1.0 "+ErrorResponseCode+" Not Found\n")
		} else {
			var starship swapi.Starship
			err = json.Unmarshal(v, &starship)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(starship.Name)
				// WriteResponse(w, SuccessResponseCode, "OK", starship)
				formatter.Text(w, http.StatusOK, "HTTP/1.0 "+SuccessResponseCode+" OK\n")
				formatter.Text(w, http.StatusOK, "Content-Type: application/json\n")
				formatter.JSON(w,http.StatusOK,starship)
			}
		}
	}
}

// func WriteResponse(w http.ResponseWriter, code string, message string, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	resp := ResponseMessage{Code: code, Message: message, Data: data}
// 	b, err := json.Marshal(resp)
// 	if err != nil {
// 		// logrus.Warnf("error when marshal response message, error:%v\n", err)
// 		fmt.Println(err)
// 	}
// 	w.Write(b)
// }
