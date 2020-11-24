
Random Numbers Generator 

Based on 
https://stackoverflow.com/questions/41632285/generating-random-numbers-concurrently-in-go

To fill array with 16 random integers, empty slice is created (which will be filled dynamically).
To fill array set of go routines is activated ahich are cinnected with separate channel between each pair.

First goroutine generates chunk of random ints and put them into channel.
The second one picks up these data from the channel and append them to target slice.

Sychronisation is performed with WaitGroup.


Colombina:5 - Random Generator 1$ go run Random.go 
Genertor:  0  data =  &[30 45 91 98]
Genertor:  1  data =  &[30 45 91 98 6 20 25 54]
Genertor:  2  data =  &[30 45 91 98 6 20 25 54 66 85 72 44]
Genertor:  3  data =  &[30 45 91 98 6 20 25 54 66 85 72 44 95 96 97 16]