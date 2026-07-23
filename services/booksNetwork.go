package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"pex.oschmid.com/model"
)

/**
* This service is for searching all book related queries
 */

func GetBooksFromGoogleBooks(searchText string) {

	//http.Get
}

/**
*This service is for searching all books based on text question
 */
func GetBooksFromOpenLibrary(searchText string) ([]model.BookResponse, error) {

	client := &http.Client{}

	fmt.Println("GetBooksFromOpenLibrary: kickstart checking books")
	req, err := http.NewRequest("GET", "https://openlibrary.org/search.json", nil)

	updatedText := strings.Join(strings.Split(searchText, " "), "+")
	if err != nil {
		fmt.Println("services.GetBooksFromOpenLibraryError", err)

	}

	q := req.URL.Query()

	q.Add("q", updatedText)
	q.Add("fields", "*,availability")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("GetBooksFromOpenLibrary err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	//Convert the body to type string
	//sb := string(body)
	//fmt.Println(string(sb))
	djs := json.NewDecoder(resp.Body)

	sb := model.OpenLibraryResponse{}
	djs.Decode(&sb)
	//fmt.Println(sb.Docs)
	returnObj := []model.BookResponse{}
	for _, obj := range sb.Docs {

		returnObj = append(returnObj, mapOpenLibraryToBookReturnResponse(obj))

	}

	return returnObj, nil
}

func mapOpenLibraryToBookReturnResponse(openLibBook model.OpenLibraryIndividualBook) model.BookResponse {
	var isbn string
	if len(openLibBook.Availability.Isbn) > 0 {
		isbn = openLibBook.Availability.Isbn

	} else if openLibBook.Isbn != nil {
		isbn = openLibBook.Isbn[0]

	}
	var coverUrl string
	if len(isbn) > 0 {
		coverUrl = "https://covers.openlibrary.org/b/isbn/" + isbn + "-M.jpg"
	}

	return model.BookResponse{
		Isbn:        isbn,
		Title:       openLibBook.Title,
		Subtitle:    openLibBook.Subtitle,
		PageCount:   openLibBook.MedianPages,
		CoverImage:  coverUrl,
		PublishYear: openLibBook.FirstPublishYear,
		Chapter:     openLibBook.Chapter,
		AuthorKey:   openLibBook.AuthorKey,
		AuthorName:  openLibBook.AuthorName,
	}
}
