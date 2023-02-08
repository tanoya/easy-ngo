// Copyright 2022 NetEase Media Technology（Beijing）Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xzk

import (
	"errors"
	"time"
)

func defaultOption() *Option {
	return &Option{}
}

func checkOptions(opt *Option) error {
	if opt.Name == "" {
		return errors.New("client name can not be nil")
	}

	if len(opt.Addr) == 0 {
		return errors.New("zk: server list must not be empty")
	}

	if opt.SessionTimeout == 0 {
		opt.SessionTimeout = 5 * time.Second
	}
	return nil
}

type Option struct {
	Name           string        // 客户端名称， 需要唯一
	Addr           []string      // 节点地址
	SessionTimeout time.Duration // 连接创建超时时间
}
