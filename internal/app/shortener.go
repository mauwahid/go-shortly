package app

import "github.com/mauwahid/go-shortly/internal/adapter/repo"

type Shortener interface {
	ShortUrl(url string) string
	CallUrl(shortUrl string) string
}

type shortener struct {
	repo repo.Shortener
}

func (s shortener) ShortUrl(url string) (hashUrl string, err error) {

	if hashUrl = s.checkIfExist(url); hashUrl == "" {
		hashUrl = s.hash64(url)
		err = s.repo.Save(url, hashUrl)
	}

	return
}

func (s shortener) CallUrl(shortUrl string) string {

}

func (s shortener) hash64(url string) string {

}

func (s shortener) base64Converter(shortUrl string) string {

}

func (s shortener) checkIfExist(url string) string {

}


