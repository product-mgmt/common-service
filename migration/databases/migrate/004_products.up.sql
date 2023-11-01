-- up

CREATE TABLE `product_category` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(32) NOT NULL,
    `description` varchar(255) NOT NULL,
    `status` ENUM('pending', 'active', 'delete') DEFAULT 'pending' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE `products` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(250) NOT NULL,
    `description` varchar(250) NOT NULL,
    `sku` varchar(250) NOT NULL UNIQUE,
    `category_id` int(32) NOT NULL,
    `price` int(64) NOT NULL DEFAULT 0,
    `status` ENUM('pending', 'active', 'delete') DEFAULT 'pending' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Foreign Key(`category_id`) references product_category(`id`)
) ENGINE=InnoDB;

CREATE TABLE `product_inventory` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `product_id` int(32) NOT NULL,
    `quantity` int(32) NOT NULL,
    `status` ENUM('pending', 'active', 'delete') DEFAULT 'pending' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Foreign Key(`product_id`) references products(`id`)
) ENGINE=InnoDB;