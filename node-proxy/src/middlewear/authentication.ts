import { Request, Response, NextFunction } from 'express';
import jwt from 'jsonwebtoken';
import { secretKey } from '../config'

interface DecodedToken {
  email: string;
  iat: number;
  exp: number;
}

const authenticateJWT = (req: Request, res: Response, next: NextFunction) => {
  const token = req.header('Authorization');

  if (!token) {
    return res.status(401).json({ message: 'Access token is missing' });
  }

  try {
    // TODO add "public" header to handle users vs guests (e.g, get projects on landing page)
    const decoded = jwt.verify(token, secretKey) as DecodedToken;
    req.email =  decoded.email ; // add more fields as needed
    next();
  } catch (err) {
    res.status(401).json({ message: 'Access token is invalid' });
  }
};

export default authenticateJWT;
