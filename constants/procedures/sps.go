package procedures

// sp_CreateUser(name, email, password)
const SP_CREATE_USER = "CALL sp_CreateUser(?, ?, ?)"

// sp_GetRecordsByTable(table_name, search_column, search_query, sort_column, sort_order, page_offset, page_limit)
const SP_GETRECORDS = "CALL sp_GetRecordsByTable(?, ?, ?, ?, ?, ?, ?)"

// sp_GetRecordByTableColAndVal(table_name, col_name, col_value)
const SP_GETRECORD = "CALL sp_GetRecordByTableColAndVal(?, ?, ?)"

// sp_SoftDeleteRecordByTableColAndVal(table_name, col_name, col_value)
const SP_SOFTDELETE = "CALL sp_SoftDeleteRecordByTableColAndVal(?, ?, ?)"

// sp_HardDeleteRecordByTableColAndVal(table_name, col_name, col_value)
const SP_HARDDELETE = "CALL sp_HardDeleteRecordByTableColAndVal(?, ?, ?)"

// sp_CreateProductCategory(p_name, p_description)
const SP_CREATE_PRODUCT_CATEGORY = "CALL sp_CreateProductCategory(?, ?)"

// sp_CreateProducts(p_name, p_description, p_sku, p_category_id, p_price)
const SP_CREATE_PRODUCTS = "CALL sp_CreateProducts(?, ?, ?, ?, ?)"

// sp_CreateProductInventory(p_product_id, p_quantity)
const SP_CREATE_PRODUCT_INVENTORY = "CALL sp_CreateProductInventory(?, ?)"
