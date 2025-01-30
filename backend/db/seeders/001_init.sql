-- Seed users
INSERT INTO users (name, email, password, created_at) VALUES
('Alice Smith', 'alice@example.com', 'hashed_password_1', NOW()),
('Bob Johnson', 'bob@example.com', 'hashed_password_2', NOW()),
('Charlie Brown', 'charlie@example.com', 'hashed_password_3', NOW()),
('Diana Prince', 'diana@example.com', 'hashed_password_4', NOW()),
('Ethan Hunt', 'ethan@example.com', 'hashed_password_5', NOW()),
('Fiona Glenanne', 'fiona@example.com', 'hashed_password_6', NOW()),
('George Washington', 'george@example.com', 'hashed_password_7', NOW()),
('Hannah Baker', 'hannah@example.com', 'hashed_password_8', NOW()),
('Ian Malcolm', 'ian@example.com', 'hashed_password_9', NOW()),
('Julia Child', 'julia@example.com', 'hashed_password_10', NOW());

-- Seed todos for each user
INSERT INTO todos (title, content, status, created_at, updated_at, user_id) VALUES
('Buy groceries', 'Milk, Bread, Cheese, Fruits', 'incomplete', NOW(), NOW(), 1),
('Read book', 'Finish reading the latest novel', 'incomplete', NOW(), NOW(), 1),
('Go for a walk', 'Walk in the park for at least 30 minutes', 'incomplete', NOW(), NOW(), 2),
('Finish project', 'Complete the project report for work', 'complete', NOW(), NOW(), 2),
('Bake a cake', 'Bake a chocolate cake for the party', 'incomplete', NOW(), NOW(), 3),
('Learn Golang', 'Finish the Golang course on Udemy', 'incomplete', NOW(), NOW(), 3),
('Call mom', 'Catch up with mom on the phone', 'incomplete', NOW(), NOW(), 4),
('Plan vacation', 'Plan a trip to Hawaii', 'incomplete', NOW(), NOW(), 4),
('Organize closet', 'Make room by organizing clothes', 'incomplete', NOW(), NOW(), 5),
('Volunteer at shelter', 'Help out at the local animal shelter', 'incomplete', NOW(), NOW(), 5),
('Submit taxes', 'Complete and submit tax returns', 'incomplete', NOW(), NOW(), 6),
('Workout at gym', 'Go to the gym three times a week', 'incomplete', NOW(), NOW(), 6),
('Fix the bike', 'Repair the broken bike tire', 'incomplete', NOW(), NOW(), 7),
('Attend a webinar', 'Participate in the tech webinar', 'complete', NOW(), NOW(), 7),
('Plant a tree', 'Plant a tree in the backyard', 'incomplete', NOW(), NOW(), 8),
('Update resume', 'Revise and update my resume', 'incomplete', NOW(), NOW(), 8),
('Arrange meeting', 'Schedule a meeting with the team', 'complete', NOW(), NOW(), 9),
('Learn a new language', 'Start learning Spanish', 'incomplete', NOW(), NOW(), 9),
('Create a website', 'Build a personal portfolio website', 'incomplete', NOW(), NOW(), 10),
('Join a cooking class', 'Enroll in a local cooking class', 'incomplete', NOW(), NOW(), 10);