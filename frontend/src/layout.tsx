import React,{Component} from "react";
import {Affix, Avatar, Col, Icon, Layout, Menu, Row} from "antd";
import {Link, Route, Switch} from "react-router-dom";
import SubMenu from "antd/es/menu/SubMenu";
import {ArticleLayout} from "./components/article/article-layout";
import {SiderMenu} from "./components/sider/sider";
import {Navigation} from "./components/sider/header";
import {ActivityLayout} from "./components/activity/activity";
import {ArticleDetailed} from "./components/article/ArticleDetailed";
import {url} from "inspector";
import {ArticleEdit} from "./components/article/article-edit";
import {ChatLayout} from "./components/chat/ChatLayout";
const {Footer, Content} = Layout;

export class Layouts extends React.Component{
    render() {
        return(
            <Layout style={{background:'#fff'}}>
                <Navigation/>
                <Content style={{minHeight: '100vh'}}>
                    {/*<p/>*/}
                    <Layout style={{minHeight: '100vh'}}>
                        <Content style={{background: '#f6f6f6'}}>
                            <Switch>
                                <Route exact path={['/article','/']} component={ArticleLayout}/>
                                <Route path={'/activity'} component={ActivityLayout}/>
                                <Route path={'/article/detailed/:id'} component={ArticleDetailed}/>
                                <Route path={'/article/edit'} component={ArticleEdit}/>
                                <Route path={'/chat'} component={ChatLayout}/>

                            </Switch>
                        </Content>
                        {/*<SiderMenu/>*/}

                    </Layout>

                </Content>
                <Footer id='footer' style={{textAlign: 'center'}}>BLOG ©2020 一个不会写代码的瘦子</Footer>
            </Layout>
        )
    }
}