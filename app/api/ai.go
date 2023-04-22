package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apicat/apicat/commom/openai"
	"github.com/apicat/apicat/commom/spec/plugin/openapi"
	"github.com/apicat/apicat/commom/translator"
	"github.com/apicat/apicat/config"
	"github.com/apicat/apicat/models"
	"github.com/gin-gonic/gin"
)

type AICreateCollectionStructure struct {
	ParentID uint   `json:"parent_id" binding:"gte=0"`        // 父级id
	Title    string `json:"title" binding:"required,lte=255"` // 名称
	SchemaID uint   `json:"schema_id" binding:"gte=0"`        // 模型id
}

func AICreateCollection(ctx *gin.Context) {
	var (
		openapiContent string
		err            error
	)

	data := &AICreateCollectionStructure{}
	if err := translator.ValiadteTransErr(ctx, ctx.ShouldBindJSON(data)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if data.SchemaID > 0 {
		schema, err := models.NewDefinitions(data.SchemaID)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": translator.Trasnlate(ctx, &translator.TT{ID: "Definitions.NotFound"}),
			})
			return
		}

		o := openai.NewOpenAI(config.SysConfig.OpenAI.Token, "zh")
		o.SetMaxTokens(2000)
		openapiContent, err = o.CreateApiBySchema(data.Title, schema.Schema)
	} else {
		// openapi := openapi.NewOpenAI(config.SysConfig.OpenAI.Token, ctx.GetHeader("Accept-Language"))
		o := openai.NewOpenAI(config.SysConfig.OpenAI.Token, "zh")
		o.SetMaxTokens(2000)
		openapiContent, err = o.CreateApi(data.Title)
		if err != nil || openapiContent == "" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": translator.Trasnlate(ctx, &translator.TT{ID: "AI.CollectionCreateFail"}),
			})
			return
		}
	}

	content, err := openapi.Decode([]byte(openapiContent))
	fmt.Println(openapiContent)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "AI.CollectionCreateFail"}),
		})
		return
	}

	byteContent, err := json.Marshal(content.Collections)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "AI.CollectionCreateFail"}),
		})
		return
	}

	currentProject, _ := ctx.Get("CurrentProject")

	collection, _ := models.NewCollections()
	collection.ProjectId = currentProject.(*models.Projects).ID
	collection.ParentId = data.ParentID
	collection.Title = data.Title
	collection.Type = "http"
	collection.Content = string(byteContent)
	if err := collection.Create(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": translator.Trasnlate(ctx, &translator.TT{ID: "AI.CollectionCreateFail"}),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         collection.ID,
		"parent_id":  collection.ParentId,
		"title":      collection.Title,
		"type":       collection.Type,
		"content":    collection.Content,
		"created_at": collection.CreatedAt.Format("2006-01-02 15:04:05"),
		"created_by": collection.Creator(),
		"updated_at": collection.UpdatedAt.Format("2006-01-02 15:04:05"),
		"updated_by": collection.Updater(),
	})
}
