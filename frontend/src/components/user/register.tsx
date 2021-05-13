import * as React from "react";
import {Alert, Button, Card, Col, Divider, Input, message, Row} from 'antd';
import './register.less';
import {UserService} from "../../service/user/user";
import {Link, Redirect} from "react-router-dom";

interface RegisterState {
    registerStatus: string;
    alertType: 'success' | 'info' | 'warning' | 'error';
    alertMsg: string;
    alertDescription: string;
    registerButtonEnable: boolean;
}

interface RegisterProp {

}

export class Register extends React.Component<RegisterProp, RegisterState> {

    schoolCard: string = "";
    user: string = "";
    mail: string = "";
    phone: string = "";
    pass: string = "";

    constructor(props: Readonly<RegisterProp>) {
        super(props);
        this.state = {
            alertDescription: "",
            alertMsg: "",
            alertType: "success",
            registerStatus: "",
            registerButtonEnable: false,
        };
    }

    userService: UserService = UserService.get();

    async onRegisterButtonClick() {
        let res = await this.userService.register(this.schoolCard, this.pass,this.user, this.phone, this.mail, );
        if (res.state !== 0) {
            message.warn(`"register failed"${res.message}`, 2.5);
            return
        }
        message.success(`register success`, 1.5);
        this.setState({registerStatus: "success"});
    }

    render() {
        return (
            <Row type="flex" justify="space-around" align="middle" className="registerRow">
                <Col span={4}>
                    <Card title="注册" extra={<Link to="/login">登陆</Link>} style={{width: '24rem'}}>
                        {this.state.registerStatus === "success" ? <Redirect to='/login'/> : null}
                        {
                            this.state.registerStatus !== "" ? (<Alert
                                message={this.state.alertMsg}
                                description={this.state.alertDescription}
                                type={this.state.alertType}
                                showIcon
                            />) : undefined
                        }

                        <p/>
                        <Input
                            placeholder="请输入您的学号"
                            onChange={input => this.schoolCard = input.target.value}
                        />
                        <p/>
                        <Input
                            placeholder="请输入您的中文名字(汉字)"
                            onChange={input => this.user = input.target.value}
                        />
                        <p/>
                        <Input
                            placeholder="请输入您的手机号码"
                            onChange={input => this.phone = input.target.value}
                        />
                        <p/>
                        <Input
                            placeholder="请输入您的邮箱"
                            onChange={input => this.mail = input.target.value}
                        />
                        <p/>
                        <Input.Password
                            placeholder="请输入您的密码"
                            onChange={input => this.pass = input.target.value}
                        />
                        <p/>
                        <Input.Password
                            placeholder="请确认您的密码"
                            onChange={input => {
                                if (input.target.value !== this.pass) {
                                    this.setState({
                                        registerStatus: "failed",
                                        alertType: "warning",
                                        alertMsg: "password not matched",
                                        alertDescription: "",
                                        registerButtonEnable: false,
                                    })
                                } else {
                                    this.setState({
                                        registerStatus: "",
                                        registerButtonEnable: true,
                                    })
                                }
                            }}
                        />
                        <Divider/>
                        <p/>
                        <Button type="primary" block onClick={() => this.onRegisterButtonClick()}
                                disabled={!this.state.registerButtonEnable}>注册</Button>
                    </Card>
                </Col>
            </Row>
        );
    }


}
