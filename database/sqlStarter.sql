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
	id VARCHAR(255),
	brotherId VARCHAR(255),
	dob VARCHAR(255),
	email VARCHAR(255),
	phoneNumber VARCHAR(255),
	altPhoneNumber VARCHAR(255), 
	altContactName VARCHAR(255), 
	altContactRelationship VARCHAR(255), 
	currentAddress VARCHAR(8080), 
	permanentAddress VARCHAR(8080)
);