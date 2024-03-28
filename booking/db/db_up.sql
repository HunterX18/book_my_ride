
CREATE TABLE rides (
    id BIGSERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    status VARCHAR(20) NOT NULL, 
    status_changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_status_changed_at() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'reserved' THEN
        NEW.status_changed_at = CURRENT_TIMESTAMP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER status_update_trigger 
BEFORE INSERT OR UPDATE ON rides
FOR EACH ROW
EXECUTE FUNCTION update_status_changed_at();