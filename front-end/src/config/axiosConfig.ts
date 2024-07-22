import axios from 'axios';

const axiosBase = axios.create({
  baseURL: process.env.REACT_APP_NODE_PROXY_URL // Set from compose for local dev or production
});

export default axiosBase;