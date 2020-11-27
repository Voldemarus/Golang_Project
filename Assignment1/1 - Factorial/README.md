
Factorial - multithreaded calculator 

Two recursive function defined -  factorial() and factorialMT()

First function is a traditional recursive implementation. 

The second one - implementation with divide/conquer algorithm where initial numbber 
is divided by two and two multiplication function are called - for 1..m and m+1..n intervals 
where m = n / 2. 

Each algorithm is started in its own thread. 
function accepts from, to border value and channel ID, used to return value. 

Example:

Colombina:1 - Factorial 1$ go run Factorial.go 
Factorial  7 =  5040
Factorial multithreaded 5040

