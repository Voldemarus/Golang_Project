
Password matching(regex)

Idea derived from Stack Overflow question
https://stackoverflow.com/questions/25837241/password-validation-with-regexp

User enters password (for simplicity is hardcoded in hthe main function)

Password is checked against several rules:

 * at least 7 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character

  Example:

 Colombina:10 - Password Validator 1$ go run Validator.go 
root123root  :-(  Password should contain at least one uppercase letter
RootRoot123  :-(  Password should contain at least one punctuation char
rootRoot_123 :-)
Root_1_root :-)
123  :-(  Too short, should contains at least 9 characters
rootrootroot  :-(  Password should contain at least 1 digit
rootroorrrrr1  :-(  Password should contain at least one uppercase letter
RootRoot**123 :-)
RootRoot,123 :-)
Colombina:10 - Password Validator 1$ 