
ErrProcessor.go

Testbed for several errors handling methods in Go

Colombina:8 - Error Handling 1$ go run ErrProcessor.go 

* Using multiple return values to check for an errors 

divideWithCheck( 73 , 0.0001 )
cannot perform divideWithCheck() -  Divisor has too small absolute value

* Error() interface implementation 

Rectangle a =  {12 40 10 10} err -  
Rectangle b =  {5 15 -5 6} err -  Invalid dimensions

* Error status escalation 

aInterval  {5 10}
bInterval  {7 5}
intersect -  intersect()  - add message from embedded error -{7 5}  Interval:: Invalid order of parameters

* Panic/Recover

Call panic with recovery activated
executePanic() started
Panic event initiated!
Recovery function completed

Call pnic without recovery
executePanic() started
panic: Panic event initiated!

goroutine 1 [running]:
main.executePanic(0x10f0000)
        /Users/1/code/src/github.com/Voldemarus/Golang_Project/Assignment1/8 - Error Handling/ErrProcessor.go:103 +0xc5
main.main()
        /Users/1/code/src/github.com/Voldemarus/Golang_Project/Assignment1/8 - Error Handling/ErrProcessor.go:85 +0x8f5
exit status 2
Colombina:8 - Error Handling 1$ 



