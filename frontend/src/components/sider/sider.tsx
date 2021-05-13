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
                            title="Active"
                            value={11.28}
                            precision={2}
                            valueStyle={{ color: '#3f8600' }}
                            prefix={<Icon type="arrow-up" />}
                            suffix="%"
                        />
                        <Statistic
                            title="Idle"
                            value={9.3}
                            precision={2}
                            valueStyle={{ color: '#cf1322' }}
                            prefix={<Icon type="arrow-down" />}
                            suffix="%"
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
                    <Meta title="Eurpose Street beat" description="www.instagram.com" />
                </Card>
                <p/>

                <Card id={"devTime"} style={{width:200}}>
                    <Timeline>
                        <Timeline.Item>Start</Timeline.Item>
                        <Timeline.Item color="green">Finish task A</Timeline.Item>
                        <Timeline.Item color="red">Finish task B</Timeline.Item>
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