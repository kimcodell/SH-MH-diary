package repositories

import (
	"database/sql"
	"fmt"

	"github.com/kimcodell/SH-MH-diary/server/database"
	"github.com/kimcodell/SH-MH-diary/server/types"
	"github.com/kimcodell/SH-MH-diary/server/utils"
)

func GetAllPosts() types.Posts {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

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
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
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
		panic(scanError)
	} else {
		utils.CatchError(utils.ErrorParams{Err: scanError, Message: "Fail to scan row"})
	}

	return post
}

func CreatePost(params types.PostCreateDto) bool {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

	result, queryError := db.Exec(`INSERT INTO post (title, content, userId) VALUES (?, ?, ?)`, params.Title, params.Content, params.UserId)
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to insert new post query"})

	affectedCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if (affectedCount != 1) {
		utils.CatchError(utils.ErrorParams{Err: fmt.Errorf("fail to create the post"), Message: "Fail to create the post"})
		return false
	}

	return true
}

func UpdatePost(params types.PostCreateDto) bool {
	db, dbConnectErr := database.ConnectToDB()
	fmt.Println("Error :", dbConnectErr)
	defer db.Close()

	//TODO

	result, queryError := db.Exec(`UPDATE `,) //TODO 쿼리 추가
	utils.CatchError(utils.ErrorParams{Err: queryError, Message: "Fail to update post query"})

	affectedCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if (affectedCount != 1) {
		utils.CatchError(utils.ErrorParams{Err: fmt.Errorf("fail to update the post"), Message: "Fail to update the post"})
		return false
	}

	return true
}
