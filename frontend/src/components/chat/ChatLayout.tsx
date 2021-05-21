import React, {Component} from "react";
import {Button, Card, Col, Icon, List, Row, Comment, Avatar, Form, message, Menu, Divider} from "antd";
import {ChatEntity, TalkEntity} from "../../service/chat/struct";
import ListItem from "antd/es/transfer/ListItem";
import TextArea from "antd/es/input/TextArea";
import {ChatService} from "../../service/chat/chatService";
import {ChatContent} from "./ChatContent";

class ChatState {
    rooms: Array<ChatEntity> = []
    data: Array<ChatEntity> = []
    mapTalkEntity:Map<string, Array<TalkEntity>> = new Map()
    mapGetMsg:Map<string,number> = new Map<string, number>()
    talkContent: Array<TalkEntity> = []
    value: string = ""
    currentRoom: string = ""
}

export class ChatLayout extends React.Component {
    state = new ChatState()
    chatService = ChatService.get()
    timer=0
    exit = (index: number, id: number, name: string) => {
        this.state.data.forEach((v) => {
            if (v.id === id) {
                v.join = false
            }
        })
        this.chatService.exitChatRoom(id, name).then(res => {
            if (res.state !== 0) {
                message.error(res.message)
            } else {
                message.success("quit success")
            }
        })

        // 关闭每秒的refresh
        window.clearInterval(this.state.mapGetMsg.get(name))
        this.setState({
            rooms: this.state.data,
            // mapTalkEntity:this.state.mapTalkEntity.delete(name)
        })


    }
    join = (index: number, id: number, name: string) => {
        this.state.data.forEach((v) => {
            if (v.id === id) {
                v.join = true
            }
        })
        this.setState({
            rooms: this.state.data,
        })

        this.chatService.joinChatRoom(id, name).then(res => {
            if (res.state !== 0) {
                message.error(res.message)
            } else {
                message.success("join success")
            }
        })
        this.setState({
            rooms: this.state.data,
            mapTalkEntity:this.state.mapTalkEntity.set(name,[])
        })
        // 开启每秒的refresh
        this.state.mapGetMsg.set(name,setInterval(this.getMessage,1000,name))
    }

    getMessage=(name:string)=>{
        console.log(name)
        this.chatService.getMessage(name).then(res=>{
            if (res.state!==0||!res.data){
                console.log(res.message)
                return
            }else{
                let temp = this.state.mapTalkEntity.get(name)
                if (temp===undefined){
                    temp=[]
                }
                this.state.mapTalkEntity.set(name,temp.concat(res.data))
                this.setState({
                    mapTalkEntity:this.state.mapTalkEntity,
                })
            }
        })
    }

    handleChange = (value: string) => {
        this.setState({
            value: value
        })
    }

    handleSubmit = () => {
        // 在前端检查是否加入了所在房间
        let join:boolean=false
        this.state.data.forEach((v, index) => {
            if (v.name===this.state.currentRoom){
                join=v.join
                return
            }
        })
        if (!join){
            message.error("you can't speak until you join into this room")
            return
        }
        this.chatService.pushMessage(this.state.currentRoom,this.state.value).then(res => {
            if (res.state!==0){
                message.error(res.message)
            }else{
                message.success("send success")
                this.setState({
                    value:"",
                })
            }
        })

    }
    // @ts-ignore
    changeChatContent = e => {
        const v = e.key
        console.log("click ",v)
        if (v === this.state.currentRoom) {
            return
        } else {
            this.setState({
                currentRoom: v,
                talkContent: [],
            })
        }
        // window.clearInterval(this.timer)
    }


    render() {
        return (
            <Card bordered={false} style={{minHeight: '100vh'}}>
                <Row>
                    <Col span={6}>
                        <Card style={{minHeight: '100vh'}}>
                            <h3 style={{textAlign: 'center'}}>Room</h3>
                            <Menu
                                onClick={this.changeChatContent}
                            >
                                {this.state.rooms.map((item, index) => {
                                    return (
                                        <Menu.Item key={item.name}>
                                            <Row>
                                                <Col span={20}>
                                                    <Icon type={"team"}/>
                                                    <span>{item.name}</span>
                                                </Col>
                                                <Col span={4}>
                                                    {item.join ? <Button type={'link'} icon={"export"}
                                                                         onClick={() => this.exit(index, item.id, item.name)}/> :
                                                        <Button type={'link'} icon={"plus"}
                                                                onClick={() => this.join(index, item.id, item.name)}/>
                                                    }
                                                </Col>

                                            </Row>
                                        </Menu.Item>
                                    )
                                })}
                            </Menu>
                        </Card>


                    </Col>
                    <Col offset={1} span={17}>
                        <ChatContent Content={this.state.mapTalkEntity.get(this.state.currentRoom)} Name={this.state.currentRoom}/>
                        <p/>
                        <Comment
                            // avatar={
                            //     <Avatar
                            //         src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                            //         alt="Han Solo"
                            //     />
                            // }
                            content={
                                <Row>
                                    <Col span={20}>
                                        <TextArea style={{height: '6vh'}} rows={3}
                                                  onChange={e => this.handleChange(e.target.value)}
                                                  value={this.state.value}/>
                                    </Col>

                                    <Col offset={1} span={3}>
                                        <Button style={{height: '6vh'}} icon={'enter'} onClick={this.handleSubmit}
                                                type="primary"/>
                                    </Col>

                                </Row>
                            }
                        />


                    </Col>
                </Row>
            </Card>
        );
    }

    componentDidMount() {
        this.chatService.getChatRooms().then(res => {
            if (res.state !== 0 || !res.data) {
                return
            } else {
                res.data.forEach((v, index) => {
                    v.join = false
                })

                this.setState({
                    rooms: res.data,
                    data: res.data,
                })
            }
        })
    }
}