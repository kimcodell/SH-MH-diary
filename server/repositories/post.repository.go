package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kimcodell/SH-MH-diary/server/database"
	"github.com/kimcodell/SH-MH-diary/server/types"
	"github.com/kimcodell/SH-MH-diary/server/utils"
)

func GetAllPosts() (types.Posts, error) {
	db, dbConnectErr := database.ConnectToDB()
	if dbConnectErr != nil {
		fmt.Println("Error :", dbConnectErr)
		return types.Posts{}, dbConnectErr
	}
	defer db.Close()

	rows, queryError := db.Query("SELECT id, title, created_at FROM post WHERE is_deleted = 0")
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to execute SQL Query."})
	defer rows.Close()

	var posts types.Posts
	for rows.Next() {
		var post types.SimplePost
		if scanRowError := rows.Scan(&post.Id, &post.Title, &post.CreatedAt); scanRowError != nil {
			utils.CatchError(utils.ErrorParams{Err: scanRowError, Message: "Fail to scan row."})
			return types.Posts{}, scanRowError
		}
		posts = append(posts, post)
	}
	queryRowsError := rows.Err()
	utils.CatchError(utils.ErrorParams{Err: queryRowsError})

	return posts, nil
}

func GetPostById(id int) (types.Post, error) {
	db, dbConnectErr := database.ConnectToDB()
	if dbConnectErr != nil {
		fmt.Println("Error :", dbConnectErr)
		return types.Post{}, dbConnectErr
	}
	defer db.Close()

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
		return types.Post{}, scanError
	} else if scanError != nil {
		utils.CatchError(utils.ErrorParams{Err: scanError, Message: "Fail to scan row"})
		return types.Post{}, scanError
	}

	return post, nil
}

func CreatePost(params types.PostCreateDto) bool {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

	result, queryError := db.Exec("INSERT INTO post (title, content, user_id) VALUES (?, ?, ?)", params.Title, params.Content, params.UserId)
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to insert new post query"})

	affectedCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if affectedCount != 1 {
		utils.CatchError(utils.ErrorParams{Err: fmt.Errorf("fail to create the post"), Message: "Fail to create the post"})
		return false
	}

	return true
}

func UpdatePost(postId int, params types.PostCreateDto) bool {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

	_, err := GetPostById(postId)
	if err != nil {
		return false
	}

	if params.Title == "" && params.Content == "" && params.UserId == 0 {
		utils.CatchError(utils.ErrorParams{Err: errors.New("no params"), Message: "At least one data is required"})
		return false
	}

	if params.Title != "" {
		_, queryError1 := db.Exec("UPDATE post SET title = ? where id = ?", params.Title, postId)
		utils.CatchError(utils.ErrorParams{Err: queryError1, Message: "Fail to update post query : title"})
	}
	if params.Content != "" {
		_, queryError2 := db.Exec("UPDATE post SET content = ? where id = ?", params.Content, postId)
		utils.CatchError(utils.ErrorParams{Err: queryError2, Message: "Fail to update post query : content"})
	}
	if params.UserId != 0 {
		_, queryError3 := db.Exec("UPDATE post SET user_id = ? where id = ?", params.UserId, postId)
		utils.CatchError(utils.ErrorParams{Err: queryError3, Message: "Fail to update post query : userId"})
	}

	return true
}

func DeletePost(postId int) bool {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

	_, err := GetPostById(postId)
	if err != nil {
		return false
	}

	result, queryError := db.Exec("UPDATE post SET is_deleted = 1 WHERE id = ?", postId)
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to delete post query"})

	affectedCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if affectedCount != 1 {
		utils.CatchError(utils.ErrorParams{Err: fmt.Errorf("fail to create the post"), Message: "Fail to create the post"})
		return false
	}

	return true
}

func GetUsersPosts() ([]types.UserPosts, error) {
	db, dbConnectErr := database.ConnectToDB()
	if dbConnectErr != nil {
		fmt.Println("Error :", dbConnectErr)
		return []types.UserPosts{}, dbConnectErr
	}
	defer db.Close()

	//TODO
	var data []types.UserPosts

	rows, queryError := db.Query("SELECT id, title, created_at FROM post WHERE is_deleted = 0 AND user_id = 1")
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to execute SQL Query."})
	defer rows.Close()

	return data, nil
}
