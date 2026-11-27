package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"pex.oschmid.com/model"
)

/*
* TODO: Need to implement this feature
* This service is for searching all book related queries
 */

func GetBooksFromGoogleBooks(searchText string, searchType string) {

	// baseUrl := os.Getenv("GOOGLE_BOOKS_API")
	// subString := "books/v1/volumes"
	// client := &http.Client{}
	// fullUrl := baseUrl + subString
	// req, err := http.NewRequest("GET", fullUrl, nil)
	// updatedText := strings.Join(strings.Split(searchText, " "), "+")

}

/**
*This service is for searching all books based on text question
 */
func GetBooksFromOpenLibrary(searchText string) ([]model.BookResponse, error) {

	client := &http.Client{}
	var openApiHeader = os.Getenv("PEX_HEADER")
	fmt.Println("GetBooksFromOpenLibrary: kickstart checking books")
	req, err := http.NewRequest("GET", "https://openlibrary.org/search.json", nil)
	req.Header.Add("User-Agent", openApiHeader)
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

/*
* mapper function which maps return response from response obtained from OpenLibrary
 */
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
