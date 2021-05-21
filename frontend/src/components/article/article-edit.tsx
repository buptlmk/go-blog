import React,{Component} from "react";
import './article-edit.less'
import {Input, Select, Button, DatePicker, Card, message, Row, Col, Result, Icon, Collapse, List, Avatar} from 'antd'
import Editor from "for-editor";
import {UserService} from "../../service/user/user";
import {ArticleService} from "../../service/article/articleService";
import ListItem from "antd/es/transfer/ListItem";


const { Panel } = Collapse;


class ArticleEditState{
    title:string=""
    value:string=""
    files:Array<File>=[]
    filePath:Array<string>=[]
    hasLogin:boolean=false
}

export class ArticleEdit extends React.Component{
    state = new ArticleEditState()
    userService = UserService.get()
    articleService = ArticleService.get()
    handleContent=(v:string)=>{
        this.setState({
            value:v
        })
        console.log(v)
    }
    handleTitle=(v:string)=>{
        this.setState({
            title:v
        })
        console.log(v)
    }

    addImg=(file:File)=>{
        this.articleService.uploadImage(file).then(res=> {
            if (res.state !== 0 || res.data === "") {
                message.error(res.message)
            } else {
                this.setState({
                    // @ts-ignore
                    filePath: this.state.filePath.concat(res.data)
                })
                message.info("图片上传成功", 1)
            }
        })
    }

    publish = ()=>{
        this.articleService.editArticle(this.state.title,this.state.value).then(res=>{
            if (res.state!==0){
                message.error(res.message,2)
            }else{
                message.success("publish success")
                this.setState({
                    title:"",
                    value:"",
                    files:[],
                    filePath:[],
                })
            }
        })
        console.log("publish")
    }

    render(){
        return(
            <div >
                {!this.state.hasLogin ? <Result
                        status="403"
                        title="403"
                        subTitle="Sorry, Please sign in."
                        extra={<Button type="primary" href={"/"}>Back</Button>}
                    /> :
                    <Card bordered={false} style={{minHeight: '100vh'}}>
                        <Row>
                            <Col span={16} offset={2}>
                                <Input
                                    placeholder="博客标题"
                                    size="large"
                                    allowClear={true}
                                    onChange={e => this.handleTitle(e.target.value)}
                                />
                            </Col>
                            <Col span={6} style={{textAlign: 'center'}}>
                                <Row>
                                    <Col span={24}>
                                        <Button size="large" shape={"round"} disabled={true}>保存草稿</Button>&nbsp;
                                        <Button type="primary" size="large" shape={"round"}
                                                onClick={this.publish}>发布文章</Button>
                                        <br/>
                                    </Col>
                                </Row>
                            </Col>
                        </Row>
                        <p/>
                        <Row>
                            <Col span={20}>
                                <Editor
                                    value={this.state.value}
                                    addImg={(file) => this.addImg(file)}
                                    onChange={value => this.handleContent(value)}
                                />
                            </Col>
                            <Col span={4}>
                                <Collapse
                                    bordered={false}
                                    defaultActiveKey={['1']}
                                    expandIcon={({ isActive }) => <Icon type="caret-right" rotate={isActive ? 90 : 0} />}
                                >
                                    <Panel header="Image URl" key="1" className='panel'>
                                        <List
                                            dataSource={this.state.filePath}
                                            renderItem={item =>
                                            <List.Item>
                                                <p>{item}</p>
                                                <Avatar size={80} shape={'square'} src={item}/>
                                            </List.Item>
                                            }
                                        >
                                        </List>
                                    </Panel>
                                </Collapse>
                            </Col>

                        </Row>

                    </Card>
                }
            </div>
        )
    }
    componentDidMount() {

        this.userService.hasLogin().then(res=>{
            if (res){
                this.setState({
                    hasLogin:true
                })
            }

        })


    }
}