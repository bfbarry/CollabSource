import axios from 'axios';

let axiosBase = axios.create({
  baseURL: process.env.REACT_APP_NODE_PROXY_URL || "http://127.0.0.1:8000", // Set from compose for local dev or production
  // headers: {
  //   'Authorization': localStorage.getItem('auth_context_state') ? localStorage.getItem('auth_context_state') : "public"
  // }
});

export default axiosBase;