package handlers

import (
	"meisterwerk/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Storage interface {
	Get(id string) (entities.Event, error)
	List(start, end string) ([]entities.Event, error)
	Delete(id string) error
	Upsert(te entities.Event) error
}

type Controller struct {
	Repository Storage
}

var _ Eventer = &Controller{}

type Eventer interface {
	Get() func(c *gin.Context)
	List() func(c *gin.Context)
	Delete() func(c *gin.Context)
	Update() func(c *gin.Context)
	Create() func(c *gin.Context)
}

func NewEventer(s Storage) Eventer {
	return &Controller{Repository: s}
}

// Get returns the event  object for the given ID
func (h *Controller) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")
		_, err := uuid.FromString(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		event, err := h.Repository.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, event)
	}
}

// List	 returns the events list for params given
func (h *Controller) List() func(c *gin.Context) {
	return func(c *gin.Context) {
		from := c.Request.URL.Query().Get("from")
		to := c.Request.URL.Query().Get("to")

		if err := timeParamsChecker(from, to); err != nil {
			zap.L().Error("time parameters check failed", zap.Error(err))
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		events, err := h.Repository.List(from, to)
		if err != nil {
			zap.L().Error("repository operation failed", zap.Error(err))
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, events)
	}
}

// Create persits the event given
func (h *Controller) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		var event entities.Event
		err := c.Bind(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		event.ID = uuid.NewV4().String()

		err = requiredFields(event)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = h.Repository.Upsert(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, event)
	}
}

// Update updates the event in the repository
func (h *Controller) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		var event entities.Event
		err := c.Bind(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err = uuid.FromString(event.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = requiredFields(event)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = h.Repository.Upsert(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, event)
	}
}

// Delete removes the event from repository
func (h *Controller) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")
		_, err := uuid.FromString(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = h.Repository.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
