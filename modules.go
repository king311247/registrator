package main

import (
	_ "github.com/king311247/registrator/consul"
	_ "github.com/king311247/registrator/consulkv"
	_ "github.com/king311247/registrator/etcd"
	_ "github.com/king311247/registrator/httpcollector"
	_ "github.com/king311247/registrator/skydns2"
	_ "github.com/king311247/registrator/zookeeper"
)
