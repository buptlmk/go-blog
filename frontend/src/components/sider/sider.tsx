import React, {Component} from 'react';
import {Card, Carousel, Icon, Statistic, Timeline, Layout, Anchor, Avatar, Divider, Affix} from "antd";
import Meta from "antd/es/card/Meta";
import {User} from "../user/user";
import {UserEntity} from "../../service/user/struct";
const {Sider}=Layout
const {Link} = Anchor


// 通用sider


export class SiderMenu extends React.Component{


    render() {
        return (
            <Sider style={{background:'#fff',padding:'0 30px'}}>
                <p/>
                <User />
                <p/>
                <Card
                    id={"other"}
                    style={{width:200}}>
                    <Carousel autoplay dots={false}>
                        <Statistic
                            title="访问量"
                            value={99999}
                            // precision={2}
                            valueStyle={{ color: '#3f8600' }}
                            prefix={<Icon type="arrow-up" />}
                            // suffix="%"
                        />
                        <Statistic
                            title="用户"
                            value={9999}
                            // precision={2}
                            valueStyle={{ color: '#3f8600' }}
                            prefix={<Icon type="arrow-up" />}
                            // suffix="%"
                        />
                    </Carousel>
                </Card>
                <p/>

                <Card
                    id = {"girl"}
                    hoverable
                    style={{width:200}}
                    cover={<img alt="example" src={"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"} />}
                >
                    <Meta title="美女" description="为枯燥的页面添加点色彩" />
                </Card>
                <p/>

                <Card id={"devTime"} style={{width:200}}>
                    <Timeline>
                        <Timeline.Item>Start</Timeline.Item>
                        <Timeline.Item color="green">Finish 前端</Timeline.Item>
                        <Timeline.Item color="red">Finish 后端</Timeline.Item>
                        <Timeline.Item dot={<Icon type="clock-circle-o" style={{ fontSize: '16px' }} />}>
                            Test
                        </Timeline.Item>
                    </Timeline>
                </Card>
                <Card style={{textAlign:'center',width:200}} bordered={false}>
                    <Affix offsetTop={800}>
                        <a href="/chat">
                            <Icon type="message" style={{ fontSize: '64px', color: '#08c' }} />
                        </a>
                    </Affix>
                </Card>


            </Sider>
        );
    }
    componentDidMount() {

    }
}