export interface Response<T> {
    state: number;
    message: string;
    data?: T;
}
