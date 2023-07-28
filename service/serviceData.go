package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/GustavoFreitas22/the_guardian/cript"
	"github.com/GustavoFreitas22/the_guardian/model"
)

func EncriptDataAndInsert(data *model.Data) error {
	encryptData, err := cript.EncriptDataInfo(*data)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(encryptData)

	return err
}

func DesencryptAndEditData(data *model.Data) model.Data {

	dataEdit := model.Data{}

	if dataEdit.ID == 0 {
		log.Println("Obejto não encontrado")
		return model.Data{}
	}

	return dataEdit
}

func GetInfoById(id string) (model.Data, error) {

	id_converted, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, id_converted)
		return model.Data{}, err
	}

	dataInfo := model.Data{}

	dataDecrypt, err := cript.DecryptDataInfo(dataInfo)
	if err != nil {
		log.Println(err)
		return dataInfo, err
	}

	return dataDecrypt, err
}

func GetSecretByContext(context string) []*model.Data {
	var secretDecrypt []*model.Data

	// for _, secret := range secrets {
	// 	dataDecrypt, err := cript.DecryptDataInfo(*secret)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return []*model.Data{}
	// 	}
	// 	secretDecrypt = append(secretDecrypt, &dataDecrypt)
	// }

	return secretDecrypt
}

func DeleteSecretById(id string) {

	id_converted, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(id_converted)
}
