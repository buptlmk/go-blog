import React,{Component} from "react";
import {Avatar, Card, Col, Icon, List, message, Result, Row, Skeleton, Tooltip} from "antd";
import {ArticleService} from "../../service/article/articleService";
import {ArticleEntity} from "../../service/article/structs";
import ReactMarkdown from "react-markdown";
import {SiderMenu} from "../sider/sider";
import '../common/comm.less'

interface DetailedProp{
}

class ArticleDetailedState{
    articleId : string="";
    article:ArticleEntity=new ArticleEntity();
    exist:boolean=false;
    loading:boolean=true;
}

export class ArticleDetailed extends React.Component<DetailedProp>{
    state = new ArticleDetailedState()
    articleService: ArticleService = ArticleService.get();
    // test = '### sdgdfs\n' +
    //     '测试图片\n' +
    //     '![img](/upload/12.png)\n' +
    //     '```go\n' +
    //     'import "fmt"\n' +
    //     '\n' +
    //     'func Article(){\n' +
    //     '       test := "test"\n' +
    //     '       fmt.Println(test)\n' +
    //     '}\n' +
    //     '```\n' +
    //     '**还有什么？**'
    render() {
        return (
            <Row className="comm-main" type="flex" justify="center">
                <Col className="comm-left" xs={23} sm={23} md={21} lg={19} xl={19}>
                    <Card bordered={false}>
                        <Skeleton loading={this.state.loading} title={false} active avatar>
                            <List.Item.Meta
                                avatar={<a href={"/person/"+this.state.article.authorId}> <Tooltip title={this.state.article.author}><Avatar src={"/upload/"+this.state.article.authorId+".png"}/> </Tooltip></a>}
                                title={this.state.article.title}
                                description={this.state.article.createdTime}
                            />
                            <ReactMarkdown
                                children={this.state.article.content}
                                skipHtml={false}
                            />
                        </Skeleton>
                    </Card>
                </Col>
                <Col className="comm-right" xs={0} sm={0} md={2} lg={4} xl={4}>
                    <SiderMenu/>
                </Col>

            </Row>

        );
    }

    componentDidMount() {
        let str = window.location.href.split('//')[1].split('/')[3]
        let id = parseInt(str,10)
        if (isNaN(id)){
            return
        }
        this.articleService.getArticle(id).then(
            res => {
                if (res.state !== 0) {
                    message.warn(res.message, 2.5);
                    return
                }
                this.setState({
                    article: res.data,
                    exist:true,
                    loading:false,
                });
            }
        )
    }
}