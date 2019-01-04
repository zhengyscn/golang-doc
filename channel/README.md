# channel 管道

- 类似unix中管道pipe
- 先进先出
- 线程安全，多个goroutine同时访问，不需加锁
- channel是有类型的，一个整数的channel只能存整数

