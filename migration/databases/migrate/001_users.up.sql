-- up
CREATE TABLE `users` (
    `id` int(32) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(250) NOT NULL,
    `email` varchar(250) NOT NULL UNIQUE,
    `password` varchar(250) NOT NULL,
    `role` ENUM('admin', 'user') DEFAULT 'user' NOT NULL,
    `status` ENUM('pending', 'active', 'blocked', 'delete') DEFAULT 'pending' NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;