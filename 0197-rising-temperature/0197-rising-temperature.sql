# Write your MySQL query statement below
select
    w.id
from Weather w
where
    w.temperature > (
        # yesterday
        select
            y.temperature
        from Weather y
        where y.recordDate = subdate(w.recordDate, 1)
        order by y.recordDate desc
        limit 1
    )