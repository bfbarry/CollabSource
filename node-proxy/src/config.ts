import axios from 'axios';

export const secretKey: string = process.env.JWT_SECRET_KEY; 

export const backendUrl: string = 'http://backend:8080'; // Might need to be replaced when live

export const axiosBase = axios.create({
  baseURL: process.env.BACK_END_URL, // Set from compose for local dev or production
});

