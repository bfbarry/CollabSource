import express, {Request, Response} from 'express';
import authenticateJWT from '../middlewear/authentication';
import axios, { AxiosResponse } from 'axios';
import { backendUrl } from '../config';

const router = express.Router();
if (process.env.USE_JWT === 'true') {
  router.use(authenticateJWT)
}
const PROJECT_BASE_PATH = '/api/v1/project'

//TODO move into types
interface Project {
  name       : string;
	description: string;
	category   : string;
	tags       : string[];
}

router.get('/:id', async (req: Request, res: Response) => {
  const id = req.params.id;
  const headers = {
    'userEmail': `${req.email}`
  }
  try {
    const response: AxiosResponse<Project> = await axios.get<Project>(`${backendUrl}${PROJECT_BASE_PATH}/${id}`, { headers });
    const project: Project = response.data;
    res.status(200).json({ project })

  } catch (error) {
    const stat: number = error.response?.status || 500
    switch (stat) {
      case 404: {
        res.status(stat).json({ message: 'Project not found with id'})
        break
      } case 500:
      default:
        res.status(stat).json({ message: 'Internal server error' });
        break;
    } 
  }



})

router.patch('/:id', async (req: Request<any, object, Project>, res: Response) => {
  const id = req.params.id;
  const headers = {
    'userEmail': `${req.email}`
  }

  try {
    const project: Project = req.body;
    await axios.patch<Project>(`${backendUrl}${PROJECT_BASE_PATH}/${id}`, project);
    res.status(200).json({ msg: "success" })

  } catch (error) {
    const stat: number = error.response?.status || 500
    switch (stat) {
      case 400: {
        res.status(stat).json({ message: 'project json is not right shape'})
        break
      } case 404: {
        res.status(stat).json({ message: 'id not found'})
        break
      } case 500:
      default: {
      res.status(stat).json({ message: 'internal server error'})
      break
      }
    } 
  }
})

export default router