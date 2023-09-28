--task_08(Инструкция SELECT, использующая скалярные подзапросы в выражениях столбцов.)
SELECT *,
       (SELECT AVG(date_departure - date_arrive)
        FROM booking b
        WHERE b.client_id = c.client_id) AS AvgDayInHotel,
        (SELECT COUNT(*) 
        FROM booking b
        WHERE b.client_id = c.client_id) AS VisitCount
FROM client c;
