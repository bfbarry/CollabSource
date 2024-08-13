import express from 'express';
import authRoutes from './routes/auth'
import userRoutes from './routes/user'
import projectRoutes from './routes/project'
import projectsRoutes from './routes/projects'
import cors from 'cors';

const app = express();
const port = 8000;

app.use(cors());
app.use(express.json());

app.use('/auth', authRoutes);
app.use('/user', userRoutes);
app.use('/project', projectRoutes);
app.use('/projects', projectsRoutes);

app.listen(port, () => {
  return console.log(`Express is listening at http://localhost:${port}`);
});