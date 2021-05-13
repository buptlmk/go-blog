import React,{Component} from "react";
import {Avatar, Card, Divider, Popover, Tooltip} from "antd";
import './user.css'
import {UserEntity} from "../../service/user/struct";
import {UserService} from "../../service/user/user";

interface userProp{
    user:UserEntity
}

class UserState{
    exist:boolean=false;
    userInfo:UserEntity=new UserEntity();
}
export class User extends Component{
    state = new UserState();
    userService = UserService.get();
    render() {

        return (

            <Card style={{width:200}}>
                <div style={{textAlign:'center'}}><Avatar size={80} src={"/upload/"+this.state.userInfo.id+".png" } /></div>
                <div style={{textAlign:'center'}}>
                    <span>{this.state.userInfo.name}</span>
                    <Divider>社交账号</Divider>
                    <a href={this.state.userInfo.github}>
                        <Avatar size={28} icon="github" className="account" />
                    </a>
                    <Popover placement={"bottomLeft"} content={<img alt="用户暂未提供" src={this.state.userInfo.qq} width={96} height={96}/>} title="QQ">
                        <Avatar size={28} icon="qq"  className="account" />
                    </Popover>
                    <Popover placement={"bottomLeft"} content={<img alt="用户暂未提供" src={this.state.userInfo.wechat} width={96} height={96}/>} title="微信">
                        <Avatar size={28} icon="wechat"  className="account"  />
                    </Popover>

                </div>
            </Card>
        );
    }

    componentDidMount() {

        this.userService.getUser().then(res=>{

            if (res.state!==0){
                // 设置预设的个人信息
               return
            }else{
                this.setState({
                    exist:true,
                    userInfo:res.data,
                })
            }
        })

    }
}