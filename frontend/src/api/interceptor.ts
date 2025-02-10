import axios, { AxiosResponse, AxiosError, InternalAxiosRequestConfig, AxiosRequestHeaders } from 'axios';

let baseURL = import.meta.env.VITE_BACKEND_URL
if (!baseURL) {
    baseURL = 'http://localhost:8080/api/v1'
}
console.log("Base URL : ", baseURL)

const instance = {
    axios: axios.create({
        baseURL: baseURL,
    }),
    token: "",
    removeToken: () => { }
}

instance.axios.interceptors.request.use(
    (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
        const token = instance.token
        if (token) {
            if (!config.headers) {
                config.headers = {} as AxiosRequestHeaders;
            }
            config.headers.Authorization = `Bearer ${token}`;
            config.headers['Content-Type'] = `application/json`;
        }
        return config;
    },
);

instance.axios.interceptors.response.use(
    (response: AxiosResponse) => response,
    (error: AxiosError) => {
        if (error.response?.status === 401) {
            instance.token = ""
        }
        instance.removeToken()
        return Promise.reject(error);
    },
);

export default instance;