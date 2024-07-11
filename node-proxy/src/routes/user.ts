import express, { Request, Response } from 'express';
import authenticateJWT from '../middlewear/authentication'
import axios, { AxiosResponse } from 'axios';
import { backendUrl } from '../config'

const router = express.Router()

if (process.env.USE_JTW === 'true') {
    router.use(authenticateJWT)
}

const USER_BASE_PATH = '/api/v1/user'

interface User {
    _id: string;
    name: string;
    email: string;
    description: string;
    password: string;
    skills: string[];
} 

router.get('/:id', async (req: Request, res: Response) => {

    const userId = req.params.id;
    const headers = {
        'userEmail':`${req.email}`
    }
    
    const response: AxiosResponse<User> = await axios.get<User>(`${backendUrl}${USER_BASE_PATH}/${userId}`, { headers });
    const user: User =  response.data;

    res.status(response.status).json({ user });
});

router.patch('/:id', async (req: Request, res: Response) => {

    const userId = req.params.id;
    const headers = {
        'userEmail':`${req.email}`
    }
    
    const response = await axios.patch<User>(`${backendUrl}${USER_BASE_PATH}/${userId}`, { headers });
    const user: User =  response.data;

    res.status(200).json({ user });
});

router.delete('/:id', async (req: Request, res: Response) => {

    const userId = req.params.id;
    const headers = {
        'userEmail':`${req.email}`
    }
    
    let response: AxiosResponse<string> 

    try {
        response = await axios.delete(`${backendUrl}${USER_BASE_PATH}/${userId}`, { headers });
        console.log(response)
    } catch(error) {
       console.log("error")
    }

    res.status(response.status).json({ data: response.data });
});
  

export default router