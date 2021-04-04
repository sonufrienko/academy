package comments

import (
	"fmt"
	"strconv"
	"time"
)

type Comment struct {
	Id     string  `json:"id" form:"id"`
	Author string  `json:"author" form:"author"`
	Text   string  `json:"text" form:"text"`
	Date   string  `json:"date" form:"date"`
	Rating float64 `json:"rating" form:"rating"`
}

type AddCommentData struct {
	CourseId string  `json:"courseId"`
	UserId   string  `json:"userId"`
	Text     string  `json:"text"`
	Rating   float64 `json:"rating"`
}

func GetComments(courseId string) ([]Comment, error) {
	comments := make([]Comment, 0)

	for i := 1; i < 11; i++ {
		comments = append(comments, Comment{
			Id:     strconv.Itoa(i),
			Date:   time.Now().Format(time.RFC3339),
			Author: fmt.Sprintf("Author #%v", i),
			Text:   fmt.Sprintf("Comment #%v.", i),
			Rating: 4,
		})
	}

	return comments, nil
}

func AddComment(data AddCommentData) (Comment, error) {
	comment := Comment{
		Id:     "11-22-33",
		Date:   time.Now().Format(time.RFC3339),
		Author: data.UserId,
		Text:   data.Text,
		Rating: data.Rating,
	}

	return comment, nil
}

func DeleteComment(commentId string) error {
	return nil
}
