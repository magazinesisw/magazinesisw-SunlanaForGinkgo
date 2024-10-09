package test_log

import (
	"Ginko/Utils"
	"Ginko/config"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega" // Gomega 断言库，用于定义期望值
	"log"
	"reflect"
)

var LogUser config.Login

// TODO BeforeSuite是一个钩子函数,它会在所有测试用例执行前运行一次
var _ = BeforeSuite(func() {

	LoginConfig := config.LoginConfig{}
	err := Utils.ReadYML("..\\..\\config\\login.yaml", &LoginConfig)
	if err != nil {
		Fail("读取配置文件失败" + err.Error())
	}
	LogUser = LoginConfig.LoginInfo
	fmt.Println(LogUser)

})

// TODO Describe用于描述一组测试用例
var _ = Describe("进行登录需求自动化测试", func() {

	seC := Utils.SetClient()
	//TODO It用于描述一个具体的测试用例
	//进行用例编写

	//TODO context的作用是描述某个特定场景或上下文
	Context("输入的账号密码匹配", func() {

		It("登录成功", func() {

			ResString, ResCode, err := Utils.SetRequest(seC, "http://Sunlana.com/api/v1/login", "POST", LogUser)
			if err != nil {
				log.Println("登录请求失败", err)
				Fail("请求失败")
				return
			}
			Expect(ResCode).To(Equal(200))
			Expect(ResString).To(ContainSubstring("登录成功"))
			Expect(ResString).To(ContainSubstring("token"))
			Expect(ResString).To(ContainSubstring("user"))
			fmt.Println("登录请求成功", ResString, ResCode)

		})

	})

	It("输入的账号密码不匹配", func() {
		LogUser.Password = "123456"

		Restring, ResCode, err := Utils.SetRequest(seC, "http://Sunlana.com/api/v1/login", "POST", LogUser)
		if err != nil {
			Fail("请求失败")
			return
		}

		log.Println(Restring, reflect.TypeOf(Restring))

		Expect(ResCode).To(Equal(400))
		Expect(Restring).To(ContainSubstring("登录失败,用户名或密码错误"))

	})

})
