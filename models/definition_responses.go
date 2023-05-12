package models

import (
	"time"
)

type DefinitionResponses struct {
	ID           uint   `gorm:"type:integer primary key autoincrement"`
	ProjectID    uint   `gorm:"type:integer;index;not null;comment:项目id"`
	Name         string `gorm:"type:varchar(255);not null;comment:响应名称"`
	Description  string `gorm:"type:varchar(255);not null;comment:状态描述"`
	Header       string `gorm:"type:mediumtext;comment:响应头"`
	Content      string `gorm:"type:mediumtext;comment:响应内容"`
	DisplayOrder int    `gorm:"type:int(11);not null;default:0;comment:显示顺序"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewDefinitionResponses(ids ...uint) (*DefinitionResponses, error) {
	definitionResponses := &DefinitionResponses{}
	if len(ids) > 0 {
		if err := Conn.Take(definitionResponses, ids[0]).Error; err != nil {
			return definitionResponses, err
		}
		return definitionResponses, nil
	}
	return definitionResponses, nil
}

func (cr *DefinitionResponses) List() ([]*DefinitionResponses, error) {
	definitionResponsesQuery := Conn.Where("project_id = ?", cr.ProjectID)

	var definitionResponses []*DefinitionResponses
	return definitionResponses, definitionResponsesQuery.Find(&definitionResponses).Error
}

func (cr *DefinitionResponses) GetCountByName() (int64, error) {
	var count int64
	return count, Conn.Model(&DefinitionResponses{}).Where("project_id = ? and name = ?", cr.ProjectID, cr.Name).Count(&count).Error
}

func (cr *DefinitionResponses) GetCountExcludeTheID() (int64, error) {
	var count int64
	return count, Conn.Model(&DefinitionResponses{}).Where("project_id = ? and name = ? and id != ?", cr.ProjectID, cr.Name, cr.ID).Count(&count).Error
}

func (cr *DefinitionResponses) Create() error {
	return Conn.Create(cr).Error
}

func (cr *DefinitionResponses) Update() error {
	return Conn.Save(cr).Error
}

func (cr *DefinitionResponses) Delete() error {
	return Conn.Delete(cr).Error
}

// func DefinitionResponsesImport(projectID uint, responses spec.HTTPResponses) nameToIdMap {
// 	var ResponsesMap nameToIdMap

// 	if responses == nil {
// 		return ResponsesMap
// 	}

// 	for i, response := range responses {
// 		header := ""
// 		if response.Header != nil {
// 			if headerByte, err := json.Marshal(response.Header); err == nil {
// 				header = string(headerByte)
// 			}
// 		}

// 		content := ""
// 		if response.Content != nil {
// 			if contentByte, err := json.Marshal(response.Content); err == nil {
// 				content = string(contentByte)
// 			}
// 		}

// 		record := &DefinitionResponses{
// 			ProjectID:    projectID,
// 			Name:         response.Name,
// 			Description:  response.Description,
// 			Header:       header,
// 			Content:      content,
// 			DisplayOrder: i,
// 		}

// 		if Conn.Create(record).Error == nil {
// 			ResponsesMap[response.Name] = record.ID
// 		}
// 	}

// 	return ResponsesMap
// }

// func DefinitionResponsesExport(projectID uint) spec.HTTPResponses {
// 	var definitionResponses []*DefinitionResponses
// 	var definitions []*DefinitionSchemas
// 	specDefinitionResponses := make(spec.HTTPResponses, 0)

// 	if err := Conn.Where("project_id = ?", projectID).Find(&definitionResponses).Error; err != nil {
// 		return specDefinitionResponses
// 	}
// 	if err := Conn.Where("project_id = ? AND type = ?", projectID, "schema").Find(&definitions).Error; err != nil {
// 		return specDefinitionResponses
// 	}

// 	idToNameMap := make(IdToNameMap)
// 	for _, definition := range definitions {
// 		idToNameMap[definition.ID] = definition.Name
// 	}

// 	for _, commonResponse := range definitionResponses {
// 		commonResponse.Content = util.ReplaceIDToName(commonResponse.Content, idToNameMap, "#/definitions/schemas/")

// 		response := spec.HTTPResponse{}
// 		response.Name = commonResponse.Name
// 		response.Description = commonResponse.Description
// 		json.Unmarshal([]byte(commonResponse.Header), &response.Header)
// 		json.Unmarshal([]byte(commonResponse.Content), &response.Content)

// 		specDefinitionResponses = append(specDefinitionResponses, response)
// 	}

// 	return specDefinitionResponses
// }
