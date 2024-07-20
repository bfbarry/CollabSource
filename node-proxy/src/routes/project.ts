import express, {Request, Response} from 'express';
import authenticateJWT from '../middlewear/authentication';
import axios, { AxiosResponse } from 'axios';
import { backendUrl } from '../config';
import { Project } from '../types/types';


const router = express.Router();
if (process.env.USE_JWT === 'true') {
  router.use(authenticateJWT)
}
const PROJECT_BASE_PATH = '/api/v1/project'


router.post('/', async (req: Request<object, object, Project>, res: Response) => {

  const headers = {
    'UUID': `${req.email}`
  }

  try {
    const project: Project = req.body;
    await axios.post<Project>(`${backendUrl}${PROJECT_BASE_PATH}/create`, project, {headers});
    res.status(200).json({ msg: "success" })

  } catch (error) {
    res.status(error.response.status).json({ message: error.response.data})
  }
})

router.get('/:id', async (req: Request, res: Response) => {
  const id = req.params.id;
  const headers = {
    'UUID': `${req.email}`
  }
  try {
    const response: AxiosResponse<Project> = await axios.get<Project>(`${backendUrl}${PROJECT_BASE_PATH}/${id}`, { headers });
    const project: Project = response.data;
    res.status(200).json({ project })

  } catch (error) {
    res.status(error.response.status).json({ message: error.response.data})
  }
});

router.patch('/:id', async (req: Request<{id: string}, object, Project>, res: Response) => {
  const id = req.params.id;
  const headers = {
    'UUID': `${req.email}`
  }
  let response : AxiosResponse<object>
  try {
    const project: Project = req.body;
    response = await axios.patch<Project>(`${backendUrl}${PROJECT_BASE_PATH}/${id}`, project, {headers});
    res.status(response.status).json({ msg: "success" })

  } catch (error) {
      res.status(error.response.status).json({ message: error.response.data})
  }
});

router.delete('/:id', async (req: Request<{id: string}, object, Project>, res: Response) => {
  const id = req.params.id;
  const headers = {
    'UUID': `${req.email}`
  }
  let response : AxiosResponse<object>
  try {
    response = await axios.delete<Project>(`${backendUrl}${PROJECT_BASE_PATH}/${id}`, { headers });
    res.status(response.status).json({ msg: "success" })

  } catch (error) {
      res.status(error.response.status).json({ message: error.response.data})
  }
});


export default router