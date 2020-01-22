CREATE TABLE [user] (
    id INT PRIMARY KEY IDENTITY (1, 1),
    subject VARCHAR (256) NOT NULL,
    name VARCHAR (256) NULL,
    created_date DATETIME DEFAULT GETDATE(),
    changed_date DATETIME DEFAULT GETDATE(),
    deleted_date DATETIME,
    CONSTRAINT user_account_subject_fk FOREIGN KEY (subject) REFERENCES account (subject)
);
