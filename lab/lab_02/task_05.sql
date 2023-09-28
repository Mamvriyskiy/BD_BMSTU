--task_05(exists - проверка на то, есть ли заказы с 5 комнатами)
SELECT *
FROM booking i
WHERE EXISTS (
    SELECT i2.booking_id
    FROM booking i2
    WHERE i.count_room = 5
);
