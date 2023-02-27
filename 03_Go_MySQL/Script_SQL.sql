select * from products
select * from clients
select * from invoice_items

SET GLOBAL time_zone = 'America/Peru'

use mlbauto; 

CREATE TABLE IF NOT EXISTS products
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP
	);
    
INSERT INTO products (name, price) VALUES ('Marlon', '25')

SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id = 2

SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id like 3