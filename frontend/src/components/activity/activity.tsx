import React,{Component} from "react";
import {Button, Card, Carousel, Col, Descriptions, Icon, List, message, Result, Row} from "antd";
import {SiderMenu} from "../sider/sider";
import '../common/comm.less'
import {ActivityEntity} from "../../service/activity/structs";
import {ActivityService} from "../../service/activity/activity-service";


class ActivityState {
    activities:Array<ActivityEntity>=[]
    exist:boolean=false
    data:Array<ActivityEntity>=[]
}

export class ActivityLayout extends React.Component{
    activityService = ActivityService.get();
    state = new ActivityState();

    PurchaseTicket=(id:number)=>{
        console.log(this.state.data)
        this.activityService.joinActivity(id).then(res=>{
            if (res.state!==0){
                message.error(res.message)
            }else{
                message.success(res.message)
                this.state.data.forEach((v,index)=>{
                    if (v.id===id){
                        v.canJoin=false
                        v.res--
                    }
                })
                console.log(this.state.data)
                this.setState({
                    activities:this.state.data,
                })
            }
        })
    }
    render() {

        const NoActivity = ()=>(
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
        )



        const Exist=()=>(
            <div>
                {/*<p style={{textAlign:'center'}}>*/}
                {/*    <Carousel autoplay >*/}
                {/*        {this.state.activities.map((v,index)=>{*/}
                {/*            return(*/}
                {/*                <p style={{textAlign:'center'}}>*/}
                {/*                    <Card bordered={false} style={{ textAlign:'center', width: 300,height:300}}*/}
                {/*                          cover={<img alt="example" src={v.img} />}/>*/}
                {/*                </p>*/}
                {/*            )*/}

                {/*        })}*/}
                {/*    </Carousel>*/}
                {/*</p>*/}


                <List
                    bordered
                    dataSource={this.state.activities}
                    renderItem={item => (
                        <List.Item>
                            <div>
                                <h2>
                                    <Row>
                                        <Col span={12}>{item.name}</Col>
                                        <Col span={8}><Button disabled={!item.canJoin} type="link" icon={'star'} onClick={()=>this.PurchaseTicket(item.id)}>我要参加</Button></Col>
                                    </Row>

                                </h2>
                                <p>内容:{item.description}</p>
                                <Row type={'flex'}>
                                    <Col span={16}><span>举办时间:{item.time}</span></Col>
                                    <Col span={4}><span>总:{item.total}</span></Col>
                                    <Col span={4}><span>余:{item.res}</span></Col>
                                </Row>
                            </div>

                            {/*<Descriptions >*/}
                            {/*    <Descriptions.Item label="内容">{item.description}</Descriptions.Item>*/}
                            {/*    <Descriptions.Item label="总票数">{item.total}</Descriptions.Item>*/}
                            {/*    <Descriptions.Item label="余票">{item.res}</Descriptions.Item>*/}
                            {/*    <Descriptions.Item label="举办时间">{item.time}</Descriptions.Item>*/}
                            {/*</Descriptions>*/}
                        </List.Item>
                    )}
                />
            </div>


        )

        return (
            <div>
                {this.state.exist ? <Exist/> : <NoActivity/>}
            </div>


        )
    }
    componentDidMount() {
        this.activityService.getAllActivities().then(res=>{
            if (res.state!==0||!res.data){
                return
            }else{
                res.data.forEach((v,index)=>{
                        v.canJoin=true
                })
                this.setState({
                    exist:true,
                    activities:res.data,
                    data:res.data,
                })
            }
        })
    }
}