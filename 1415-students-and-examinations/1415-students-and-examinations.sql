# Write your MySQL query statement below
select
    s.student_id,
    s.student_name,
    sub.subject_name,
    coalesce(e.exams, 0) attended_exams
from Students s
cross join Subjects sub
left join (
    select
        student_id,
        subject_name,
        count(*) exams
    from Examinations
    group by student_id, subject_name
) e
    using(student_id, subject_name)
    # on e.student_id = s.student_id
    # and e.subject_name = sub.subject_name
order by s.student_id, sub.subject_name