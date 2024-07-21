import express, { Request, Response } from 'express';
import jwt from 'jsonwebtoken';
import axios from 'axios';
import { secretKey, backendUrl } from '../config'
import { UserRegisterRequestBody } from '../types/types'

const router = express.Router()

const BASE_PATH = '/api/v1'
  
router.post('/register', async (req: Request<object, object, UserRegisterRequestBody>, res: Response) => {
    const reqBody: UserRegisterRequestBody  = req.body;

    if (!reqBody.email || !reqBody.password) {
        return res.status(400).json({ message: 'Username and password are required' });
    }

    try {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const response = await axios.post(`${backendUrl}${BASE_PATH}/register`, reqBody);

        res.status(201).json({ message: 'User registered successfully' });
   
    } catch (error) {
        if (error.response.status === 422) {
            res.status(error.response.status).json({ message: 'Failed to register user' });
        } 
    }  
});

router.post('/login', async (req: Request, res: Response) => {
    const { email, password }: { email: string; password: string } = req.body;

    // TODO make sure username is a valid email

    if (!email || !password) {
        return res.status(400).json({ message: 'Username and password are required' });
    }

    try {
        await axios.post(`${backendUrl}${BASE_PATH}/login`, { email, password });
    } catch (error) {
        console.log(error)
            return res.status(error.response.status).json({ message: error.response.data });
    }  

    const token = jwt.sign({ email }, secretKey, { expiresIn: '1h' });
    res.status(200).json({ token });
});

export default router