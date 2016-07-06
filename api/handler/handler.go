package handler

import (
	"encoding/json"
	"fmt"
	elastic "github.com/kazoup/elastic/srv/proto/elastic"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"net/http"
)

// Elastic struct
type Elastic struct{}

// Create API handler
func (es *Elastic) Create(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var input map[string]interface{}
	var data []byte

	// Unmarshal unknown JSON
	if err = json.Unmarshal([]byte(req.Body), &input); err != nil {
		return errors.BadRequest("go.micro.api.elastic", err.Error())
	}

	// Marshal unknown JSON (data)
	data, err = json.Marshal(input["data"])

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Create",
		&elastic.CreateRequest{
			Index: fmt.Sprintf("%v", input["index"]),
			Type:  fmt.Sprintf("%v", input["type"]),
			Id:    fmt.Sprintf("%v", input["id"]),
			Data:  string(data),
		},
	)
	srvRsp := &elastic.CreateResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = `{}`

	return nil
}

// Read API handler
func (es *Elastic) Read(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var readRequest *elastic.ReadRequest

	if err = json.Unmarshal([]byte(req.Body), &readRequest); err != nil {
		return errors.InternalServerError("go.micro.api.elastic", err.Error())
	}

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Read",
		readRequest,
	)
	srvRsp := &elastic.ReadResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = srvRsp.Result

	return nil
}

// Update API handler
func (es *Elastic) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var input map[string]interface{}
	var data []byte

	// Unmarshal unknown JSON
	if err = json.Unmarshal([]byte(req.Body), &input); err != nil {
		return errors.BadRequest("go.micro.api.elastic", err.Error())
	}

	// Marshal unknown JSON (data)
	data, err = json.Marshal(input["data"])

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Update",
		&elastic.UpdateRequest{
			Index: fmt.Sprintf("%v", input["index"]),
			Type:  fmt.Sprintf("%v", input["type"]),
			Id:    fmt.Sprintf("%v", input["id"]),
			Data:  string(data),
		},
	)
	srvRsp := &elastic.UpdateResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = `{}`

	return nil
}

// Delete API handler
func (es *Elastic) Delete(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var deleteRequest *elastic.DeleteRequest

	if err = json.Unmarshal([]byte(req.Body), &deleteRequest); err != nil {
		return errors.InternalServerError("go.micro.api.elastic", err.Error())
	}

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Delete",
		deleteRequest,
	)
	srvRsp := &elastic.DeleteResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = `{}`

	return nil
}

// Search API handler
func (es *Elastic) Search(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var input map[string]interface{}

	// Unmarshal unknown JSON
	if err = json.Unmarshal([]byte(req.Body), &input); err != nil {
		return errors.BadRequest("go.micro.api.elastic", err.Error())
	}

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Search",
		&elastic.SearchRequest{
			Index:  fmt.Sprintf("%v", input["index"]),
			Type:   fmt.Sprintf("%v", input["type"]),
			Limit:  int64((input["limit"]).(float64)),
			Offset: int64((input["offset"]).(float64)),
			Query:  fmt.Sprintf("%v", input["query"]),
		},
	)
	srvRsp := &elastic.SearchResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = srvRsp.Result

	return nil
}

// Query API handler
func (es *Elastic) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var input map[string]interface{}
	var query []byte

	// Unmarshal unknown JSON
	if err = json.Unmarshal([]byte(req.Body), &input); err != nil {
		return errors.BadRequest("go.micro.api.elastic", err.Error())
	}

	query, err = json.Marshal(input["query"])

	srvReq := client.NewRequest(
		"go.micro.srv.elastic",
		"Elastic.Query",
		&elastic.QueryRequest{
			Index: fmt.Sprintf("%v", input["index"]),
			Type:  fmt.Sprintf("%v", input["type"]),
			Query: string(query),
		},
	)
	srvRsp := &elastic.SearchResponse{}
	if err = client.Call(ctx, srvReq, srvRsp); err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = srvRsp.Result

	return nil
}
