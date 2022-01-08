CREATE DATABASE IF NOT EXISTS user;

USE user;

CREATE TABLE IF NOT EXISTS user_table1 (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  firstname varchar(30) NOT NULL,
  lastname varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_table2 (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  firstname varchar(30) NOT NULL,
  lastname varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_table3 (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  firstname varchar(30) NOT NULL,
  lastname varchar(30) NOT NULL
);

DELIMITER //

CREATE PROCEDURE sp_fakecreateuser(IN firstname varchar(225), IN lastname varchar(225), out insert_message varchar(225))

BEGIN 
-- Declare statements ..

DECLARE EXIT HANDLER FOR SQLEXCEPTION 
BEGIN

      ROLLBACK;
END;

START TRANSACTION;

IF (length(firstname) >= 2 && length(firstname) <= 20) && (length(lastname) >= 2 && length(lastname) <= 20) THEN

		INSERT INTO user_table1 (firstname, lastname ) VALUES ( firstname, lastname );
		INSERT INTO user_table2 (firstname, lastname ) VALUES ( firstname, lastname );
		INSERT INTO user_table3 (firstname, lastname ) VALUES ( firstname, lastname );

		SET insert_message = 'inserted into user_table1, user_table2 & user_table3';
ELSE 
		INSERT INTO user_table1 (firstname, lastname ) VALUES ( firstname, lastname );
		INSERT INTO user_table2 (firstname, lastname ) VALUES ( firstname, lastname );
		SET insert_message = 'data incorrect, inserted into user_table1 & user_table2 but not inside user_table3';
		END IF;

COMMIT;

END //

DELIMITER ;