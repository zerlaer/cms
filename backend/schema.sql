-- 创建数据库
CREATE DATABASE IF NOT EXISTS `cms` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `cms`;

-- 元件表
CREATE TABLE IF NOT EXISTS `components` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `productCode` varchar(100) NOT NULL COMMENT '商品编号',
  `brand` varchar(100) NOT NULL COMMENT '品牌',
  `category` varchar(100) NOT NULL COMMENT '分类',
  `model` varchar(200) NOT NULL COMMENT '厂家型号',
  `package` varchar(100) NOT NULL COMMENT '封装',
  `name` varchar(200) NOT NULL COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT 0 COMMENT '订购数量',
  `price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '商品金额',
  `stockIn` int NOT NULL DEFAULT 0 COMMENT '入库数量',
  `stockOut` int NOT NULL DEFAULT 0 COMMENT '出库数量',
  `currentStock` int NOT NULL DEFAULT 0 COMMENT '当前库存',
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_productCode` (`productCode`),
  KEY `idx_brand` (`brand`),
  KEY `idx_category` (`category`),
  KEY `idx_model` (`model`),
  KEY `idx_name` (`name`),
  KEY `idx_currentStock` (`currentStock`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='电子元件表';

-- 库存记录表
CREATE TABLE IF NOT EXISTS `stockRecords` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `componentId` bigint unsigned NOT NULL COMMENT '元件 ID',
  `type` varchar(20) NOT NULL COMMENT '类型：in/out',
  `quantity` int NOT NULL COMMENT '数量',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `createdAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_componentId` (`componentId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='库存记录表';

-- 分类表
CREATE TABLE IF NOT EXISTS `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

-- 分类关键词表
CREATE TABLE IF NOT EXISTS `categoryKeywords` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `categoryId` bigint unsigned NOT NULL COMMENT '分类ID',
  `keyword` varchar(100) NOT NULL COMMENT '关键词',
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_categoryId` (`categoryId`),
  CONSTRAINT `fk_categoryKeywords_categoryId` FOREIGN KEY (`categoryId`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类关键词表';
