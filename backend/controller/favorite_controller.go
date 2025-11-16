package controller

import (
	"myapp-backend/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	favoriteService service.FavoriteService
}

func NewFavoriteController(favoriteService service.FavoriteService) *FavoriteController {
	return &FavoriteController{
		favoriteService: favoriteService,
	}
}

func (fc *FavoriteController) RegisterRoutes(e *echo.Echo) {
	favorites := e.Group("/favorites")

	favorites.POST("", fc.AddFavorite)
	favorites.GET("", fc.ListFavorites)
	favorites.DELETE("/:movie_id", fc.RemoveFavorite)
	favorites.GET("/is-favorited/:movie_id", fc.IsFavorited)
}

func getUserIDFromContext(c echo.Context) (int64, bool) {
	v := c.Get("user_id")
	if v == nil {
		return 0, false
	}
	switch t := v.(type) {
	case int64:
		return t, true
	case int:
		return int64(t), true
	case float64:
		return int64(t), true
	case string:
		id, err := strconv.ParseInt(t, 10, 64)
		if err == nil {
			return id, true
		}
	}
	return 0, false
}

func (h *FavoriteController) AddFavorite(c echo.Context) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var body struct {
		MovieID int64 `json:"movie_id"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	if body.MovieID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "movie_id required"})
	}

	fav, err := h.favoriteService.AddFavorite(userID, body.MovieID)
	if err != nil {
		if err == service.ErrAlreadyFavorited {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message":  "already favorited",
				"favorite": fav,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, fav)
}

func (h *FavoriteController) ListFavorites(c echo.Context) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	limitQ := c.QueryParam("limit")
	offsetQ := c.QueryParam("offset")
	limit := 20
	offset := 0

	if limitQ != "" {
		if v, err := strconv.Atoi(limitQ); err == nil && v > 0 {
			limit = v
		}
	}
	if offsetQ != "" {
		if v, err := strconv.Atoi(offsetQ); err == nil && v >= 0 {
			offset = v
		}
	}

	list, err := h.favoriteService.ListFavorites(userID, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func (h *FavoriteController) RemoveFavorite(c echo.Context) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	movieIDStr := c.Param("movie_id")
	movieID, err := strconv.ParseInt(movieIDStr, 10, 64)
	if err != nil || movieID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid movie_id"})
	}

	err = h.favoriteService.RemoveFavorite(userID, movieID)
	if err != nil {
		if err == service.ErrFavoriteNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "favorite not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *FavoriteController) IsFavorited(c echo.Context) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	movieIDStr := c.Param("movie_id")
	movieID, err := strconv.ParseInt(movieIDStr, 10, 64)
	if err != nil || movieID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid movie_id"})
	}

	exists, err := h.favoriteService.IsFavorited(userID, movieID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]bool{"favorited": exists})
}
