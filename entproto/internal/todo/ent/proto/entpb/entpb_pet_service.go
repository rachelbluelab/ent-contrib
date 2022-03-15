// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	ent "entgo.io/contrib/entproto/internal/todo/ent"
	pet "entgo.io/contrib/entproto/internal/todo/ent/pet"
	user "entgo.io/contrib/entproto/internal/todo/ent/user"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strconv "strconv"
)

// PetService implements PetServiceServer
type PetService struct {
	client *ent.Client
	UnimplementedPetServiceServer
}

// NewPetService returns a new PetService
func NewPetService(client *ent.Client) *PetService {
	return &PetService{
		client: client,
	}
}

// toProtoPet transforms the ent type to the pb type
func toProtoPet(e *ent.Pet) (*Pet, error) {
	v := &Pet{}
	id := int64(e.ID)
	v.Id = id
	if edg := e.Edges.Owner; edg != nil {
		id := int64(edg.ID)
		v.Owner = &User{
			Id: id,
		}
	}
	return v, nil
}

// Create implements PetServiceServer.Create
func (svc *PetService) Create(ctx context.Context, req *CreatePetRequest) (*Pet, error) {
	pet := req.GetPet()
	m := svc.client.Pet.Create()
	if pet.GetOwner() != nil {
		petOwner := int(pet.GetOwner().GetId())
		m.SetOwnerID(petOwner)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPet(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements PetServiceServer.Get
func (svc *PetService) Get(ctx context.Context, req *GetPetRequest) (*Pet, error) {
	var (
		err error
		get *ent.Pet
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetPetRequest_VIEW_UNSPECIFIED, GetPetRequest_BASIC:
		get, err = svc.client.Pet.Get(ctx, id)
	case GetPetRequest_WITH_EDGE_IDS:
		get, err = svc.client.Pet.Query().
			Where(pet.ID(id)).
			WithOwner(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoPet(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements PetServiceServer.Update
func (svc *PetService) Update(ctx context.Context, req *UpdatePetRequest) (*Pet, error) {
	pet := req.GetPet()
	petID := int(pet.GetId())
	m := svc.client.Pet.UpdateOneID(petID)
	if pet.GetOwner() != nil {
		petOwner := int(pet.GetOwner().GetId())
		m.SetOwnerID(petOwner)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPet(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements PetServiceServer.Delete
func (svc *PetService) Delete(ctx context.Context, req *DeletePetRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Pet.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements PetServiceServer.List
func (svc *PetService) List(ctx context.Context, req *ListPetRequest) (*ListPetResponse, error) {
	var (
		err      error
		entList  []*ent.Pet
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Pet.Query().
		Order(ent.Desc(pet.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(pet.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListPetRequest_VIEW_UNSPECIFIED, ListPetRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListPetRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithOwner(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		var pbList []*Pet
		for _, entEntity := range entList {
			pbEntity, err := toProtoPet(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListPetResponse{
			PetList:       pbList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}
