package adminProduct

import (
	"context"
	"log/slog"

	"github.com/wigata-intech/logres/internal/api/repository"
	"github.com/wigata-intech/logres/internal/api/service"
	"github.com/wigata-intech/logres/internal/shared/model"
)

type adminProductService struct {
	logger             *slog.Logger
	categoryRepository repository.ICategoryRepository
	fileRepository     repository.IFileRepository
	productRepository  repository.IProductRepository
	tagRepository      repository.ITagRepository
}

func New(
	logger *slog.Logger,
	categoryRepository repository.ICategoryRepository,
	fileRepository repository.IFileRepository,
	productRepository repository.IProductRepository,
	tagRepository repository.ITagRepository,
) service.IAdminProductService {
	return &adminProductService{
		logger:             logger,
		categoryRepository: categoryRepository,
		fileRepository:     fileRepository,
		productRepository:  productRepository,
		tagRepository:      tagRepository,
	}
}

func (s *adminProductService) Create(ctx context.Context, req *model.AdminProductCreateRequest) (*model.AdminProductResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) List(ctx context.Context, req *model.AdminProductListRequest) (*model.AdminProductListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) Detail(ctx context.Context, req *model.AdminProductDetailRequest) (*model.AdminProductResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) Update(ctx context.Context, req *model.AdminProductUpdateRequest) (*model.AdminProductResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) Delete(ctx context.Context, req *model.AdminProductDeleteRequest) error {
	panic("not implemented") // TODO: Implement
}

// Category and Tags
func (s *adminProductService) CategoryCreate(ctx context.Context, req *model.AdminCategoryCreateRequest) (*model.AdminCategoryResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) CategoryList(ctx context.Context, req *model.AdminCategoryListRequest) (*model.AdminCategoryListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) CategoryUpdate(ctx context.Context, req *model.AdminCategoryUpdateRequest) (*model.AdminCategoryResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) CategoryDelete(ctx context.Context, req *model.AdminCategoryDeleteRequest) error {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) TagCreate(ctx context.Context, req *model.AdminTagCreateRequest) (*model.AdminTagResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) TagList(ctx context.Context, req *model.AdminTagListRequest) (*model.AdminTagListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) TagUpdate(ctx context.Context, req *model.AdminTagUpdateRequest) (*model.AdminTagResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *adminProductService) TagDelete(ctx context.Context, req *model.AdminTagDeleteRequest) error {
	panic("not implemented") // TODO: Implement
}
