import axios from 'axios';

// 定义接口返回类型
interface ResponseData {
    code: number;
    message: string;
    data?: any;
}

// 封装发送 POST 请求的函数
export const sendPost = async (url: string, data: any): Promise<ResponseData> => {
    try {
        const response = await axios.post(url, data);
        return response.data;
    } catch (error) {
        console.error(error);
        throw new Error(error.message);
    }
};

// 封装发送 GET 请求的函数
export const CacheUpdate = async (url: string): Promise<ResponseData> => {
    try {
        const response = await axios.get(url);
        return response.data;
    } catch (error) {
        console.error(error);
        throw new Error(error.message);
    }
};