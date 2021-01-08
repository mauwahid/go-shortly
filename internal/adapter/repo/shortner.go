package repo

type Shortener interface {
	Save(longUrl string) (string, error)
	Get(shortUrl string) (string, error)
}
