-- up

CREATE TABLE `orders` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` int(32) NOT NULL,
    `total` int(64) NOT NULL,
    `status` ENUM('success', 'failed', 'processing', 'processed') DEFAULT 'processing' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Foreign Key(`user_id`) references users(`id`)
) ENGINE=InnoDB;

CREATE TABLE `order_items` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `order_id` int(32) NOT NULL,
    `product_id` int(32) NOT NULL,
    `quantity` int(32) NOT NULL,
    `status` ENUM('success', 'failed', 'processing', 'processed') DEFAULT 'processing' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    Foreign Key(`order_id`) references orders(`id`),
    Foreign Key(`product_id`) references products(`id`)
) ENGINE=InnoDB;