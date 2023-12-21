import axios, { AxiosRequestConfig } from "axios";
import { useState, useEffect } from 'react';



const baseURL = "http://localhost:9090";

const httpClient = axios.create({
  baseURL,
  timeout: 10000,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getReq = async (url: string, data:any, config?: AxiosRequestConfig) => {
  return httpClient.get(url, config);
};

export const postReq = async (url: string, data: any, config?: AxiosRequestConfig) => {
  return httpClient.post(url, data, config);
};

export const patchReq = async (url: string, data: any, config?: AxiosRequestConfig) => {
  return httpClient.patch(url, data, config);
};

export const putReq = async (url: string, data: any, config?: AxiosRequestConfig) => {
  return httpClient.put(url, data, config);
};

export const deleteReq = async (url: string, data:any, config?: AxiosRequestConfig) => {
  return httpClient.delete(url, config);
};


const useData = (url: string, method:string, data: any, config?: AxiosRequestConfig) => {
  const [loading, setLoading] = useState(false);
  const [response, setResponse] = useState(null);
  const [error, setError] = useState<any>(null);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        var requestorFunc;
        switch(method){
          case "GET":
            requestorFunc = getReq;
            break;
          case "POST":
            requestorFunc = postReq;
            break;
          case "PUT":
            requestorFunc = putReq;
            break;
          case "PATCH":
            requestorFunc = patchReq;
            break;
          case "DELETE":
            requestorFunc = deleteReq;
            break;
          default:
            requestorFunc = getReq;
            break;
        }
        const res = await requestorFunc(url, data, config);
        setResponse(res.data);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [url, data, config, method]);

  return { loading, response, error };
};

export default useData;
