
Server and Workers 

Two pools of workers is created which should perform two different tasks. They are synchronised with 
data, supplied by channels. 

Example: 

Colombina:4- Server and Workers 1$ go run WorkerPool.go 
Worker  1 Operation 2 started as job # 1
Worker  3 Operation 1 started as job # 3
Worker  4 Operation 2 started as job # 3
Worker  2 Operation 1 started as job # 2
Worker  2 Operation 2 started as job # 4
Worker  3 Operation 2 started as job # 2
Worker  4 Operation 1 started as job # 4
Worker  6 Operation 2 started as job # 6
Worker  1 Operation 1 started as job # 1
Worker  5 Operation 2 started as job # 5
Worker  5 Operation 2 finished as job # 5
Worker  3 Operation 2 finished as job # 2
Worker  1 Operation 2 finished as job # 1
Worker  3 Operation 2 started as job # 7
Worker  2 Operation 2 finished as job # 4
Worker  2 Operation 2 started as job # 8
Worker  1 Operation 2 started as job # 9
Worker  6 Operation 2 finished as job # 6
Worker  4 Operation 2 finished as job # 3
Operation 2 -  {3 2 Worker #3 :: job : 2 result = 4}
Operation 2 -  {2 4 Worker #2 :: job : 4 result = 16}
Operation 2 -  {1 1 Worker #1 :: job : 1 result = 1}
Operation 2 -  {6 6 Worker #6 :: job : 6 result = 36}
Operation 2 -  {4 3 Worker #4 :: job : 3 result = 9}
Operation 2 -  {5 5 Worker #5 :: job : 5 result = 25}
Worker  3 Operation 1 finished as job # 3
Worker  4 Operation 1 finished as job # 4
Worker  4 Operation 1 started as job # 6
Worker  3 Operation 1 started as job # 5
Worker  1 Operation 1 finished as job # 1
Worker  1 Operation 1 started as job # 7
Worker  2 Operation 1 finished as job # 2
Worker  2 Operation 1 started as job # 8
Operation 1 -  {3 3 303}
Operation 1 -  {4 4 404}
Operation 1 -  {1 1 101}
Operation 1 -  {2 2 202}
Worker  2 Operation 2 finished as job # 8
Worker  1 Operation 2 finished as job # 9
Operation 2 -  {2 8 Worker #2 :: job : 8 result = 64}
Operation 2 -  {1 9 Worker #1 :: job : 9 result = 81}
Worker  3 Operation 2 finished as job # 7
Operation 2 -  {3 7 Worker #3 :: job : 7 result = 49}
Worker  1 Operation 1 finished as job # 7
Worker  3 Operation 1 finished as job # 5
Worker  3 Operation 1 started as job # 10
Worker  2 Operation 1 finished as job # 8
Worker  1 Operation 1 started as job # 9
Worker  4 Operation 1 finished as job # 6
Operation 1 -  {1 7 107}
Operation 1 -  {3 5 305}
Operation 1 -  {2 8 208}
Operation 1 -  {4 6 406}
Worker  1 Operation 1 finished as job # 9
Worker  3 Operation 1 finished as job # 10
Operation 1 -  {1 9 109}
Operation 1 -  {3 10 310}
Colombina:4- Server and Workers 1$ 