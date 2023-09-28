
SELECT *,
  AVG(price) OVER(PARTITION BY price) AS AvgPrice
FROM room where price > 5000;