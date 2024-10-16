import './App.css';
import './Pages/HomePage/HomePage';
import HomePage from './Pages/HomePage/HomePage';
import NavBar from './Pages/NavBar/NavBar';
import CreateProject from './Pages/CreateProject/CreateProject';
import ProjectPage from './Pages/ProjectPage/ProjectPage';
import ErrorPage from './error-page';
import User from './Pages/User/User'


import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom";
import { AuthContextProvider } from './context/AuthContext';
import Explore from './Pages/Explore/Explore';
import AboutPage from './Pages/AboutPage/AboutPage';

const AppLayout = () => (
  <div>
    <NavBar />
    <Outlet />
  </div>
);

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
  {
    path: "/",
    element: <HomePage/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/explore",
    element:<Explore/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/about",
    element: <AboutPage/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/project/:id",
    element: <ProjectPage/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/create_project",
    element: <CreateProject/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/user/:id",
    element: <User/>,
    errorElement: <ErrorPage/>
  }
]}]);

function App() {

  return (
    <div>
      <AuthContextProvider>
        {/* <NavBar/> */}
        <RouterProvider router={router}/>
      </AuthContextProvider>
    </div>
  );
}

export default App;
