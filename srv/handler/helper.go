package handler

import (
	elasticseach "github.com/Rakanixu/elasticsearch/srv/proto/elasticsearch"
	"github.com/micro/go-micro/errors"
)

// RequiredRecordFieldsExists returns an error if required fields has zero value
func RequiredRecordFieldsExists(r *elasticseach.Record) error {
	if len(r.Index) <= 0 {
		return errors.BadRequest("go.micro.srv.elasticsearch", "Index required")
	}

	if len(r.Type) <= 0 {
		return errors.BadRequest("go.micro.srv.elasticsearch", "Type required")
	}

	if len(r.Id) <= 0 {
		return errors.BadRequest("go.micro.srv.elasticsearch", "Id required")
	}

	return nil
}
