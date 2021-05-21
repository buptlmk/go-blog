import React from "react";
import {TalkEntity} from "../../service/chat/struct";
import {Card, Comment, List} from "antd";
import {ChatService} from "../../service/chat/chatService";


interface ChatContentProp{
    Content:Array<TalkEntity>|undefined
    Name:string
}

class ChatContentState {
    content:Array<TalkEntity>=[]
    name:string=""
}

export class ChatContent extends React.Component<ChatContentProp>{
    // state = new ChatContentState()
    chatService = ChatService.get()

    // getMessage=(name:string)=>{
    //     this.chatService.getMessage(name).then(res=>{
    //         if (res.state!==0||!res.data){
    //             console.log(res.message)
    //             return
    //         }else{
    //             this.setState({
    //                 Content:this.props.Content.concat(res.data),
    //             })
    //         }
    //     })
    // }

    render (){
        return(
            <Card style={{minHeight: '70vh'}} title={<li style={{textAlign:'center'}}>{this.props.Name}</li>}>
                <List
                    itemLayout="horizontal"
                    dataSource={this.props.Content}
                    renderItem={(item,index) => (
                        <li>
                            <Comment
                                author={item.user}
                                // avatar={item.cardId}
                                content={item.text}
                            />
                        </li>
                    )}
                />
            </Card>
        )
    }
    componentDidMount() {
        console.log(this.state)
        // if (this.props.Name===""){}
    }

    // componentDidUpdate(prevProps: Readonly<ChatContentProp>, prevState: Readonly<{}>, snapshot?: any) {
    //     console.log("---",prevProps.Name,this.props.Name)
    // }


}