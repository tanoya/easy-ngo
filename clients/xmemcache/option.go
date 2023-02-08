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

package xmemcache

import (
	"errors"
	"time"
)

// Options 可配置的数据
type Option struct {
	Name         string        // 客户端名称，需要唯一
	Timeout      time.Duration // 客户端连接超时时间
	MaxIdleConns int           // 最大空闲连接
	Addr         []string      // 集群地址
	EnableTracer bool
}

func defaultOption() *Option {
	return &Option{}
}

func checkOptions(opt *Option) error {
	if opt.Name == "" {
		return errors.New("client name can not be nil")
	}
	return nil
}
