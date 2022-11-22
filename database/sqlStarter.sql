CREATE TABLE brothers (
	id VARCHAR(255) primary key,
	firstName VARCHAR(255),
	lastName VARCHAR(255),
	sudoName VARCHAR(255),
	referralCode VARCHAR(255),
	referredBy VARCHAR(255), 
	approvalStatus VARCHAR(255)
);

CREATE TABLE personalDetails (
	Id VARCHAR(255),
	UserId VARCHAR(255),
	DoB VARCHAR(255),
	Email VARCHAR(255),
	PhoneNumber VARCHAR(255),
	AltPhoneNumber VARCHAR(255), 
	CurrentAddress VARCHAR(8080), 
	PermanentAddress VARCHAR(8080)
);