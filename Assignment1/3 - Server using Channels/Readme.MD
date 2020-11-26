
ServerChannel

Init web server aon 9999 port, and init two handlers.

Handler 1. Uses closure to process data, placed into channel by handler. Copact, but not much clear.
Handler 2. Structure is defined, which contains channel. Also two delegate methods for this structure is defined: the first to be used as handler, and the second one - as logger. So both uses "internal" for this structure channel. 


Inspired by 

https://stackoverflow.com/questions/28419876/golang-handlefunc-with-channel
