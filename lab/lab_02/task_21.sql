ALTER TABLE bookingroom
DROP CONSTRAINT bookingroom_room_id_fkey;  -- Удаляем текущее ограничение

ALTER TABLE bookingroom
ADD CONSTRAINT bookingroom_room_id_fkey
FOREIGN KEY (room_id) REFERENCES room (room_id)
ON DELETE CASCADE;

delete from room 
where price > (select AVG(price) from room) 
and room_number > 100 AND room_number < 150;
