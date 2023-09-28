CREATE TABLE public.MyEmployees (
    EmployeeID smallint NOT NULL,
    FirstName varchar(30) NOT NULL,
    LastName varchar(40) NOT NULL,
    Title varchar(50) NOT NULL,
    DeptID smallint NOT NULL,
    ManagerID int NULL,
    CONSTRAINT PK_EmployeeID PRIMARY KEY (EmployeeID)
);

-- Заполнение таблицы значениями.
INSERT INTO public.MyEmployees
VALUES (1, N'Иван', N'Петров', N'Главный исполнительный директор', 16, NULL);

-- Использование WITH RECURSIVE для определения рекурсивного CTE
WITH RECURSIVE DirectReports (ManagerID, EmployeeID, Title, DeptID, Level) AS
(
    -- Определение начального элемента
    SELECT e.ManagerID, e.EmployeeID, e.Title, e.DeptID, 0 AS Level
    FROM public.MyEmployees AS e
    WHERE ManagerID IS NULL
    UNION ALL
    -- Определение рекурсивного элемента
    SELECT e.ManagerID, e.EmployeeID, e.Title, e.DeptID, dr.Level + 1
    FROM public.MyEmployees AS e
    INNER JOIN DirectReports AS dr ON e.ManagerID = dr.EmployeeID
)
-- Инструкция, использующая рекурсивный CTE
SELECT ManagerID, EmployeeID, Title, DeptID, Level
FROM DirectReports;
