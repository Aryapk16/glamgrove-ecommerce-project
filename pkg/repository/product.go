package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"glamgrove/pkg/domain"
	interfaces "glamgrove/pkg/repository/interfaces"
	"glamgrove/pkg/utils/request"
	"glamgrove/pkg/utils/response"

	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: db}
}

func (p *productDatabase) FindBrand(ctx context.Context, brand request.Category) (request.Category, error) {
	query := `SELECT * FROM categories category_name=?`
	if p.DB.Raw(query, brand.CategoryName).Scan(&brand).Error != nil {
		return request.Category{}, errors.New("failed to get brand")
	}
	return brand, nil
}

// To add brand
func (p *productDatabase) AddCategory(ctx context.Context, brand request.Category) (err error) {

	err = p.DB.Create(&brand).Error

	if err != nil {
		return errors.New("failed to save brand")
	}
	return nil
}

func (p *productDatabase) GetAllBrand(ctx context.Context) (brand []response.Brand, err error) {

	query := `SELECT c.id, c.category_name FROM categories as c`
	if p.DB.Raw(query).Scan(&brand).Error != nil {
		return brand, fmt.Errorf("failed to get brands data from db")
	}
	fmt.Println(brand)

	return brand, nil
}
func (p *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {
	query := `INSERT INTO products (name, description, category_id, price, created_at) VALUES ($1, $2, $3, $4, $5)`

	createdAt := time.Now()
	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID, product.Price, createdAt).Error != nil {
		return errors.New("failed to save product on database")
	}
	return nil
}

// Add Image
func (pd *productDatabase) AddImage(c context.Context, pid int, filename string) (domain.ProductImage, error) {

	// Store the image record in the database
	image := domain.ProductImage{ProductId: uint(pid), Image: filename}
	if err := pd.DB.Create(&image).Error; err != nil {

		return domain.ProductImage{}, errors.New("failed to store image record")
	}

	return image, nil
}

func (pd *productDatabase) AddItemImage(c context.Context, pid int, filename string) (domain.ProductItemImage, error) {

	// Store the image record in the database
	image := domain.ProductItemImage{ProductItemID: uint(pid), Image: filename}
	if err := pd.DB.Create(&image).Error; err != nil {

		return domain.ProductItemImage{}, errors.New("failed to store image record")
	}

	return image, nil
}

func (p *productDatabase) GetProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products where id = ? product_name = ?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("failure to get product")
	}
	return product, nil
}
func (p *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products WHERE id = ? OR name=?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("failed to get product")
	}
	return product, nil
}
func (p *productDatabase) FindProductByID(ctx context.Context, productID uint) (product domain.Product, err error) {
	query := `SELECT * FROM products WHERE id = $1`
	err = p.DB.Raw(query, productID).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("failed find product with prduct_id %v", productID)
	}
	return product, nil
}
func (p *productDatabase) GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT p.id, p.name, p.description, c.category_name, p.price, p.discount_price, p.created_at, p.updated_at, pi.image FROM products p LEFT JOIN categories c ON p.category_id = c.id LEFT JOIN product_images pi ON p.id = pi.product_id ORDER BY p.created_at DESC LIMIT $1 OFFSET $2;
	`

	if p.DB.Raw(query, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("failed to get products from database")
	}

	fmt.Println(products[0].Image)

	return products, nil
}

// update product
func (p *productDatabase) UpdateProduct(ctx context.Context, product domain.Product) error {
	existingProduct, err := p.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	}
	if product.Name == "" {
		product.Name = existingProduct.Name
	}
	if product.Description == "" {
		product.Description = existingProduct.Description
	}
	if product.Price == 0 {
		product.Price = existingProduct.Price
	}
	// if product.Image == "" {
	// 	product.Image = existingProduct.Image
	// }
	if product.CategoryID == 0 {
		product.CategoryID = existingProduct.CategoryID
	}
	query := `UPDATE products SET name = $1, description = $2, category_id = $3,
	price = $4,  updated_at = $5 WHERE id = $6`

	updatedAt := time.Now()

	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID,
		product.Price, updatedAt, product.ID).Error != nil {
		return errors.New("failed to update product")
	}

	return nil
}

func (p *productDatabase) DeleteProduct(ctx context.Context, productID uint) (domain.Product, error) {
	// Check requested product is exist or not
	var existingProduct domain.Product
	existingProduct, err := p.FindProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	} else if existingProduct.Name == "" {
		return domain.Product{}, errors.New("invalid product_id")
	}

	//delete query
	query := `DELETE FROM products WHERE id = $1`
	if err := p.DB.Exec(query, productID).Error; err != nil {
		return domain.Product{}, fmt.Errorf("failed to delete error : %v", err)
	}
	return existingProduct, nil
}

func (p *productDatabase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	//tnx := p.DB.Begin()
	var product_item domain.ProductItem

	// Check if the product already exists
	existingProduct, err := p.FindProductByID(ctx, productItem.ProductID)
	if err != nil {
		return err
	}
	if existingProduct.ID != productItem.ProductID {

		return errors.New("product does not exist for the requested product item")
	}

	// Save the product item to the database
	query := `INSERT INTO product_items (product_id, qty_in_stock, price, discount_price, created_at) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	createdAt := time.Now()
	if err := p.DB.Raw(query, productItem.ProductID, productItem.QtyInStock, productItem.Price,
		productItem.DiscountPrice, createdAt).Scan(&product_item).Error; err != nil {

		return fmt.Errorf("failed to add product item: %v", err)
	}

	return nil
}

