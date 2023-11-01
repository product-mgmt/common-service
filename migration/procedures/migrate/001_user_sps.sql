CREATE PROCEDURE IF NOT EXISTS sp_CreateUser(
    IN p_name VARCHAR(50),
    IN p_email VARCHAR(250),
    IN p_password VARCHAR(250)
)
BEGIN
    DECLARE error_message VARCHAR(255);
    DECLARE user_id INT;
    DECLARE p_message VARCHAR(255);

    -- Check if the record with the specified ID exists
    IF EXISTS (SELECT 1 FROM users WHERE email = p_email) THEN
        -- Record exists, perform the update or delete operation
        -- Record does not exist, send an error message
        SET error_message = CONCAT(p_email, ' already registered');
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = error_message;

    ELSE
        -- Insert the user data into table
        INSERT INTO users (name, email, password)
        VALUES (p_name, p_email, p_password);

        -- Get the last inserted user ID
        SET user_id = LAST_INSERT_ID();

        -- SELECT 'User registered successfully' AS message;
        SET p_message = 'user registered successfully';

        -- Return inserted data along with the message
        SELECT user_id AS insertedID, p_message AS message;

    END IF;
END;