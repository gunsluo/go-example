CREATE TABLE [account] (
    id INT PRIMARY KEY IDENTITY (1, 1),
    subject VARCHAR (256) NOT NULL,
    email VARCHAR (256) NOT NULL,
    created_date DATETIME DEFAULT GETDATE(),
    changed_date DATETIME DEFAULT GETDATE(),
    deleted_date DATETIME,
    CONSTRAINT account_subject_ak UNIQUE(subject)
);
