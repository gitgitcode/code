# 类型

## 声明

    type 声明用户定义的类型
    一个结构类型
    ```
    type user struct {
        name string
        ext int
        privileged bool
    }
    //声明user 类型的变量 并初始化其为零值
    var bill user
    //声明一个suer 类型的变量
    lisa := user {
        name: "lisa,
        ext : 1,
        privileged: true
    }
    lisa := user {"lisa",1,true}

    type admin struct{
        persion user
        level string
    }
    fred := admin {
        person :user {
            name       : "lisa",
            ext        : 123,
            privileged : true,
        },
        level   :"super",
    }
    type Duration int64 //使用的是内置in64 但是不能等同于他就是int64
    var dur Duration
    dur = int64(1000)//报错 编译器不会做隐式转换
    ```

## 函数

关键字func 和函数名之间的参数被称为接收者 将函数与接收者的类型绑定在一起
如果一个函数有接收者，这个函数被称为方法
    ```
    func (u user) notify() {
        //func 和函数之间的参数 成为接收者
        fmt.Printf("Sending User Email To %s<%s>\n",
            u.name,
            u.email)
    }
    ```
使用指针接收者
    ```
    func (u *user) changeEmail(email string) {
        //指向user类型的指针 来调用notify方法
        u.email = email
    }
    指针调用 notify方法  为了支持这种方法 go语言调整了指针的来符合接受者的定义  (*lisa).notify()
    lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
    //值类型也可以调用指针接收者
    bill.changeEmail("bill@newdomain.com")
	//(&bill),changeEmail(bill@newdomain.com)
    bill 值得到一个指针，这个指针就能够匹配方法的接收者类型，在进行调用，
    Go 即允许使用值，也允许使用指针来调用方法，不必严格符合接收者的类型
    ```

## 类型的本质

    如果给类型增加或者删除某个值，是要创建一个新的值，还是要更改当前的值
    - 如果要创建一个新值使用值接收者
    - 如果修改当前值使用指针接收者
    * 是按值传递还是按指针传递

### 内置类型

    - 原始类型 的增加删除会创建一个新的值
    * 数值类型 
    * 字符串类型
    * 布尔类型 
    把这些类型传递给方法函数时，传递一个对应的值的副本

字符串string 原始的数据值，在函数或方法传递时 传递字符串的副本 例内置 Trim 方法

### 引用类型

    - 引用类型 创建变量时 创建的变量被称为标头 header的值 （从技术细节上说 字符串也算一种引用类型）
    * 切片
    * 映射
    * 通道
    * 接口
    * 函数类型

每个引用类型创建的标头值是包含一个指向底层数据结构的指针。
每个引用类型还包含一组独特的字段，用于管理底层的数据结构。
因为标头是为赋值而设计的，所以永远不需要共享一个引用类型的值。
标头值里包含一个指针，因此通过复制来传递一个引用类型的副本，本质上就是在共享底层数据结构。

### 结构类型

```
type File struct{
    *file
}
//file 是*File的实际表示
//额外的已成数据结构保证咩有哪个os的客户端能够覆盖这些数据
//如果覆盖这些数据，可能在变量终结时关闭错误的文件描述符
type file struct{
    fd int
    name string
    dirinfo *dirInfo //除了目录结构 此字段为nil
    nepipe int32 //Write 操作时遇到连续EPIPE的次数
}
```
通过嵌套指针的方式阻止复制。
```
func Open(name sting)(file *File, err error){
    return OpenFile(name, O_RDONLY, 0)
}
```
Opne 函数 调用者得到是是一个指向File类型的指针， *File
Open创建了File类型的值，并返回值的指针。
- 如果一个创建用的工厂函数返回了一个指针，就表示这个被返回的值的本质是非原始的。
- 即便函数或者方法没有直接改变原始的值的状态，依旧应该是攻心啊高的方式传递
- 例外 让类型值符合某个接口的时候，即便类型的本质是非原始本质的，也可以选择使用接收者声明方法。

## 接口

- 多态 根据类型的具体实现采取不同的行为的 能力

    如果一个类型实现了某个接口，所有使用这个接口的方法，都可以支持这种类型的值。

