
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_dataset
-- ----------------------------
DROP TABLE IF EXISTS `tb_dataset`;
CREATE TABLE `tb_dataset` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL COMMENT '数据集名称',
  `center_id` int NOT NULL COMMENT '分中心id',
  `mid` int NOT NULL COMMENT '数据集主中心id',
  `sid` int NOT NULL COMMENT '数据集分中心id',
  `non_public` varchar(255) NOT NULL COMMENT '是否是公开数据集',
  `provider_type` varchar(255) NOT NULL COMMENT '所在的存储类型 (1 obs 2 自定义)',
  `volume_id` int DEFAULT NULL COMMENT 'provider_type 为自定义是,此字段为必填',
  `sub_path` varchar(255) DEFAULT NULL COMMENT 'provider_type 为自定义是,此字段为必填',
  `user_id` int NOT NULL COMMENT '用户di',
  `updated_at` datetime DEFAULT NULL COMMENT '修改是哪',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tb_model
-- ----------------------------
DROP TABLE IF EXISTS `tb_model`;
CREATE TABLE `tb_model` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL COMMENT '模型名称',
  `center_id` int NOT NULL COMMENT '分中心id',
  `mid` int NOT NULL COMMENT '模型主中心id',
  `sid` int NOT NULL COMMENT '模型分中心id',
  `non_public` varchar(255) NOT NULL COMMENT '是否是公开模型',
  `provider_type` varchar(255) NOT NULL COMMENT '所在的存储类型 (1 obs 2 自定义)',
  `volume_id` int DEFAULT NULL COMMENT 'provider_type 为自定义是,此字段为必填',
  `sub_path` varchar(255) DEFAULT NULL COMMENT 'provider_type 为自定义是,此字段为必填',
  `user_id` int NOT NULL COMMENT '用户id',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;
