CREATE SEQUENCE product_id_seq;

CREATE TABLE IF NOT EXISTS Product (
    product_id     VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT nextval('product_id_seq'),
    name            VARCHAR(255) NOT NULL, 
    price           INTEGER NOT NULL,
    imageUrl        VARCHAR(255) NOT NULL,
    stock           INTEGER NOT NULL,
    condition       VARCHAR(255) NOT NULL, -- Assuming Condition is a string type
    is_Purchaseable  BOOLEAN NOT NULL,
    user_id          VARCHAR(255) NOT NULL
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

