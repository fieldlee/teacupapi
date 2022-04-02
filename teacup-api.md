1、验证码接口
http://127.0.0.1:17315/teacup/api/v1/user/getSmsCode
入参：
{
"user_phone":"123-345-6789",   
"nation":"us",   // cn 中国  us 美国
"user_phone_type":1,   // 1 iOS  2 Android
"type":1   // 1 登录  2注册 
}

返回：
    a)、用户已注册
        {
        "data": {},
        "message": "user exist",
        "status_code": 6021
        }
    b)、验证码已发送
        {
        "data": {},
        "message": "success",
        "status_code": 6000
        }

2、检验验证码
http://127.0.0.1:17315/teacup/api/v1/user/checkSmsCode
入参：
{
"user_phone":"18616017950",
"nation":"cn",    //cn 中国  us 美国
"smscode":"888888"  //测试环境 验证码都是 888888
}
返回：
    a)、验证码错误
        {
        "data": {},
        "message": "验证码错误",
        "status_code": 6019
        }
    b)、验证码正确
    {
    "data": {},
    "message": "success",
    "status_code": 6000
    }

3、注册
http://127.0.0.1:17315/teacup/api/v1/user/register
入参：
{
"user_phone":"18616017954",
"nation":"cn",  //cn 中国  us 美国
"user_phone_type":1,  // 1 iOS  2 Android
"smscode":"888888"
}
返回类型：
    a)、 验证码错误例子
        {
        "data": {},
        "message": "验证码错误",
        "status_code": 6019
        }
    b)、用户已存在，重复注册
        {
        "data": {},
        "message": "用户已存在",
        "status_code": 6021
        }
    c)、注册成功例子
        {
        "data": {
        "token": "1V9CniiRg7BNGQhR+zm6evlBaITvUacKSUeMGnUGs6rsW1Xo3u5vaAPIBicrkDDJRw5TiuRIngsIWlbsiIUSdA=="
        },
        "message": "注册完成",
        "status_code": 6000
        }

4、登录
http://127.0.0.1:17315/teacup/api/v1/user/login
入参：
{
"user_phone":"18616017954",
"nation":"cn",
"user_phone_type":1,
"smscode":"888888"
}
返回：
{
"data": {
"token": "CcjDseria9QUyjCqhAdD/LFkewE7GXsac4qPEWBGpJz0cmwrntI1iuYQfxlINSEJpE3+bFE6Q0anMMdHjMRjjw=="
},
"message": "登录成功",
"status_code": 6000
}

5、登录后token
请求 Headers 添加 key value 
X-API-TOKEN ： CcjDseria9QUyjCqhAdD/LFkewE7GXsac4qPEWBGpJz0cmwrntI1iuYQfxlINSEJpE3+bFE6Q0anMMdHjMRjjw==

6、上传图片
http://34.213.138.197:17315/teacup/api/v1/common/upload
--header 'X-API-TOKEN: h+3p7DhaVO0r60R0GthV9FhuJvGAuVJFo9+J48wc02mQfpjiJaktiqbsC0/a6WlebvMTxOKhXlHurQ6rniOqhw==' \
--header 'Content-Type: multipart/form-data' \
入参：
--form 'file=@"/Users/fieldlee/Downloads/fieldlee/欢乐送.jpeg"'
返回：
{
"data": "http://34.213.138.197:17315/images/1648123709782627091.jpeg",
"message": "success",
"status_code": 6000
}

7、编辑个人信息
http://34.213.138.197:17315/teacup/api/v1/user/fillInformation
入参：
{
"id":3,  //必须传
"user_phone":"18616017957",
"nation":"cn",
"gender":1,  //用户性别 1 男 2女
"user_phone_type":1, // 手机类型 1 iOS 2 android
"birthday":"2006-01-02",
"user_name":"fieldlee",
"user_avator":"depeng",  //昵称
"user_union":"cn",  // cn 中国 us 美国
"uuid":"",   //手机udid  手机通知消息用
"user_image":"http://34.213.138.197:17315/images/1648123709782627091.jpeg",    //头像图片地址
"user_topics":"专家;社区;泡茶;赛车"
}

8、创建社区
http://34.213.138.197:17315/teacup/api/v1/community/newCommunity
入参：
{
"community_name":"test1",
"community_comment":"comment1",
"community_tags":"大学;出国;留学",
"community_image":"/test.img",
"is_only_invite":1,  //  是否只有邀请人加入 1不是 2是
"is_any_join":2   //  是否任何人都可以加入 1不可以 2可以
}

9、创建茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/newTeacup
入参：
{
"community_id":1,
"teacup_name":"newteacup for test",
"teacup_comment":"聊聊出国那些事",
"teacup_image":"",
"teacup_tags":"出国;留学;旅游"
}

