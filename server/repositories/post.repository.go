package repositories

import (
	"database/sql"
	"fmt"

	"github.com/kimcodell/SH-MH-diary/server/database"
	"github.com/kimcodell/SH-MH-diary/server/types"
	"github.com/kimcodell/SH-MH-diary/server/utils"
)

func GetAllPosts() types.Posts {
	db := database.GetConnectedDB()
	err := db.Ping()
	utils.CatchError(utils.ErrorParams{Err: err})

	rows, queryError := db.Query("SELECT id, title, created_at FROM post WHERE is_deleted = 0")
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to execute SQL Query."})

	var posts types.Posts
	for rows.Next() {
		var post types.SimplePost
		if scanRowError := rows.Scan(&post.Id, &post.Title, &post.CreatedAt); scanRowError != nil {
			utils.CatchError(utils.ErrorParams{Err: scanRowError, Message: "Fail to scan row."})
		}
		posts = append(posts, post)
	}
	queryRowsError := rows.Err()
	utils.CatchError(utils.ErrorParams{Err: queryRowsError})

	defer db.Close()
	defer rows.Close()

	return posts
}

func GetPostById(id int) types.Post {
	db := database.GetConnectedDB()
	err := db.Ping()
	utils.CatchError(utils.ErrorParams{Err: err})

	var post types.Post

	scanError := db.QueryRow(
		`
			SELECT p.id, p.title, p.content, u.name, p.created_at, p.updated_at
			FROM post p join user u on p.user_id = u.id
			WHERE p.id = ? AND p.is_deleted = 0;
		`,
		id,
	).Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.UpdatedAt)
	if scanError == sql.ErrNoRows {
		fmt.Println(fmt.Errorf("err : %v\nmessage: There are no SQL query result", scanError))
		panic(scanError)
	}
	utils.CatchError(utils.ErrorParams{Err: scanError, Message: "Fail to scan row"})

	defer db.Close()

	return post
}

func CreatePost() {

}
