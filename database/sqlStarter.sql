CREATE TABLE brothers (
    id VARCHAR(255) PRIMARY KEY,
    firstName VARCHAR(255) DEFAULT '',
    lastName VARCHAR(255) DEFAULT '',
    sudoName VARCHAR(255) DEFAULT '',
    referralCode VARCHAR(255) DEFAULT '',
    referredBy VARCHAR(255) DEFAULT '',
    emailVerificationStatus VARCHAR(255) DEFAULT '',
    approvalStatus VARCHAR(255) DEFAULT '',
    approvedBy VARCHAR(255) DEFAULT ''
);

CREATE TABLE personalDetails (
    id VARCHAR(255) PRIMARY KEY,
    brotherId VARCHAR(255) DEFAULT '',
    dob VARCHAR(255) DEFAULT '',
    email VARCHAR(255) DEFAULT '',
    phoneNumber VARCHAR(255) DEFAULT '',
    altPhoneNumber VARCHAR(255) DEFAULT '',
    altContactName VARCHAR(255) DEFAULT '',
    altContactRelationship VARCHAR(255) DEFAULT '',
    currentAddress VARCHAR(8080) DEFAULT '',
    currentFileType VARCHAR(255) DEFAULT '',
    currentContentType VARCHAR(255) DEFAULT '',
    permanentAddress VARCHAR(8080) DEFAULT '',
    permanentFileType VARCHAR(255) DEFAULT '',
    permanentContentType VARCHAR(255) DEFAULT ''
);
