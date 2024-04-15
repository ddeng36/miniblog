// Copyright 2024 ddeng36 <dingchaodeng@vip.qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ddeng36.

package main

import (
	"os"
	"github.com/ddeng36/miniblog/internal/miniblog"
	_ "go.uber.org/automaxprocs"
)

/*
miniblog程序的默认入口函数
*/
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
