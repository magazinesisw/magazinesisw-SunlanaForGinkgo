package test_log_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

// 测试套件的入口函数
func TestTestLogin(t *testing.T) {
	RegisterFailHandler(Fail) // 注册一个失败处理程序（Ginkgo 的标准用法）
	RunSpecs(t, "TestLogin Suite")

}
