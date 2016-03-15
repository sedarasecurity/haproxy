DELETE FROM plugin WHERE id = "81875";
DELETE FROM plugin_sid where plugin_id = "81875";
INSERT INTO plugin (id, type, name, description, product_type, vendor) VALUES (81875, 4, 'haproxy', 'HAProxy', 4, 'HAProxy');
INSERT INTO `plugin_sid` (`plugin_id`,`sid`,`category_id`,`subcategory_id`,`reliability`,`priority`,`name`) VALUES (81875, 1, 3, 30, 1, 1, 'HAProxy TCP Session');
