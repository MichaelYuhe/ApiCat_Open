package models

import (
	"time"
)

type GlobalParameters struct {
	ID        uint   `gorm:"type:integer primary key autoincrement"`
	ProjectID uint   `gorm:"type:integer;index;not null;comment:项目id"`
	In        string `gorm:"type:varchar(255);not null;comment:位置:header,cookie,query,path"`
	Name      string `gorm:"type:varchar(255);not null;comment:参数名称"`
	Required  int    `gorm:"type:tinyint(1);not null;comment:是否必传:0-否,1-是"`
	Schema    string `gorm:"type:mediumtext;comment:参数内容"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewGlobalParameters(ids ...uint) (*GlobalParameters, error) {
	globalParameters := &GlobalParameters{}
	if len(ids) > 0 {
		if err := Conn.Take(globalParameters, ids[0]).Error; err != nil {
			return globalParameters, err
		}
		return globalParameters, nil
	}
	return globalParameters, nil
}

func (gp *GlobalParameters) List() ([]*GlobalParameters, error) {
	globalParametersQuery := Conn.Where("project_id = ?", gp.ProjectID)

	var globalParameters []*GlobalParameters
	return globalParameters, globalParametersQuery.Order("id desc").Find(&globalParameters).Error
}

func (gp *GlobalParameters) GetCountByName() (int64, error) {
	var count int64
	return count, Conn.Model(&GlobalParameters{}).Where("project_id = ? and name = ?", gp.ProjectID, gp.Name).Count(&count).Error
}

func (gp *GlobalParameters) GetCountExcludeTheID() (int64, error) {
	var count int64
	return count, Conn.Model(&GlobalParameters{}).Where("project_id = ? and name = ? and id != ?", gp.ProjectID, gp.Name, gp.ID).Count(&count).Error
}

func (gp *GlobalParameters) Create() error {
	return Conn.Create(gp).Error
}

func (gp *GlobalParameters) Update() error {
	return Conn.Save(gp).Error
}
