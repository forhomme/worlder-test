package database

import (
	"fmt"
	"gorm.io/gorm"
	"worlder-test/app/core/models"
	"worlder-test/app/core/ports"
	"worlder-test/app/utils"
)

type dbConnection struct {
	db *gorm.DB
}

func NewDbConnection(db *gorm.DB) ports.DatabaseOutputPort {
	return &dbConnection{db: db}
}

func (d *dbConnection) ListSensors(req *models.RequestSensorParam) (resp []models.Sensors, count int64, err error) {
	query := d.db.Model(&models.Sensors{})
	if req.ID1 != "" && req.ID2 != "" {
		query = query.Where("id1 = ? and id2 = ?", req.ID1, req.ID2)
	}
	if req.Duration != "" {
		q := fmt.Sprintf(" TIMESTAMPDIFF(SECOND, created_at, CONVERT_TZ(now(), '+00:00','+07:00')) < %s", req.Duration)
		query = query.Where(q)
	}
	if req.TimeFrom != "" && req.TimeTo != "" {
		query = query.Where("created_at between ? and ?", req.TimeFrom, req.TimeTo)
	}

	if err = query.Count(&count).Error; err != nil {
		err = fmt.Errorf("Failed to count data sensors : %s", err.Error())
		return
	}

	limit, offset := utils.TranslatePagination(req.Page, req.PerPage)
	query = query.Limit(limit).Offset(offset)

	_ = OrderAndSort(query, req.OrderBy, req.SortType)

	err = query.Find(&resp).Error
	return
}

func (d *dbConnection) GetSensorByID(id string) (resp *models.Sensors, err error) {
	err = d.db.Where("id = ?", id).Find(&resp).Error
	return
}

func (d *dbConnection) InsertSensor(in *models.Sensors) (resp *models.Sensors, err error) {
	err = d.db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (d *dbConnection) UpdateSensor(id string, in *models.Sensors) (resp *models.Sensors, err error) {
	err = d.db.Where("id = ?", id).Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (d *dbConnection) DeleteSensor(id string) (err error) {
	err = d.db.Delete(&models.Sensors{}, id).Error
	return
}

func OrderAndSort(query *gorm.DB, orderBy, sortType string) (err error) {
	if orderBy != "" {
		if sortType != "" {
			query = query.Order(orderBy + " " + sortType)
		} else {
			query = query.Order(orderBy + " asc")
		}
	}
	return
}
