import React,{Component} from "react";
import {CommentEntity} from "../../service/article/structs";
import {List, Comment, Avatar, Form,Button,Input} from "antd";
import moment from 'moment';
import {UserService} from "../../service/user/user";
const { TextArea } = Input;

interface CommentProp {
    comments : Array<CommentEntity>
}

class CommentState{
    comments:Array<CommentEntity>=[]
    value:string=""
    submitting:boolean=false
}

export class CommentStatus extends React.Component<CommentProp>{
    state = new CommentState()
    userService = UserService.get()

    componentDidMount() {
        this.setState({
            comments:this.props.comments,
        })
    }

    handleChange = (e:any) => {
        this.setState({
            value: e.target.value,
        });
    };


    handleSubmit = () => {
        if (!this.state.value) {
            return;
        }

        this.setState({
            submitting: true,
        });

        setTimeout(() => {
            this.setState({
                submitting: false,
                value: '',
                comments: [
                    ...this.state.comments,
                    {
                        user: this.userService.name,
                        userId:this.userService.id,
                        content: <p>{this.state.value}</p>,
                        createdTime: moment().fromNow(),
                    },

                ],
            });
        }, 1000);
    };


    render(){
        return(

            <div>
                <List
                    className="comment-list"
                    header={`${this.props.comments.length} replies`}
                    itemLayout="horizontal"
                    dataSource={this.state.comments}
                    renderItem={item => (
                        <li style={{height:60}}>
                            <Comment
                                // actions={item.actions}
                                author={item.user}
                                avatar={<Avatar src={"/upload/"+item.userId+".png"}/>}
                                content={item.content}
                                datetime={item.createdTime}

                            />
                        </li>
                    )}
                />

                <Comment
                    avatar={
                        <Avatar
                            src={"/upload/"+this.userService.getId()+".png"}
                            alt={this.userService.getName()}
                        />
                    }
                    content={
                        <div>
                            <Form.Item>
                                <TextArea rows={3} onChange={this.handleChange} value={this.state.value} />
                            </Form.Item>
                            <Form.Item>
                                <Button htmlType="submit" loading={this.state.submitting} onClick={this.handleSubmit} type="primary">
                                    Add Comment
                                </Button>
                            </Form.Item>

                        </div>
                    }
                />
            </div>
        )
    }
}