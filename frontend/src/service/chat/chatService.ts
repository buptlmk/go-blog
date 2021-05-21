import {BaseService} from "../base-service";
import axios from "axios";
import {Response} from "../response";
import {ArticleEntity} from "../article/structs";
import {ChatEntity, TalkEntity} from "./struct";


export class ChatService extends BaseService{
    public static instance: ChatService

    public static get = ()=>{
        if (ChatService.instance){
            return ChatService.instance
        }else{
            ChatService.instance = new ChatService()
            return ChatService.instance
        }
    }

    async getChatRooms(){
        let res
        try {
            let response = await axios.get<Response<Array<ChatEntity>>>("/chat/rooms",{
                headers:{
                    "token":this.getToken(),
                }
            })
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:[],
            }
        }
        console.log(res)
        return res

    }


    async joinChatRoom(id:number,name:string){
        let res
        try {
            let response = await axios.post<Response<Array<ChatEntity>>>("/chat/join",
                {
                    "id":id,
                    "name":name,
                },
                {
                    headers: {
                        "token": this.getToken(),
                    }
                })
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:[],
            }
        }
        console.log(res)
        return res
    }
    async exitChatRoom(id:number,name:string){
        let res
        try {
            let response = await axios.post<Response<null>>("/chat/exit",
                {
                    "id":id,
                    "name":name,
                },
                {
                    headers: {
                        "token": this.getToken(),
                    }
                })
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:null,
            }
        }
        console.log(res)
        return res
    }

    async pushMessage(name:string,content:string){
        let res
        try {
            let response = await axios.post<Response<null>>("/chat/push/"+name,
                {
                    "type":1,
                    "user":this.getName(),
                    "cardId":this.getCardId(),
                    "text":content,
                },
                {
                    headers: {
                        "token": this.getToken(),
                    }
                })
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:null,
            }
        }
        console.log(res)
        return res
    }


    async getMessage(name:string){
        let res
        try {
            let response = await axios.post<Response<Array<TalkEntity>>>("/chat/pull",
                {
                    "room_name":name,
                    "cardId":this.getCardId()
                },
                {
                    headers: {
                        "token": this.getToken(),
                    }
                })
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:[],
            }
        }
        console.log(res)
        return res
    }

}