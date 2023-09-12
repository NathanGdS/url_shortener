package models

import (
	"time"

	"github.com/lucsky/cuid"
)

type UrlShortener struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid()"`
	Slug         string    `json:"slug"`
	Url          string    `json:"url"`
	Opened       bool      `json:"opened"`
	OpenedTimes  int       `json:"opened_times"`
	LastOpenedAt time.Time `json:"last_opened_at"`
	ShortenedAt  time.Time `json:"shortened_at"`
}

func (u *UrlShortener) Open() {
	u.Opened = true
	u.OpenedTimes++
	u.LastOpenedAt = time.Now()
}

func (u *UrlShortener) Shortener() {
	u.OpenedTimes = 0
	u.Slug = cuid.Slug()
	u.ShortenedAt = time.Now()
}

func (u *UrlShortener) Create(url string) {
	u.Url = url
	u.Shortener()
}

func checkDate(date time.Time) interface{} {
	if date.IsZero() {
		return nil
	} else {
		return date.Format("2006-01-02 15:04:05")
	}
}

type UrlShortenerDto struct {
	Id           string      `json:"id"`
	Slug         string      `json:"slug"`
	Url          string      `json:"url"`
	Opened       bool        `json:"opened"`
	OpenedTimes  int         `json:"opened_times"`
	LastOpenedAt interface{} `json:"last_opened_at"`
	ShortenedAt  interface{} `json:"shortened_at"`
}

func (u *UrlShortenerDto) ResponseDto(p *UrlShortener) UrlShortenerDto {
	lastOpenedAt := checkDate(p.LastOpenedAt)
	shortenedAt := checkDate(p.ShortenedAt)

	return UrlShortenerDto{
		Id:           p.ID,
		Slug:         p.Slug,
		Url:          p.Url,
		Opened:       p.Opened,
		OpenedTimes:  p.OpenedTimes,
		LastOpenedAt: lastOpenedAt,
		ShortenedAt:  shortenedAt,
	}
}
