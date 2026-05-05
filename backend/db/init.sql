-- Sample assets for demo day
INSERT INTO assets (id, asset_code, name, category, serial_number,
  purchase_date, purchase_value, current_value, location, status)
VALUES
  ('asset-001','ASSET-0001','Dell Laptop XPS 15','IT','DL-2024-001',
   '2024-01-15',250000,200000,'IT Department','ACTIVE'),
  ('asset-002','ASSET-0002','Office Chair - Ergonomic','Furniture','OC-2024-001',
   '2024-02-10',45000,40000,'Floor 2 - Room 201','ACTIVE'),
  ('asset-003','ASSET-0003','Toyota Corolla 2023','Vehicle','TC-WP-2023-01',
   '2023-06-01',7500000,6800000,'Company Garage','ACTIVE'),
  ('asset-004','ASSET-0004','HP LaserJet Pro','IT','HP-2024-002',
   '2024-03-01',85000,75000,'Admin Office','MAINTENANCE');