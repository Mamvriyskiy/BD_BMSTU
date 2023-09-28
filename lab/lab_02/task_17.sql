INSERT INTO room (room_number, categorie, price)
SELECT room_number, categorie, price
FROM room
WHERE categorie = 'standard';
