import express, {Request, Response} from 'express';
import authenticateJWT from '../middlewear/authentication';
import { AxiosResponse } from 'axios';
import { axiosBase } from '../config';
import { Project, ProjectRequest, ProjectResponse } from '../types/types';
import cors from 'cors';


const router = express.Router();
router.use(cors());

if (process.env.USE_JWT === 'true') {
  router.use(authenticateJWT)
}
const PROJECT_BASE_PATH = '/api/v1/project'


router.post('/', async (req: Request<object, object, Project>, res: Response) => {

  if (req.email ===  "public") {
    res.status(401).json({ message: 'Access token is invalid' });
  }
  const headers = {
    'UUID': `${req.email}`
  }

  try {
    const project: Project = req.body;
    const response: AxiosResponse<string> = await axiosBase.post<string>(`${PROJECT_BASE_PATH}/create`, project, {headers});
    res.status(200).json(response.data)

  } catch (error) {
    res.status(error.response.status).json({ message: error.response.data})
  }
})

router.get('/:id', async (req: Request, res: Response) => {
  const id = req.params.id;
  const headers = {
    'UUID': `${req.email}`,
    'userId': req.header("userId")
  }
  try {
    const response: AxiosResponse<Project> = await axiosBase.get<Project>(`${PROJECT_BASE_PATH}/${id}`, { headers });
    const project: Project = response.data;
    res.status(200).json({ ...project })

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
    response = await axiosBase.patch<Project>(`${PROJECT_BASE_PATH}/${id}`, project, {headers});
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
    response = await axiosBase.delete(`${PROJECT_BASE_PATH}/${id}`, { headers });
    res.status(response.status).json({ msg: "success" })

  } catch (error) {
      res.status(error.response.status).json({ message: error.response.data})
  }
});

router.post('/project_request/dummy', async (req: Request<object, object, ProjectRequest>, res: Response) => {

  if (req.email ===  "public") {
    res.status(401).json({ message: 'Access token is invalid' });
  }
  const headers = {
    'UUID': `${req.email}`
  }

  try {
    const projectRequest: ProjectRequest = req.body;
    const response: AxiosResponse<string> = await axiosBase.post<string>(`/api/v1/project_request/dummy`, projectRequest, {headers});
    res.status(response.status).json(response.data)

  } catch (error) {
    res.status(error.response.status).json({ message: error.response.data})
  }
})

router.patch('/project_request/:id', async (req: Request<{id: string}, object, ProjectResponse>, res: Response) => {
  const id = req.params.id;
  const headers = {
    'UUID': `${req.email}`
  }

  try {
    const projectResponse: ProjectResponse = req.body;
    const response: AxiosResponse<string> = await axiosBase.patch<string>(`/api/v1/project_request/${id}`, projectResponse, {headers});
    res.status(response.status).json({ msg: "success" })

  } catch (error) {
      res.status(error.response.status).json({ message: error.response.data})
  }
});


export default router