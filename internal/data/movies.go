package data

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time // timestamps for when the movie is add to our database
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}