### 实现

    - 接口定义行为的类型。不由接口直接实现，通过方法由用户一的类型实现。

如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口类型的值。这个赋值会把定义的类型的值存入接口类型的值

对接口值的方法调用会执行接口值里存储的用户定义的类型的值对应方法。因为任何用户定义的类型都可以实现接口，所以接口值方法的调用自然是一种多态。
用户定义的类型通常叫做实体。

    ```
            var n notifier
            n = user{"Bill"}
    //实体复制后接口
    notifier 接口的值              iTable
    iTabe的地址 -------------------->user的类型
    user的地址  ------>User [存储的值] 方法集

    ```
接口值是一个两个字段长度的数据结构[notifier 接口的值] ,iTabe的地址,
存储值的类型信息。iTable包含了已经存储的值的类型信息以及相关联的一组方法。
第二个指向所有存储的指针。将类型信息和指针组合在一起。

    ```
            var n notifier
            n = &user{"Bill"}
    notifier 接口的值              iTable
    iTabe的地址 -------------------->*user的类型
    user的地址  ------>User [存储的值] 方法集
    ```
类型信息会存储一个指向保存的类型的指针，而接口值第二个字依旧保存指向指向实体值的指针。

### 方法集

    方法集定义了一组关联到给定类型的值或者指针的方法。定义方法时候使用的接收者的类型决定了这个方法是关联到值，还是关联到指针，还是两个都关联

（t T） t是变量，T是类型
    ```
    values             metherds receivers
    -----------------------------------
    T                 (t T)
    *T                (t T) and (t *T)
    ```
T类型值方法集只包含值接收者声明的方法。而指向T类型的指针的方法既包含值接收者声明的方法，也包含指针接收者声明的方法。值角度看很复杂。

接收者角度
    ```
     metherds receivers     values
    -----------------------------------
    (t T)                   T and *T
   （t *T)                  *T
    ```
- 如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。
- 如果使用值接收者来实现一个接口，那么哪个类型的值和指针都能够实现对应的接口
- 因为不能总能获取一个值的地址所以值的方法集值包括了使用值接收者实现的方法

### 多态

    ```
        type notifier interface{
            notify
        }
        type user {}
        type admin{}
        func (a *admin) notify{}//使用指针接收者
        func (u *user) notify{} //各自实现了接口
        //多态函数 实现了notifiter 接口的值作为参数
        func sendNotification(n notifier){ n.notify() } 
        //任意一个实体类型都能实现该接口，函数指针可以真对任意实体来执行
        bill :=user{}
        lisa :=admin{}
        sendNotification(&bill)//可以同时执行user admin实现的行为
        sendNotification(&lisa)
    ```

### 嵌入类型

    嵌入类型是将已有的类型直接声明在新的结构类型里。被嵌入的类型被称为新的外部类型的内部类型。

- 内部类型相关的标识符会被提升到外部类型上
- 被提升的标识符像直接声明在外部类型里面的标识符一样，也是外部类型的一部分。
- 外部类型组合了内部类型包含的所有属性，并且可以添加新的字段和方法
- 外部类型也可以通过声明与内部类型标识符同名的标识符来覆盖内部标识符的字段或者方法

    ```
        type notifter interface{
            notify()
        }
        type user {name string }
        type admin { user level string }
        func sendMian(n notifter){ notify}
        func (u *user) notify(){}
        ad:= admin { 
            user : user {name : "test"},
            level :"super",
            }
        ad.user.notify()
        ad.notify()
        sendMian(&ad)
    ```

    ad 变量外部类型是admin 嵌入了user 类型 。内部类型的提升，内部类型的接口会自动提升到外部类型。内部类型的实现，外部类型也同样实现了这个接口

- 如果外部也实现了方法，那么内部就不会提升 func (ad *admin)notify(){}

### 公开或未公开的标识符

    - 控制可见性
    * 小写开头的标识符 对外部不可见
    * 大写则公开

工厂函数命名为New是习惯
创建一个未公开的类型的值，并将这个值返回给调用者
