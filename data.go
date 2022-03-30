package ami

import (
	"encoding/json"
	"fmt"
)

type IdData struct {
	TitleCode             int    `json:"titleCode"`
	TitleDesc             string `json:"titleDesc"`
	TitleName             string `json:"titleName"`
	TitleSex              int    `json:"titleSex"`
	FirstName             string `json:"firstName"`
	MiddleName            string `json:"middleName"`
	LastName              string `json:"lastName"`
	GenderCode            int    `json:"genderCode"`
	GenderDesc            string `json:"genderDesc"`
	DateOfBirth           int    `json:"dateOfBirth"`
	NationalityCode       int    `json:"nationalityCode"`
	NationalityDesc       string `json:"nationalityDesc"`
	OwnerStatusDesc       string `json:"ownerStatusDesc"`
	StatusOfPersonCode    int    `json:"statusOfPersonCode"`
	StatusOfPersonDesc    string `json:"statusOfPersonDesc"`
	DateOfMoveIn          int    `json:"dateOfMoveIn"`
	Age                   int    `json:"age"`
	FatherPersonalID      int64  `json:"fatherPersonalID"`
	FatherName            string `json:"fatherName"`
	FatherNationalityCode int    `json:"fatherNationalityCode"`
	FatherNational        string `json:"fatherNational"`
	MotherPersonalID      int64  `json:"motherPersonalID"`
	MotherName            string `json:"motherName"`
	MotherNationalityCode int    `json:"motherNationalityCode"`
	MotherNational        string `json:"motherNational"`
	FullnameAndRank       string `json:"fullnameAndRank"`
	EnglishTitleDesc      string `json:"englishTitleDesc"`
	EnglishFirstName      string `json:"englishFirstName"`
	EnglishMiddleName     string `json:"englishMiddleName"`
	EnglishLastName       string `json:"englishLastName"`
}

func BindIdData(input string) (data *IdData, err error) {
	data = &IdData{}
	err = json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Printf("BindIdData err: %v\n", err.Error())
		return
	}
	return
}

type HomeData struct {
	HouseID         int64  `json:"houseID"`
	HouseNo         string `json:"houseNo"`
	HouseType       int    `json:"houseType"`
	HouseTypeDesc   string `json:"houseTypeDesc"`
	VillageNo       int    `json:"villageNo"`
	AlleyWayCode    int    `json:"alleyWayCode"`
	AlleyWayDesc    string `json:"alleyWayDesc"`
	AlleyCode       int    `json:"alleyCode"`
	AlleyDesc       string `json:"alleyDesc"`
	RoadCode        int    `json:"roadCode"`
	RoadDesc        string `json:"roadDesc"`
	SubdistrictCode int    `json:"subdistrictCode"`
	SubdistrictDesc string `json:"subdistrictDesc"`
	DistrictCode    int    `json:"districtCode"`
	DistrictDesc    string `json:"districtDesc"`
	ProvinceCode    int    `json:"provinceCode"`
	ProvinceDesc    string `json:"provinceDesc"`
	RcodeCode       string `json:"rcodeCode"`
	RcodeDesc       string `json:"rcodeDesc"`
	DateOfTerminate int    `json:"dateOfTerminate"`
}

func BindHomeData(input string) (data *HomeData, err error) {
	data = &HomeData{}
	err = json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Printf("BindHomeData err: %v\n", err.Error())
		return
	}
	return
}
