package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// Product represents a product in the inventory.
type Product struct {
	ID       int
	Name     string
	Quantity int
}

// InventoryDB represents the database interface for inventory operations.
type InventoryDB interface {
	GetProductByID(productID int) (*Product, error)
	UpdateProductQuantity(productID, newQuantity int) error
}

// NotificationService represents a service for sending notifications.
type NotificationService interface {
	SendNotification(message string) error
}

// InventoryManager represents the manager for inventory operations.
type InventoryManager struct {
	DB     InventoryDB
	Notify NotificationService
}

// NewInventoryManager creates a new InventoryManager with the specified database and notification service.
func NewInventoryManager(db InventoryDB, notify NotificationService) *InventoryManager {
	return &InventoryManager{DB: db, Notify: notify}
}

// UpdateInventory updates the inventory for the specified product.
func (im *InventoryManager) UpdateInventory(productID, newQuantity int) error {
	// Start a transaction
	tx, err := im.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			// Rollback the transaction if an error occurred
			tx.Rollback()
			return
		}

		// Commit the transaction if no error occurred
		if tx.Commit() != nil {
			log.Printf("failed to commit transaction: %v\n", err)
		}
	}()

	product, err := im.validateProduct(productID, newQuantity)
	if err != nil {
		return err
	}

	// Update the product quantity
	err = im.DB.UpdateProductQuantity(productID, newQuantity)
	if err != nil {
		return fmt.Errorf("failed to update product quantity: %w", err)
	}

	// Log the inventory update
	log.Printf("Inventory updated for product '%s' (ID: %d). New quantity: %d\n", product.Name, product.ID, newQuantity)

	// Handle different scenarios based on inventory changes
	return im.updateQuantity(product, newQuantity)
}

func (im *InventoryManager) validateProduct(productID, newQuantity int) (*Product, error) {
	// Get the product from the database
	product, err := im.DB.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	// Validate the new quantity
	if newQuantity < 0 {
		return nil, errors.New("invalid quantity. Quantity must be non-negative")
	}

	return product, nil
}

func (im *InventoryManager) updateQuantity(product *Product, newQuantity int) error {
	if newQuantity == 0 {
		message := fmt.Sprintf("Product '%s' (ID: %d) is out of stock", product.Name, product.ID)
		return im.Notify.SendNotification(message)
	}

	if newQuantity < product.Quantity {
		message := fmt.Sprintf("Inventory decreased for product '%s' (ID: %d). New quantity: %d", product.Name, product.ID, newQuantity)
		return im.Notify.SendNotification(message)
	}

	if newQuantity > product.Quantity {
		message := fmt.Sprintf("Inventory increased for product '%s' (ID: %d). New quantity: %d", product.Name, product.ID, newQuantity)
		return im.Notify.SendNotification(message)
	}

	return nil
}

// SQLInventoryDB represents a SQL-based implementation of the InventoryDB interface.
type SQLInventoryDB struct {
	DB *sql.DB
}

// GetProductByID retrieves a product from the database by ID.
func (sidb *SQLInventoryDB) GetProductByID(productID int) (*Product, error) {
	// Placeholder for SQL query to retrieve product by ID
	// This is a simplified example for demonstration purposes
	return &Product{ID: productID, Name: "Sample Product", Quantity: 10}, nil
}

// UpdateProductQuantity updates the quantity of a product in the database.
func (sidb *SQLInventoryDB) UpdateProductQuantity(productID, newQuantity int) error {
	// Placeholder for SQL update query to update product quantity
	// This is a simplified example for demonstration purposes
	return nil
}

// EmailNotificationService represents an email notification service.
type EmailNotificationService struct {
	// Placeholder for email notification service configuration
}

// SendNotification sends a notification via email.
func (ens *EmailNotificationService) SendNotification(message string) error {
	// Placeholder for sending email notification
	return nil
}

func main() {
	// Initialize the SQL database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/inventory")
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}
	defer db.Close()

	// Create a new instance of SQLInventoryDB
	sqlDB := &SQLInventoryDB{DB: db}

	// Create a new instance of EmailNotificationService
	emailNotify := &EmailNotificationService{}

	// Create a new instance of InventoryManager
	im := NewInventoryManager(sqlDB, emailNotify)

	// Update the inventory for a product
	if err := im.UpdateInventory(123, 20); err != nil {
		log.Println("Error:", err)
	}
}
