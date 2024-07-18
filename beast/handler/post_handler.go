package handler

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/beast/data"
	"github.com/HsiaoCz/beast-clone/beast/types"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	post data.PostStorer
}

func NewPostHandler(post data.PostStorer) *PostHandler {
	return &PostHandler{
		post: post,
	}
}

func (p *PostHandler) HandleCreatePost(c *fiber.Ctx) error {
	var post_params types.CreatePostParams
	if err := c.BodyParser(&post_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	post, err := p.post.CreatePost(types.NewPostFromParams(post_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create post success",
		"post":    post,
	})
}

func (p *PostHandler) HandleGetPostByID(c *fiber.Ctx) error {
	post_id := c.Query("post_id")
	post, err := p.post.GetPostByID(post_id)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "get post success",
		"post":    post,
	})
}

func (p *PostHandler) HandleDeletePostByID(c *fiber.Ctx) error {
	post_id := c.Query("post_id")
	if err := p.post.DeletePostByID(post_id); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "delete post success",
	})
}
