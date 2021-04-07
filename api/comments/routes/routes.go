package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonufrienko/academy/api/comments/comments"
)

func GetComments(c *fiber.Ctx) error {
	query := new(comments.CommentsQuery)
	c.QueryParser(query)

	comments, err := comments.GetComments(query)
	if err != nil {
		return err
	}

	c.JSON(comments)
	return nil
}

func AddComments(c *fiber.Ctx) error {
	data := new(comments.AddCommentData)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	comment, err := comments.AddComment(data)
	if err != nil {
		return err
	}

	c.JSON(comment)
	return nil
}

func DeleteComments(c *fiber.Ctx) error {
	commentId := c.Params("id")

	err := comments.DeleteComment(commentId)
	if err != nil {
		return err
	}

	c.SendStatus(200)
	return nil
}
