\copy client FROM './mnt/data/client.csv' DELIMITER ',' CSV;

\copy booking FROM './mnt/data/booking.csv' DELIMITER ',' CSV;

\copy room FROM './mnt/data/room.csv' DELIMITER ',' CSV;

\copy bookingroom FROM './mnt/data/bookingroom.csv' DELIMITER ',' CSV;

\copy invoiceForPayment FROM './mnt/data/invoiceforpayment.csv' DELIMITER ',' CSV;

\copy services FROM './mnt/data/services.csv' DELIMITER ',' CSV;

\copy invoiceServices FROM './mnt/data/invoiceservices.csv' DELIMITER ',' CSV;


