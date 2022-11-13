package service

import (
	"context"
	"golang-united-homework/pkg/api"
	"golang-united-homework/pkg/repositories"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Homework struct {
	api.UnimplementedHomeworkServer
}

func (l *Homework) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	homework := &repositories.Homework{
		LectureId:   request.GetLectureId(),
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		CreatedBy:   request.GetCreatedBy(),
		UpdatedBy:   request.GetCreatedBy(),
	}

	err := repositories.CreateHomework(homework)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &api.CreateResponse{Id: homework.Id}, nil

}

func (l *Homework) Get(ctx context.Context, request *api.GetRequest) (*api.GetResponse, error) {

	homework, err := repositories.GetHomework(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	response := &api.GetResponse{
		Id:          homework.Id,
		LectureId:   homework.LectureId,
		Title:       homework.Title,
		Description: homework.Description,
		CreatedBy:   homework.CreatedBy,
		UpdatedBy:   homework.UpdatedBy,
		DeletedBy:   homework.DeletedBy,
		CreatedAt:   timestamppb.New(homework.CreatedAt),
		UpdatedAt:   timestamppb.New(homework.UpdatedAt),
	}

	if !homework.DeletedAt.IsZero() {
		response.DeletedAt = timestamppb.New(homework.DeletedAt)
	}

	return response, nil

}

func (l *Homework) Update(ctx context.Context, request *api.UpdateRequest) (*emptypb.Empty, error) {

	homework, err := repositories.GetHomework(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	if !homework.DeletedAt.IsZero() {
		return nil, status.New(codes.Aborted, "homework is deleted").Err()
	}

	homework.Title = request.GetTitle()
	homework.Description = request.GetDescription()
	homework.UpdatedBy = request.GetUpdatedBy()

	err = repositories.UpdateHomework(homework)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil

}

func (l *Homework) Delete(ctx context.Context, request *api.DeleteRequest) (*emptypb.Empty, error) {

	homework, err := repositories.GetHomework(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	if !homework.DeletedAt.IsZero() {
		return nil, status.New(codes.Aborted, "homework is deleted").Err()
	}

	homework.DeletedAt = time.Now()
	homework.DeletedBy = request.GetDeletedBy()

	err = repositories.UpdateHomework(homework)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil

}

func (l *Homework) List(ctx context.Context, request *api.ListRequest) (*api.ListResponse, error) {

	homework, err := repositories.GetHomeworkList(request.GetLectureId(), request.GetShowDeleted(), request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &api.ListResponse{}

	response.Homework = make([]*api.GetResponse, 0, len(*homework))

	for _, homework := range *homework {

		homeworkResponse := &api.GetResponse{
			Id:          homework.Id,
			LectureId:   homework.LectureId,
			Title:       homework.Title,
			Description: homework.Description,
			CreatedAt:   timestamppb.New(homework.CreatedAt),
			CreatedBy:   homework.CreatedBy,
			UpdatedAt:   timestamppb.New(homework.UpdatedAt),
			UpdatedBy:   homework.UpdatedBy,
			DeletedBy:   homework.DeletedBy,
		}

		if !homework.DeletedAt.IsZero() {
			homeworkResponse.DeletedAt = timestamppb.New(homework.DeletedAt)
		}

		response.Homework = append(response.Homework, homeworkResponse)

	}

	return response, nil

}
