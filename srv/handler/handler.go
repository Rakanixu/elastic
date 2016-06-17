package handler

import (
	"github.com/Rakanixu/elasticsearch/srv/elastic"
	elasticseach "github.com/Rakanixu/elasticsearch/srv/proto/elasticsearch"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

// Elasticsearch struct
type Elasticsearch struct{}

// Create srv handler
func (es *Elasticsearch) Create(ctx context.Context, req *elasticseach.CreateRequest, rsp *elasticseach.CreateResponse) error {
	var err error

	if req.Record == nil {
		return errors.BadRequest("go.micro.srv.elasticsearch.Elasticsearch.Create", "Record required")
	}

	if err = RequiredRecordFieldsExists(req.Record); err != nil {
		return err
	}

	if err = elastic.Create(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Create", err.Error())
	}

	return nil
}

// Read srv handler
func (es *Elasticsearch) Read(ctx context.Context, req *elasticseach.ReadRequest, rsp *elasticseach.ReadResponse) error {
	var err error

	rHelper := &elasticseach.Record{
		Index: req.Index,
		Type:  req.Type,
		Id:    req.Id,
	}
	if err = RequiredRecordFieldsExists(rHelper); err != nil {
		return err
	}

	record, err := elastic.Read(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Read", err.Error())
	}

	rsp.Record = record

	return nil
}

// Update srv handler
func (es *Elasticsearch) Update(ctx context.Context, req *elasticseach.UpdateRequest, rsp *elasticseach.UpdateResponse) error {
	var err error

	if req.Record == nil {
		return errors.BadRequest("go.micro.srv.elasticsearch.Elasticsearch.Update", "Record required")
	}

	if err = RequiredRecordFieldsExists(req.Record); err != nil {
		return err
	}

	if err = elastic.Update(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Update", err.Error())
	}

	return nil
}

// Delete srv handler
func (es *Elasticsearch) Delete(ctx context.Context, req *elasticseach.DeleteRequest, rsp *elasticseach.DeleteResponse) error {
	var err error

	rHelper := &elasticseach.Record{
		Index: req.Index,
		Type:  req.Type,
		Id:    req.Id,
	}
	if err = RequiredRecordFieldsExists(rHelper); err != nil {
		return err
	}

	if err = elastic.Delete(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Delete", err.Error())
	}

	return nil
}

// Search srv handler
func (es *Elasticsearch) Search(ctx context.Context, req *elasticseach.SearchRequest, rsp *elasticseach.SearchResponse) error {
	result, err := elastic.Search(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Search", err.Error())
	}

	rsp.Result = result

	return nil
}

// Query srv handler
func (es *Elasticsearch) Query(ctx context.Context, req *elasticseach.QueryRequest, rsp *elasticseach.QueryResponse) error {
	result, err := elastic.Query(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elasticsearch.Elasticsearch.Query", err.Error())
	}

	rsp.Result = result

	return nil
}
