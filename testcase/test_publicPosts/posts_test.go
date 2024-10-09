package test_publicPosts

import (
	"Ginko/Utils"
	"Ginko/public"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega" // Gomega 断言库，用于定义期望值
	"gorm.io/gorm"
)

var db *gorm.DB

var _ = BeforeSuite(func() {
	var err error
	db, err = Utils.InitDataBase()
	if err != nil {
		Fail("数据库连接失败,请检查配置文件并确认数据库设置" + err.Error())
	}

})

var _ = Describe("查询帖子需求自动化测试", func() {
	seC := Utils.SetClient()

	It("查询帖子成功", func() {
		var postslist []public.Posts
		var res public.PostsObj

		Restring, ResCode, err := Utils.SetRequest(seC, "http://sunlana.com/api/v1/posts?page=1", "GET", nil)
		if err != nil {
			Fail("请求失败")
		}

		Expect(ResCode).To(Equal(200))

		err = db.Debug().Find(&postslist).Error
		if err != nil {
			Fail("数据库查询失败" + err.Error())
		}

		err = Utils.JsonDeserialization(Restring, &res)

		Expect(ResCode).To(Equal(200))
		Expect(Restring).To(Equal("查询成功"))

	})

})
