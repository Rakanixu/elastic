package handler

import (
	elasticseach "github.com/Rakanixu/elasticsearch/srv/proto/elasticsearch"
	"github.com/micro/go-micro/errors"
)

// DocRefFieldsExists returns an error if DocRef struct has zero value
func DocRefFieldsExists(dr *elasticseach.DocRef) error {
	if len(dr.Index) <= 0 {
		return errors.BadRequest("go.micro.srv.elasticsearch", "Index required")
	}

	if len(dr.Type) <= 0 {
		return errors.BadRequest("go.micro.srv.elasticsearch", "Type required")
	}

	return nil
}