10、修改茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/updateTeacup
入参：
{
"id":1,
"community_id":1,
"teacup_name":"newteacup for test",
"teacup_comment":"聊聊出国那些事",
"teacup_image":"1111.jpg",
"teacup_tags":"出国;留学;旅游"
}
11、我创建的茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/myTeacup
入参：
{
"page":1,
"pageSize":5
}
返回：
{
"data": {
"page": 1,
"pageSize": 5,
"totalRecord": 1,
"list": [
{
"id": 1,
"user_id": 3,
"community_id": 1,
"teacup_name": "newteacup for test",
"teacup_comment": "聊聊出国那些事",
"teacup_image": "",
"teacup_tags": "出国;留学;旅游"
}
]
},
"message": "",
"status_code": 6000
}

12、根据id查茶室信息
http://34.213.138.197:17315/teacup/api/v1/teacup/getTeacup
入参：
{
"id":1
}
返回：
{
"data": {
"id": 1,
"user_id": 3,
"community_id": 1,
"teacup_name": "newteacup for test",
"teacup_comment": "聊聊出国那些事",
"teacup_image": "",
"teacup_tags": "出国;留学;旅游"
},
"message": "成功",
"status_code": 6000
}

13、邀请加入茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/invite2Teacup
入参：
{
"teacup_id":1,
"invited_user_id":6
}
14、邀请加入茶室是否同意审批
http://34.213.138.197:17315/teacup/api/v1/teacup/update2Teacup
入参：
{
"id":1,  //邀请函id
"invited_status":2  //1是不同意  2是同意
}

15、当前人邀请函的信息
Get 方法：
http://34.213.138.197:17315/teacup/api/v1/teacup/getInviteTeacup


16、申请加入茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/joinTeacup
{
"teacup_id":1  //茶室id
}

17、离开茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/leaveTeacup
{
"teacup_id":1  //茶室id
}

18、开始茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/startTeacup
{
"id":1,  //茶室id
"speech_room":"232323", //茶室的socket room
"speech_channel":"rrrrr", //茶室的语音channel
}

19、开始茶室
http://34.213.138.197:17315/teacup/api/v1/teacup/closeTeacup
{
"id":1  //茶室id
}

20、茶室记录
http://34.213.138.197:17315/teacup/api/v1/teacup/logTeacup
{
"teacup_id":1,  // 茶室id
"community_id":1,  //社区的id
"teacup_start_at":"2022-02-05 12:00:00", //茶室开始
"teacup_end_at":"2022-02-05 14:00:00"  //茶室结束
}

21、茶室打分评价
http://34.213.138.197:17315/teacup/api/v1/teacup/markTeacup
{
"teacup_id":1,  //茶室id
"mark_value":10, //茶室打分
"mark_comment":"test"//茶室评价
}
返回：
{
"data": {},
"message": "",
"status_code": 6000
}
22、茶室嘉宾查询
http://34.213.138.197:17315/teacup/api/v1/teacup/vipsOfTeacup
{
"id":1,  //茶室id
"page":1,  //当前第几页
"pageSize":5 //每页几个
}
返回：
{
"data": {
"page": 1,
"pageSize": 5,
"totalRecord": 2,
"list": [
{
"id": 9,
"user_phone": "18616017954",
"gender": 1,
"birthday": "2006-01-02T00:00:00+08:00",
"nation": "cn",
"user_phone_type": 1,
"uuid": "",
"user_name": "fieldlee",
"user_avator": "depeng",
"user_union": "cn",
"user_image": "",
"user_lvl": 1,
"user_badges": "",
"user_topics": "专家;社区;泡茶;赛车"
},
{
"id": 6,
"user_phone": "18616017952",
"gender": 1,
"birthday": "2022-03-08T09:31:28+08:00",
"nation": "cn",
"user_phone_type": 1,
"uuid": "",
"user_name": "",
"user_avator": "",
"user_union": "",
"user_image": "",
"user_lvl": 1,
"user_badges": "",
"user_topics": ""
}
]
},
"message": "",
"status_code": 6000
}
23、我作为嘉宾的茶室查询
http://34.213.138.197:17315/teacup/api/v1/teacup/teacupOfVips
{
"page":1,
"pageSize":5
}
返回：
{
"data": {
"page": 1,
"pageSize": 5,
"totalRecord": 1,
"list": [
{
"id": 1,
"user_id": 9,
"community_id": 1,
"teacup_name": "newteacup for test",
"teacup_comment": "聊聊出国那些事",
"teacup_image": "1111.jpg",
"teacup_tags": "出国;留学;旅游"
}
]
},
"message": "",
"status_code": 6000
}