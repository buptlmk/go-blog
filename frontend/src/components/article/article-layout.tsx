import * as React from "react";
import {ArticleService} from "../../service/article/articleService";
import {Button, List, message, Card, Skeleton, Avatar, Tooltip, Row, Col, Affix, Icon} from "antd";
import {ArticleEntity} from "../../service/article/structs";
import {ArticleStatus} from "./article-status";
import {SiderMenu} from "../sider/sider";
import ReactMarkdown from 'react-markdown';
// 不适合ts
import MarkdownNavbar from 'markdown-navbar/index.js';
// The default style of markdown-navbar should be imported additionally
import 'markdown-navbar/dist/navbar.css';
// import '../common/comm.less'

class ArticleProp {
}

class ArticleState {
    articles: Array<ArticleEntity> = [];
    data: Array<ArticleEntity> = [];
    initLoading: boolean = true;
    loading = false;

    addEvent = false
}

export class ArticleLayout extends React.Component<ArticleProp, ArticleState> {
    articleService: ArticleService = ArticleService.get();
    state: ArticleState = new ArticleState();
    getNumber:number=0

    newArticle = () => {
        let temp = new ArticleEntity()
        temp.loading = true
        return temp
    }

    onLoadMore = () => {
        this.setState({
            loading: true,
            articles: this.state.data.concat([...new Array(1)].map(() => (this.newArticle()))),
        });
        message.info("loading...", 1)

        this.articleService.getRecentArticles().then(res => {
            if (res.state !== 0 || !res.data) {
                message.warn("get articles failed", 2.5);
                this.setState({
                    loading: false,
                })
                return
            }
            const data = this.state.data.concat(res.data)
            this.setState({
                    data: data,
                    articles: data,
                    loading: false,

                },
                () => {
                    window.dispatchEvent(new Event('resize'));
                });
        });
    };

    onscroll = () => {
        //变量scrollTop是滚动条滚动时，距离顶部的距离，兼容ie
        let scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
        //变量windowHeight是可视区的高度，兼容ie
        let windowHeight = document.documentElement.clientHeight || document.body.clientHeight;
        //变量scrollHeight是滚动条的总高度，兼容ie
        let scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight;
        //滚动条到底部的条件，兼容ie

        let href = window.location.href
        let path = href.split('//')[1].split('/')[1]
        if (scrollTop + windowHeight >= scrollHeight && (path.startsWith("article") || path ==="")) {
            if (new Date().getTime() - this.getNumber>1500){
                this.getNumber = new Date().getTime()
                console.log(this.getNumber)
                this.onLoadMore()

            }
        }
    };

        render() {
            return (
                <Row className="comm-main" type="flex" justify="center">
                    <Col className="comm-left" xs={24} sm={24} md={22} lg={19} xl={19}>
                        <List
                            itemLayout="vertical"
                            className="demo-loadmore-list"
                            loading={this.state.initLoading}
                            // size="large"
                            bordered={false}
                            split ={false}
                            // loadMore={loadMore}
                            dataSource={this.state.articles}
                            renderItem={item => (
                                <div>
                                    <Card>
                                        <List.Item
                                            key={item.title}
                                            // @ts-ignore
                                            // actions={
                                            //     !this.state.loading && [
                                            //         <ArticleStatus article={item}/>,
                                            //     ]
                                            // }
                                            // extra={
                                            //     !this.state.loading && (
                                            //         <img
                                            //             width={128}
                                            //             height={128}
                                            //             alt="logo"
                                            //             src="/upload/bear.png"
                                            //         />
                                            //     )
                                            // }
                                        >
                                            <Skeleton loading={item.loading} title={false} active avatar>
                                                <List.Item.Meta
                                                    avatar={<a href={"/person/"+item.authorId}> <Tooltip title={item.author}><Avatar src={"/upload/"+item.authorId+".png"}/> </Tooltip></a>}
                                                    title={item.title}
                                                    description={item.createdTime}
                                                />

                                                <ReactMarkdown
                                                    children={item.content.substr(0,100)}
                                                    skipHtml={false}
                                                />
                                                {/*<Row type={'flex'}>*/}
                                                {/*    <Col span={21}>*/}
                                                {/*        */}

                                                {/*    </Col>*/}
                                                {/*    <Col >*/}
                                                {/*        */}

                                                {/*    </Col>*/}
                                                {/*</Row>*/}
                                                <div style={{textAlign:'right'}}>
                                                    <a href={"/article/detailed/"+item.id}>
                                                        <Icon type="read" style={{ fontSize: '16px', color: '#08c' }}/>
                                                        <span style={{ fontSize: '14px', color: '#08c' }}>详细阅读</span>
                                                    </a>
                                                </div>


                                                {/*<a href={"/article/detailed/"+item.id} style={{textAlign:'center'}}>*/}
                                                {/*    <Icon type="read" style={{ fontSize: '32px', color: '#08c' }}/>*/}
                                                {/*</a>*/}
                                                {/*<Button icon={'read'} type={'link'} style={{ fontSize: '32px', color: '#08c',textAlign:'center' }}>详细阅读</Button>*/}
                                            </Skeleton>

                                        </List.Item>
                                        {this.state.loading?null:<ArticleStatus article={item}/>}
                                    </Card>
                                    <p/>
                                </div>
                            )}
                        />
                    </Col>
                    <Col className="comm-right" xs={0} sm={0} md={2} lg={4} xl={4}>
                        <SiderMenu/>
                    </Col>

                </Row>

            );
        }
        componentDidMount(): void {
            window.addEventListener("scroll",this.onscroll)


            this.articleService.getRecentArticles().then(
                res => {
                    if (res.state !== 0 || !res.data) {
                        message.warn("get articles failed", 2.5);
                        return
                    }
                    this.setState({
                        articles: res.data,
                        initLoading: false,
                        data: res.data,
                    });
                }
            )
       }
       componentWillUnmount() {
           window.removeEventListener("scroll",this.onscroll)
       }
}
