CREATE DATABASE amsa;

-- COMPANY TABLE

CREATE TABLE company (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    country VARCHAR(50) NULL,
    address VARCHAR(150) NULL,
    description VARCHAR(150) NULL,
    phone VARCHAR(50) NULL,
    createdAt DATETIME,
    updatedAt DATETIME
);

CREATE TRIGGER before_insert_company
BEFORE INSERT ON company
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_company
BEFORE UPDATE ON company
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END COMPANY TABLE


-- USERS TABLE

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NULL,
    second_last_name VARCHAR(50) NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    company_id INT NOT NULL,
    activated BOOLEAN NOT NULL DEFAULT true,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_users
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_users
BEFORE UPDATE ON users
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END USERS TABLE


-- CLIENTS TABLE

CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(50) NULL,
    address VARCHAR(50) NULL,
    email VARCHAR(50)  NULL,
    description VARCHAR(50)  NULL,
    activated VARCHAR(50) NOT NULL,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_clients
BEFORE INSERT ON clients
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_clients
BEFORE UPDATE ON clients
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END USERS TABLE


-- SUPPLIERS TABLE

CREATE TABLE suppliers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    executive VARCHAR(100) NULL,
    phone VARCHAR(50) NULL,
    category VARCHAR(50) NULL,
    address VARCHAR(50) NULL,
    email VARCHAR(50)  NULL,
    description VARCHAR(50)  NULL,
    activated VARCHAR(50) NOT NULL,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_suppliers
BEFORE INSERT ON suppliers
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_suppliers
BEFORE UPDATE ON suppliers
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END USERS TABLE



-- CATEGORY TABLE

CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT  NULL,
    description VARCHAR(100) NULL,
    activated BOOLEAN NOT NULL DEFAULT true,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_categories
BEFORE INSERT ON categories
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_categories
BEFORE UPDATE ON categories
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END CATEGORIE TABLE



-- TYPE PRODUCT TABLE

CREATE TABLE product_type (
    id INT AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(100) NOT  NULL,
    description VARCHAR(100) NULL,
    activated BOOLEAN NOT NULL DEFAULT true,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_product_type
BEFORE INSERT ON product_type
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_product_type
BEFORE UPDATE ON product_type
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END TYPE PRODUCT TABLE


-- TYPE PRODUCT INSERT
INSERT INTO `product_type` (`type`, `description`, `activated`)
 VALUES ('Venta', 'Productos de Venta', True);

 INSERT INTO `product_type` (`type`, `description`, `activated`)
 VALUES ('Compra', 'Productos de Compra', True); 

-- END TYPE PRODUCT INSERT




-- Units TABLE

CREATE TABLE measurement_units (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT  NULL,
    description VARCHAR(100) NULL,
    activated BOOLEAN NOT NULL DEFAULT true,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_measurement_units
BEFORE INSERT ON measurement_units
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_measurement_units
BEFORE UPDATE ON measurement_units
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END CATEGORIE TABLE





-- PRODUCTS TABLE

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    clave VARCHAR(100)  NULL,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL DEFAULT "N/A",
    unit VARCHAR(50)  NOT NULL DEFAULT "N/A",
    costo DECIMAL(15,4)  NULL,
    precio DECIMAL(15,4)  NULL,
    iva_percentage DECIMAL(15,4) NULL;
    note VARCHAR(50) NULL,
    activated BOOLEAN NOT NULL DEFAULT true,
    company_id INT NOT NULL,
    category_id INT NOT NULL,
    unit_id INT NOT NULL,
    type_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (unit_id) REFERENCES measurementUnits(id),
    FOREIGN KEY (type_id) REFERENCES product_type(id)
);

CREATE TRIGGER before_insert_products
BEFORE INSERT ON products
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_products
BEFORE UPDATE ON products
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

-- END PRODUCTS TABLE


-- Creación de la tabla Ventas
CREATE TABLE sales (
    id INT PRIMARY KEY AUTO_INCREMENT,
    date DATE NOT NULL,
    client_id INT,
    total DECIMAL(10, 2),
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (clientID) REFERENCES clients(id)
);

CREATE TRIGGER before_insert_sales
BEFORE INSERT ON sales
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_sales
BEFORE UPDATE ON sales
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;




-- Creación de la tabla Detalles de Venta
CREATE TABLE sale_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    sale_id INT,
    product_id INT,
    product_name VARCHAR(100),
    product_clave VARCHAR(100),
    amount INT NOT NULL,
    precio DECIMAL(10, 2) NOT NULL,
    iva DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (sale_id) REFERENCES sales(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TRIGGER before_insert_sale_details
BEFORE INSERT ON saleDetails
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_sale_details
BEFORE UPDATE ON saleDetails
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;



----- ORDENES DE COMPRA ----

-- Creación de la tabla de Órdenes de Compra
CREATE TABLE purchase_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    folio VARCHAR(50),
    supplier_id INT,
    order_date DATE,
    total DECIMAL(10, 2),
    status VARCHAR(50) NULL,
    date_linked DATETIME,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (company_id) REFERENCES company(id)
);

CREATE TRIGGER before_insert_purchase_orders
BEFORE INSERT ON purchase_orders
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_purchase_orders
BEFORE UPDATE ON purchase_orders
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;


-- Creación de la tabla de Detalles de Orden de Compra
CREATE TABLE purchase_order_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT,
    product_name VARCHAR(100) NOT NULL,
    product_clave VARCHAR(100) NOT NULL,
    unitary_price DECIMAL(10, 2),
    iva DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (order_id) REFERENCES purchase_orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TRIGGER before_insert_purchase_order_details
BEFORE INSERT ON purchase_order_details
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_purchase_order_details
BEFORE UPDATE ON purchase_order_details
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;

















-- Creación de la tabla Facturas
CREATE TABLE bills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    folio VARCHAR(50),
    date DATE NOT NULL,
    client_id INT,
    total DECIMAL(10, 2),
    status VARCHAR(50) NULL,
    date_linked DATETIME NUll,
    company_id INT NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TRIGGER before_insert_bills
BEFORE INSERT ON bills
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_bills
BEFORE UPDATE ON bills
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;


-- Creación de la tabla Detalles de Factura
CREATE TABLE bill_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    bill_id INT,
    product_id INT,
    quantity INT NOT NULL,
    unitary_price DECIMAL(10, 2) NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    product_clave VARCHAR(100) NOT NULL,
    iva DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY (bill_id) REFERENCES bills(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TRIGGER before_insert_bill_details
BEFORE INSERT ON bill_details
FOR EACH ROW
BEGIN
    SET NEW.createdAt = NOW();
END;

CREATE TRIGGER before_update_bill_details
BEFORE UPDATE ON bill_details
FOR EACH ROW
BEGIN
    SET NEW.updatedAt = NOW();
END;
