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
const PROJECT_BASE_PATH = '/api/v1/project'


router.post('/', async (req: Request<object, object, Project>, res: Response) => {

  const headers = {
    'UUID': `${req.email}`
  }

  try {
    const project: Project = req.body;
    await axiosBase.post<Project>(`${PROJECT_BASE_PATH}/create`, project, {headers});
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
    const response: AxiosResponse<Project> = await axiosBase.get<Project>(`${PROJECT_BASE_PATH}/${id}`, { headers });
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
    response = await axiosBase.delete<Project>(`${PROJECT_BASE_PATH}/${id}`, { headers });
    res.status(response.status).json({ msg: "success" })

  } catch (error) {
      res.status(error.response.status).json({ message: error.response.data})
  }
});


export default router