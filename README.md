- [支持平台](#支持平台)
- [安装](#安装)
- [Docker](#docker)
- [依赖模块](#依赖模块)

Golang 短视频去水印, 视频目前支持 20 个平台, 图集目前支持 4 个平台, 欢迎各位 Star。

> ps: 使用时, 请尽量使用 app 分享链接, 电脑网页版未做测试.

# 其他语言版本

- [Python 版本](https://github.com/wujunwei928/parse-video-py)

# 支持平台

## 图集

| 平台   | 状态 | 域名                            |
| ------ | ---- | ------------------------------- |
| 抖音   | ✔    | https://p3-sign.douyinpic.com;  |
| 快手   | ✔    | https://p4.a.yximgs.com;         |
| 小红书 | ✔    | https://ci.xiaohongshu.com;     |
| 皮皮虾 | ✔    | https://p3-ppx-sign.byteimg.com; |

## 视频

| 平台           | 状态 | 域名                              |
| -------------- | ---- | --------------------------------- |
| 小红书         | ✔    | https://sns-video-hw.xhscdn.com;  |
| 皮皮虾         | ✔    | https://v3-cdn-tos.ppxvod.com;    |
| 抖音短视频     | ✔    | https://v26-cold.douyinvod.com;    |
| 火山短视频     | ✔    |                                   |
| 皮皮搞笑       | ✔    |                                   |
| 快手短视频     | ✔    | https://alimov2.a.kwimgs.com;     |
| 微视短视频     | ✔    |                                   |
| 西瓜视频       | ✔    |                                   |
| 最右           | ✔    |                                   |
| 梨视频         | ✔    |                                   |
| 度小视(原全民) | ✔    |                                   |
| 逗拍           | ✔    |                                   |
| 微博           | ✔    | https://f.video.weibocdn.com;     |
| 绿洲           | ✔    |                                   |
| 全民 K 歌      | ✔    |                                   |
| 6 间房         | ✔    |                                   |
| 美拍           | ✔    |                                   |
| 新片场         | ✔    |                                   |
| 好看视频       | ✔    | https://vd2.bdstatic.com;         |
| 虎牙           | ✔    |                                   |
| AcFun          | ✔    | https://tx-safety-video.acfun.cn; |

# 安装

```go
// 根据分享链接解析
res, _ := parser.ParseVideoShareUrl("分享链接")
fmt.Printf("%#v", res)

// 根据视频id解析
res2, _ := parser.ParseVideoId(parser.SourceDouYin, "视频id")
fmt.Printf("%#v", res2)
```

# Docker

获取 docker image

```bash
docker pull wujunwei928/parse-video
```

运行 docker 容器, 端口 8080

```bash
docker run -d -p 8080:8080 wujunwei928/parse-video
```

查看前端页面  
访问: http://127.0.0.1:8080/

请求接口, 查看 json 返回

```bash
curl 'http://127.0.0.1:8080/video/share/url/parse?url=视频分享链接' | jq
```

返回格式

```json
{
  "author": {
    "uid": "uid",
    "name": "name",
    "avatar": "https://xxx"
  },
  "title": "记录美好生活#峡谷天花板",
  "video_url": "https://xxx",
  "music_url": "https://yyy",
  "cover_url": "https://zzz"
}
```

| 字段名        | 说明           |
| ------------- | -------------- |
| author.uid    | 视频作者 id    |
| author.name   | 视频作者名称   |
| author.avatar | 视频作者头像   |
| title         | 视频标题       |
| video_url     | 视频无水印链接 |
| music_url     | 视频音乐链接   |
| cover_url     | 视频封面       |

> 字段除了视频地址, 其他字段可能为空

# 依赖模块

| 模块                                                                     | 作用                       |
| ------------------------------------------------------------------------ | -------------------------- |
| [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)             | web 框架                   |
| [github.com/go-resty/resty/v2](https://github.com/go-resty/resty/v2)     | HTTP 和 REST 客户端        |
| [github.com/tidwall/gjson](https://github.com/tidwall/gjson)             | 使用一行代码获取 JSON 的值 |
| [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) | jQuery 语法解析 html 页面  |

```bash
go get github.com/gin-gonic/gin
go get github.com/go-resty/resty/v2
go get github.com/tidwall/gjson
go get github.com/PuerkitoBio/goquery
```
