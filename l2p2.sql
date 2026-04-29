-- Table: Users
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    last_login_date TIMESTAMP,
    jwt_token TEXT
);

-- Table: Customers
CREATE TABLE Customers (
    customer_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15),
    address TEXT
);

-- Table: Tours
CREATE TABLE Tours (
    tour_id SERIAL PRIMARY KEY,
    tour_name VARCHAR(255) NOT NULL,
    tour_description TEXT,
    tour_price DECIMAL(10, 2) NOT NULL,
    tour_duration VARCHAR(50) NOT NULL
);

-- Table: Bookings
CREATE TABLE Bookings (
    booking_id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES Customers(customer_id),
    booking_date DATE NOT NULL,
    booking_status VARCHAR(50) NOT NULL
);

-- Table: Tour_Bookings
CREATE TABLE Tour_Bookings (
    tour_booking_id SERIAL PRIMARY KEY,
    booking_id INT REFERENCES Bookings(booking_id),
    tour_id INT REFERENCES Tours(tour_id),
    tour_start_date DATE NOT NULL,
    tour_end_date DATE NOT NULL,
    num_of_people INT NOT NULL
);

-- Table: Payments
CREATE TABLE Payments (
    payment_id SERIAL PRIMARY KEY,
    booking_id INT REFERENCES Bookings(booking_id),
    payment_date DATE NOT NULL,
    payment_amount DECIMAL(10, 2) NOT NULL
);


-- Insert data into Users
INSERT INTO Users (email, password_hash, last_login_date, jwt_token) VALUES
('john.doe@example.com', 'hashed_password1', '2024-08-16 10:00:00', 'jwt_token_value1'),
('jane.smith@example.com', 'hashed_password2', '2024-08-16 10:15:00', 'jwt_token_value2'),
('mark.brown@example.com', 'hashed_password3', '2024-08-16 10:30:00', 'jwt_token_value3'),
('lucy.jones@example.com', 'hashed_password4', '2024-08-16 10:45:00', 'jwt_token_value4'),
('michael.taylor@example.com', 'hashed_password5', '2024-08-16 11:00:00', 'jwt_token_value5'),
('sarah.davis@example.com', 'hashed_password6', '2024-08-16 11:15:00', 'jwt_token_value6'),
('david.martin@example.com', 'hashed_password7', '2024-08-16 11:30:00', 'jwt_token_value7'),
('anna.jackson@example.com', 'hashed_password8', '2024-08-16 11:45:00', 'jwt_token_value8'),
('peter.moore@example.com', 'hashed_password9', '2024-08-16 12:00:00', 'jwt_token_value9'),
('emma.white@example.com', 'hashed_password10', '2024-08-16 12:15:00', 'jwt_token_value10');

-- Insert data into Customers
INSERT INTO Customers (user_id, name, email, phone_number, address) VALUES
(1, 'John Doe', 'john.doe@example.com', '0123456789', '123 Elm Street'),
(2, 'Jane Smith', 'jane.smith@example.com', '0123456790', '456 Oak Avenue'),
(3, 'Mark Brown', 'mark.brown@example.com', '0123456791', '789 Pine Road'),
(4, 'Lucy Jones', 'lucy.jones@example.com', '0123456792', '321 Maple Lane'),
(5, 'Michael Taylor', 'michael.taylor@example.com', '0123456793', '654 Birch Blvd'),
(6, 'Sarah Davis', 'sarah.davis@example.com', '0123456794', '987 Cedar Court'),
(7, 'David Martin', 'david.martin@example.com', '0123456795', '159 Oak Street'),
(8, 'Anna Jackson', 'anna.jackson@example.com', '0123456796', '753 Willow Way'),
(9, 'Peter Moore', 'peter.moore@example.com', '0123456797', '951 Pine Terrace'),
(10, 'Emma White', 'emma.white@example.com', '0123456798', '258 Maple Plaza');

-- Insert data into Tours
INSERT INTO Tours (tour_name, tour_description, tour_price, tour_duration) VALUES
('Adventure to Mount Everest', 'Experience the thrill of a lifetime with a trek to Mount Everest.', 5000.00, '7 days'),
('Trip to Bali', 'Relax and unwind in the tropical paradise of Bali.', 2000.00, '5 days'),
('Safari in Kenya', 'Witness the majestic wildlife of Kenya on a safari adventure.', 3500.00, '6 days'),
('Cruise to the Caribbean', 'Sail through the Caribbean on a luxury cruise.', 4500.00, '8 days'),
('Tour of Rome', 'Explore the ancient wonders of Rome.', 3000.00, '4 days'),
('Great Wall of China Tour', 'Walk along the historic Great Wall of China.', 2500.00, '5 days'),
('Tokyo City Tour', 'Discover the vibrant culture of Tokyo.', 4000.00, '3 days'),
('Northern Lights in Iceland', 'Witness the spectacular Northern Lights in Iceland.', 5500.00, '6 days'),
('Amazon Rainforest Expedition', 'Explore the diverse ecosystem of the Amazon Rainforest.', 6000.00, '10 days'),
('Grand Canyon Adventure', 'Hike through the stunning Grand Canyon.', 3000.00, '3 days');

-- Insert data into Bookings
INSERT INTO Bookings (customer_id, booking_date, booking_status) VALUES
(1, '2024-08-10', 'Completed'),
(2, '2024-08-11', 'Completed'),
(3, '2024-08-12', 'Pending'),
(4, '2024-08-13', 'Pending'),
(5, '2024-08-14', 'Completed'),
(6, '2024-08-15', 'Completed'),
(7, '2024-08-16', 'Pending'),
(8, '2024-08-17', 'Pending'),
(9, '2024-08-18', 'Completed'),
(10, '2024-08-19', 'Pending');

-- Insert data into Tour_Bookings
INSERT INTO Tour_Bookings (booking_id, tour_id, tour_start_date, tour_end_date, num_of_people) VALUES
(1, 1, '2024-08-16', '2024-08-20', 2),
(2, 2, '2024-08-17', '2024-08-21', 4),
(3, 3, '2024-08-18', '2024-08-22', 1),
(4, 4, '2024-08-19', '2024-08-23', 3),
(5, 5, '2024-08-20', '2024-08-24', 2),
(6, 6, '2024-08-21', '2024-08-25', 5),
(7, 7, '2024-08-22', '2024-08-26', 3),
(8, 8, '2024-08-23', '2024-08-27', 4),
(9, 9, '2024-08-24', '2024-08-28', 2),
(10, 10, '2024-08-25', '2024-08-29', 1);

-- Insert data into Payments
INSERT INTO Payments (booking_id, payment_date, payment_amount) VALUES
(1, '2024-08-11', 5000.00),
(2, '2024-08-12', 2000.00),
(3, '2024-08-13', 3500.00),
(4, '2024-08-14', 4500.00),
(5, '2024-08-15', 3000.00),
(6, '2024-08-16', 2500.00),
(7, '2024-08-17', 4000.00),
(8, '2024-08-18', 5500.00),
(9, '2024-08-19', 6000.00),
(10, '2024-08-20', 3000.00);
