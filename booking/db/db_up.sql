
CREATE TABLE rides (
    id bigserial PRIMARY KEY,
    owner_id INT NOT NULL,
    status varchar(20) NOT NULL, 
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);