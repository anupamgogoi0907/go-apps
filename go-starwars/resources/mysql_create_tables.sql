CREATE TABLE Inventory
(
    id int PRIMARY KEY NOT NULL
)
;



CREATE TABLE Item
(
    id int PRIMARY KEY NOT NULL,
    name varchar(45) NOT NULL,
    quantity int NOT NULL,
    points int NOT NULL,
    inventory_id int NOT NULL
)
;
ALTER TABLE Item
    ADD CONSTRAINT FK_Item_Inventory
        FOREIGN KEY (inventory_id)
            REFERENCES Inventory(id);

CREATE TABLE Localization
(
    id int PRIMARY KEY NOT NULL,
    latitude varchar(45) NOT NULL,
    longitude varchar(45) NOT NULL,
    name varchar(45) NOT NULL
)
;


CREATE TABLE Rebel
(
    id int NOT NULL,
    name varchar(45) NOT NULL,
    age int,
    gender varchar(45),
    traitor int NOT NULL,
    localization_id int NOT NULL,
    inventory_id int NOT NULL
)
;
ALTER TABLE Rebel
    ADD CONSTRAINT FK_Rebel_Inventory
        FOREIGN KEY (inventory_id)
            REFERENCES Inventory(id)
;
ALTER TABLE Rebel
    ADD CONSTRAINT FK_Rebel_Localization
        FOREIGN KEY (localization_id)
            REFERENCES Localization(id)
;

