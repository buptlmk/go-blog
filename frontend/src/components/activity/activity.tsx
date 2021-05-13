import React,{Component} from "react";
import {Col, Icon, Result, Row} from "antd";
import {SiderMenu} from "../sider/sider";
import '../common/comm.less'

export class ActivityLayout extends React.Component{
    render() {
        return (
            <Row className="comm-main" type="flex" justify="center">
                <Col className="comm-left" xs={23} sm={23} md={21} lg={19} xl={19}>
                    <Result
                        icon={<Icon type="smile" theme="twoTone" />}
                        title="Sorry, we have no more activities!"
                    />
                </Col>
                <Col className="comm-right" xs={0} sm={0} md={2} lg={4} xl={4}>
                    <SiderMenu/>
                </Col>

            </Row>
        );
    }
}