# Before (Cyclomatic Complexity of UpdateInventory = 13)

~~~
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
		} else {
			// Commit the transaction if no error occurred
			err = tx.Commit()
			if err != nil {
				log.Printf("failed to commit transaction: %v\n", err)
			}
		}
	}()

	// Get the product from the database
	product, err := im.DB.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	// Validate the new quantity
	if newQuantity < 0 {
		return errors.New("invalid quantity. Quantity must be non-negative")
	}

	// Update the product quantity
	err = im.DB.UpdateProductQuantity(productID, newQuantity)
	if err != nil {
		return fmt.Errorf("failed to update product quantity: %w", err)
	}

	// Log the inventory update
	log.Printf("Inventory updated for product '%s' (ID: %d). New quantity: %d\n", product.Name, product.ID, newQuantity)

	// Handle different scenarios based on inventory changes
	switch {
	case newQuantity == 0:
		message := fmt.Sprintf("Product '%s' (ID: %d) is out of stock", product.Name, product.ID)
		if err := im.Notify.SendNotification(message); err != nil {
			log.Printf("failed to send notification: %v\n", err)
		}
	case newQuantity < product.Quantity:
		message := fmt.Sprintf("Inventory decreased for product '%s' (ID: %d). New quantity: %d", product.Name, product.ID, newQuantity)
		if err := im.Notify.SendNotification(message); err != nil {
			log.Printf("failed to send notification: %v\n", err)
		}
	case newQuantity > product.Quantity:
		message := fmt.Sprintf("Inventory increased for product '%s' (ID: %d). New quantity: %d", product.Name, product.ID, newQuantity)
		if err := im.Notify.SendNotification(message); err != nil {
			log.Printf("failed to send notification: %v\n", err)
		}
	}

	// Return nil if the operation was successful
	return nil
}

~~~

# After (Cyclomatic Complexity of UpdateInventory = 6 )

~~~
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
~~~

## Methods used

- removed else cases
- removed switch/cases
- encapsulated processing logic into separated functions