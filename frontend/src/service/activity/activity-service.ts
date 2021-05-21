import {BaseService} from "../base-service";
import axios from 'axios'
import {Response} from "../response";
import {ActivityEntity} from "./structs";

export class ActivityService extends BaseService{

    private static instance : ActivityService;

    public static  get = ()=>{
        if (ActivityService.instance){
            return ActivityService.instance
        }else{
            ActivityService.instance = new ActivityService();
            return ActivityService.instance;
        }
    };

    async getAllActivities(){

        let res
        try {
            let response = await axios.get<Response<Array<ActivityEntity>>>("/activity")
            res = response.data;
        }catch (e) {
            res = {
                state:e.response.status,
                message:e.response.statusText,
                data:[]
            }
        }
        console.log(res)
        return res
    }

    async joinActivity(id:number){
        let res
        try {
            let response = await axios.get<Response<null>>("/activity/join/"+id,{
                headers:{
                    "token":this.getToken(),
                }
            })

            res = response.data
        }catch (e) {
            res = {
                state:e.response.status,
                message:e.response.statusText,
                data:null,
            }
        }
        console.log(res)
        return res
    }

}