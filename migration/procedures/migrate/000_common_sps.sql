-- GET ONE RECORD
CREATE PROCEDURE IF NOT EXISTS sp_GetRecordByTableColAndVal(
    IN table_name VARCHAR(250),
    IN col_name VARCHAR(250),
    IN col_value VARCHAR(250)
)
BEGIN
        DECLARE recordFound INT DEFAULT 0;

        -- Construct the dynamic SQL query
        SET @sql = CONCAT(
            'SELECT 1 INTO @recordFound FROM ', table_name,
            ' WHERE status != ? AND ', col_name, ' = ?');
        PREPARE stmt FROM @sql;
        
        -- Execute the dynamic query with the parameter
        SET @paramValue = col_value;
        SET @status = 'delete';
        EXECUTE stmt USING @status, @paramValue;

        -- Deallocate the prepared statement
        DEALLOCATE PREPARE stmt;

        IF @recordFound = 0 THEN
            SET @errorMessage = CONCAT('Record not found for ', col_name, ' = ', col_value);
            SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = @errorMessage;
        ELSE
            SET @query = CONCAT(
                'SELECT * FROM ',table_name,
                ' WHERE status != ? AND ', col_name, ' = ?');
            PREPARE stmt FROM @query;
            SET @paramValue = col_value;
            SET @status = 'delete';
            EXECUTE stmt USING @status, @paramValue;
            DEALLOCATE PREPARE stmt;
        END IF;
END;

-- GET ALL RECORDS
CREATE PROCEDURE IF NOT EXISTS sp_GetRecordsByTable(
    table_name VARCHAR(255),
    search_column VARCHAR(255),
    search_query VARCHAR(255),
    sort_column VARCHAR(255),
    sort_order VARCHAR(255),
    page_offset INT,
    page_limit INT
)
BEGIN
    SET @sql = CONCAT('SELECT * FROM ', table_name, ' WHERE status != ? AND ', search_column, ' LIKE ''%', search_query, '%'' ORDER BY ', sort_column, ' ',sort_order, ' LIMIT ', page_offset, ', ', page_limit);

    PREPARE stmt FROM @sql;
    SET @status = 'delete';
    EXECUTE stmt USING @status;
    DEALLOCATE PREPARE stmt;

END;

-- HARD DELETE
CREATE PROCEDURE IF NOT EXISTS sp_HardDeleteRecordByTableColAndVal(
    IN table_name VARCHAR(250),
    IN col_name VARCHAR(250),
    IN col_value VARCHAR(250)
)
BEGIN
    SET @query = CONCAT(
        'DELETE FROM ',table_name,
        ' WHERE ', col_name, ' = ?');
    PREPARE stmt FROM @query;
    SET @paramValue = col_value;
    EXECUTE stmt USING @paramValue;
    DEALLOCATE PREPARE stmt;
END;

-- SOFT DELETE
CREATE PROCEDURE IF NOT EXISTS sp_SoftDeleteRecordByTableColAndVal(
    IN table_name VARCHAR(250),
    IN col_name VARCHAR(250),
    IN col_value VARCHAR(250)
)
BEGIN
    SET @query = CONCAT(
        'UPDATE ',table_name,
        ' SET status=delete'
        ' WHERE ', col_name, ' = ?');
    PREPARE stmt FROM @query;
    SET @paramValue = col_value;
    EXECUTE stmt USING @paramValue;
    DEALLOCATE PREPARE stmt;
END;