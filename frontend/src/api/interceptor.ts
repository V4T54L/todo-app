import axios, { AxiosResponse, AxiosError, InternalAxiosRequestConfig, AxiosRequestHeaders } from 'axios';

const instance = {
    axios: axios.create({
        baseURL: 'http://localhost:8080/api/v1',
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