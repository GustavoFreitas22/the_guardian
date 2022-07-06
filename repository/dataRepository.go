package repository

import (
	"log"

	"github.com/GustavoFreitas22/the_guardian/config"
	"github.com/GustavoFreitas22/the_guardian/cript"
	"github.com/GustavoFreitas22/the_guardian/model"
)

var (
	data model.Data

	secrets []*model.Data
)

func InsertDataOnDatabase(data *model.Data) {
	config.Database.Create(&data)
}

func GetDataById(id int) model.Data {
	config.Database.First(&data, id)
	return data
}

func GetDataByContext(context string) []*model.Data {
	config.Database.Where(model.Data{Context: context}).Find(&secrets)
	return secrets
}

func EditData(dataEdit model.Data, id int) model.Data {
	config.Database.First(&data, id)

	if dataEdit.ID == 0 {
		log.Println("Obejto não encontrado")
		return model.Data{}
	}

	data.Context = dataEdit.Context
	data.Key = dataEdit.Key
	data.Value = dataEdit.Value

	encryptData, err := cript.EncriptDataInfo(data)
	if err != nil {
		log.Println(err)
		return model.Data{}
	}

	config.Database.Save(&encryptData)

	return data
}

func DeleteDataById(id int) error {
	err := config.Database.Delete(&data, id)
	if err != nil {
		log.Println(err)
		return err.Error
	}
	return err.Error
}
