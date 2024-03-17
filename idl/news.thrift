// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go news_gorm
namespace py news_gorm
namespace java news_gorm

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

enum State {
    Background = 0
    Publish    = 1
    Processing  = 2
}

struct News {
    1: i64 id
    2: string title
    3: State state
    4: i64    cid
    5: string content
}

struct CreateNewsRequest{
    1: string title      (api.body="title", api.form="title",api.vd="(len($) > 0 && len($) < 100)")
    2: State state    (api.body="state", api.form="state",api.vd="($ == 1||$ == 2)")
    3: i64    cid       (api.body="cid", api.form="cid",api.vd="$>0")
    4: string content (api.body="content", api.form="content",api.vd="(len($) > 0 && len($) < 1000)")
}

struct CreateNewsResponse{
   1: Code code
   2: string msg
}

struct QueryNewsRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryNewsResponse{
   1: Code code
   2: string msg
   3: list<News> news
   4: i64 totoal
}

struct DeleteNewsRequest{
   1: i64    id   (api.path="id",api.vd="$>0")
}

struct DeleteNewsResponse{
   1: Code code
   2: string msg
}

struct UpdateNewsRequest{
    1: i64    id   (api.path="id",api.vd="$>0")
    2: string title      (api.body="title", api.form="title",api.vd="(len($) > 0 && len($) < 100)")
    3: State state    (api.body="state", api.form="state",api.vd="($ == 1||$ == 2)")
    4: i64    cid       (api.body="cid", api.form="cid",api.vd="$>0")
    5: string content (api.body="content", api.form="content",api.vd="(len($) > 0 && len($) < 1000)")
}

struct UpdateNewsResponse{
   1: Code code
   2: string msg
}


service NewsService {
   UpdateNewsResponse UpdateNews(1:UpdateNewsRequest req)(api.post="/v1/news/update/:id")
   DeleteNewsResponse DeleteNews(1:DeleteNewsRequest req)(api.post="/v1/news/delete/:id")
   QueryNewsResponse  QueryNews(1: QueryNewsRequest req)(api.post="/v1/news/list/")
   CreateNewsResponse CreateNews(1:CreateNewsRequest req)(api.post="/v1/news/create/")
}