import React, {Component} from "react";
import {Avatar, Col, Icon, Layout, Menu, Row} from "antd";
import {Link} from "react-router-dom";
import SubMenu from "antd/es/menu/SubMenu";
import {UserService} from "../../service/user/user";

const {Header} = Layout

class NavigationState {
    hasLogin : boolean = false
}

export class Navigation extends React.Component {

    state = new NavigationState()
    userService = UserService.get()

    logout = ()=>{
        this.userService.cookies.set('id',0)
        this.userService.cookies.set('name',"")
        this.userService.cookies.set('token',"")
    }

    render() {

        const NoLogin = ()=>(
            <Menu theme="dark">
                <Menu.Item key="4">
                    <Link to="/login">
                        <span>登录</span>
                    </Link>
                </Menu.Item>
                <Menu.Item key="5">
                    <Link to="/register">
                        <span>注册</span>
                    </Link>
                </Menu.Item>
            </Menu>
        )
        const Login = ()=>(
            <Menu theme="dark">
                <Menu.Item key="4">
                    <Link to="/person">
                        <span>个人中心</span>
                    </Link>
                </Menu.Item>
                <Menu.Item key="5">
                    <a href="/" onClick={()=>this.logout()}>
                        <span>退出</span>
                    </a>
                </Menu.Item>
            </Menu>
        )

        return (
            <Header id='header' style={{ zIndex: 1, width: '100%'}}>
                <Row>
                    <Col style={{float: "left"}}>
                        <Link to="/" className="logo-text">
                            <span>
                                BLOG
                            </span>
                        </Link>
                    </Col>

                    <Col style={{float: "right"}}>
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            style={{lineHeight: '64px'}}
                            // inlineCollapsed={true}
                        >
                            <Menu.Item key='0'>
                                <Link to="/write">
                                    <Icon type="edit" theme="filled" style={{color:'#08ff00'}}/>
                                    <span>创作</span>
                                </Link>
                            </Menu.Item>
                            <Menu.Item key="1">
                                <Link to="/article">
                                    <Icon type="fire" theme="filled" style={{color:'#ff5900'}}/>
                                    <span>文章</span>
                                </Link>
                            </Menu.Item>
                            <Menu.Item key="2">
                                <Link to="/activity">
                                    <Icon type="profile" theme="filled"style={{color:'#0066ff'}} />
                                    <span>活动</span>
                                </Link>
                            </Menu.Item>
                            <Menu.Item key="3">
                                <Link to="/message">
                                    <span>消息</span>
                                </Link>
                            </Menu.Item>
                            <SubMenu
                                title={!this.state.hasLogin?<Avatar style={{color: '#f56a00', backgroundColor: '#fde3cf'}}>游客</Avatar>:<Avatar src={"/upload/"+this.userService.getId()+".png"}/>}>
                                {this.state.hasLogin? <Login/>:<NoLogin/>}
                                {/*<Menu.Item key="4">*/}
                                {/*    <Link to="/login">*/}
                                {/*        <span>登录</span>*/}
                                {/*    </Link>*/}
                                {/*</Menu.Item>*/}
                                {/*<Menu.Item key="5">*/}
                                {/*    <Link to="/register">*/}
                                {/*        <span>注册</span>*/}
                                {/*    </Link>*/}
                                {/*</Menu.Item>*/}
                            </SubMenu>
                        </Menu>
                    </Col>
                    <Col span={1}></Col>
                </Row>
            </Header>
        );
    }

    componentDidMount() {

        this.userService.hasLogin().then(
            res=>{
                this.setState({
                    hasLogin:res,
                })
            }
        )

    }
}