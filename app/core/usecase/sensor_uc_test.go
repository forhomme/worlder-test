package usecase

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"worlder-test/app/core/mocks"
	"worlder-test/app/core/models"
)

var (
	sensorInput *models.Sensors
	reqParam    *models.RequestSensorParam
	sensorData  []models.Sensors
)

func setupFail() {
	sensorInput = &models.Sensors{
		Value: "",
		ID1:   "",
		ID2:   "",
	}
	reqParam = &models.RequestSensorParam{
		ID1:              "",
		ID2:              "",
		Duration:         "",
		TimeFrom:         "",
		TimeTo:           "",
		FilterPagination: models.FilterPagination{},
	}
}

func setupSuccess() {
	sensorInput = &models.Sensors{
		Value: "10",
		ID1:   "10",
		ID2:   "10",
	}
	reqParam = &models.RequestSensorParam{
		ID1:              "",
		ID2:              "",
		Duration:         "",
		TimeFrom:         "",
		TimeTo:           "",
		FilterPagination: models.FilterPagination{},
	}
	sensorData = make([]models.Sensors, 0)
	sensorData = append(sensorData, *sensorInput)
}

func TestSensorUsecase_InsertSensor(t *testing.T) {
	Convey("Test insert data sensor", t, func() {
		dbRepo := &mocks.DatabaseOutputPort{}

		Convey("Negative Scenario", func() {
			Convey("Error when insert to db", func() {
				setupFail()
				dbRepo.On("InsertSensor", sensorInput).Return(nil, errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.InsertSensor(sensorInput)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenario", func() {
			Convey("Return response", func() {
				setupSuccess()
				dbRepo.On("InsertSensor", sensorInput).Return(sensorInput, nil).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.InsertSensor(sensorInput)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSensorUsecase_UpdateSensor(t *testing.T) {
	Convey("Test update data sensor", t, func() {
		dbRepo := &mocks.DatabaseOutputPort{}

		Convey("Negative Scenario", func() {
			Convey("Error when get detail", func() {
				setupFail()
				dbRepo.On("GetSensorByID", "id").Return(nil, errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.UpdateSensor("id", sensorInput)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("Error when update", func() {
				setupFail()
				dbRepo.On("GetSensorByID", "id").Return(sensorInput, nil).Once()
				dbRepo.On("UpdateSensor", "id", sensorInput).Return(nil, errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.UpdateSensor("id", sensorInput)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive scenario", func() {
			Convey("Return response", func() {
				setupSuccess()
				dbRepo.On("GetSensorByID", "id").Return(sensorInput, nil).Once()
				dbRepo.On("UpdateSensor", "id", sensorInput).Return(sensorInput, nil).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.UpdateSensor("id", sensorInput)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSensorUsecase_GetSensorByID(t *testing.T) {
	Convey("Test get data sensor", t, func() {
		dbRepo := &mocks.DatabaseOutputPort{}

		Convey("Negative Scenario", func() {
			Convey("Error when get detail", func() {
				dbRepo.On("GetSensorByID", "id").Return(nil, errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.GetSensorByID("id")
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive scenario", func() {
			Convey("Return response", func() {
				setupSuccess()
				dbRepo.On("GetSensorByID", "id").Return(sensorInput, nil).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.GetSensorByID("id")
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSensorUsecase_ListSensors(t *testing.T) {
	Convey("Test list data sensor", t, func() {
		dbRepo := &mocks.DatabaseOutputPort{}

		Convey("Negative Scenario", func() {
			Convey("Error when list data", func() {
				dbRepo.On("ListSensors", reqParam).Return(nil, int64(0), errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.ListSensors(reqParam)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive scenario", func() {
			Convey("Return response", func() {
				setupSuccess()
				dbRepo.On("ListSensors", reqParam).Return(sensorData, int64(0), nil).Once()
				uc := NewSensorUsecase(dbRepo)
				resp, err := uc.ListSensors(reqParam)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSensorUsecase_DeleteSensor(t *testing.T) {
	Convey("Test delete data sensor", t, func() {
		dbRepo := &mocks.DatabaseOutputPort{}

		Convey("Negative Scenario", func() {
			Convey("Error when get data", func() {
				dbRepo.On("GetSensorByID", "id").Return(nil, errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				err := uc.DeleteSensor("id")
				So(err, ShouldNotBeNil)
			})
			Convey("Error when delete data", func() {
				setupFail()
				dbRepo.On("GetSensorByID", "id").Return(sensorInput, nil).Once()
				dbRepo.On("DeleteSensor", "id").Return(errors.New("error database")).Once()
				uc := NewSensorUsecase(dbRepo)
				err := uc.DeleteSensor("id")
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive scenario", func() {
			Convey("Return response", func() {
				setupSuccess()
				dbRepo.On("GetSensorByID", "id").Return(sensorInput, nil).Once()
				dbRepo.On("DeleteSensor", "id").Return(nil).Once()
				uc := NewSensorUsecase(dbRepo)
				err := uc.DeleteSensor("id")
				So(err, ShouldBeNil)
			})
		})
	})
}
