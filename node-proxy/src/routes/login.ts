import express, { Request, Response } from 'express';
import jwt from 'jsonwebtoken';
import axios from 'axios';
import { secretKey, backendUrl } from '../config'
import { UserRequestBody } from '../types/types'

const router = express.Router()

const BASE_PATH = '/api/v1'
  
router.post('/register', async (req: Request<object, object, UserRequestBody>, res: Response) => {
    const { email, password, name, description  }: { email: string; password: string; description: string; name: string } = req.body;

    if (!email || !password) {
        return res.status(400).json({ message: 'Username and password are required' });
    }

    try {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const response = await axios.post(`${backendUrl}${BASE_PATH}/register`, { name: name, password, email, description });

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

    const response = await axios.post(`${backendUrl}${BASE_PATH}/login`, { email, password });

    if (response.status !== 200) {
      return res.status(401).json({ message: 'Invalid username or password' });
    }

    const token = jwt.sign({ email }, secretKey, { expiresIn: '1h' });
    res.status(200).json({ token });
});

export default router