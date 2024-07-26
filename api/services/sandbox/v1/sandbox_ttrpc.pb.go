// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/sandbox/v1/sandbox.proto
package sandbox

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
)

type TTRPCStoreService interface {
	Create(context.Context, *StoreCreateRequest) (*StoreCreateResponse, error)
	Update(context.Context, *StoreUpdateRequest) (*StoreUpdateResponse, error)
	Delete(context.Context, *StoreDeleteRequest) (*StoreDeleteResponse, error)
	List(context.Context, *StoreListRequest) (*StoreListResponse, error)
	Get(context.Context, *StoreGetRequest) (*StoreGetResponse, error)
}

func RegisterTTRPCStoreService(srv *ttrpc.Server, svc TTRPCStoreService) {
	srv.RegisterService("containerd.services.sandbox.v1.Store", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StoreCreateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StoreUpdateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StoreDeleteRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"List": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StoreListRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.List(ctx, &req)
			},
			"Get": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StoreGetRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Get(ctx, &req)
			},
		},
	})
}

type ttrpcstoreClient struct {
	client *ttrpc.Client
}

func NewTTRPCStoreClient(client *ttrpc.Client) TTRPCStoreService {
	return &ttrpcstoreClient{
		client: client,
	}
}

func (c *ttrpcstoreClient) Create(ctx context.Context, req *StoreCreateRequest) (*StoreCreateResponse, error) {
	var resp StoreCreateResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Store", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpcstoreClient) Update(ctx context.Context, req *StoreUpdateRequest) (*StoreUpdateResponse, error) {
	var resp StoreUpdateResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Store", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpcstoreClient) Delete(ctx context.Context, req *StoreDeleteRequest) (*StoreDeleteResponse, error) {
	var resp StoreDeleteResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Store", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpcstoreClient) List(ctx context.Context, req *StoreListRequest) (*StoreListResponse, error) {
	var resp StoreListResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Store", "List", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpcstoreClient) Get(ctx context.Context, req *StoreGetRequest) (*StoreGetResponse, error) {
	var resp StoreGetResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Store", "Get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TTRPCControllerService interface {
	Create(context.Context, *ControllerCreateRequest) (*ControllerCreateResponse, error)
	Start(context.Context, *ControllerStartRequest) (*ControllerStartResponse, error)
	Platform(context.Context, *ControllerPlatformRequest) (*ControllerPlatformResponse, error)
	Stop(context.Context, *ControllerStopRequest) (*ControllerStopResponse, error)
	Wait(context.Context, *ControllerWaitRequest) (*ControllerWaitResponse, error)
	Status(context.Context, *ControllerStatusRequest) (*ControllerStatusResponse, error)
	Shutdown(context.Context, *ControllerShutdownRequest) (*ControllerShutdownResponse, error)
	Metrics(context.Context, *ControllerMetricsRequest) (*ControllerMetricsResponse, error)
	Update(context.Context, *ControllerUpdateRequest) (*ControllerUpdateResponse, error)
}

func RegisterTTRPCControllerService(srv *ttrpc.Server, svc TTRPCControllerService) {
	srv.RegisterService("containerd.services.sandbox.v1.Controller", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerCreateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerStartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Platform": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerPlatformRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Platform(ctx, &req)
			},
			"Stop": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerStopRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Stop(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerWaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
			"Status": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerStatusRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Status(ctx, &req)
			},
			"Shutdown": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerShutdownRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Shutdown(ctx, &req)
			},
			"Metrics": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerMetricsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Metrics(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ControllerUpdateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
		},
	})
}

type ttrpccontrollerClient struct {
	client *ttrpc.Client
}

func NewTTRPCControllerClient(client *ttrpc.Client) TTRPCControllerService {
	return &ttrpccontrollerClient{
		client: client,
	}
}

func (c *ttrpccontrollerClient) Create(ctx context.Context, req *ControllerCreateRequest) (*ControllerCreateResponse, error) {
	var resp ControllerCreateResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Start(ctx context.Context, req *ControllerStartRequest) (*ControllerStartResponse, error) {
	var resp ControllerStartResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Platform(ctx context.Context, req *ControllerPlatformRequest) (*ControllerPlatformResponse, error) {
	var resp ControllerPlatformResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Platform", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Stop(ctx context.Context, req *ControllerStopRequest) (*ControllerStopResponse, error) {
	var resp ControllerStopResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Wait(ctx context.Context, req *ControllerWaitRequest) (*ControllerWaitResponse, error) {
	var resp ControllerWaitResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Status(ctx context.Context, req *ControllerStatusRequest) (*ControllerStatusResponse, error) {
	var resp ControllerStatusResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Shutdown(ctx context.Context, req *ControllerShutdownRequest) (*ControllerShutdownResponse, error) {
	var resp ControllerShutdownResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Shutdown", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Metrics(ctx context.Context, req *ControllerMetricsRequest) (*ControllerMetricsResponse, error) {
	var resp ControllerMetricsResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Metrics", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontrollerClient) Update(ctx context.Context, req *ControllerUpdateRequest) (*ControllerUpdateResponse, error) {
	var resp ControllerUpdateResponse
	if err := c.client.Call(ctx, "containerd.services.sandbox.v1.Controller", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}