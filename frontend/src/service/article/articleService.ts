import {BaseService} from "../base-service";
import axios from 'axios'
import {ArticleEntity, CommentEntity} from "./structs";
import {Response} from "../response";
import moment from "moment";
import React from "react";

export class ArticleService extends BaseService{

    private static instance : ArticleService;

    public static  get = ()=>{
        if (ArticleService.instance){
            return ArticleService.instance
        }else{
            ArticleService.instance = new ArticleService();
            return ArticleService.instance;
        }
    };

    async getRecentArticles(){
        let res
        try{
            let response = await axios.get<Response<Array<ArticleEntity>>>("/article/recent/1");
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

    async getArticle(id:number){
        let res
        try {
            let response = await axios.get<Response<ArticleEntity>>("/article/"+id)
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

    async editArticle(title:string,content:string){
        let res
        let token = this.cookies.get("token")
        let a:number = this.getId()
        console.log(a)
        try {
            let response = await axios.post<Response<null>>(
                "/article/edit",
                {
                    "id":a,
                    "title":title,
                    "content":content,
                },
                {
                    "headers":{
                        "token":token,
                    }
                }
            )
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText
            }
        }
        console.log(res)
        return res

    }

    async starArticle(id:number){
        let res
        let token = this.cookies.get("token")
        try {
            let response = await axios.get<Response<null>>(
                "/article/star/"+id,
                {
                    "headers":{
                        "token":token,
                    }
                }
            )
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText
            }
        }
        console.log(res)
        return res

    }

    async addComment(id:number,content:string){
        let res
        let token = this.cookies.get("token")
        try {
            let response = await axios.post<Response<null>>(
                "/comment/"+id,
                {
                    "user": this.getName(),
                    "userId":this.getId(),
                    "articleId":id,
                    "content": content,
                },
                {
                    "headers":{
                        "token":token,
                    }
                }
            )
            res = response.data
        }catch (err){
            res = {
                state:err.response.status,
                message:err.response.statusText
            }
        }
        console.log(res)
        return res

    }

    async getComments(id:number){
        let res
        let token = this.getToken()

        try {
            let response = await axios.get<Response<Array<CommentEntity>>>("/comment/"+id,
                {
                    "headers":{
                        "token":token,
                    }
                }
            );
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