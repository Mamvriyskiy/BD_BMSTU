CREATE TABLE IF NOT EXISTS client (
    client_id int,
    first_name varchar(50),
    last_name varchar(50),
    phone_number varchar(20),
    date_of_birth date,
    sex varchar(2),
    passport varchar(15)
);

CREATE TABLE IF NOT EXISTS booking (
   booking_id int,
   date_arrive date,
   date_departure date,
   count_room int,
   invoice_id int,
   client_id int
);

CREATE TABLE IF NOT EXISTS invoiceForPayment (
 invoice_id integer,
 booking_id integer,
 date_of_issue date,
 total_price integer
);

CREATE TABLE IF NOT EXISTS invoiceServices (
    invoice_id int,
    services_id int
);

CREATE TABLE IF NOT EXISTS BookingRoom (
    booking_id int,
    room_id int
);

CREATE TABLE IF NOT EXISTS services (
    services_id int,
    name_services varchar(50),
    price int
);

CREATE TABLE IF NOT EXISTS room (
    room_id int,
    room_number int,
    categorie varchar(50),
    price int
);