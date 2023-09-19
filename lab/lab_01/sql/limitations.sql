ALTER TABLE client
ALTER COLUMN client_id SET NOT NULL,
ALTER COLUMN first_name SET NOT NULL,
ALTER COLUMN last_name SET NOT NULL,
ALTER COLUMN phone_number SET NOT NULL,
ALTER COLUMN date_of_birth SET NOT NULL,
ALTER COLUMN sex SET NOT NULL,
ALTER COLUMN passport SET NOT NULL,
ADD check (first_name != ''),
ADD check (last_name != ''),
ADD check (phone_number != ''),
ADD check (sex != ''),
ADD check (passport != ''),
ADD primary key (client_id);

ALTER TABLE booking
ALTER COLUMN booking_id SET NOT NULL,
ALTER COLUMN date_arrive SET NOT NULL,
ALTER COLUMN date_departure SET NOT NULL,
ALTER COLUMN count_room SET NOT NULL,
ALTER COLUMN invoice_id SET NOT NULL,
ALTER COLUMN client_id SET NOT NULL,
ADD check (count_room != 0),
ADD primary key (booking_id),
ADD FOREIGN KEY (client_id) REFERENCES client(client_id);


ALTER TABLE invoiceForPayment
ALTER COLUMN invoice_id SET NOT NULL,
ALTER COLUMN booking_id SET NOT NULL,
ALTER COLUMN date_of_issue SET NOT NULL,
ALTER COLUMN total_price SET NOT NULL,
ADD check (total_price > 0),
ADD FOREIGN KEY (booking_id) REFERENCES booking(booking_id),
ADD primary key (invoice_id);

ALTER TABLE services
ALTER COLUMN services_id SET NOT NULL,
ALTER COLUMN name_services SET NOT NULL,
ALTER COLUMN price SET NOT NULL,
ADD check (name_services != ''),
ADD check (price > 0),
ADD primary key (services_id);

ALTER TABLE room
ALTER COLUMN room_id SET NOT NULL,
ALTER COLUMN room_number SET NOT NULL,
ALTER COLUMN categorie SET NOT NULL,
ALTER COLUMN price SET NOT NULL,
ADD check (categorie != ''),
ADD check (price > 0),
ADD primary key (room_id);

ALTER TABLE invoiceServices
ALTER COLUMN invoice_id SET NOT NULL,
ALTER COLUMN services_id SET NOT NULL,
ADD FOREIGN KEY (invoice_id) REFERENCES invoiceForPayment(invoice_id),
ADD FOREIGN KEY (services_id) REFERENCES services(services_id);

ALTER TABLE BookingRoom
ALTER COLUMN booking_id SET NOT NULL,
ALTER COLUMN room_id SET NOT NULL,
ADD FOREIGN KEY (booking_id) REFERENCES booking(booking_id),
ADD FOREIGN KEY (room_id) REFERENCES room(room_id);

