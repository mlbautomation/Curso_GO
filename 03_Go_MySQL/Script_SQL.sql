select * from products

use mlbauto; 

CREATE TABLE IF NOT EXISTS invoice_items(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	client_id INT NOT NULL,
	product_id INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP,
	CONSTRAINT invoice_items_client_id_fk FOREIGN KEY (client_id) REFERENCES clients (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)