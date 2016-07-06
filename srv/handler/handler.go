package handler

import (
	"github.com/kazoup/elastic/srv/elastic"
	proto "github.com/kazoup/elastic/srv/proto/elastic"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

// Elastic struct
type Elastic struct{}

// Create srv handler
func (es *Elastic) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	var err error

	if err = DocRefFieldsExists(&proto.DocRef{
		Index: req.Index,
		Type:  req.Type,
		Id:    req.Id,
	}); err != nil {
		return err
	}

	if err = elastic.Create(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Create", err.Error())
	}

	return nil
}

// Read srv handler
func (es *Elastic) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	var err error

	if err = DocRefFieldsExists(&proto.DocRef{
		Index: req.Index,
		Type:  req.Type,
		Id:    req.Id,
	}); err != nil {
		return err
	}

	record, err := elastic.Read(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Read", err.Error())
	}

	rsp.Result = record

	return nil
}

// Update srv handler
func (es *Elastic) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	var err error

	if err = DocRefFieldsExists(&proto.DocRef{
		Id:    req.Id,
		Index: req.Index,
		Type:  req.Type,
	}); err != nil {
		return err
	}

	if err = elastic.Update(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Update", err.Error())
	}

	return nil
}

// Delete srv handler
func (es *Elastic) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	var err error

	if err = DocRefFieldsExists(&proto.DocRef{
		Index: req.Index,
		Type:  req.Type,
		Id:    req.Id,
	}); err != nil {
		return err
	}

	if err = elastic.Delete(req); err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Delete", err.Error())
	}

	return nil
}

// Search srv handler
func (es *Elastic) Search(ctx context.Context, req *proto.SearchRequest, rsp *proto.SearchResponse) error {
	result, err := elastic.Search(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Search", err.Error())
	}

	rsp.Result = result

	return nil
}

// Query srv handler
func (es *Elastic) Query(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {
	result, err := elastic.Query(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.elastic.Elastic.Query", err.Error())
	}

	rsp.Result = result

	return nil
}
