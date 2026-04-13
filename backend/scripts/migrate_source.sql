-- 添加购买渠道字段
ALTER TABLE components ADD COLUMN source VARCHAR(50) COMMENT '购买渠道' AFTER productCode;

-- 更新现有数据，根据商品编号自动设置购买渠道
UPDATE components SET source = '立创商城' WHERE productCode LIKE 'C%';
UPDATE components SET source = '淘宝' WHERE source IS NULL OR source = '';

-- 设置默认值
ALTER TABLE components MODIFY COLUMN source VARCHAR(50) NOT NULL DEFAULT '淘宝' COMMENT '购买渠道';
