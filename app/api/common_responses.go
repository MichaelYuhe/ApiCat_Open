package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/apicat/apicat/commom/spec"
	"github.com/apicat/apicat/commom/translator"
	"github.com/apicat/apicat/models"
	"github.com/gin-gonic/gin"
)

type CommonResponsesID struct {
	CommonResponsesID uint `uri:"response-id" binding:"required,gt=0"`
}

func (cr *CommonResponsesID) CheckCommonResponses(ctx *gin.Context) (*models.CommonResponses, error) {
	if err := translator.ValiadteTransErr(ctx, ctx.ShouldBindUri(&cr)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "GlobalParameters.NotFound"}),
		})
		return nil, err
	}

	definitionsResponses, err := models.NewCommonResponses(cr.CommonResponsesID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "response not found",
		})
		return nil, err
	}

	return definitionsResponses, nil
}

type ResponseDetailData struct {
	Name        string                 `json:"name" binding:"required,lte=255"`
	Code        int                    `json:"code" binding:"required"`
	Description string                 `json:"description" binding:"required,lte=255"`
	Header      []*HeaderData          `json:"header,omitempty" binding:"omitempty,dive"`
	Content     map[string]spec.Schema `json:"content,omitempty" binding:"required"`
	Ref         string                 `json:"$ref,omitempty" binding:"omitempty,lte=255"`
}

type HeaderData struct {
	Name        string      `json:"name" binding:"required,lte=255"`
	Description string      `json:"description" binding:"omitempty,lte=255"`
	Example     string      `json:"example" binding:"omitempty,lte=255"`
	Default     string      `json:"default" binding:"omitempty,lte=255"`
	Required    bool        `json:"required"`
	Schema      spec.Schema `json:"schema"`
}

func CommonResponsesList(ctx *gin.Context) {
	currentProject, _ := ctx.Get("CurrentProject")
	project, _ := currentProject.(*models.Projects)

	definitionsResponses, _ := models.NewCommonResponses()
	definitionsResponses.ProjectID = project.ID
	definitionsResponsesList, err := definitionsResponses.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := []map[string]interface{}{}
	for _, v := range definitionsResponsesList {
		result = append(result, map[string]interface{}{
			"id":          v.ID,
			"code":        v.Code,
			"description": v.Description,
			"name":        v.Name,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func CommonResponsesDetail(ctx *gin.Context) {
	cr := CommonResponsesID{}
	definitionsResponses, err := cr.CheckCommonResponses(ctx)
	if err != nil {
		return
	}

	header := []*HeaderData{}
	if err := json.Unmarshal([]byte(definitionsResponses.Header), &header); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	content := map[string]spec.Schema{}
	if err := json.Unmarshal([]byte(definitionsResponses.Content), &content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":          definitionsResponses.ID,
		"name":        definitionsResponses.Name,
		"code":        definitionsResponses.Code,
		"description": definitionsResponses.Description,
		"header":      header,
		"content":     content,
	})
}

func CommonResponsesCreate(ctx *gin.Context) {
	currentProject, _ := ctx.Get("CurrentProject")
	project, _ := currentProject.(*models.Projects)

	data := ResponseDetailData{}
	if err := translator.ValiadteTransErr(ctx, ctx.ShouldBindJSON(&data)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	definitionsResponses, _ := models.NewCommonResponses()
	definitionsResponses.ProjectID = project.ID
	definitionsResponses.Name = data.Name

	count, err := definitionsResponses.GetCountByName()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "CommonResponses.NameExists"}),
		})
		return
	}

	definitionsResponses.Code = data.Code
	definitionsResponses.Description = data.Description

	responseHeader := make([]*HeaderData, 0)
	if len(data.Header) > 0 {
		responseHeader = data.Header
	}

	header, err := json.Marshal(responseHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	definitionsResponses.Header = string(header)

	content, err := json.Marshal(data.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	definitionsResponses.Content = string(content)

	if err := definitionsResponses.Create(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":          definitionsResponses.ID,
		"name":        definitionsResponses.Name,
		"code":        definitionsResponses.Code,
		"description": definitionsResponses.Description,
		"header":      data.Header,
		"content":     data.Content,
	})
}

func CommonResponsesUpdate(ctx *gin.Context) {
	data := ResponseDetailData{}
	if err := translator.ValiadteTransErr(ctx, ctx.ShouldBindJSON(&data)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	cr := CommonResponsesID{}
	definitionsResponses, err := cr.CheckCommonResponses(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	definitionsResponses.Name = data.Name
	count, err := definitionsResponses.GetCountExcludeTheID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "CommonResponses.NameExists"}),
		})
		return
	}

	definitionsResponses.Code = data.Code
	definitionsResponses.Description = data.Description

	header, err := json.Marshal(data.Header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	definitionsResponses.Header = string(header)

	content, err := json.Marshal(data.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	definitionsResponses.Content = string(content)

	if err := definitionsResponses.Update(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func CommonResponsesDelete(ctx *gin.Context) {
	currentProject, _ := ctx.Get("CurrentProject")
	project, _ := currentProject.(*models.Projects)

	cr := CommonResponsesID{}
	definitionsResponses, err := cr.CheckCommonResponses(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	header := []*HeaderData{}
	if err := json.Unmarshal([]byte(definitionsResponses.Header), &header); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	content := map[string]spec.Schema{}
	if err := json.Unmarshal([]byte(definitionsResponses.Content), &content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	responseDetail := ResponseDetailData{
		Name:        definitionsResponses.Name,
		Code:        definitionsResponses.Code,
		Description: definitionsResponses.Description,
		Header:      header,
		Content:     content,
	}

	collections, _ := models.NewCollections()
	collections.ProjectId = project.ID
	collectionList, err := collections.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ref := "{$ref:#/commons/responses/" + strconv.FormatUint(uint64(definitionsResponses.ID), 10) + "}"
	responseDetailJson, err := json.Marshal(responseDetail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, collection := range collectionList {
		if collection.Type == "http" {
			if strings.Contains(collection.Content, ref) {
				newContent := strings.Replace(collection.Content, ref, string(responseDetailJson), -1)
				collection.Content = newContent
				if err := collection.Update(); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
					})
					return
				}
			}
		}
	}

	if err := definitionsResponses.Delete(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
