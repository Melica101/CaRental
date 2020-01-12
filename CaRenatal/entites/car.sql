-- CREATE TABLE cars (
--     ID VARCHAR(20) PRIMARY KEY,
--     Name VARCHAR(100),
--     Model VARCHAR(100),
--     Year TIMESTAMP,
--     Color VARCHAR(100),
--     Transmission VARCHAR(100),
--     FuelType VARCHAR(100),
--     FuelUsage VARCHAR(100),
--     SeatsNumber int,
--     PlateNumber VARCHAR(100),
--     Price FLOAT,
--     Photo VARCHAR(100)
-- )


-- CREATE TABLE users(
--  ID INT AUTO_INCREMENT PRIMARY KEY,
-- 	Firstname   VARCHAR(50),
-- 	Lastname    VARCHAR(50),
-- 	Password    VARCHAR(50),
-- 	Phonenumber VARCHAR(50),
-- 	Email       VARCHAR(50),
-- 	Address     VARCHAR(50),
-- 	City        VARCHAR(50)  
-- )


-- CREATE TABLE sessions(
--  ID     VARCHAR(10),
-- 	Email  VARCHAR(50),
-- 	LoggedIn    boolean	
-- )



-- CREATE TABLE admins(
--  ID INT AUTO_INCREMENT PRIMARY KEY,
-- 	Name     VARCHAR(50),
-- 	Password VARCHAR(50)
-- )

-- CREATE TABLE posted_cars(
--     carID VARCHAR(20),
--     userID int,
--     FOREIGN KEY (carID) REFERENCES cars(ID),
--     FOREIGN KEY (userID) REFERENCES users(ID)
-- )