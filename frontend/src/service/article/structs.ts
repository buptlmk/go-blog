export class ArticleEntity {
    id: number = 0;
    title: string = "";
    content: string = "";
    author: string = "";
    authorId:number = 0;
    createdTime: string = "2019-08-29T08:10:27.180940909Z";
    visitedNumber: number = 0;
    starNumber: number = 0;
    likeNumber: number = 0;
    commentNumber:number = 0;
    loading: boolean = false
}

export class CommentEntity {
    userId: number = 0;
    user: string = "";
    articleId: number = 0;
    content: string = "";
    createdTime: string = "2019-08-29T08:10:27.180940909Z";
    // createdTime : Date = new Date()
}