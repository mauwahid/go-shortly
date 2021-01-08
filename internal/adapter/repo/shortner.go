package repo

type Shortener interface {
	Save(longUrl string, shortUrl string) error
	GetByLongUrl(longUrl string) (string, error)
	GetByShortUrl(shortUrl string) (string, error)
}
