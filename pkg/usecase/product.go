package usecase

import (
	"context"
	"errors"
	"fmt"
	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	service "glamgrove/pkg/usecase/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type productUseCase struct {
	ProductRepository interfaces.ProductRepository
}

func NewProductUseCase(ProdRepo interfaces.ProductRepository) service.ProductService {
	return &productUseCase{ProductRepository: ProdRepo}
}

func (p *productUseCase) AddCategory(ctx context.Context, brand request.Category) error {
	// check if req brand already exists in db
	// dbBrand, _ := p.ProductRepository.FindBrand(ctx, brand)
	// if dbBrand.ID == 0 {
	// 	return fmt.Errorf("brand already exist with %s name", brand.CategoryName)
	// }
	if err := p.ProductRepository.AddCategory(ctx, brand); err != nil {
		return err
	}
	return nil
}

// to get all brands
func (p *productUseCase) GetAllBrands(ctx context.Context) ([]response.Brand, error) {
	allBrands, err := p.ProductRepository.GetAllBrand(ctx)
	if err != nil {
		return []response.Brand{}, err
	}
	fmt.Println(allBrands)

	return allBrands, nil

}
func (p *productUseCase) AddProduct(ctx context.Context, product domain.Product) error {
	// Check the product already exists in databse
	if dbProd, err := p.ProductRepository.FindProduct(ctx, product); err != nil {
		return err
	} else if dbProd.ID != 0 {
		return fmt.Errorf("product already exist with %s product name", dbProd.Name)
	}
	return p.ProductRepository.SaveProduct(ctx, product)

}

func (p *productUseCase) AddImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.ProductImage, error) {
	var images []domain.ProductImage
	for _, file := range files {
		// Generate a unique filename for the image

		ext := filepath.Ext(file.Filename)
		filename := uuid.New().String() + ext

		image, err := p.ProductRepository.AddImage(c, pid, filename)
		if err != nil {
			return []domain.ProductImage{}, err
		}

		src, err := file.Open()
		if err != nil {
			return []domain.ProductImage{}, err
		}
		defer src.Close()

		// Create the destination file
		dst, err := os.Create(filepath.Join("images", filename)) // Replace "path/to/save/images" with your desired directory
		if err != nil {
			return []domain.ProductImage{}, err
		}
		defer dst.Close()

		// Copy the uploaded file's content to the destination file
		_, err = io.Copy(dst, src)
		if err != nil {
			return []domain.ProductImage{}, err
		}
		// product, _ := pu.productRepo.GetProductByID(c, pid)
		// image.Product = product

		images = append(images, image)
	}

	return images, nil
}

func (p *productUseCase) AddItemImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.ProductItemImage, error) {
	var images []domain.ProductItemImage
	for _, file := range files {
		// Generate a unique filename for the image

		ext := filepath.Ext(file.Filename)
		filename := uuid.New().String() + ext

		image, err := p.ProductRepository.AddItemImage(c, pid, filename)
		if err != nil {
			return []domain.ProductItemImage{}, err
		}

		src, err := file.Open()
		if err != nil {
			return []domain.ProductItemImage{}, err
		}
		defer src.Close()

		// Create the destination file
		dst, err := os.Create(filepath.Join("images", filename)) // Replace "path/to/save/images" with your desired directory
		if err != nil {
			return []domain.ProductItemImage{}, err
		}
		defer dst.Close()

		// Copy the uploaded file's content to the destination file
		_, err = io.Copy(dst, src)
		if err != nil {
			return []domain.ProductItemImage{}, err
		}
		// product, _ := pu.productRepo.GetProductByID(c, pid)
		// image.Product = product

		images = append(images, image)
	}

	return images, nil
}

func (p *productUseCase) GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {
	return p.ProductRepository.GetAllProducts(ctx, page)
}

func (p *productUseCase) UpdateProduct(ctx context.Context, product domain.Product) error {
	// Check requested product is exist or not
	existingProduct, err := p.ProductRepository.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	} else if existingProduct.Name == "" {
		return errors.New("invalid product_id")
	}

	// check the given product_name already exist or not
	existingProduct, err = p.ProductRepository.FindProduct(ctx, domain.Product{Name: product.Name})
	if err != nil {
		return err
	} else if existingProduct.ID != 0 && existingProduct.ID != product.ID {
		return errors.New("can't update the product \nrequested product_name already existing in database")
	}

	return p.ProductRepository.UpdateProduct(ctx, product)
}

func (p *productUseCase) DeleteProduct(ctx context.Context, productID uint) (domain.Product, error) {

	existingProduct, err := p.ProductRepository.DeleteProduct(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}
	return existingProduct, nil
}

func (p *productUseCase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	if err := p.ProductRepository.AddProductItem(ctx, productItem); err != nil {
		return fmt.Errorf("failed to add product item: %v", err)
	}
	return nil
}

func (p *productUseCase) GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error) {

	productItems, err := p.ProductRepository.GetProductItems(ctx, productId)
	if err != nil {
		return productItems, count, err
	}

	count = len(productItems)

	return productItems, count, nil
}

func (pu *productUseCase) SalesData(sDate, Edate time.Time) (response.SalesResponse, error) {
	salesData, err := pu.ProductRepository.SalesData(sDate, Edate)
	if err != nil {
		return response.SalesResponse{}, err
	}

	return salesData, nil
}
