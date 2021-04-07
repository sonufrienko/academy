package comments

import (
	"context"
	"strings"
	"time"

	"github.com/sonufrienko/academy/api/comments/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	CourseId  primitive.ObjectID `json:"courseId" bson:"courseId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Text      string             `json:"text" bson:"text"`
	Rating    float64            `json:"rating" bson:"rating"`
}

type AddCommentData struct {
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	CourseId  primitive.ObjectID `json:"courseId" bson:"courseId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Text      string             `json:"text" bson:"text"`
	Rating    float64            `json:"rating" bson:"rating"`
}

type CommentsQuery struct {
	Limit  int64  `query:"limit"`
	Skip   int64  `query:"skip"`
	Sort   string `query:"sort"`
	Course string `query:"course"`
}

func getCommentsFilter(query *CommentsQuery) primitive.M {
	filter := make(map[string]interface{})

	courseIdObjectID, err := primitive.ObjectIDFromHex(query.Course)
	if err == nil {
		filter["courseId"] = courseIdObjectID
	}

	return filter
}

func getCommentsOptions(query *CommentsQuery) *options.FindOptions {
	sortBy := query.Sort
	sortOrder := 1

	if strings.HasPrefix(query.Sort, "-") {
		sortBy = strings.Trim(query.Sort, "-")
		sortOrder = -1
	}

	opts := options.Find()
	if sortBy != "" {
		opts.SetSort(bson.M{sortBy: sortOrder})
	}

	opts.SetLimit(query.Limit)
	opts.SetSkip(query.Skip)

	return opts
}

func GetComments(query *CommentsQuery) ([]Comment, error) {
	var comments []Comment

	filter := getCommentsFilter(query)
	opts := getCommentsOptions(query)
	collection := db.GetCollection("comments")
	cursor, err := collection.Find(context.Background(), filter, opts)

	if err != nil {
		return comments, err
	}

	if err = cursor.All(context.Background(), &comments); err != nil {
		return comments, err
	}

	return comments, nil
}

func AddComment(data *AddCommentData) (Comment, error) {
	data.CreatedAt = time.Now()

	collection := db.GetCollection("comments")
	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return Comment{}, err
	}

	comment, err := GetComment(res.InsertedID)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func GetComment(courseId interface{}) (Comment, error) {
	var comment = new(Comment)
	collection := db.GetCollection("comments")
	if err := collection.FindOne(context.Background(), bson.M{"_id": courseId}).Decode(comment); err != nil {
		return Comment{}, err
	}

	return *comment, nil
}

func DeleteComment(commentId string) error {
	return nil
}
