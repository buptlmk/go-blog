import React,{Component} from "react";
import {CommentEntity} from "../../service/article/structs";
import {List, Comment, Avatar, Form, Button, Input, message} from "antd";
import moment from 'moment';
import {UserService} from "../../service/user/user";
import {ArticleService} from "../../service/article/articleService";
const { TextArea } = Input;

interface CommentProp {
    comments : Array<CommentEntity>
    articleId:number
}

class CommentState{
    comments:Array<CommentEntity>=[]
    value:string=""
    submitting:boolean=false
    articleId:number=0
}

export class CommentStatus extends React.Component<CommentProp>{
    state = new CommentState()
    articleService = ArticleService.get()

    componentDidMount() {
        this.setState({
            comments:this.props.comments,
            articleId:this.props.articleId,
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
        this.articleService.addComment(this.state.articleId,this.state.value).then(res=>{
            if (res.state!==0){
                message.error(res.message)
            }else{
                message.success("comment success")
                this.setState({
                    submitting: false,
                    value: '',
                    comments: [
                        ...this.state.comments,
                        {
                            user: this.articleService.name,
                            userId:this.articleService.id,
                            content: <p>{this.state.value}</p>,
                            createdTime: moment().fromNow(),
                        },

                    ],
                });
            }
        });

        // setTimeout(() => {
        //
        // }, 1000);
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
                            src={"/upload/"+this.articleService.getId()+".png"}
                            alt={this.articleService.getName()}
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