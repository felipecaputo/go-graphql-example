package main

import "github.com/zebresel-com/mongodm"

//User represents anyone registered on this blog engine
type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	ID       int64  `json:"id,omitempty" bson:"id"`
	Username string `json:"username,omitempty" bson:"username"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Bio      string `json:"bio,omitempty" bson:"bio,omitempty"`
}

//Post represents a blog post writen by some user
type Post struct {
	ID           int64      `json:"id,omitempty"`
	AuthourID    int64      `json:"authour_id,omitempty"`
	Author       *User      `json:"author,omitempty"`
	Title        string     `json:"title,omitempty"`
	Excerpt      string     `json:"excerpt,omitempty"`
	Description  string     `json:"description,omitempty"`
	MainImageURL string     `json:"main_image_url,omitempty"`
	RawBody      string     `json:"raw_body,omitempty"`
	RenderedBody string     `json:"rendered_body,omitempty"`
	SharesCount  int64      `json:"shares_count,omitempty"`
	LikesCount   int64      `json:"likes_count,omitempty"`
	CommentsIDs  []int64    `json:"comments_i_ds,omitempty"`
	Comments     []*Comment `json:"comments,omitempty"`
}

//Comment representes any comment made by any user in any blog post
type Comment struct {
	ID              string   `json:"id,omitempty"`
	AuthorID        int64    `json:"author_id,omitempty"`
	Author          *User    `json:"author,omitempty"`
	PostID          int64    `json:"post_id,omitempty"`
	Post            *Post    `json:"post,omitempty"`
	ParentCommentID int64    `json:"parent_comment_id,omitempty"`
	ParentComment   *Comment `json:"parent_comment,omitempty"`
}
