package models

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
)

const articlesFile = "data/articles.json"

func LoadArticles() (map[uuid.UUID]Article, error) {
	f, err := os.Open(articlesFile)
	if os.IsNotExist(err) {
		return make(map[uuid.UUID]Article), nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var articles map[uuid.UUID]Article
	err = json.NewDecoder(f).Decode(&articles)
	return articles, err
}

func SaveArticles(articles map[uuid.UUID]Article) error {
	f, err := os.Create(articlesFile)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(articles)
}

func AddArticle(article *Article) error {
	articles, err := LoadArticles()
	if err != nil {
		return err
	}
	if _, ok := articles[article.Id]; ok {
		return errors.New("article with id=" + article.Id.String() + " already exists")
	}
	articles[article.Id] = *article
	return SaveArticles(articles)
}

func UpdateArticle(article *Article) error {
	articles, err := LoadArticles()
	if err != nil {
		return err
	}
	if _, ok := articles[article.Id]; !ok {
		return errors.New("article with id=" + article.Id.String() + " doesn't exist")
	}
	articles[article.Id] = *article
	return SaveArticles(articles)
}

func DeleteArticle(article *Article) error {
	articles, err := LoadArticles()
	if err != nil {
		return err
	}
	if _, ok := articles[article.Id]; !ok {
		return errors.New("article with id=" + article.Id.String() + " doesn't exist")
	}
	delete(articles, article.Id)
	return SaveArticles(articles)
}
