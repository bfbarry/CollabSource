import { Response, NextFunction } from 'express';
import { IRequest } from '../types/types';
import jwt from 'jsonwebtoken';
import { secretKey } from '../config'

interface DecodedToken {
  email: string;
  id: string;
  iat: number;
  exp: number;
}

const authenticateJWT = (req: IRequest, res: Response, next: NextFunction) => {
  const token = req.header('Authorization')
  if (token === undefined || token === "public") {
    req.id = "public"
    req.email = "public";
    next();
    return
  }

  try {
    const decoded = jwt.verify(token, secretKey) as DecodedToken;
    req.id =  decoded.id ; // add more fields as needed
    req.email = decoded.email;
    next();
  } catch (err) {
    res.status(401).json({ message: 'Access token is invalid' });
  }
};

export default authenticateJWT;
