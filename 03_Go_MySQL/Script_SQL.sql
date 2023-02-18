select * from products

use mlbauto; 

CREATE TABLE IF NOT EXISTS Alumnos3
 (
 id int  primary key not null,
 name varchar (25) not null,
 observations Varchar (100),
 price int not null,
 created_at TIMESTAMP NOT NULL DEFAULT now(),
 update_at TIMESTAMP
 );