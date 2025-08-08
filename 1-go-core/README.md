## 设计哲学:
- 1.显式优于隐式
- 2.简单优于复杂
- 3.组合优于继承
- 4.少就是多 - 精简的标准库，不盲目加特性

## 特色:
- 1.函数可以返回多个值
- 2.空接口：任意类型的值都可以赋值给空接口（如果一个函数的参数定义的是空接口interface{}类型，那么他可以接受任何类型的值,相当于是Object）
- 3.空白符_ ：用来忽略返回值或者参数
- 4.切片, 切片是引用类型，数组是值类型（函数数组/切片是存在的）
- 
- 11.go function(), 使用协程处理异步处理任务
- 12.sync.RWMutex 读写锁控制并发
- 13.select + channel控制协程间数据传递

```
sync.RWMutex 读写锁，这能有效避免并发读写 map 的问题，其工作原理如下：
Join 和 Leave 方法（写操作）
Join 和 Leave 方法在修改 players map 时使用了写锁（Lock 和 Unlock）：
写锁是独占锁，当一个 goroutine 获取写锁后，other goroutine 既不能获取写锁，也不能获取读锁，直到写锁被释放。这意味着在 Join 或 Leave 方法执行期间，Broadcast 方法无法对 players 进行遍历。

Broadcast 方法（读操作）
Broadcast 方法在遍历 players map 时使用了读锁（RLock 和 RUnlock）：
读锁允许多个 goroutine 同时获取，也就是说多个 Broadcast 方法可以并发执行。但在有 goroutine 持有读锁时，other goroutine 无法获取写锁，直到所有读锁都被释放。这保证了在 Broadcast 方法遍历 players 期间，Join 和 Leave 方法无法修改 players
*/

```


