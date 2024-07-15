import express from 'express';
import loginRoutes from './routes/login'
import userRoutes from './routes/user'
import projectRoutes from './routes/project'
import cors from 'cors';

const app = express();
const port = 8000;

app.use(cors());
app.use(express.json());

app.use('/auth', loginRoutes);
app.use('/user', userRoutes);
app.use('/project', projectRoutes);

app.listen(port, () => {
  return console.log(`Express is listening at http://localhost:${port}`);
});