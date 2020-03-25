// Copyright 2020 The Pipe Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kapetaniosci/pipe/pkg/app/api/service"
)

// RunnerAPI implements the behaviors for the gRPC definitions of WebAPI.
type WebAPI struct {
	logger *zap.Logger
}

// NewWebAPIService creates a new service instance.
func NewWebAPIService(logger *zap.Logger) *WebAPI {
	a := &WebAPI{
		logger: logger.Named("web-api"),
	}
	return a
}

// Register registers all handling of this service into the specified gRPC server.
func (a *WebAPI) Register(server *grpc.Server) {
	service.RegisterWebAPIServer(server, a)
}

func (a *WebAPI) RegisterRunner(ctx context.Context, req *service.RegisterRunnerRequest) (*service.RegisterRunnerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) DisableRunner(ctx context.Context, req *service.DisableRunnerRequest) (*service.DisableRunnerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) ListRunners(ctx context.Context, req *service.ListRunnersRequest) (*service.ListRunnersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) AddEnvironment(ctx context.Context, req *service.AddEnvironmentRequest) (*service.AddEnvironmentResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) RenameEnvironment(ctx context.Context, req *service.RenameEnvironmentRequest) (*service.RenameEnvironmentResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) ListEnvironments(ctx context.Context, req *service.ListEnvironmentsRequest) (*service.ListEnvironmentsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) AddApplication(ctx context.Context, req *service.AddApplicationRequest) (*service.AddApplicationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) DisableApplication(ctx context.Context, req *service.DisableApplicationRequest) (*service.DisableApplicationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) ListApplications(ctx context.Context, req *service.ListApplicationsRequest) (*service.ListApplicationsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) SyncApplication(ctx context.Context, req *service.SyncApplicationRequest) (*service.SyncApplicationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) ListPipelines(ctx context.Context, req *service.ListPipelinesRequest) (*service.ListPipelinesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) GetPipeline(ctx context.Context, req *service.GetPipelineRequest) (*service.GetPipelineResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) CancelPipeline(ctx context.Context, req *service.CancelPipelineRequest) (*service.CancelPipelineResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) ApprovePipelineStage(ctx context.Context, req *service.ApprovePipelineStageRequest) (*service.ApprovePipelineStageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) RetryPipelineStage(ctx context.Context, req *service.RetryPipelineStageRequest) (*service.RetryPipelineStageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) GetApplicationState(ctx context.Context, req *service.GetApplicationStateRequest) (*service.GetApplicationStateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) GetProject(ctx context.Context, req *service.GetProjectRequest) (*service.GetProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (a *WebAPI) GetMe(ctx context.Context, req *service.GetMeRequest) (*service.GetMeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}