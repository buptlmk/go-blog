import {ArticleEntity, CommentEntity} from "../../service/article/structs";
import React from "react";
import {Avatar, Button, Col, Icon, List, message, Row, Skeleton} from "antd";
import {ArticleService} from "../../service/article/articleService";
import {CommentStatus} from "./comment";


interface ArticleStatusProp {
    article:ArticleEntity
}

class ArticleStatusState{
    starNumber : number=0
    likeNumber : number=0
    visitedNumber:number=0
    commentNumber:number=0


    comments : Array<CommentEntity> = []
    // starBool : boolean =false
    // likeBool : boolean = false
    commentBool: boolean = false


}


export class ArticleStatus extends React.Component<ArticleStatusProp,ArticleStatusState>{

    state = new ArticleStatusState();
    articleService: ArticleService = ArticleService.get();

    componentDidMount() {
        this.setState({
            starNumber:this.props.article.starNumber,
            likeNumber:this.props.article.likeNumber,
            visitedNumber:this.props.article.visitedNumber,
            commentNumber:this.props.article.commentNumber,
        })
    }

    star=()=>{
        this.articleService.starArticle(this.props.article.id).then(res=>{
            if (res.state !== 0 ) {
                message.warn(res.message, 1.5);
                return
            }else{
                message.success("star article success",1.5)
                this.setState({
                    starNumber:this.state.starNumber+1
                })
            }
        })
    }
    openComment = ()=>{
        if (!this.state.commentBool){
            // this.setState({
            //     comments:this.state.comments.concat([...new Array(1)])
            // })
            console.log("---->")
            this.articleService.getComments(this.props.article.id).then(res=>{
                if (res.state !== 0) {
                    message.warn(res.message, 1.5);
                    return
                }else{
                    // message.success("star article success",1.5)
                    // const data = this.state.comments.concat(res.data)
                    if (!res.data){
                        this.setState({
                            commentBool:true
                        },
                        () => {
                            window.dispatchEvent(new Event('resize'));
                        })
                    }else{
                        this.setState({
                            commentBool:true,
                            comments:res.data,
                        },
                        () => {
                            window.dispatchEvent(new Event('resize'));
                        })
                    }
                }
            })
        }else{
            this.setState({
                commentBool:false,
                comments:[],
            })
        }
    }

    render(){
        return(
            <div>
                <Row type='flex' justify={"start"}>
                    <Col>
                        <span>
                            <Button type="link" icon={"eye"} />
                            {this.state.visitedNumber > 1000 ?  (this.state.visitedNumber/1000).toFixed(1)+"k":this.state.visitedNumber}
                        </span>
                    </Col>
                    <Col>
                        <span>
                            <Button type={"link"} icon={"star-o"} onClick={()=>{this.star()}}/>
                            { this.state.starNumber > 1000 ?  (this.state.starNumber/1000).toFixed(1)+"k":this.state.starNumber}
                        </span>
                    </Col>
                    <Col>
                        <span>
                            <Button type={"link"} icon={"like-o"}/>
                            {this.state.likeNumber > 1000 ?  (this.state.likeNumber/1000).toFixed(1)+"k":this.state.likeNumber}
                        </span>
                    </Col>
                    <Col>
                        <span>
                            <Button type={"link"} icon={"message"} onClick={()=>this.openComment()}/>
                            {this.state.commentNumber > 1000 ?  (this.state.commentNumber/1000).toFixed(1)+"k":this.state.commentNumber}
                        </span>
                    </Col>
                </Row>
                {!this.state.commentBool?null:<CommentStatus articleId={this.props.article.id} comments={this.state.comments}/>}
            </div>

        )
    }

}