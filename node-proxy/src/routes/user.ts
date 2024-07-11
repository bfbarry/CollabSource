import express, { Request, Response } from 'express';
import authenticateJWT from '../middlewear/authentication'
import axios from 'axios';
import { backendUrl } from '../config'

const router = express.Router()

router.use(authenticateJWT)

const USER_BASE_PATH = '/api/v1/user'

interface User {
    name: string;
    email: string;
} 

router.get('/:id', async (req: Request, res: Response) => {

    const userId = req.params.id;
    const headers = {
        'userEmail':`${req.email}`
    }
    
    const response = await axios.get<User>(`${backendUrl}${USER_BASE_PATH}/${userId}`, { headers });
    const user: User =  response.data;

    res.status(200).json({ user });
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
    
    try {
       const response =  await axios.delete(`${backendUrl}${USER_BASE_PATH}/${userId}`, { headers });
    } catch(error) {
        res.status(response).json({ user });
    }

    res.status(200).json({ user });
});
  

export default router