# message-board

RedRockHomework: message-board

<h2>api</h2>
<p>
存放路由router.go，以及将不同的接口函数按照性质放在不同的文件里，
比如user.go里面存放有关user相关的接口。comment.go中存放comment
相关接口。这一层是负责参数校验，请求转发。
</p>
<h2>cmd</h2>
<p>
仅存放main.go文件，作为整个项目的入口。
</p>
<h2>dao</h2>
<p>
数据访问层，全称为Data Access Object，所有数据库相关操作全部
封装在这一层。将下层存储以更简单的函数、接口形式暴露给 service 层来使用。
</p>
<h2>model</h2>
<p>
存放项目中所有的struct，按照文件名归类。
</p>
<h2>service</h2>
<p>
将各个服务封装之后供各个接口调用。可以认为从这里开始，
所有的请求参数一定是合法的，业务逻辑和业务流程也都在
这一层中。举例来说，比如注册接口需要判断用户名是否重复
，那么我们需要写一个
func IsRepeatUserName(userName string)(bool,error){}
的函数，放在service文件夹下的user.go文件中。
</p>
<h2>util</h2>
<p>

</p>
<h2>go.mod</h2>
<p>mysql</p>
<p>gin</p>


