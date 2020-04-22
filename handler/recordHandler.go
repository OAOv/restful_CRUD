package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/OAOv/restful_CRUD/service"
	"github.com/OAOv/restful_CRUD/types"
	"github.com/gin-gonic/gin"
)

type RecordAPI struct {
	recordService service.Service
}

func (r *RecordAPI) CreateRecord(c *gin.Context) {
	record := types.Record{}
	c.BindJSON(&record)
	if record.UserID == "" || record.Subject == "" || record.Score == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrEmptyInput.Error(),
		})
		return
	} else if val, err := strconv.ParseInt(record.UserID, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	}

	err = r.recordService.CreateRecord(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "create completed",
	})
}

func (r *RecordAPI) GetRecords(c *gin.Context) {
	records, err := r.recordService.GetRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    records,
		"message": "readAll completed",
	})
}

func (r *RecordAPI) GetRecord(c *gin.Context) {
	id := c.Param("id")
	if val, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	}

	record, err := r.recordService.GetRecord(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	var records []types.Record
	records = append(records, record)
	c.JSON(http.StatusOK, gin.H{
		"data":    records,
		"message": "readOne completed",
	})
}

func (r *RecordAPI) GetRecordByUser(c *gin.Context) {
	id := c.Param("id")
	if val, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	}

	records, err := r.recordService.GetRecordByUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    records,
		"message": "readOne completed",
	})
}

func (r *RecordAPI) UpdateRecord(c *gin.Context) {
	record := types.Record{}
	c.BindJSON(&record)
	if val, err := strconv.ParseInt(record.ID, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	} else if record.UserID == "" || record.Subject == "" || record.Score == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrEmptyInput.Error(),
		})
		return
	}

	err = r.recordService.UpdateRecord(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "update completed",
	})
}

func (r *RecordAPI) DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	if val, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	}

	err = r.recordService.DeleteRecord(id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "delete completed",
	})
}
