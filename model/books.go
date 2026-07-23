package model

type OpenLibraryResponse struct {
	Numfound string                      `json:"numFound"`
	Count    int                         `json:"num_found"`
	Docs     []OpenLibraryIndividualBook `json:"docs"`
}

type OpenLibraryIndividualBook struct {
	Availability     OpenLibraryAvaialibility `json:"availability"`
	Title            string                   `json:"title"`
	Subtitle         string                   `json:"subtitle"`
	Subject          []string                 `json:"subject"  `
	AuthorKey        []string                 `json:"author_key"`
	AuthorName       []string                 `json:"author_name"`
	Chapter          []string                 `json:"chapter"`
	FirstPublishYear int                      `json:"first_publish_year"`
	MedianPages      int                      `json:"number_of_pages_median"`
	Isbn             []string                 `json:"isbn"`
}

type OpenLibraryAvaialibility struct {
	Isbn string `json:"isbn"`
}

type BookResponse struct {
	Isbn        string   `json:"isbn"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	CoverImage  string   `json:"coverImage"`
	Subject     []string `json:"subject"  `
	AuthorKey   []string `json:"authorsKey"`
	AuthorName  []string `json:"authorsName"`
	Chapter     []string `json:"chapters"`
	PublishYear int      `json:"publishYear"`
	PageCount   int      `json:"pageCount"`
	Id          string   `json:"id"`
	Source      string   `json:"source"`
}
