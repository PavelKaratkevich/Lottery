create table Lottery (
id int primary key auto_increment,
first_name varchar(30),
last_name varchar(30),
id_number varchar(10) unique
)
DELIMITER $$
CREATE TRIGGER max_number_of_rows_Lottery
BEFORE INSERT ON Lottery
FOR EACH ROW
BEGIN
    DECLARE cnt INT;

    SELECT count(*) INTO cnt FROM Lottery;

    IF cnt = 200000 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'All the 10 ticket were sold out';
    END IF;
END
$$

DELIMITER ;
