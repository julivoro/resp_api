package main

/*
	HTTP RESTapi example
	on localhost:10000
	artciles in JSON format
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id`
	Title   string `json:"Title"`
	Desc    string `json:"desc`
	Content string `json:"content"`
}

type Articles []Article

var articles = Articles{
	Article{Id: "0", Title: "Test Title", Desc: "test Desc", Content: "test content"},
	Article{Id: "1", Title: "Test Title1", Desc: "test Desc1", Content: "test content1"},
	Article{Id: "2", Title: "Test Title2", Desc: "test Desc2", Content: "test content2"},
	Article{Id: "3", Title: "Test Title3", Desc: "test Desc3", Content: "test content3"},
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: allArticles")
	json.NewEncoder(w).Encode(articles)
}

func testAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint reached!")
	fmt.Println("Endpoint Hit: homePage")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]
	//fmt.Fprintf(w, "Key: "+key)

	//Loop over all of Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)

	fmt.Println("Endpoint Hit: createNewArticle")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
	fmt.Println("Endpoint Hit: deleteArticle")
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	for _, article := range articles {
		if article.Id == id {
			i, _ := strconv.Atoi(id)
			reqBody, _ := ioutil.ReadAll(r.Body)

			var article Article
			json.Unmarshal(reqBody, &article)
			// update our global Articles array to include
			// our new Article
			articles[i].Title = article.Title
			articles[i].Content = article.Content
			articles[i].Desc = article.Desc

			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Println("Endpoint Hit: updateArticle")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testAllArticles).Methods("POST")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{Id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{Id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
