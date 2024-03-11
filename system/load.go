package system

import (
	"encoding/json"
	"log"

	"github.com/codebdy/entify-core/model/meta"
)

var SystemMeta *meta.UMLMeta

func init() {
	content := meta.UMLMeta{}
	err := json.Unmarshal([]byte(systemSeed), &content)
	if nil != err {
		log.Panic(err.Error())
	}

	SystemMeta = &content

}
