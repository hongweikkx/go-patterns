Golang设计模式思想
-----
* 前前言 <br>
    fork from sevenelevenlee/go-patterns. 修改如下：
    * 改变目录结构。 分成两大块，面向对象的设计模式和goroutine的组织模式
    * 面向对象的设计模式中增删改设计模式, 并修改了对应的TEST代码(便于同种类型中模式的横向比较)
    * 面向对象的设计模式中增加应用实例在*_example_test. 方便学习如何在实际中使用。
    * goroutine 组织模式进行的全新的修改和增加
* 前言
    *   一切设计模式都是灵活应用struct的组合模式，以及go隐形继承接口的特性
    *   go中的interface就是一些方法装饰, 而struct并不依赖于接口

* 设计模式类型
  
[创建模式(creational)](./design-pattern/01-creational)
----
- [简单工厂模式(Simple Factory)](./design-pattern/01-creational/01-simple-factory-pattern)
- [工厂方法模式(Factory Method)](./design-pattern/01-creational/02-factory-method-patterns)
    -     使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化
- [抽象工厂模式(Abstract Factory)](./design-pattern/01-creational/03-abstract-factory)
    -     提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类
- [单例模式(singleton)](./design-pattern/01-creational/04-singleton-pattern)
    -     保证一个类仅有一个实例, 并提供一个全局的访问接口
- [建造者模式(Builder Pattern)](./design-pattern/01-creational/05-builder-patterns)
    -     将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
- [原型模式(Prototype Pattern)](./design-pattern/01-creational/06-prototype-pattern)
    -     也叫生成器模式,复制一个已存在的实例
    
[结构模式](./design-pattern/02-structural)
----
- [装饰模式(Decorator Pattern)](./design-pattern/02-structural/01-decorator-pattern)
    -     装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能
- [适配器模式(Adapter Pattern)](./design-pattern/02-structural/02-adapter-pattern)
    -     将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
- [桥接模式(Bridge Pattern)](./design-pattern/02-structural/03-bridge-pattern)
    -     桥接模式分离抽象部分和实现部分，使得两部分可以独立扩展
- [组合模式(Composite)](./design-pattern/02-structural/04-composite-pattern)
    -     组合模式有助于表达数据结构, 将对象组合成树形结构以表示"部分-整体"的层次结构, 常用于树状的结构
- [享元模式(Flyweight Pattern)](./design-pattern/02-structural/05-flyweight-pattern)
    -     把多个实例对象共同需要的数据，独立出一个享元，从而减少对象数量和节省内存
- [代理模式(Proxy Pattern)](./design-pattern/02-structural/06-proxy-pattern)
    -     代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理。
- [外观模式(Facade Pattern)](./design-pattern/02-structural/07-facade-pattern)
    -     外观模式在客户端和现有系统之间加入一个外观对象, 为子系统提供一个统一的接入接口, 类似与委托
    
行为模式
----
- [观察者模式(Observer)](./design-pattern/03-behavioural/01-observer-pattern)
    -     定义对象间的一种一对多的依赖关系,以便当一个对象的状态发生改变时,所有依赖于它的对象都得到通知并自动刷新
- [策略模式(Strategy)](./design-pattern/03-behavioural/02-strategy-pattern)
    -     定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则
- [状态模式(State Pattern)](./design-pattern/03-behavioural/03-state-pattern)
    -     用于系统中复杂对象的状态转换以及不同状态下行为的封装问题
- [访问者模式(Visitor Pattern)](./design-pattern/03-behavioural/04-visitor-pattern)
    -     访问者模式是将对象的数据和操作分离
- [模板方法模式(Template Method Pattern)](./design-pattern/03-behavioural/05-template-method-pattern)
    -     模版方法使用匿名组合的继承机制, 将通用的方法和属性放在父类中, 具体的实现放在子类中延迟执行
- [备忘录模式(Memento Pattern)](./design-pattern/03-behavioural/06-memento-pattern)
    -     备忘录模式捕获一个对象的内部状态，并在对象之外保存这个状态
- [中介模式(Mediator Pattern)](./design-pattern/03-behavioural/07-mediator-pattern)
    -     中介者模式用一个中介对象来封装一系列对象交互，将多对多关联转换成一对多，构成星状结构
- [迭代器模式(Iterator Pattern)](./design-pattern/03-behavioural/08-iterator-pattern)
    -     可以配合访问者模式，将不同的数据结构，使用迭代器遍历
- [解释器模式(Interpreter Pattern)](./design-pattern/03-behavioural/09-interpreter-pattern)
    -     解释器模式实现一个表达式接口，该接口解释一个特定的上下文。通常用于SQL解析和符号处理引擎
- [命令模式(Command Pattern)](./design-pattern/03-behavioural/10-command-pattern)
    -     命令模式是一种数据驱动模式，将请求封装成一个对象，从而可以用不同的请求对客户进行参数化，实现调用者和接收者的解藕
- [责任链模式(Chain of Responsibility)](./design-pattern/03-behavioural/11-chain-of-responsibility-pattern)
    -     责任链模式是将处理请求的多个对象连成一条链(类似队列)，每个对象都包含下一个对象的引用，请求沿着链传递，直到被处理

同步模式(synchronization patterns)
----
- [信号量模式(Semaphore)](./09-semaphore-pattern)
    -       信号量是一种同步模式，对有限数量的资源同步互斥
- [发布订阅模式(publish-subscribe)](./27-publish-and-subscribe)
    -       有别于传统的生产者消费者，pubsub模型将消费发布给一个主题
附录(设计模式彩图)
-
   ![设计模式彩图](./go-design-image.jpg)

