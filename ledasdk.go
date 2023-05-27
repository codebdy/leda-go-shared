package ledasdk

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify/db"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/codebdy/leda-service-sdk/consts"
	"github.com/codebdy/leda-service-sdk/system"
	"github.com/mitchellh/mapstructure"
)

type App struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Meta struct {
	Content meta.UMLMeta `json:"content"`
}

type AppJson struct {
	App  App  `json:"app"`
	Meta Meta `json:"meta"`
}

type Service struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type ServiceJson struct {
	Service Service `json:"service"`
	Meta    Meta    `json:"meta"`
}

func GetAppMata(appName string, dbConfig db.DbConfig) (*system.Meta, error) {
	return nil, nil
}

func GetServiceMata(serviceName string, dbConfig db.DbConfig) (*system.Meta, error) {
	repo := entify.New(dbConfig)
	repo.Init(*system.SystemMeta, 0)
	session, err := repo.OpenSession()
	if err != nil {
		return nil, err
	}

	data := session.QueryOne(consts.SERVICE_ENTITY_NAME,
		graph.QueryArg{
			shared.ARG_WHERE: graph.QueryArg{
				"name": graph.QueryArg{
					shared.ARG_EQ: serviceName,
				},
			},
		},
	)

	if data == nil {
		panic("Can not find service")
	}
	metaId := data.(map[string]interface{})["metaId"]
	if metaId == 0 {
		panic("Can not find metaId")
	}

	metaData := session.QueryOne(consts.META_ENTITY_NAME,
		graph.QueryArg{
			shared.ARG_WHERE: graph.QueryArg{
				shared.ID_NAME: graph.QueryArg{
					shared.ARG_EQ: metaId,
				},
			},
		},
	)

	if metaData == nil {
		panic("Can not find meta")
	}
	var m system.Meta
	if err := mapstructure.Decode(metaData, &m); err != nil {
		panic(err.Error())
	}
	return &m, nil
}

func GetAppMataById(appId shared.ID, dbConfig db.DbConfig) (*system.Meta, error) {
	return nil, nil
}

func GetServiceMataById(serviceId shared.ID, dbConfig db.DbConfig) (*system.Meta, error) {
	return nil, nil
}

func ReadContentFromJson(fileName string) meta.UMLMeta {
	data, err := ioutil.ReadFile(fileName)
	content := meta.UMLMeta{}
	if nil != err {
		log.Panic(err.Error())
	} else {
		err = json.Unmarshal(data, &content)
	}

	return content
}

func ReadAppFromJson(fileName string) ServiceJson {
	data, err := ioutil.ReadFile(fileName)
	service := ServiceJson{}
	if nil != err {
		log.Panic(err.Error())
	} else {
		err = json.Unmarshal(data, &service)
	}

	return service
}

func ReadServieFromJson(fileName string) AppJson {
	data, err := ioutil.ReadFile(fileName)
	app := AppJson{}
	if nil != err {
		log.Panic(err.Error())
	} else {
		err = json.Unmarshal(data, &app)
	}

	return app
}
