import {BaseService} from "../base-service";
import axios from "axios";
import {Response} from "../response";
import {ArticleEntity} from "../article/structs";
import {createHash} from "crypto";
import {UserEntity} from "./struct";

export class UserService extends BaseService{
    private static instance : UserService;

    public static  get = ()=>{
        if (UserService.instance){
            return UserService.instance
        }else{
            UserService.instance = new UserService();
            return UserService.instance;
        }
    };


    async login(cardId:string,password:string){
        let res
        console.log(createHash('sha256').update(password).digest('hex'))
        try{
            let response = await axios.post<Response<number>>("/login",{
                card_id:cardId,
                password:createHash('sha256').update(password).digest('hex'),
            });
            res = response.data

        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:[],
            }
        }
        return res
    }

    async register(cardId:string,password:string,user:string,phone:string,email:string){
        let res
        try{
            let response = await axios.post<Response<null>>("/register",{
                card_id:cardId,
                password:createHash('sha256').update(password).digest('hex'),
                user:user,
                phone:phone,
                email:email,
            });
            res = response.data

        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText,
                data:[],
            }
        }
        return res
    }

    async hasLogin(){
        // console.log(this.getToken(),this.getId(),this.getName())
        let token = this.getToken()
        console.log(token)
        if (token===""||token===undefined){
            return false
        }
        // TODO:这里应进一步检查
        return true
    };

    async getUser(){
        let res
        let token = this.getToken()
        if (token===""||token===undefined){
            return {
                state:1,
                message:"未登录",
                data:null,
            }
        }
        try{
            let response = await axios.get<Response<UserEntity>>("/user/detail",{
                "headers":{
                    "token":token
                }
            });
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