CREATE PROCEDURE IF NOT EXISTS sp_CreateProductCategory(
    IN p_name VARCHAR(50),
    IN p_description VARCHAR(250)
)
BEGIN
    DECLARE insertedID INT;
    DECLARE p_message VARCHAR(255);

    -- Insert the user data into table
    INSERT INTO product_category (name, description)
    VALUES (p_name, p_description);

    -- Get the last inserted user ID
    SET insertedID = LAST_INSERT_ID();

    -- SELECT 'User registered successfully' AS message;
    SET p_message = 'record added successfully';

    -- Return inserted data along with the message
    SELECT insertedID, p_message AS message;

END;

CREATE PROCEDURE IF NOT EXISTS sp_CreateProducts(
    IN p_name VARCHAR(50),
    IN p_description VARCHAR(250),
    IN p_sku VARCHAR(250),
    IN p_category_id INT,
    IN p_price INT
)
BEGIN
    DECLARE error_message VARCHAR(255);
    DECLARE insertedID INT;
    DECLARE p_message VARCHAR(255);

    -- Check if the record with the specified ID exists
    IF EXISTS (SELECT 1 FROM products WHERE sku = p_sku) THEN
        -- Record exists, perform the update or delete operation
        -- Record does not exist, send an error message
        SET error_message = CONCAT(p_sku, ' already exists');
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = error_message;

    ELSE
        -- Insert the user data into table
        INSERT INTO products (name, description, sku, category_id, price)
        VALUES (p_name, p_description, p_sku, p_category_id, p_price);

        -- Get the last inserted user ID
        SET insertedID = LAST_INSERT_ID();

        -- SELECT 'User registered successfully' AS message;
        SET p_message = 'record added successfully';

        -- Return inserted data along with the message
        SELECT insertedID, p_message AS message;

    END IF;
END;

CREATE PROCEDURE IF NOT EXISTS sp_CreateProductInventory(
    IN p_product_id VARCHAR(50),
    IN p_quantity VARCHAR(250)
)
BEGIN
    DECLARE error_message VARCHAR(255);
    DECLARE insertedID INT;
    DECLARE p_message VARCHAR(255);

    -- Check if the record with the specified ID exists
    IF EXISTS (SELECT 1 FROM product_inventory WHERE product_id = p_product_id) THEN
        -- Record exists, perform the update or delete operation
        -- Record does not exist, send an error message
        SET error_message = CONCAT(p_product_id, ' already exists');
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = error_message;

    ELSE
        -- Insert the user data into table
        INSERT INTO product_inventory (product_id, quantity)
        VALUES (p_product_id, p_quantity);

        -- Get the last inserted user ID
        SET insertedID = LAST_INSERT_ID();

        -- SELECT 'User registered successfully' AS message;
        SET p_message = 'record added successfully';

        -- Return inserted data along with the message
        SELECT insertedID, p_message AS message;

    END IF;
END;