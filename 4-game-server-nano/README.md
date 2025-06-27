## Nano

nano是go生态里的一个游戏框架，可以用来开发中小规模游戏,比如棋牌，卡牌

## 核心组件
- Component
  - 在struct中嵌入component.Base // 内嵌Base组件，提供组件生命周期管理
    ```
      type RoomManager struct {
      component.Base // 内嵌Base组件，提供组件生命周期管理
      timer          *scheduler.Timer
      rooms          map[int]*pojo.Room // map结构-存储所有房间
      }
    ```
  - 在nano中把这个组件注册进去
    ```
    components := &component.Components{}
    components.Register(
        NewRoomManager(),
        component.WithName("room"), // rewrite component and handler name
        component.WithNameFunc(strings.ToLower),
    )
    ```
  - 客户端就可以通过room.xxx来访问RoomManager里面的方法  
- Group
  - 它里面有一个map结构的sessions字段，用来管理多个Session
  - 提供了常用的Broadcast方法，用来向所有的Session发送消息
- Session
  - 它有2个字段比较重要，一个是uid, 一个是data
  - uid是用来标识这个Session的，data可以用来在这个Session上存储一些数据
- Scheduler
  - nano自己封装的一个定时器
- pipeline
  - 用来处理消息的管道 提供了常用的Outbound/Inbound方法, PushBack/PushFront添加自定义处理器的方法