// to list product item
func (p *productDatabase) GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error) {
	// Check if the product ID exists
	var productItems []response.ProductItemResp

	dbProduct, err := p.FindProductByID(ctx, productId)
	if err != nil {
		return productItems, err
	}
	if dbProduct.ID == 0 {
		return productItems, errors.New("invalid product ID")
	}

	// Get product items from the database
	query := `SELECT
    p.id AS product_id,
    pi.id AS product_item_id,
    pi.qty_in_stock AS stock_available,
    c.category_name AS brand,
    
    pi.price,
    pi.discount_price AS offer_price
   
FROM
    products p
    LEFT JOIN categories c ON c.id = p.category_id
    LEFT JOIN product_items pi ON pi.product_id = p.id
    LEFT JOIN product_item_images im ON p.id = im.product_item_id
WHERE
    p.id = $1;
`
	fmt.Println(productItems)
	if err := p.DB.Raw(query, productId).Scan(&productItems).Error; err != nil {
		return productItems, fmt.Errorf("failed to get product items: %v", err)
	}
	fmt.Println("Product Items: ", productItems)

	// Fetch product item images
	query = `SELECT
		pimg.image
	FROM
		product_item_images pimg
	WHERE product_item_id = $1`
	for i := range productItems {
		productItems[i].Images = []string{}
		p.DB.Raw(query, productItems[i].ProductItemID).Scan(&productItems[i].Images)
	}
	fmt.Println("product Id: ", productId)

	return productItems, nil
}

// sales
func (pd *productDatabase) SalesData(sDate, eDate time.Time) (response.SalesResponse, error) {
	var salesData response.SalesResponse
	query := `SELECT COUNT(*) FROM orders WHERE order_date >=$1 AND order_date <= $2`
	err := pd.DB.Raw(query, sDate, eDate).Scan(&salesData.TotalOrder).Error
	if err != nil {
		return response.SalesResponse{}, errors.New("failed to count total orders")
	}

	status := "delivered"
	query1 := `SELECT COUNT(*) FROM orders WHERE order_date >= $1 AND order_date <= $2 AND delivery_status = $3`
	err1 := pd.DB.Raw(query1, sDate, eDate, status).Scan(&salesData.DeliveredOrder).Error
	if err1 != nil {
		return response.SalesResponse{}, errors.New("failed to count delivered orders")
	}

	status1 := "Pending"
	query2 := `SELECT COUNT(*) FROM orders WHERE order_date >= $1 AND order_date <= $2 AND delivery_status = $3 AND order_status != 'order cancelled'`
	err2 := pd.DB.Raw(query2, sDate, eDate, status1).Scan(&salesData.PendingOrder).Error
	if err2 != nil {
		return response.SalesResponse{}, errors.New("failed to count pending orders")
	}

	status2 := "order cancelled"
	query3 := `SELECT COUNT(*) FROM orders WHERE order_date >= $1 AND order_date <= $2 AND order_status = $3`
	err3 := pd.DB.Raw(query3, sDate, eDate, status2).Scan(&salesData.CancelledOrder).Error
	if err3 != nil {
		return response.SalesResponse{}, errors.New("failed to count cancelled orders")
	}

	return salesData, nil

}

// verify and delete cart items
func (pd *productDatabase) DeleteCart(c context.Context, usr_id uint) error {
	var cartItems domain.CartItems
	query := `DELETE FROM carts WHERE user_id=?`
	err := pd.DB.Raw(query, usr_id).Scan(&cartItems).Error
	if err != nil {
		return errors.New("failed to delete cart items")
	}
	return nil
}

func (pd *productDatabase) UpdateStatusRazorpay(c context.Context, order_id uint, order_status string, payment_status string) (response.OrderResponse, error) {
	var order domain.Order
	var orderResp response.OrderResponse
	query := `update orders set order_status=?,payment_status=?  where order_id=?`
	err := pd.DB.Raw(query, order_status, payment_status, order_id).Scan(&order).Error
	if err != nil {
		return response.OrderResponse{}, errors.New("failed to update order status")
	}
	query1 := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id where o.order_id=?`
	err1 := pd.DB.Raw(query1, order_id).Scan(&orderResp).Error
	if err1 != nil {
		return response.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderResp, nil
}
