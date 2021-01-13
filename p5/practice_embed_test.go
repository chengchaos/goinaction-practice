package p5

import "testing"

func TestEmbedUser(t *testing.T) {

    ad := EmbedAdmin{
        EmbedUser: EmbedUser{
            Name:  "John Smith",
            Email: "john.smith@126.com",
        },
        level: "super",
    }

    // 我们可以直接访问内部类型的方法
    ad.EmbedUser.notify()

    // 内部类型的方法也被提升到了外部类型
    ad.notify()

    add := &ad
    add.EmbedUser.notify()

    (&ad).notify()

}
