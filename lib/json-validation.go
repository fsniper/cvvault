package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	embedcontent "github.com/fsniper/cvvault/emb"
	"github.com/qri-io/jsonschema"
)

func JsonValidate(schemaName string, data []byte) {
	log.Println("Validating json")
	schemaData, err := embedcontent.EmbeddedContent.ReadFile(schemaName)
	if err != nil {
		fmt.Printf("Error reading embedded file: %v\n", err)
		return
	}

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	errs, err := rs.ValidateBytes(context.Background(), data)
	if err != nil {
		log.Fatal("Error validating data ", err)
	}
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println("Validation Error: ", e.Error())
		}
		log.Fatal("Exiting for validation errors in Project.Basics")
	}
}
