package p5

import "fmt"

type EmbedUser struct {
    Name  string
    Email string
}

func (u *EmbedUser) notify() {
    fmt.Printf("Sending user email to %s<%s>\n",
        u.Name, u.Email)
}

type EmbedAdmin struct {
    EmbedUser // 嵌入类型
    level string
}


