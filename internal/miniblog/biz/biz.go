
package biz

import (
	"github.com/ddeng36/miniblog/internal/miniblog/biz/user"
	"github.com/ddeng36/miniblog/internal/miniblog/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Users() user.UserBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
