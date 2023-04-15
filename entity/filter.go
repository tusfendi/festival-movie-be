package entity

type SearchByKeyword struct {
	Keyword string `form:"keyword"`
}

type SearchMovies struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	ActorID     string `form:"artist_id"`
	GenreID     string `form:"genre_id"`
	PerPage     string `form:"per_page"`
	PageNumber  string `form:"page_number"`
}
