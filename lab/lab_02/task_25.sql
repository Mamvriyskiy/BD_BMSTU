WITH RankedRooms AS (
  SELECT *,
         AVG(price) OVER(PARTITION BY price) AS AvgPrice,
         ROW_NUMBER() OVER(PARTITION BY price ORDER BY room_id) AS RowNum
  FROM room
  WHERE price > 5000
)
SELECT room_id, price
  FROM RankedRooms
  WHERE RowNum > 1;
