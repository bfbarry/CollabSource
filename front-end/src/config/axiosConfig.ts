import axios from 'axios';

const axiosBase = axios.create({
  baseURL: process.env.REACT_APP_NODE_PROXY_URL || "https://node-proxy-4jjvxz6spq-uc.a.run.app", // Set from compose for local dev or production
  headers: {
    'Authorization': localStorage.getItem('access_token') ? localStorage.getItem('access_token') : "public"
  }
});

export default axiosBase;