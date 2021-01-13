package p5

import "fmt"

type User struct {
    name       string
    email      string
    ext        int
    privileged bool
}

/*
值接受者。
调用是会使用这个值的一个副本来执行。
*/
func (u User) notify() {
    fmt.Printf("Sending User Email to %s<%s>\n",
        u.name, u.email)
}

func (u *User) changeEmail(email string) {
    u.email = email
}

type Manager User

func InspectUser() {

    // 任何时候，创建一个变量并初始化为其零值，习惯是使用 var 关键字。
    // 这种用法是为了更明确的表示一个变量被设置为零值。
    var u User
    fmt.Println("u => ", u)

    u.name = "程超"
    fmt.Println("u =>", u)

    aMap := map[string]string{
        "id":    "1",
        "name":  "chengchao",
        "email": "chengchaos@126.com",
    }

    // 字面量就是使用一对大括号括住内部字段的初始值。
    // 就像上面的 Map 那样
    bob := User{
        name:       "Bob",
        email:      "bob@126.com",
        ext:        134,
        privileged: true,
    }

    // 第二种形式，按顺序：
    // 一般写在一行里。
    tom := User{"Tom", "tom@126.com", 343, true}
    fmt.Println("aMap =>", aMap)
    fmt.Println("bob =>", bob)
    fmt.Println("tom =>", tom)

    type Admin struct {
        person User
        level  string
    }

    fred := Admin{
        person: User{
            name:       "Fred",
            email:      "fred@126.com",
            ext:        233,
            privileged: true,
        },
        level: "super",
    }

    fmt.Println("fred =>", fred)

    // 基于一个已有的类型，将其作为新类型的说明：
    // 我们把 int64 类型叫做 Duration 的基础类型。
    type Duration int64

    var dur Duration
    var i64 int64 = 1000
    // Cannot use 'i64' (type int64) as type Duration
    // dur = i64

    fmt.Printf("dur => %v, i64 => %v\n",
        dur, i64)

    bill := User{
        name:  "Bill",
        email: "bill@126.com",
    }
    bill.notify()

    alice := &User{name: "Alice", email: "alice@126.com"}
    alice.notify()

    bill.changeEmail("bill@microsoft.com")
    bill.notify()

    alice.changeEmail("alice@microsoft.com")
    alice.notify()

    master := Manager{
        name:  "The Master",
        email: "the-master@126.com",
    }

    // master 就没有 notify 方法了。
    //master.notify()
    fmt.Println("master =>", master)

}
