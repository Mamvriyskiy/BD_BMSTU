SELECT c.first_name, c.last_name, b.date_arrive, r.categorie from client c join booking b on b.client_id = c.client_id 
join bookingroom br on br.booking_id = b.booking_id join room r on r.room_id = r.room_id
where b.date_arrive < now() and r.categorie = 'standard'