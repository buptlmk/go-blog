import * as React from "react";
import {Alert, Button, Card, Col, Divider, Icon, Input, Row} from 'antd';
import './login.less';
import {UserService} from "../../service/user/user"
import {Link, Redirect} from "react-router-dom";

interface LoginState {
    loginStatus: string;
    alertType: 'success' | 'info' | 'warning' | 'error';
    alertMsg: string;
    alertDescription: string;
    redirectToIndex: boolean;
}

interface LoginProp {

}

export class Login extends React.Component<LoginProp, LoginState> {


    user: string = "";
    pass: string = "";

    constructor(props: Readonly<LoginProp>) {
        super(props);
        this.state = {
            alertDescription: "",
            alertMsg: "",
            alertType: "success",
            loginStatus: "",
            redirectToIndex: false,
        };
    }

    userService: UserService = UserService.get();

    onLoginButtonClick() {
        this.userService.login(this.user, this.pass).then(
            res=>{
                if (res.state !== 0) {
                    this.setState({
                        loginStatus: "failed",
                        alertType: "warning",
                        alertMsg: "Login Failed",
                        alertDescription: res.message,
                    })
                } else {
                    this.setState({
                        loginStatus: "success",
                        alertType: "success",
                        alertMsg: "Login Success",
                        alertDescription: res.message,
                        redirectToIndex: true,
                    })
                }
            }
        )

    }

    render() {
        return (
            <Row type="flex" justify="space-around" align="middle" className="loginRow">
                <Col>
                    <Card title="登录" extra={<Link to="/register">注册</Link>} style={{width: '24rem'}}>
                        {
                            this.state.loginStatus !== "" ? (<Alert
                                message={this.state.alertMsg}
                                description={this.state.alertDescription}
                                type={this.state.alertType}
                                showIcon
                            />) : null
                        }
                        {
                            this.state.redirectToIndex ? <Redirect to={"/"}/> : null
                        }

                        <p/>
                        <Input
                            placeholder="请输入您的学号"
                            allowClear
                            prefix={<Icon type="user" style={{color: 'rgba(0,0,0,.25)'}}/>}
                            onChange={input => this.user = input.target.value}
                        />

                        <p/>
                        <Input.Password
                            placeholder="请输入您的密码"
                            allowClear
                            prefix={<Icon type="key" style={{color: 'rgba(0,0,0,.25)'}}/>}
                            onChange={input => this.pass = input.target.value}
                        />
                        <Divider/>
                        <p/>
                        <Button type="primary" block onClick={() => this.onLoginButtonClick()}>登陆</Button>
                        <p/>
                        <Link className="forgetWord" to="/updatePassword">忘记密码?</Link>
                    </Card>
                </Col>
            </Row>
        );
    }

    // componentDidMount() {
    //
    // }
}
