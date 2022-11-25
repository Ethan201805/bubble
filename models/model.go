package models

import (
	"bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Substance struct {
	Name       string `json:"name"`
	Cas_no     string `json:"cas_no"`
	List_no    string `json:"ec_no"`
	Info_key   string `json:"info_key"`
	Info_value string `json:"info_value"`
	Sid        string `json:"sid"`
	Parent_id  string `json:"parent_id"`
}

type SubstanceInfo struct {
	ID         string `json:"id"`
	Parent_id  string `json:"parent_id"`
	Info_key   string `json:"info_key"`
	Info_value string `json:"info_value"`
}

var AllowedEndpoint = []string{"Melting/freezing point", "Boiling point", "Relative density", "Vapour pressure", "Surface tension", "Water solubility", "Partition coefficient", "pH value", "Dissociation constant", "Henry's constant", "Particle size distribution (Granulometry)", "Viscosity", "Flash point", "Flammability", "Explosivness", "Auto flammability", "Oxidising properties", "Stability in organic solvents and identity of relevant degradation products", "Corrosiveness to metals", "Acute toxicity", "Irritation / corrosion", "Sensitisation", "Mutagenicity/Genotoxicity", "Repeated dose toxicity", "Reproductive/developmental toxicity", "Carcinogenicity", "Acute aquatic toxicity", "Chronic aquatic toxicity", "Terrestrial toxicity", "Sediment toxicity", "Adsorption/Desorption", "Biodegradation", "Stability", "Bioaccumulation"}

//var AllowedEndpoint = []string{"Biodegradation"}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todo []*Todo, err error) {
	err = dao.DB.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo map[string]interface{}, id string) (err error) {
	err = dao.DB.Table("todos").Where("id = ?", id).Updates(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}

func GetEndpointList() (endpoint []*Substance, err error) {
	err = dao.DB.Table("substance").Select("substance.name, substance.cas_no,substance.list_no,substance_info.id as sid,substance_info.parent_id,substance_info.info_key,substance_info.info_value").Joins("left join substance_info on substance_info.substance_id = substance.id").Where("substance.id = ? and substance_info.info_key IN (?) ", 100, AllowedEndpoint).Scan(&endpoint).Error
	if err != nil {
		return nil, err
	}
	return
}

//递归查询
func CheckEndpoint(id string) (flag bool) {
	var substanceinfo = []*SubstanceInfo{}
	var fields = []interface{}{"Description of key information", "Applicant's summary and conclusion", "Key value for chemical safety assessment"}
	_ = dao.DB.Table("substance_info").Select("id, parent_id,info_key,info_value").Where("parent_id = ? ", id).Scan(&substanceinfo).Error

	if len(substanceinfo) > 0 {
		for _, info := range substanceinfo {
			if InArray(info.Info_key, fields) {
				if info.Info_value != "" && !InArray(info.Info_value, fields) {
					return true
				}
			} else {
				if info.Info_value != "" {
					flag = CheckEndpoint(info.ID)
					//如果最终找到结果，则立即返回不继续循环，如果没有则继续循环
					if flag {
						return true
					}
				}
			}
		}

	}
	return
}

func InArray(need interface{}, needArr []interface{}) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}
