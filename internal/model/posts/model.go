package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashtags string    `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)
type (
	PostResponse struct {
		ID           int64     `json:"id"`
		UserID       int64     `json:"user_id"`
		PostTitle    string    `json:"post_title"`
		PostContent  string    `json:"post_content"`
		PostHashtags []string  `json:"post_hashtags"`
		CreatedAt    time.Time `json:"created_at"`
		IsLiked      bool      `json:"is_liked"`
		IsSaved      bool      `json:"is_saved"`
		UpdatedAt    time.Time `json:"updated_at"`
		CreatedBy    string    `json:"created_by"`
		UpdatedBy    string    `json:"updated_by"`
	}

	Post struct {
		PostDetail PostResponse      `json:"post_detail"`
		Comments   []CommentResponse `json:"comments"`
		LikeCount  int64             `json:"like_count"`
	}
)
