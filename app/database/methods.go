package database

import (
	"log"
	"postapi/app/models"
)


func (d *DB) CreatePost(p *models.Post) error {
	res, err := d.db.Exec(insertPostSchema, p.Title, p.Content, p.Author)
	if err != nil {
		log.Println("An Error Has Occured: ", err)
		return err
	}

	res.LastInsertId()
	return err	
}

func (d *DB) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		log.Println("An Error Has Occured: ", err)
		return posts, err
	}

	return posts, nil
}