import axios from "axios";
import {Response} from "./response";
import Cookies from "universal-cookie/es6";

class CachedData {
    time = Date.now();
    data: string;

    constructor(data: string) {
        this.data = data;
    }
}

export class BaseService {

    cookies = new Cookies();
    id: number = 0;
    cardId:string="";
    name: string = "";
    operator: number = this.id;
    // session: string = "1d5a59fd33fd1488723642b0cacc2ff04d3e8438348fa2162d4a1cbff5adaff5";
    // token: string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA4Nzc5NDIsImlkIjoxMiwiY2FyZF9pZCI6IjIwMTYyMTA3OTUiLCJleHBpcmUiOmZhbHNlfQ.Yc3A_RENNp1C3DVBAWaipFBM1x-hDA8vN4WPIFYZUEE"
    // userService: UserService = UserService.get();
    token: string = ""
    cache = new Map<string, Map<string, CachedData>>();

    getName(): string {
        if (this.name === "" || this.name === undefined) {
            this.name = this.cookies.get("name");
        }
        return this.name
    }
    setToken(){
        this.token=""
    }

    getId(): number {
        if (this.id === 0 || this.id === undefined) {
            console.log(this.id)
            this.id = Number(this.cookies.get("id"));
        }
        return this.id
    }

    getCardId(): string {
        if (this.cardId === "" || this.cardId === undefined) {
            this.cardId = this.cookies.get("cardId");
        }
        return this.cardId
    }

    getOperator(): number {
        if (!this.operator) {
            this.operator = Number(this.cookies.get("id"));
        }
        return this.operator
    }

    //
    // getSession(): string {
    //     if (!this.session) {
    //         this.session = this.cookies.get("session");
    //     }
    //     return this.session
    // }
    getToken(): string {
        if (this.token === "" || this.token === undefined) {
            this.token = this.cookies.get("token");
        }
        // this.cookies.set("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA4Nzc5NDIsImlkIjoxMiwiY2FyZF9pZCI6IjIwMTYyMTA3OTUiLCJleHBpcmUiOmZhbHNlfQ.Yc3A_RENNp1C3DVBAWaipFBM1x-hDA8vN4WPIFYZUEE")
        return this.token
    }

    addCache(url: string, data: any, res: any) {
        let tempCache = this.cache.get(url);
        if (!tempCache) {
            tempCache = new Map<string, CachedData>();
        }
        tempCache.set(JSON.stringify(data), new CachedData(res));
        this.cache.set(url, tempCache);
    }

    fetchCache(url: string, data: any, cacheTime: number): any | undefined {
        let cache = this.cache.get(url);
        if (cache) {
            let cacheData = cache.get(JSON.stringify(data));
            return cacheData && Date.now() - cacheData.time < cacheTime ? cacheData.data : undefined
        }
        return undefined
    }
    async uploadImage(image:File){
        let res
        let token = this.getToken()
        let param = new FormData()
        param.append('file',image)
        try{
            let response = await axios.post<Response<string>>("/upload/"+image.name,
                param,
                {
                    "headers":{
                        "Content-Type":"multipart/form-data",
                        "token":token,
                    },
                },
            );
            res = response.data
        }catch (e) {
            res = {
                state:e.status,
                message:e.statusText,
                data:'',
            }
        }
        return res
    }

    async cachedPost<T>(url: string, data: any, cacheTime: number = 2000): Promise<Response<T>> {
        let operator = this.getOperator();
        let res: Response<T>;
        let sendData: any = {
            operator: operator,
            user: operator,
            // session: this.getSession(),
        };
        let caller = this.getCallerFile(2);
        // 扩展sendData属性
        for (let i in data) {
            if (data.hasOwnProperty(i)) {
                sendData[i] = data[i]
            }
        }

        let cache = this.fetchCache(url, sendData, cacheTime);
        if (cache) {
            res = JSON.parse(cache);
            console.log(`${caller}(${url}, ${JSON.stringify(data)})=>cache=>`, res);
            return new Promise((resolve, reject) => {
                resolve(res);
            });
        }

        return axios.post<any>(url, sendData).then(
            response => {
                res = response.data;
                this.addCache(url, sendData, JSON.stringify(response.data));
                return res
            }
        ).catch(
            err => {
                res = {
                    state: 2,
                    data: undefined,
                    message: err.response.statusText
                };
                if (err.response) {
                    res = err.response.data
                }
                return res
            }
        ).finally(
            () => {
                console.log(`${caller}(${url}, ${JSON.stringify(data)})=>`, res);
            }
        )


    }

    getCallerFile(position = 2) {
        if (position >= Error.stackTraceLimit) {
            throw new TypeError('getCallerFile(position) requires position be less then Error.stackTraceLimit but position was: `' + position + '` and Error.stackTraceLimit was: `' + Error.stackTraceLimit + '`');
        }

        const oldPrepareStackTrace = Error.prepareStackTrace;
        Error.prepareStackTrace = (_, stack) => stack;
        const stack = new Error().stack;
        Error.prepareStackTrace = oldPrepareStackTrace;
        if (stack !== null && typeof stack === 'object') {
            return stack[position] ? (stack[position] as any).getFunctionName() : undefined;
        }
    }

}


