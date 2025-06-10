DROP TABLE students;
DROP TABLE courses;
DROP TABLE enrollments;


CREATE TABLE students
(
    id    serial PRIMARY KEY,
    name  VARCHAR,
    phone VARCHAR
);

CREATE TABLE courses
(
    id          SERIAL PRIMARY KEY,
    course_name VARCHAR,
    start_date  DATE
);

CREATE TABLE enrollments
(
    student_id INT,
    course_id  INT,
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES students (id),
    FOREIGN KEY (course_id) REFERENCES courses (id)
);

-- Студенты
INSERT INTO students (id, name, phone)
VALUES (1, 'Ivan', 12345),
       (2, 'Ali', 54321),
       (3, 'Sara', 32145);

-- Курсы
INSERT INTO courses (id, course_name, start_date)
VALUES (1, 'Mathematics', null),
       (2, 'English', null),
       (3, 'Programming', null);

-- Записи в enrollments (5 записей)
INSERT INTO enrollments (student_id, course_id)
VALUES (1, 1),
       (1, 2),
       (2, 2),
       (2, 3),
       (3, 1);

SELECT s.id, s.name, c.id, c.course_name
FROM enrollments e
         JOIN students s ON e.student_id = s.id
         JOIN courses c ON e.course_id = c.id;

SELECT s.id, s.name
FROM students s
         LEFT JOIN enrollments e ON s.id = e.student_id
WHERE e.course_id IS NULL;

