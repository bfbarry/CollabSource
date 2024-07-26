import express, {Request, Response} from 'express';
import authenticateJWT from '../middlewear/authentication';
import { AxiosResponse } from 'axios';
import { axiosBase } from '../config';
import { Project } from '../types/types';
import cors from 'cors';

const router = express.Router();
router.use(cors());

if (process.env.USE_JWT === 'true') {
  router.use(authenticateJWT)
}
const PROJECT_BASE_PATH = '/api/v1/projects'


router.get('/', async (req: Request, res: Response) => {
  console.log(req.query)
  const headers = {
    'userEmail': `${req.email}`
  }
  const page = req.query.page
  const size = req.query.size
  try {
    const response: AxiosResponse<Project[]> = await axiosBase.get<Project[]>(`${PROJECT_BASE_PATH}?page=${page}&size=${size}`, { headers });
    const project: Project[] = response.data;
    res.status(200).json(project)

  } catch (error) {
    res.status(error.response.status).json({ message: error.response.data})
  }
});
export default router
