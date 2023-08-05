# Write your MySQL query statement below
select
    w.id
from Weather w, Weather y
where
    datediff(w.recordDate, y.recordDate) = 1
    and w.temperature > y.temperature